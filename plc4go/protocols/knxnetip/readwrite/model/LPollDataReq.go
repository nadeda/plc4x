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

// LPollDataReq is the corresponding interface of LPollDataReq
type LPollDataReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// LPollDataReqExactly can be used when we want exactly this type and not a type which fulfills LPollDataReq.
// This is useful for switch cases.
type LPollDataReqExactly interface {
	LPollDataReq
	isLPollDataReq() bool
}

// _LPollDataReq is the data-structure of this message
type _LPollDataReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LPollDataReq) GetMessageCode() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LPollDataReq) InitializeParent(parent CEMI) {}

func (m *_LPollDataReq) GetParent() CEMI {
	return m._CEMI
}

// NewLPollDataReq factory function for _LPollDataReq
func NewLPollDataReq(size uint16) *_LPollDataReq {
	_result := &_LPollDataReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLPollDataReq(structType any) LPollDataReq {
	if casted, ok := structType.(LPollDataReq); ok {
		return casted
	}
	if casted, ok := structType.(*LPollDataReq); ok {
		return *casted
	}
	return nil
}

func (m *_LPollDataReq) GetTypeName() string {
	return "LPollDataReq"
}

func (m *_LPollDataReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_LPollDataReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LPollDataReqParse(ctx context.Context, theBytes []byte, size uint16) (LPollDataReq, error) {
	return LPollDataReqParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func LPollDataReqParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (LPollDataReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LPollDataReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LPollDataReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LPollDataReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LPollDataReq")
	}

	// Create a partially initialized instance
	_child := &_LPollDataReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LPollDataReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LPollDataReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollDataReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LPollDataReq")
		}

		if popErr := writeBuffer.PopContext("LPollDataReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LPollDataReq")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LPollDataReq) isLPollDataReq() bool {
	return true
}

func (m *_LPollDataReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
