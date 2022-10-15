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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNameValue is the corresponding interface of BACnetNameValue
type BACnetNameValue interface {
	utils.LengthAware
	utils.Serializable
	// GetName returns Name (property field)
	GetName() BACnetContextTagCharacterString
	// GetValue returns Value (property field)
	GetValue() BACnetConstructedData
}

// BACnetNameValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNameValue.
// This is useful for switch cases.
type BACnetNameValueExactly interface {
	BACnetNameValue
	isBACnetNameValue() bool
}

// _BACnetNameValue is the data-structure of this message
type _BACnetNameValue struct {
	Name  BACnetContextTagCharacterString
	Value BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNameValue) GetName() BACnetContextTagCharacterString {
	return m.Name
}

func (m *_BACnetNameValue) GetValue() BACnetConstructedData {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNameValue factory function for _BACnetNameValue
func NewBACnetNameValue(name BACnetContextTagCharacterString, value BACnetConstructedData) *_BACnetNameValue {
	return &_BACnetNameValue{Name: name, Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetNameValue(structType interface{}) BACnetNameValue {
	if casted, ok := structType.(BACnetNameValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNameValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNameValue) GetTypeName() string {
	return "BACnetNameValue"
}

func (m *_BACnetNameValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNameValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits()

	// Optional Field (value)
	if m.Value != nil {
		lengthInBits += m.Value.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetNameValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNameValueParse(readBuffer utils.ReadBuffer) (BACnetNameValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNameValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNameValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of BACnetNameValue")
	}
	name := _name.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	// Optional Field (value) (Can be skipped, if a given expression evaluates to false)
	var value BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("value"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for value")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(1), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'value' field of BACnetNameValue")
		default:
			value = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for value")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetNameValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNameValue")
	}

	// Create the instance
	return &_BACnetNameValue{
		Name:  name,
		Value: value,
	}, nil
}

func (m *_BACnetNameValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNameValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNameValue")
	}

	// Simple Field (name)
	if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for name")
	}
	_nameErr := writeBuffer.WriteSerializable(m.GetName())
	if popErr := writeBuffer.PopContext("name"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for name")
	}
	if _nameErr != nil {
		return errors.Wrap(_nameErr, "Error serializing 'name' field")
	}

	// Optional Field (value) (Can be skipped, if the value is null)
	var value BACnetConstructedData = nil
	if m.GetValue() != nil {
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		value = m.GetValue()
		_valueErr := writeBuffer.WriteSerializable(value)
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetNameValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNameValue")
	}
	return nil
}

func (m *_BACnetNameValue) isBACnetNameValue() bool {
	return true
}

func (m *_BACnetNameValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
