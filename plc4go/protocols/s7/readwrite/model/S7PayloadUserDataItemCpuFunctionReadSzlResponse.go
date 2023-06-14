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

// S7PayloadUserDataItemCpuFunctionReadSzlResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlResponse
type S7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetItems returns Items (property field)
	GetItems() []byte
}

// S7PayloadUserDataItemCpuFunctionReadSzlResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionReadSzlResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionReadSzlResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionReadSzlResponse
	isS7PayloadUserDataItemCpuFunctionReadSzlResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionReadSzlResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
	*_S7PayloadUserDataItem
	Items []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetItems() []byte {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionReadSzlResponse factory function for _S7PayloadUserDataItemCpuFunctionReadSzlResponse
func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(items []byte, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		Items:                  items,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType any) S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		lengthInBits += 8 * uint16(len(m.Items))
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(ctx context.Context, theBytes []byte, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	return S7PayloadUserDataItemCpuFunctionReadSzlResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (items)
	numberOfBytesitems := int(dataLength)
	items, _readArrayErr := readBuffer.ReadByteArray("items", numberOfBytesitems)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'items' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Items:                  items,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}

		// Array Field (items)
		// Byte Array field (items)
		if err := writeBuffer.WriteByteArray("items", m.GetItems()); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) isS7PayloadUserDataItemCpuFunctionReadSzlResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
