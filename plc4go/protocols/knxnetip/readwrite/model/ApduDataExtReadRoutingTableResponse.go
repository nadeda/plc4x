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

// ApduDataExtReadRoutingTableResponse is the corresponding interface of ApduDataExtReadRoutingTableResponse
type ApduDataExtReadRoutingTableResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtReadRoutingTableResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtReadRoutingTableResponse.
// This is useful for switch cases.
type ApduDataExtReadRoutingTableResponseExactly interface {
	ApduDataExtReadRoutingTableResponse
	isApduDataExtReadRoutingTableResponse() bool
}

// _ApduDataExtReadRoutingTableResponse is the data-structure of this message
type _ApduDataExtReadRoutingTableResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRoutingTableResponse) GetExtApciType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRoutingTableResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtReadRoutingTableResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtReadRoutingTableResponse factory function for _ApduDataExtReadRoutingTableResponse
func NewApduDataExtReadRoutingTableResponse(length uint8) *_ApduDataExtReadRoutingTableResponse {
	_result := &_ApduDataExtReadRoutingTableResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRoutingTableResponse(structType any) ApduDataExtReadRoutingTableResponse {
	if casted, ok := structType.(ApduDataExtReadRoutingTableResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRoutingTableResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRoutingTableResponse) GetTypeName() string {
	return "ApduDataExtReadRoutingTableResponse"
}

func (m *_ApduDataExtReadRoutingTableResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRoutingTableResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtReadRoutingTableResponseParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtReadRoutingTableResponse, error) {
	return ApduDataExtReadRoutingTableResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtReadRoutingTableResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtReadRoutingTableResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRoutingTableResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRoutingTableResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRoutingTableResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRoutingTableResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtReadRoutingTableResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtReadRoutingTableResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRoutingTableResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRoutingTableResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRoutingTableResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRoutingTableResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRoutingTableResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRoutingTableResponse) isApduDataExtReadRoutingTableResponse() bool {
	return true
}

func (m *_ApduDataExtReadRoutingTableResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
