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

// BACnetConstructedDataIntegerValueFaultLowLimit is the corresponding interface of BACnetConstructedDataIntegerValueFaultLowLimit
type BACnetConstructedDataIntegerValueFaultLowLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataIntegerValueFaultLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegerValueFaultLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataIntegerValueFaultLowLimitExactly interface {
	BACnetConstructedDataIntegerValueFaultLowLimit
	isBACnetConstructedDataIntegerValueFaultLowLimit() bool
}

// _BACnetConstructedDataIntegerValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueFaultLowLimit struct {
	*_BACnetConstructedData
	FaultLowLimit BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagSignedInteger {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueFaultLowLimit factory function for _BACnetConstructedDataIntegerValueFaultLowLimit
func NewBACnetConstructedDataIntegerValueFaultLowLimit(faultLowLimit BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueFaultLowLimit {
	_result := &_BACnetConstructedDataIntegerValueFaultLowLimit{
		FaultLowLimit:          faultLowLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueFaultLowLimit(structType interface{}) BACnetConstructedDataIntegerValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueFaultLowLimit"
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegerValueFaultLowLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueFaultLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultLowLimit)
	if pullErr := readBuffer.PullContext("faultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultLowLimit")
	}
	_faultLowLimit, _faultLowLimitErr := BACnetApplicationTagParse(readBuffer)
	if _faultLowLimitErr != nil {
		return nil, errors.Wrap(_faultLowLimitErr, "Error parsing 'faultLowLimit' field of BACnetConstructedDataIntegerValueFaultLowLimit")
	}
	faultLowLimit := _faultLowLimit.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("faultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultLowLimit")
	}

	// Virtual field
	_actualValue := faultLowLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueFaultLowLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueFaultLowLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultLowLimit: faultLowLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueFaultLowLimit")
		}

		// Simple Field (faultLowLimit)
		if pushErr := writeBuffer.PushContext("faultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultLowLimit")
		}
		_faultLowLimitErr := writeBuffer.WriteSerializable(m.GetFaultLowLimit())
		if popErr := writeBuffer.PopContext("faultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultLowLimit")
		}
		if _faultLowLimitErr != nil {
			return errors.Wrap(_faultLowLimitErr, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueFaultLowLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) isBACnetConstructedDataIntegerValueFaultLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
