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

// BACnetLightingOperationTagged is the corresponding interface of BACnetLightingOperationTagged
type BACnetLightingOperationTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLightingOperation
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetLightingOperationTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLightingOperationTagged.
// This is useful for switch cases.
type BACnetLightingOperationTaggedExactly interface {
	BACnetLightingOperationTagged
	isBACnetLightingOperationTagged() bool
}

// _BACnetLightingOperationTagged is the data-structure of this message
type _BACnetLightingOperationTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLightingOperation
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingOperationTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLightingOperationTagged) GetValue() BACnetLightingOperation {
	return m.Value
}

func (m *_BACnetLightingOperationTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLightingOperationTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingOperationTagged factory function for _BACnetLightingOperationTagged
func NewBACnetLightingOperationTagged(header BACnetTagHeader, value BACnetLightingOperation, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLightingOperationTagged {
	return &_BACnetLightingOperationTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingOperationTagged(structType any) BACnetLightingOperationTagged {
	if casted, ok := structType.(BACnetLightingOperationTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingOperationTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingOperationTagged) GetTypeName() string {
	return "BACnetLightingOperationTagged"
}

func (m *_BACnetLightingOperationTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetLightingOperationTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingOperationTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLightingOperationTagged, error) {
	return BACnetLightingOperationTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLightingOperationTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLightingOperationTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingOperationTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingOperationTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLightingOperationTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, header.GetActualLength(), BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLightingOperationTagged")
	}
	var value BACnetLightingOperation
	if _value != nil {
		value = _value.(BACnetLightingOperation)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetLightingOperationTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
		proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingOperationTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingOperationTagged")
	}

	// Create the instance
	return &_BACnetLightingOperationTagged{
		TagNumber:        tagNumber,
		TagClass:         tagClass,
		Header:           header,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetLightingOperationTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingOperationTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLightingOperationTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingOperationTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingOperationTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingOperationTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingOperationTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLightingOperationTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLightingOperationTagged) isBACnetLightingOperationTagged() bool {
	return true
}

func (m *_BACnetLightingOperationTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
