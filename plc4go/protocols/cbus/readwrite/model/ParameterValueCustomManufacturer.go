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

// ParameterValueCustomManufacturer is the corresponding interface of ParameterValueCustomManufacturer
type ParameterValueCustomManufacturer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() CustomManufacturer
}

// ParameterValueCustomManufacturerExactly can be used when we want exactly this type and not a type which fulfills ParameterValueCustomManufacturer.
// This is useful for switch cases.
type ParameterValueCustomManufacturerExactly interface {
	ParameterValueCustomManufacturer
	isParameterValueCustomManufacturer() bool
}

// _ParameterValueCustomManufacturer is the data-structure of this message
type _ParameterValueCustomManufacturer struct {
	*_ParameterValue
	Value CustomManufacturer
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueCustomManufacturer) GetParameterType() ParameterType {
	return ParameterType_CUSTOM_MANUFACTURER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueCustomManufacturer) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueCustomManufacturer) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueCustomManufacturer) GetValue() CustomManufacturer {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueCustomManufacturer factory function for _ParameterValueCustomManufacturer
func NewParameterValueCustomManufacturer(value CustomManufacturer, numBytes uint8) *_ParameterValueCustomManufacturer {
	_result := &_ParameterValueCustomManufacturer{
		Value:           value,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueCustomManufacturer(structType any) ParameterValueCustomManufacturer {
	if casted, ok := structType.(ParameterValueCustomManufacturer); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueCustomManufacturer); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueCustomManufacturer) GetTypeName() string {
	return "ParameterValueCustomManufacturer"
}

func (m *_ParameterValueCustomManufacturer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ParameterValueCustomManufacturer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueCustomManufacturerParse(ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueCustomManufacturer, error) {
	return ParameterValueCustomManufacturerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueCustomManufacturerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueCustomManufacturer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueCustomManufacturer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueCustomManufacturer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := CustomManufacturerParseWithBuffer(ctx, readBuffer, uint8(numBytes))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueCustomManufacturer")
	}
	value := _value.(CustomManufacturer)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueCustomManufacturer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueCustomManufacturer")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueCustomManufacturer{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueCustomManufacturer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueCustomManufacturer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueCustomManufacturer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueCustomManufacturer")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueCustomManufacturer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueCustomManufacturer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueCustomManufacturer) isParameterValueCustomManufacturer() bool {
	return true
}

func (m *_ParameterValueCustomManufacturer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
