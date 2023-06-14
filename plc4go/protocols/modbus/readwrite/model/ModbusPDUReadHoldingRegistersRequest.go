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

// ModbusPDUReadHoldingRegistersRequest is the corresponding interface of ModbusPDUReadHoldingRegistersRequest
type ModbusPDUReadHoldingRegistersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
}

// ModbusPDUReadHoldingRegistersRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadHoldingRegistersRequest.
// This is useful for switch cases.
type ModbusPDUReadHoldingRegistersRequestExactly interface {
	ModbusPDUReadHoldingRegistersRequest
	isModbusPDUReadHoldingRegistersRequest() bool
}

// _ModbusPDUReadHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUReadHoldingRegistersRequest struct {
	*_ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x03
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadHoldingRegistersRequest factory function for _ModbusPDUReadHoldingRegistersRequest
func NewModbusPDUReadHoldingRegistersRequest(startingAddress uint16, quantity uint16) *_ModbusPDUReadHoldingRegistersRequest {
	_result := &_ModbusPDUReadHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		_ModbusPDU:      NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadHoldingRegistersRequest(structType any) ModbusPDUReadHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadHoldingRegistersRequest"
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadHoldingRegistersRequestParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadHoldingRegistersRequest, error) {
	return ModbusPDUReadHoldingRegistersRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadHoldingRegistersRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadHoldingRegistersRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUReadHoldingRegistersRequest")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUReadHoldingRegistersRequest")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadHoldingRegistersRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadHoldingRegistersRequest{
		_ModbusPDU:      &_ModbusPDU{},
		StartingAddress: startingAddress,
		Quantity:        quantity,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadHoldingRegistersRequest")
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.GetStartingAddress())
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.GetQuantity())
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadHoldingRegistersRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadHoldingRegistersRequest) isModbusPDUReadHoldingRegistersRequest() bool {
	return true
}

func (m *_ModbusPDUReadHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
