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

// BACnetLiftCarDoorCommandTagged is the corresponding interface of BACnetLiftCarDoorCommandTagged
type BACnetLiftCarDoorCommandTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftCarDoorCommand
}

// BACnetLiftCarDoorCommandTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftCarDoorCommandTagged.
// This is useful for switch cases.
type BACnetLiftCarDoorCommandTaggedExactly interface {
	BACnetLiftCarDoorCommandTagged
	isBACnetLiftCarDoorCommandTagged() bool
}

// _BACnetLiftCarDoorCommandTagged is the data-structure of this message
type _BACnetLiftCarDoorCommandTagged struct {
	Header BACnetTagHeader
	Value  BACnetLiftCarDoorCommand

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarDoorCommandTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftCarDoorCommandTagged) GetValue() BACnetLiftCarDoorCommand {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftCarDoorCommandTagged factory function for _BACnetLiftCarDoorCommandTagged
func NewBACnetLiftCarDoorCommandTagged(header BACnetTagHeader, value BACnetLiftCarDoorCommand, tagNumber uint8, tagClass TagClass) *_BACnetLiftCarDoorCommandTagged {
	return &_BACnetLiftCarDoorCommandTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarDoorCommandTagged(structType any) BACnetLiftCarDoorCommandTagged {
	if casted, ok := structType.(BACnetLiftCarDoorCommandTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarDoorCommandTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarDoorCommandTagged) GetTypeName() string {
	return "BACnetLiftCarDoorCommandTagged"
}

func (m *_BACnetLiftCarDoorCommandTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetLiftCarDoorCommandTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarDoorCommandTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDoorCommandTagged, error) {
	return BACnetLiftCarDoorCommandTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLiftCarDoorCommandTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDoorCommandTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarDoorCommandTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarDoorCommandTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLiftCarDoorCommandTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetLiftCarDoorCommand_NONE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLiftCarDoorCommandTagged")
	}
	var value BACnetLiftCarDoorCommand
	if _value != nil {
		value = _value.(BACnetLiftCarDoorCommand)
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftCarDoorCommandTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarDoorCommandTagged")
	}

	// Create the instance
	return &_BACnetLiftCarDoorCommandTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetLiftCarDoorCommandTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarDoorCommandTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLiftCarDoorCommandTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarDoorCommandTagged")
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

	if popErr := writeBuffer.PopContext("BACnetLiftCarDoorCommandTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarDoorCommandTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftCarDoorCommandTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLiftCarDoorCommandTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLiftCarDoorCommandTagged) isBACnetLiftCarDoorCommandTagged() bool {
	return true
}

func (m *_BACnetLiftCarDoorCommandTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
