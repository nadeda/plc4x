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

// BACnetConstructedDataBelongsTo is the corresponding interface of BACnetConstructedDataBelongsTo
type BACnetConstructedDataBelongsTo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBelongsTo returns BelongsTo (property field)
	GetBelongsTo() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataBelongsToExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBelongsTo.
// This is useful for switch cases.
type BACnetConstructedDataBelongsToExactly interface {
	BACnetConstructedDataBelongsTo
	isBACnetConstructedDataBelongsTo() bool
}

// _BACnetConstructedDataBelongsTo is the data-structure of this message
type _BACnetConstructedDataBelongsTo struct {
	*_BACnetConstructedData
	BelongsTo BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBelongsTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BELONGS_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBelongsTo) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBelongsTo) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetBelongsTo() BACnetDeviceObjectReference {
	return m.BelongsTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetBelongsTo())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBelongsTo factory function for _BACnetConstructedDataBelongsTo
func NewBACnetConstructedDataBelongsTo(belongsTo BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBelongsTo {
	_result := &_BACnetConstructedDataBelongsTo{
		BelongsTo:              belongsTo,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBelongsTo(structType any) BACnetConstructedDataBelongsTo {
	if casted, ok := structType.(BACnetConstructedDataBelongsTo); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBelongsTo); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBelongsTo) GetTypeName() string {
	return "BACnetConstructedDataBelongsTo"
}

func (m *_BACnetConstructedDataBelongsTo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (belongsTo)
	lengthInBits += m.BelongsTo.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBelongsTo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBelongsToParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBelongsTo, error) {
	return BACnetConstructedDataBelongsToParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBelongsToParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBelongsTo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBelongsTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBelongsTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (belongsTo)
	if pullErr := readBuffer.PullContext("belongsTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for belongsTo")
	}
	_belongsTo, _belongsToErr := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
	if _belongsToErr != nil {
		return nil, errors.Wrap(_belongsToErr, "Error parsing 'belongsTo' field of BACnetConstructedDataBelongsTo")
	}
	belongsTo := _belongsTo.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("belongsTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for belongsTo")
	}

	// Virtual field
	_actualValue := belongsTo
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBelongsTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBelongsTo")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBelongsTo{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BelongsTo: belongsTo,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBelongsTo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBelongsTo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBelongsTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBelongsTo")
		}

		// Simple Field (belongsTo)
		if pushErr := writeBuffer.PushContext("belongsTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for belongsTo")
		}
		_belongsToErr := writeBuffer.WriteSerializable(ctx, m.GetBelongsTo())
		if popErr := writeBuffer.PopContext("belongsTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for belongsTo")
		}
		if _belongsToErr != nil {
			return errors.Wrap(_belongsToErr, "Error serializing 'belongsTo' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBelongsTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBelongsTo")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBelongsTo) isBACnetConstructedDataBelongsTo() bool {
	return true
}

func (m *_BACnetConstructedDataBelongsTo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
