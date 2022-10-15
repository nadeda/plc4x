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

// BACnetConstructedDataProportionalConstant is the corresponding interface of BACnetConstructedDataProportionalConstant
type BACnetConstructedDataProportionalConstant interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProportionalConstant returns ProportionalConstant (property field)
	GetProportionalConstant() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataProportionalConstantExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProportionalConstant.
// This is useful for switch cases.
type BACnetConstructedDataProportionalConstantExactly interface {
	BACnetConstructedDataProportionalConstant
	isBACnetConstructedDataProportionalConstant() bool
}

// _BACnetConstructedDataProportionalConstant is the data-structure of this message
type _BACnetConstructedDataProportionalConstant struct {
	*_BACnetConstructedData
	ProportionalConstant BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProportionalConstant) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROPORTIONAL_CONSTANT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProportionalConstant) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProportionalConstant) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetProportionalConstant() BACnetApplicationTagReal {
	return m.ProportionalConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetProportionalConstant())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProportionalConstant factory function for _BACnetConstructedDataProportionalConstant
func NewBACnetConstructedDataProportionalConstant(proportionalConstant BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProportionalConstant {
	_result := &_BACnetConstructedDataProportionalConstant{
		ProportionalConstant:   proportionalConstant,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProportionalConstant(structType interface{}) BACnetConstructedDataProportionalConstant {
	if casted, ok := structType.(BACnetConstructedDataProportionalConstant); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProportionalConstant); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProportionalConstant) GetTypeName() string {
	return "BACnetConstructedDataProportionalConstant"
}

func (m *_BACnetConstructedDataProportionalConstant) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProportionalConstant) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (proportionalConstant)
	lengthInBits += m.ProportionalConstant.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProportionalConstant) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProportionalConstantParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProportionalConstant, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProportionalConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProportionalConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (proportionalConstant)
	if pullErr := readBuffer.PullContext("proportionalConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for proportionalConstant")
	}
	_proportionalConstant, _proportionalConstantErr := BACnetApplicationTagParse(readBuffer)
	if _proportionalConstantErr != nil {
		return nil, errors.Wrap(_proportionalConstantErr, "Error parsing 'proportionalConstant' field of BACnetConstructedDataProportionalConstant")
	}
	proportionalConstant := _proportionalConstant.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("proportionalConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for proportionalConstant")
	}

	// Virtual field
	_actualValue := proportionalConstant
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProportionalConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProportionalConstant")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProportionalConstant{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProportionalConstant: proportionalConstant,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProportionalConstant) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProportionalConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProportionalConstant")
		}

		// Simple Field (proportionalConstant)
		if pushErr := writeBuffer.PushContext("proportionalConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for proportionalConstant")
		}
		_proportionalConstantErr := writeBuffer.WriteSerializable(m.GetProportionalConstant())
		if popErr := writeBuffer.PopContext("proportionalConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for proportionalConstant")
		}
		if _proportionalConstantErr != nil {
			return errors.Wrap(_proportionalConstantErr, "Error serializing 'proportionalConstant' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProportionalConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProportionalConstant")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProportionalConstant) isBACnetConstructedDataProportionalConstant() bool {
	return true
}

func (m *_BACnetConstructedDataProportionalConstant) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
