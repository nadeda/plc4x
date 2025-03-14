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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// NPDUControl is the corresponding interface of NPDUControl
type NPDUControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageTypeFieldPresent returns MessageTypeFieldPresent (property field)
	GetMessageTypeFieldPresent() bool
	// GetDestinationSpecified returns DestinationSpecified (property field)
	GetDestinationSpecified() bool
	// GetSourceSpecified returns SourceSpecified (property field)
	GetSourceSpecified() bool
	// GetExpectingReply returns ExpectingReply (property field)
	GetExpectingReply() bool
	// GetNetworkPriority returns NetworkPriority (property field)
	GetNetworkPriority() NPDUNetworkPriority
}

// NPDUControlExactly can be used when we want exactly this type and not a type which fulfills NPDUControl.
// This is useful for switch cases.
type NPDUControlExactly interface {
	NPDUControl
	isNPDUControl() bool
}

// _NPDUControl is the data-structure of this message
type _NPDUControl struct {
	MessageTypeFieldPresent bool
	DestinationSpecified    bool
	SourceSpecified         bool
	ExpectingReply          bool
	NetworkPriority         NPDUNetworkPriority
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NPDUControl) GetMessageTypeFieldPresent() bool {
	return m.MessageTypeFieldPresent
}

func (m *_NPDUControl) GetDestinationSpecified() bool {
	return m.DestinationSpecified
}

func (m *_NPDUControl) GetSourceSpecified() bool {
	return m.SourceSpecified
}

func (m *_NPDUControl) GetExpectingReply() bool {
	return m.ExpectingReply
}

func (m *_NPDUControl) GetNetworkPriority() NPDUNetworkPriority {
	return m.NetworkPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNPDUControl factory function for _NPDUControl
func NewNPDUControl(messageTypeFieldPresent bool, destinationSpecified bool, sourceSpecified bool, expectingReply bool, networkPriority NPDUNetworkPriority) *_NPDUControl {
	return &_NPDUControl{MessageTypeFieldPresent: messageTypeFieldPresent, DestinationSpecified: destinationSpecified, SourceSpecified: sourceSpecified, ExpectingReply: expectingReply, NetworkPriority: networkPriority}
}

// Deprecated: use the interface for direct cast
func CastNPDUControl(structType any) NPDUControl {
	if casted, ok := structType.(NPDUControl); ok {
		return casted
	}
	if casted, ok := structType.(*NPDUControl); ok {
		return *casted
	}
	return nil
}

func (m *_NPDUControl) GetTypeName() string {
	return "NPDUControl"
}

func (m *_NPDUControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (messageTypeFieldPresent)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (destinationSpecified)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (sourceSpecified)
	lengthInBits += 1

	// Simple field (expectingReply)
	lengthInBits += 1

	// Simple field (networkPriority)
	lengthInBits += 2

	return lengthInBits
}

func (m *_NPDUControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NPDUControlParse(ctx context.Context, theBytes []byte) (NPDUControl, error) {
	return NPDUControlParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NPDUControlParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NPDUControl, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NPDUControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NPDUControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (messageTypeFieldPresent)
	_messageTypeFieldPresent, _messageTypeFieldPresentErr := readBuffer.ReadBit("messageTypeFieldPresent")
	if _messageTypeFieldPresentErr != nil {
		return nil, errors.Wrap(_messageTypeFieldPresentErr, "Error parsing 'messageTypeFieldPresent' field of NPDUControl")
	}
	messageTypeFieldPresent := _messageTypeFieldPresent

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NPDUControl")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (destinationSpecified)
	_destinationSpecified, _destinationSpecifiedErr := readBuffer.ReadBit("destinationSpecified")
	if _destinationSpecifiedErr != nil {
		return nil, errors.Wrap(_destinationSpecifiedErr, "Error parsing 'destinationSpecified' field of NPDUControl")
	}
	destinationSpecified := _destinationSpecified

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NPDUControl")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (sourceSpecified)
	_sourceSpecified, _sourceSpecifiedErr := readBuffer.ReadBit("sourceSpecified")
	if _sourceSpecifiedErr != nil {
		return nil, errors.Wrap(_sourceSpecifiedErr, "Error parsing 'sourceSpecified' field of NPDUControl")
	}
	sourceSpecified := _sourceSpecified

	// Simple Field (expectingReply)
	_expectingReply, _expectingReplyErr := readBuffer.ReadBit("expectingReply")
	if _expectingReplyErr != nil {
		return nil, errors.Wrap(_expectingReplyErr, "Error parsing 'expectingReply' field of NPDUControl")
	}
	expectingReply := _expectingReply

	// Simple Field (networkPriority)
	if pullErr := readBuffer.PullContext("networkPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkPriority")
	}
	_networkPriority, _networkPriorityErr := NPDUNetworkPriorityParseWithBuffer(ctx, readBuffer)
	if _networkPriorityErr != nil {
		return nil, errors.Wrap(_networkPriorityErr, "Error parsing 'networkPriority' field of NPDUControl")
	}
	networkPriority := _networkPriority
	if closeErr := readBuffer.CloseContext("networkPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkPriority")
	}

	if closeErr := readBuffer.CloseContext("NPDUControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NPDUControl")
	}

	// Create the instance
	return &_NPDUControl{
		MessageTypeFieldPresent: messageTypeFieldPresent,
		DestinationSpecified:    destinationSpecified,
		SourceSpecified:         sourceSpecified,
		ExpectingReply:          expectingReply,
		NetworkPriority:         networkPriority,
		reservedField0:          reservedField0,
		reservedField1:          reservedField1,
	}, nil
}

