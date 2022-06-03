/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetTimerStateChangeValueCharacterString is the data-structure of this message
type BACnetTimerStateChangeValueCharacterString struct {
	*BACnetTimerStateChangeValue
	CharacterStringValue *BACnetApplicationTagCharacterString

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

// IBACnetTimerStateChangeValueCharacterString is the corresponding interface of BACnetTimerStateChangeValueCharacterString
type IBACnetTimerStateChangeValueCharacterString interface {
	IBACnetTimerStateChangeValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetTimerStateChangeValueCharacterString) InitializeParent(parent *BACnetTimerStateChangeValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetTimerStateChangeValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetTimerStateChangeValueCharacterString) GetParent() *BACnetTimerStateChangeValue {
	return m.BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTimerStateChangeValueCharacterString) GetCharacterStringValue() *BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueCharacterString factory function for BACnetTimerStateChangeValueCharacterString
func NewBACnetTimerStateChangeValueCharacterString(characterStringValue *BACnetApplicationTagCharacterString, peekedTagHeader *BACnetTagHeader, objectTypeArgument BACnetObjectType) *BACnetTimerStateChangeValueCharacterString {
	_result := &BACnetTimerStateChangeValueCharacterString{
		CharacterStringValue:        characterStringValue,
		BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetTimerStateChangeValueCharacterString(structType interface{}) *BACnetTimerStateChangeValueCharacterString {
	if casted, ok := structType.(BACnetTimerStateChangeValueCharacterString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetTimerStateChangeValue); ok {
		return CastBACnetTimerStateChangeValueCharacterString(casted.Child)
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValue); ok {
		return CastBACnetTimerStateChangeValueCharacterString(casted.Child)
	}
	return nil
}

func (m *BACnetTimerStateChangeValueCharacterString) GetTypeName() string {
	return "BACnetTimerStateChangeValueCharacterString"
}

func (m *BACnetTimerStateChangeValueCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimerStateChangeValueCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimerStateChangeValueCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueCharacterStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (*BACnetTimerStateChangeValueCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueCharacterString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterStringValue)
	if pullErr := readBuffer.PullContext("characterStringValue"); pullErr != nil {
		return nil, pullErr
	}
	_characterStringValue, _characterStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _characterStringValueErr != nil {
		return nil, errors.Wrap(_characterStringValueErr, "Error parsing 'characterStringValue' field")
	}
	characterStringValue := CastBACnetApplicationTagCharacterString(_characterStringValue)
	if closeErr := readBuffer.CloseContext("characterStringValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTimerStateChangeValueCharacterString{
		CharacterStringValue:        CastBACnetApplicationTagCharacterString(characterStringValue),
		BACnetTimerStateChangeValue: &BACnetTimerStateChangeValue{},
	}
	_child.BACnetTimerStateChangeValue.Child = _child
	return _child, nil
}

func (m *BACnetTimerStateChangeValueCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueCharacterString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (characterStringValue)
		if pushErr := writeBuffer.PushContext("characterStringValue"); pushErr != nil {
			return pushErr
		}
		_characterStringValueErr := m.CharacterStringValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("characterStringValue"); popErr != nil {
			return popErr
		}
		if _characterStringValueErr != nil {
			return errors.Wrap(_characterStringValueErr, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTimerStateChangeValueCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
