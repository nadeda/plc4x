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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLimitMonitoringInterval is the corresponding interface of BACnetConstructedDataLimitMonitoringInterval
type BACnetConstructedDataLimitMonitoringInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLimitMonitoringInterval returns LimitMonitoringInterval (property field)
	GetLimitMonitoringInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLimitMonitoringIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLimitMonitoringInterval.
// This is useful for switch cases.
type BACnetConstructedDataLimitMonitoringIntervalExactly interface {
	BACnetConstructedDataLimitMonitoringInterval
	isBACnetConstructedDataLimitMonitoringInterval() bool
}

// _BACnetConstructedDataLimitMonitoringInterval is the data-structure of this message
type _BACnetConstructedDataLimitMonitoringInterval struct {
	*_BACnetConstructedData
	LimitMonitoringInterval BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIMIT_MONITORING_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLimitMonitoringInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetLimitMonitoringInterval() BACnetApplicationTagUnsignedInteger {
	return m.LimitMonitoringInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLimitMonitoringInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLimitMonitoringInterval factory function for _BACnetConstructedDataLimitMonitoringInterval
func NewBACnetConstructedDataLimitMonitoringInterval(limitMonitoringInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLimitMonitoringInterval {
	_result := &_BACnetConstructedDataLimitMonitoringInterval{
		LimitMonitoringInterval: limitMonitoringInterval,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLimitMonitoringInterval(structType any) BACnetConstructedDataLimitMonitoringInterval {
	if casted, ok := structType.(BACnetConstructedDataLimitMonitoringInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLimitMonitoringInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetTypeName() string {
	return "BACnetConstructedDataLimitMonitoringInterval"
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (limitMonitoringInterval)
	lengthInBits += m.LimitMonitoringInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLimitMonitoringIntervalParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLimitMonitoringInterval, error) {
	return BACnetConstructedDataLimitMonitoringIntervalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLimitMonitoringIntervalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLimitMonitoringInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLimitMonitoringInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLimitMonitoringInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (limitMonitoringInterval)
	if pullErr := readBuffer.PullContext("limitMonitoringInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for limitMonitoringInterval")
	}
	_limitMonitoringInterval, _limitMonitoringIntervalErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _limitMonitoringIntervalErr != nil {
		return nil, errors.Wrap(_limitMonitoringIntervalErr, "Error parsing 'limitMonitoringInterval' field of BACnetConstructedDataLimitMonitoringInterval")
	}
	limitMonitoringInterval := _limitMonitoringInterval.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("limitMonitoringInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for limitMonitoringInterval")
	}

	// Virtual field
	_actualValue := limitMonitoringInterval
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLimitMonitoringInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLimitMonitoringInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLimitMonitoringInterval{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LimitMonitoringInterval: limitMonitoringInterval,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLimitMonitoringInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLimitMonitoringInterval")
		}

		// Simple Field (limitMonitoringInterval)
		if pushErr := writeBuffer.PushContext("limitMonitoringInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for limitMonitoringInterval")
		}
		_limitMonitoringIntervalErr := writeBuffer.WriteSerializable(ctx, m.GetLimitMonitoringInterval())
		if popErr := writeBuffer.PopContext("limitMonitoringInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for limitMonitoringInterval")
		}
		if _limitMonitoringIntervalErr != nil {
			return errors.Wrap(_limitMonitoringIntervalErr, "Error serializing 'limitMonitoringInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLimitMonitoringInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLimitMonitoringInterval")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) isBACnetConstructedDataLimitMonitoringInterval() bool {
	return true
}

func (m *_BACnetConstructedDataLimitMonitoringInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
