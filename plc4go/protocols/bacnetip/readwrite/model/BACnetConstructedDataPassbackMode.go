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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataPassbackMode is the corresponding interface of BACnetConstructedDataPassbackMode
type BACnetConstructedDataPassbackMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPassbackMode returns PassbackMode (property field)
	GetPassbackMode() BACnetAccessPassbackModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessPassbackModeTagged
}

// BACnetConstructedDataPassbackModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPassbackMode.
// This is useful for switch cases.
type BACnetConstructedDataPassbackModeExactly interface {
	BACnetConstructedDataPassbackMode
	isBACnetConstructedDataPassbackMode() bool
}

// _BACnetConstructedDataPassbackMode is the data-structure of this message
type _BACnetConstructedDataPassbackMode struct {
	*_BACnetConstructedData
	PassbackMode BACnetAccessPassbackModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassbackMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassbackMode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPassbackMode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetPassbackMode() BACnetAccessPassbackModeTagged {
	return m.PassbackMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetActualValue() BACnetAccessPassbackModeTagged {
	return CastBACnetAccessPassbackModeTagged(m.GetPassbackMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassbackMode factory function for _BACnetConstructedDataPassbackMode
func NewBACnetConstructedDataPassbackMode(passbackMode BACnetAccessPassbackModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassbackMode {
	_result := &_BACnetConstructedDataPassbackMode{
		PassbackMode:           passbackMode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassbackMode(structType interface{}) BACnetConstructedDataPassbackMode {
	if casted, ok := structType.(BACnetConstructedDataPassbackMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassbackMode) GetTypeName() string {
	return "BACnetConstructedDataPassbackMode"
}

func (m *_BACnetConstructedDataPassbackMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataPassbackMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passbackMode)
	lengthInBits += m.PassbackMode.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassbackMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassbackModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPassbackMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passbackMode)
	if pullErr := readBuffer.PullContext("passbackMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passbackMode")
	}
	_passbackMode, _passbackModeErr := BACnetAccessPassbackModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _passbackModeErr != nil {
		return nil, errors.Wrap(_passbackModeErr, "Error parsing 'passbackMode' field of BACnetConstructedDataPassbackMode")
	}
	passbackMode := _passbackMode.(BACnetAccessPassbackModeTagged)
	if closeErr := readBuffer.CloseContext("passbackMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passbackMode")
	}

	// Virtual field
	_actualValue := passbackMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPassbackMode{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PassbackMode: passbackMode,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPassbackMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackMode")
		}

		// Simple Field (passbackMode)
		if pushErr := writeBuffer.PushContext("passbackMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passbackMode")
		}
		_passbackModeErr := writeBuffer.WriteSerializable(m.GetPassbackMode())
		if popErr := writeBuffer.PopContext("passbackMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passbackMode")
		}
		if _passbackModeErr != nil {
			return errors.Wrap(_passbackModeErr, "Error serializing 'passbackMode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassbackMode) isBACnetConstructedDataPassbackMode() bool {
	return true
}

func (m *_BACnetConstructedDataPassbackMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
