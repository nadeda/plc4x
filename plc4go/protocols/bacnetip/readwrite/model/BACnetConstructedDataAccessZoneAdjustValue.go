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

// BACnetConstructedDataAccessZoneAdjustValue is the corresponding interface of BACnetConstructedDataAccessZoneAdjustValue
type BACnetConstructedDataAccessZoneAdjustValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataAccessZoneAdjustValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessZoneAdjustValue.
// This is useful for switch cases.
type BACnetConstructedDataAccessZoneAdjustValueExactly interface {
	BACnetConstructedDataAccessZoneAdjustValue
	isBACnetConstructedDataAccessZoneAdjustValue() bool
}

// _BACnetConstructedDataAccessZoneAdjustValue is the data-structure of this message
type _BACnetConstructedDataAccessZoneAdjustValue struct {
	*_BACnetConstructedData
	AdjustValue BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_ZONE
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetAdjustValue() BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessZoneAdjustValue factory function for _BACnetConstructedDataAccessZoneAdjustValue
func NewBACnetConstructedDataAccessZoneAdjustValue(adjustValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessZoneAdjustValue {
	_result := &_BACnetConstructedDataAccessZoneAdjustValue{
		AdjustValue:            adjustValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessZoneAdjustValue(structType any) BACnetConstructedDataAccessZoneAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAccessZoneAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessZoneAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAccessZoneAdjustValue"
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessZoneAdjustValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessZoneAdjustValue, error) {
	return BACnetConstructedDataAccessZoneAdjustValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessZoneAdjustValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessZoneAdjustValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessZoneAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessZoneAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (adjustValue)
	if pullErr := readBuffer.PullContext("adjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for adjustValue")
	}
	_adjustValue, _adjustValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _adjustValueErr != nil {
		return nil, errors.Wrap(_adjustValueErr, "Error parsing 'adjustValue' field of BACnetConstructedDataAccessZoneAdjustValue")
	}
	adjustValue := _adjustValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("adjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for adjustValue")
	}

	// Virtual field
	_actualValue := adjustValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessZoneAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessZoneAdjustValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessZoneAdjustValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AdjustValue: adjustValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessZoneAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessZoneAdjustValue")
		}

		// Simple Field (adjustValue)
		if pushErr := writeBuffer.PushContext("adjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for adjustValue")
		}
		_adjustValueErr := writeBuffer.WriteSerializable(ctx, m.GetAdjustValue())
		if popErr := writeBuffer.PopContext("adjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for adjustValue")
		}
		if _adjustValueErr != nil {
			return errors.Wrap(_adjustValueErr, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessZoneAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessZoneAdjustValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) isBACnetConstructedDataAccessZoneAdjustValue() bool {
	return true
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
