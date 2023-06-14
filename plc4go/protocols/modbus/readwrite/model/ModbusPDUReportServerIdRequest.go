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

// ModbusPDUReportServerIdRequest is the corresponding interface of ModbusPDUReportServerIdRequest
type ModbusPDUReportServerIdRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
}

// ModbusPDUReportServerIdRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReportServerIdRequest.
// This is useful for switch cases.
type ModbusPDUReportServerIdRequestExactly interface {
	ModbusPDUReportServerIdRequest
	isModbusPDUReportServerIdRequest() bool
}

// _ModbusPDUReportServerIdRequest is the data-structure of this message
type _ModbusPDUReportServerIdRequest struct {
	*_ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReportServerIdRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReportServerIdRequest) GetFunctionFlag() uint8 {
	return 0x11
}

func (m *_ModbusPDUReportServerIdRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReportServerIdRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReportServerIdRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

// NewModbusPDUReportServerIdRequest factory function for _ModbusPDUReportServerIdRequest
func NewModbusPDUReportServerIdRequest() *_ModbusPDUReportServerIdRequest {
	_result := &_ModbusPDUReportServerIdRequest{
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReportServerIdRequest(structType any) ModbusPDUReportServerIdRequest {
	if casted, ok := structType.(ModbusPDUReportServerIdRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReportServerIdRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReportServerIdRequest) GetTypeName() string {
	return "ModbusPDUReportServerIdRequest"
}

func (m *_ModbusPDUReportServerIdRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ModbusPDUReportServerIdRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReportServerIdRequestParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReportServerIdRequest, error) {
	return ModbusPDUReportServerIdRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReportServerIdRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReportServerIdRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReportServerIdRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReportServerIdRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReportServerIdRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReportServerIdRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReportServerIdRequest{
		_ModbusPDU: &_ModbusPDU{},
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReportServerIdRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReportServerIdRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReportServerIdRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReportServerIdRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReportServerIdRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReportServerIdRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReportServerIdRequest) isModbusPDUReportServerIdRequest() bool {
	return true
}

func (m *_ModbusPDUReportServerIdRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