func (m *_NPDUControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NPDUControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NPDUControl"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NPDUControl")
	}

	// Simple Field (messageTypeFieldPresent)
	messageTypeFieldPresent := bool(m.GetMessageTypeFieldPresent())
	_messageTypeFieldPresentErr := writeBuffer.WriteBit("messageTypeFieldPresent", (messageTypeFieldPresent))
	if _messageTypeFieldPresentErr != nil {
		return errors.Wrap(_messageTypeFieldPresentErr, "Error serializing 'messageTypeFieldPresent' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 1, uint8(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (destinationSpecified)
	destinationSpecified := bool(m.GetDestinationSpecified())
	_destinationSpecifiedErr := writeBuffer.WriteBit("destinationSpecified", (destinationSpecified))
	if _destinationSpecifiedErr != nil {
		return errors.Wrap(_destinationSpecifiedErr, "Error serializing 'destinationSpecified' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0)
		if m.reservedField1 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteUint8("reserved", 1, uint8(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (sourceSpecified)
	sourceSpecified := bool(m.GetSourceSpecified())
	_sourceSpecifiedErr := writeBuffer.WriteBit("sourceSpecified", (sourceSpecified))
	if _sourceSpecifiedErr != nil {
		return errors.Wrap(_sourceSpecifiedErr, "Error serializing 'sourceSpecified' field")
	}

	// Simple Field (expectingReply)
	expectingReply := bool(m.GetExpectingReply())
	_expectingReplyErr := writeBuffer.WriteBit("expectingReply", (expectingReply))
	if _expectingReplyErr != nil {
		return errors.Wrap(_expectingReplyErr, "Error serializing 'expectingReply' field")
	}

	// Simple Field (networkPriority)
	if pushErr := writeBuffer.PushContext("networkPriority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkPriority")
	}
	_networkPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetNetworkPriority())
	if popErr := writeBuffer.PopContext("networkPriority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkPriority")
	}
	if _networkPriorityErr != nil {
		return errors.Wrap(_networkPriorityErr, "Error serializing 'networkPriority' field")
	}

	if popErr := writeBuffer.PopContext("NPDUControl"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NPDUControl")
	}
	return nil
}

func (m *_NPDUControl) isNPDUControl() bool {
	return true
}

func (m *_NPDUControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
