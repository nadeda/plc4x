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

// StatusByte is the corresponding interface of StatusByte
type StatusByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetGav3 returns Gav3 (property field)
	GetGav3() GAVState
	// GetGav2 returns Gav2 (property field)
	GetGav2() GAVState
	// GetGav1 returns Gav1 (property field)
	GetGav1() GAVState
	// GetGav0 returns Gav0 (property field)
	GetGav0() GAVState
}

// StatusByteExactly can be used when we want exactly this type and not a type which fulfills StatusByte.
// This is useful for switch cases.
type StatusByteExactly interface {
	StatusByte
	isStatusByte() bool
}

// _StatusByte is the data-structure of this message
type _StatusByte struct {
	Gav3 GAVState
	Gav2 GAVState
	Gav1 GAVState
	Gav0 GAVState
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusByte) GetGav3() GAVState {
	return m.Gav3
}

func (m *_StatusByte) GetGav2() GAVState {
	return m.Gav2
}

func (m *_StatusByte) GetGav1() GAVState {
	return m.Gav1
}

func (m *_StatusByte) GetGav0() GAVState {
	return m.Gav0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusByte factory function for _StatusByte
func NewStatusByte(gav3 GAVState, gav2 GAVState, gav1 GAVState, gav0 GAVState) *_StatusByte {
	return &_StatusByte{Gav3: gav3, Gav2: gav2, Gav1: gav1, Gav0: gav0}
}

// Deprecated: use the interface for direct cast
func CastStatusByte(structType any) StatusByte {
	if casted, ok := structType.(StatusByte); ok {
		return casted
	}
	if casted, ok := structType.(*StatusByte); ok {
		return *casted
	}
	return nil
}

func (m *_StatusByte) GetTypeName() string {
	return "StatusByte"
}

func (m *_StatusByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (gav3)
	lengthInBits += 2

	// Simple field (gav2)
	lengthInBits += 2

	// Simple field (gav1)
	lengthInBits += 2

	// Simple field (gav0)
	lengthInBits += 2

	return lengthInBits
}

func (m *_StatusByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusByteParse(ctx context.Context, theBytes []byte) (StatusByte, error) {
	return StatusByteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusByteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StatusByte, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (gav3)
	if pullErr := readBuffer.PullContext("gav3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for gav3")
	}
	_gav3, _gav3Err := GAVStateParseWithBuffer(ctx, readBuffer)
	if _gav3Err != nil {
		return nil, errors.Wrap(_gav3Err, "Error parsing 'gav3' field of StatusByte")
	}
	gav3 := _gav3
	if closeErr := readBuffer.CloseContext("gav3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for gav3")
	}

	// Simple Field (gav2)
	if pullErr := readBuffer.PullContext("gav2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for gav2")
	}
	_gav2, _gav2Err := GAVStateParseWithBuffer(ctx, readBuffer)
	if _gav2Err != nil {
		return nil, errors.Wrap(_gav2Err, "Error parsing 'gav2' field of StatusByte")
	}
	gav2 := _gav2
	if closeErr := readBuffer.CloseContext("gav2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for gav2")
	}

	// Simple Field (gav1)
	if pullErr := readBuffer.PullContext("gav1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for gav1")
	}
	_gav1, _gav1Err := GAVStateParseWithBuffer(ctx, readBuffer)
	if _gav1Err != nil {
		return nil, errors.Wrap(_gav1Err, "Error parsing 'gav1' field of StatusByte")
	}
	gav1 := _gav1
	if closeErr := readBuffer.CloseContext("gav1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for gav1")
	}

	// Simple Field (gav0)
	if pullErr := readBuffer.PullContext("gav0"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for gav0")
	}
	_gav0, _gav0Err := GAVStateParseWithBuffer(ctx, readBuffer)
	if _gav0Err != nil {
		return nil, errors.Wrap(_gav0Err, "Error parsing 'gav0' field of StatusByte")
	}
	gav0 := _gav0
	if closeErr := readBuffer.CloseContext("gav0"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for gav0")
	}

	if closeErr := readBuffer.CloseContext("StatusByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusByte")
	}

	// Create the instance
	return &_StatusByte{
		Gav3: gav3,
		Gav2: gav2,
		Gav1: gav1,
		Gav0: gav0,
	}, nil
}

func (m *_StatusByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("StatusByte"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StatusByte")
	}

	// Simple Field (gav3)
	if pushErr := writeBuffer.PushContext("gav3"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for gav3")
	}
	_gav3Err := writeBuffer.WriteSerializable(ctx, m.GetGav3())
	if popErr := writeBuffer.PopContext("gav3"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for gav3")
	}
	if _gav3Err != nil {
		return errors.Wrap(_gav3Err, "Error serializing 'gav3' field")
	}

	// Simple Field (gav2)
	if pushErr := writeBuffer.PushContext("gav2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for gav2")
	}
	_gav2Err := writeBuffer.WriteSerializable(ctx, m.GetGav2())
	if popErr := writeBuffer.PopContext("gav2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for gav2")
	}
	if _gav2Err != nil {
		return errors.Wrap(_gav2Err, "Error serializing 'gav2' field")
	}

	// Simple Field (gav1)
	if pushErr := writeBuffer.PushContext("gav1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for gav1")
	}
	_gav1Err := writeBuffer.WriteSerializable(ctx, m.GetGav1())
	if popErr := writeBuffer.PopContext("gav1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for gav1")
	}
	if _gav1Err != nil {
		return errors.Wrap(_gav1Err, "Error serializing 'gav1' field")
	}

	// Simple Field (gav0)
	if pushErr := writeBuffer.PushContext("gav0"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for gav0")
	}
	_gav0Err := writeBuffer.WriteSerializable(ctx, m.GetGav0())
	if popErr := writeBuffer.PopContext("gav0"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for gav0")
	}
	if _gav0Err != nil {
		return errors.Wrap(_gav0Err, "Error serializing 'gav0' field")
	}

	if popErr := writeBuffer.PopContext("StatusByte"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StatusByte")
	}
	return nil
}

func (m *_StatusByte) isStatusByte() bool {
	return true
}

func (m *_StatusByte) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
