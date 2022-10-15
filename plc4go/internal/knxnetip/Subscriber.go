/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package knxnetip

import (
	"context"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	driverModel "github.com/apache/plc4x/plc4go/protocols/knxnetip/readwrite/model"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	values2 "github.com/apache/plc4x/plc4go/spi/values"
	"time"
)

type Subscriber struct {
	connection *Connection
	consumers  map[*spiModel.DefaultPlcConsumerRegistration]apiModel.PlcSubscriptionEventConsumer
}

func NewSubscriber(connection *Connection) *Subscriber {
	return &Subscriber{
		connection: connection,
		consumers:  make(map[*spiModel.DefaultPlcConsumerRegistration]apiModel.PlcSubscriptionEventConsumer),
	}
}

func (m *Subscriber) Subscribe(ctx context.Context, subscriptionRequest apiModel.PlcSubscriptionRequest) <-chan apiModel.PlcSubscriptionRequestResult {
	// TODO: handle context
	result := make(chan apiModel.PlcSubscriptionRequestResult)
	go func() {
		internalPlcSubscriptionRequest := subscriptionRequest.(*spiModel.DefaultPlcSubscriptionRequest)

		// Add this subscriber to the connection.
		m.connection.addSubscriber(m)

		// Just populate all requests with an OK
		responseCodes := map[string]apiModel.PlcResponseCode{}
		subscriptionValues := make(map[string]apiModel.PlcSubscriptionHandle)
		for _, fieldName := range internalPlcSubscriptionRequest.GetFieldNames() {
			responseCodes[fieldName] = apiModel.PlcResponseCode_OK
			fieldType := internalPlcSubscriptionRequest.GetType(fieldName)
			subscriptionValues[fieldName] = NewSubscriptionHandle(m, fieldName, internalPlcSubscriptionRequest.GetField(fieldName), fieldType, internalPlcSubscriptionRequest.GetInterval(fieldName))
		}

		result <- &spiModel.DefaultPlcSubscriptionRequestResult{
			Request:  subscriptionRequest,
			Response: spiModel.NewDefaultPlcSubscriptionResponse(subscriptionRequest, responseCodes, subscriptionValues),
			Err:      nil,
		}
	}()
	return result
}

func (m *Subscriber) Unsubscribe(ctx context.Context, unsubscriptionRequest apiModel.PlcUnsubscriptionRequest) <-chan apiModel.PlcUnsubscriptionRequestResult {
	// TODO: handle context
	result := make(chan apiModel.PlcUnsubscriptionRequestResult)

	// TODO: As soon as we establish a connection, we start getting data...
	// subscriptions are more an internal handling of which values to pass where.

	return result
}

/*
 * Callback for incoming value change events from the KNX bus
 */
func (m *Subscriber) handleValueChange(destinationAddress []byte, payload []byte, changed bool) {
	// Decode the group-address according to the settings in the driver
	// Group addresses can be 1, 2 or 3 levels (3 being the default)
	garb := utils.NewReadBufferByteBased(destinationAddress)
	groupAddress, err := driverModel.KnxGroupAddressParse(garb, m.connection.getGroupAddressNumLevels())
	if err != nil {
		return
	}

	// TODO: aggregate fields and send it to a consumer which want's all of them
	for registration, consumer := range m.consumers {
		for _, subscriptionHandle := range registration.GetSubscriptionHandles() {
			subscriptionHandle := subscriptionHandle.(*SubscriptionHandle)
			groupAddressField, ok := subscriptionHandle.field.(GroupAddressField)
			if !ok || !groupAddressField.matches(groupAddress) {
				continue
			}
			if subscriptionHandle.fieldType != spiModel.SubscriptionChangeOfState || !changed {
				continue
			}
			fields := map[string]apiModel.PlcField{}
			types := map[string]spiModel.SubscriptionType{}
			intervals := map[string]time.Duration{}
			responseCodes := map[string]apiModel.PlcResponseCode{}
			addresses := map[string][]byte{}
			plcValues := map[string]values.PlcValue{}
			fieldName := subscriptionHandle.fieldName
			rb := utils.NewReadBufferByteBased(payload)
			if groupAddressField.GetFieldType() == nil {
				responseCodes[fieldName] = apiModel.PlcResponseCode_INVALID_DATATYPE
				plcValues[fieldName] = nil
				continue
			}
			// If the size of the field is greater than 6, we have to skip the first byte
			if groupAddressField.GetFieldType().GetLengthInBits() > 6 {
				_, _ = rb.ReadUint8("groupAddress", 8)
			}
			elementType := *groupAddressField.GetFieldType()
			numElements := groupAddressField.GetQuantity()

			fields[fieldName] = groupAddressField
			types[fieldName] = subscriptionHandle.fieldType
			intervals[fieldName] = subscriptionHandle.interval
			addresses[fieldName] = destinationAddress

			var plcValueList []values.PlcValue
			responseCode := apiModel.PlcResponseCode_OK
			for i := uint16(0); i < numElements; i++ {
				// If we don't know the datatype, we'll create a RawPlcValue instead
				// so the application can decode the content later on.
				if elementType == driverModel.KnxDatapointType_DPT_UNKNOWN {
					// If this is an unknown 1 byte payload, we need the first byte.
					if !rb.HasMore(1) {
						rb.Reset(0)
					}
					plcValue := values2.NewRawPlcValue(rb, NewValueDecoder(rb))
					plcValueList = append(plcValueList, plcValue)
				} else {
					plcValue, err2 := driverModel.KnxDatapointParse(rb, elementType)
					if err2 == nil {
						plcValueList = append(plcValueList, plcValue)
					} else {
						// TODO: Do a little more here ...
						responseCode = apiModel.PlcResponseCode_INTERNAL_ERROR
						break
					}
				}
			}
			responseCodes[fieldName] = responseCode
			if responseCode == apiModel.PlcResponseCode_OK {
				if len(plcValueList) == 1 {
					plcValues[fieldName] = plcValueList[0]
				} else {
					plcValues[fieldName] = values2.NewPlcList(plcValueList)
				}
			}
			event := NewSubscriptionEvent(fields, types, intervals, responseCodes, addresses, plcValues)
			consumer(&event)
		}
	}
}

func (m *Subscriber) Register(consumer apiModel.PlcSubscriptionEventConsumer, handles []apiModel.PlcSubscriptionHandle) apiModel.PlcConsumerRegistration {
	consumerRegistration := spiModel.NewDefaultPlcConsumerRegistration(m, consumer, handles...)
	m.consumers[consumerRegistration] = consumer
	return consumerRegistration
}

func (m *Subscriber) Unregister(registration apiModel.PlcConsumerRegistration) {
	delete(m.consumers, registration.(*spiModel.DefaultPlcConsumerRegistration))
}
