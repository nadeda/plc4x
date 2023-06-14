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

// SerialNumber is the corresponding interface of SerialNumber
type SerialNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOctet1 returns Octet1 (property field)
	GetOctet1() byte
	// GetOctet2 returns Octet2 (property field)
	GetOctet2() byte
	// GetOctet3 returns Octet3 (property field)
	GetOctet3() byte
	// GetOctet4 returns Octet4 (property field)
	GetOctet4() byte
}

// SerialNumberExactly can be used when we want exactly this type and not a type which fulfills SerialNumber.
// This is useful for switch cases.
type SerialNumberExactly interface {
	SerialNumber
	isSerialNumber() bool
}

// _SerialNumber is the data-structure of this message
type _SerialNumber struct {
	Octet1 byte
	Octet2 byte
	Octet3 byte
	Octet4 byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SerialNumber) GetOctet1() byte {
	return m.Octet1
}

func (m *_SerialNumber) GetOctet2() byte {
	return m.Octet2
}

func (m *_SerialNumber) GetOctet3() byte {
	return m.Octet3
}

func (m *_SerialNumber) GetOctet4() byte {
	return m.Octet4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSerialNumber factory function for _SerialNumber
func NewSerialNumber(octet1 byte, octet2 byte, octet3 byte, octet4 byte) *_SerialNumber {
	return &_SerialNumber{Octet1: octet1, Octet2: octet2, Octet3: octet3, Octet4: octet4}
}

// Deprecated: use the interface for direct cast
func CastSerialNumber(structType any) SerialNumber {
	if casted, ok := structType.(SerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*SerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_SerialNumber) GetTypeName() string {
	return "SerialNumber"
}

func (m *_SerialNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (octet1)
	lengthInBits += 8

	// Simple field (octet2)
	lengthInBits += 8

	// Simple field (octet3)
	lengthInBits += 8

	// Simple field (octet4)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SerialNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SerialNumberParse(ctx context.Context, theBytes []byte) (SerialNumber, error) {
	return SerialNumberParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SerialNumberParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SerialNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octet1)
	_octet1, _octet1Err := readBuffer.ReadByte("octet1")
	if _octet1Err != nil {
		return nil, errors.Wrap(_octet1Err, "Error parsing 'octet1' field of SerialNumber")
	}
	octet1 := _octet1

	// Simple Field (octet2)
	_octet2, _octet2Err := readBuffer.ReadByte("octet2")
	if _octet2Err != nil {
		return nil, errors.Wrap(_octet2Err, "Error parsing 'octet2' field of SerialNumber")
	}
	octet2 := _octet2

	// Simple Field (octet3)
	_octet3, _octet3Err := readBuffer.ReadByte("octet3")
	if _octet3Err != nil {
		return nil, errors.Wrap(_octet3Err, "Error parsing 'octet3' field of SerialNumber")
	}
	octet3 := _octet3

	// Simple Field (octet4)
	_octet4, _octet4Err := readBuffer.ReadByte("octet4")
	if _octet4Err != nil {
		return nil, errors.Wrap(_octet4Err, "Error parsing 'octet4' field of SerialNumber")
	}
	octet4 := _octet4

	if closeErr := readBuffer.CloseContext("SerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SerialNumber")
	}

	// Create the instance
	return &_SerialNumber{
		Octet1: octet1,
		Octet2: octet2,
		Octet3: octet3,
		Octet4: octet4,
	}, nil
}

func (m *_SerialNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SerialNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SerialNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SerialNumber")
	}

	// Simple Field (octet1)
	octet1 := byte(m.GetOctet1())
	_octet1Err := writeBuffer.WriteByte("octet1", (octet1))
	if _octet1Err != nil {
		return errors.Wrap(_octet1Err, "Error serializing 'octet1' field")
	}

	// Simple Field (octet2)
	octet2 := byte(m.GetOctet2())
	_octet2Err := writeBuffer.WriteByte("octet2", (octet2))
	if _octet2Err != nil {
		return errors.Wrap(_octet2Err, "Error serializing 'octet2' field")
	}

	// Simple Field (octet3)
	octet3 := byte(m.GetOctet3())
	_octet3Err := writeBuffer.WriteByte("octet3", (octet3))
	if _octet3Err != nil {
		return errors.Wrap(_octet3Err, "Error serializing 'octet3' field")
	}

	// Simple Field (octet4)
	octet4 := byte(m.GetOctet4())
	_octet4Err := writeBuffer.WriteByte("octet4", (octet4))
	if _octet4Err != nil {
		return errors.Wrap(_octet4Err, "Error serializing 'octet4' field")
	}

	if popErr := writeBuffer.PopContext("SerialNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SerialNumber")
	}
	return nil
}

func (m *_SerialNumber) isSerialNumber() bool {
	return true
}

func (m *_SerialNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
