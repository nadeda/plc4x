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

// Checksum is the corresponding interface of Checksum
type Checksum interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() byte
}

// ChecksumExactly can be used when we want exactly this type and not a type which fulfills Checksum.
// This is useful for switch cases.
type ChecksumExactly interface {
	Checksum
	isChecksum() bool
}

// _Checksum is the data-structure of this message
type _Checksum struct {
	Value byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Checksum) GetValue() byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChecksum factory function for _Checksum
func NewChecksum(value byte) *_Checksum {
	return &_Checksum{Value: value}
}

// Deprecated: use the interface for direct cast
func CastChecksum(structType any) Checksum {
	if casted, ok := structType.(Checksum); ok {
		return casted
	}
	if casted, ok := structType.(*Checksum); ok {
		return *casted
	}
	return nil
}

func (m *_Checksum) GetTypeName() string {
	return "Checksum"
}

func (m *_Checksum) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *_Checksum) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChecksumParse(ctx context.Context, theBytes []byte) (Checksum, error) {
	return ChecksumParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ChecksumParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Checksum, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Checksum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Checksum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadByte("value")
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of Checksum")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("Checksum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Checksum")
	}

	// Create the instance
	return &_Checksum{
		Value: value,
	}, nil
}

func (m *_Checksum) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Checksum) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Checksum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Checksum")
	}

	// Simple Field (value)
	value := byte(m.GetValue())
	_valueErr := writeBuffer.WriteByte("value", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("Checksum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Checksum")
	}
	return nil
}

func (m *_Checksum) isChecksum() bool {
	return true
}

func (m *_Checksum) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
