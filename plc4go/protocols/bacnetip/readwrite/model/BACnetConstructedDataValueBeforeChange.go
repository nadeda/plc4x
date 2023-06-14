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

// BACnetConstructedDataValueBeforeChange is the corresponding interface of BACnetConstructedDataValueBeforeChange
type BACnetConstructedDataValueBeforeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetValuesBeforeChange returns ValuesBeforeChange (property field)
	GetValuesBeforeChange() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataValueBeforeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataValueBeforeChange.
// This is useful for switch cases.
type BACnetConstructedDataValueBeforeChangeExactly interface {
	BACnetConstructedDataValueBeforeChange
	isBACnetConstructedDataValueBeforeChange() bool
}

// _BACnetConstructedDataValueBeforeChange is the data-structure of this message
type _BACnetConstructedDataValueBeforeChange struct {
	*_BACnetConstructedData
	ValuesBeforeChange BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValueBeforeChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_BEFORE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataValueBeforeChange) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetValuesBeforeChange() BACnetApplicationTagUnsignedInteger {
	return m.ValuesBeforeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetValuesBeforeChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValueBeforeChange factory function for _BACnetConstructedDataValueBeforeChange
func NewBACnetConstructedDataValueBeforeChange(valuesBeforeChange BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValueBeforeChange {
	_result := &_BACnetConstructedDataValueBeforeChange{
		ValuesBeforeChange:     valuesBeforeChange,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueBeforeChange(structType any) BACnetConstructedDataValueBeforeChange {
	if casted, ok := structType.(BACnetConstructedDataValueBeforeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueBeforeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueBeforeChange) GetTypeName() string {
	return "BACnetConstructedDataValueBeforeChange"
}

func (m *_BACnetConstructedDataValueBeforeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (valuesBeforeChange)
	lengthInBits += m.ValuesBeforeChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValueBeforeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataValueBeforeChangeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValueBeforeChange, error) {
	return BACnetConstructedDataValueBeforeChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataValueBeforeChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValueBeforeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueBeforeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (valuesBeforeChange)
	if pullErr := readBuffer.PullContext("valuesBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valuesBeforeChange")
	}
	_valuesBeforeChange, _valuesBeforeChangeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _valuesBeforeChangeErr != nil {
		return nil, errors.Wrap(_valuesBeforeChangeErr, "Error parsing 'valuesBeforeChange' field of BACnetConstructedDataValueBeforeChange")
	}
	valuesBeforeChange := _valuesBeforeChange.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("valuesBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valuesBeforeChange")
	}

	// Virtual field
	_actualValue := valuesBeforeChange
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueBeforeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataValueBeforeChange{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ValuesBeforeChange: valuesBeforeChange,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataValueBeforeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValueBeforeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueBeforeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueBeforeChange")
		}

		// Simple Field (valuesBeforeChange)
		if pushErr := writeBuffer.PushContext("valuesBeforeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for valuesBeforeChange")
		}
		_valuesBeforeChangeErr := writeBuffer.WriteSerializable(ctx, m.GetValuesBeforeChange())
		if popErr := writeBuffer.PopContext("valuesBeforeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for valuesBeforeChange")
		}
		if _valuesBeforeChangeErr != nil {
			return errors.Wrap(_valuesBeforeChangeErr, "Error serializing 'valuesBeforeChange' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueBeforeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueBeforeChange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValueBeforeChange) isBACnetConstructedDataValueBeforeChange() bool {
	return true
}

func (m *_BACnetConstructedDataValueBeforeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
