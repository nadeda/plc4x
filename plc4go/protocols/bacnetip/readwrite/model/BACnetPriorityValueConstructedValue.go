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

// BACnetPriorityValueConstructedValue is the corresponding interface of BACnetPriorityValueConstructedValue
type BACnetPriorityValueConstructedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() BACnetConstructedData
}

// BACnetPriorityValueConstructedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueConstructedValue.
// This is useful for switch cases.
type BACnetPriorityValueConstructedValueExactly interface {
	BACnetPriorityValueConstructedValue
	isBACnetPriorityValueConstructedValue() bool
}

// _BACnetPriorityValueConstructedValue is the data-structure of this message
type _BACnetPriorityValueConstructedValue struct {
	*_BACnetPriorityValue
	ConstructedValue BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueConstructedValue) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueConstructedValue) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueConstructedValue) GetConstructedValue() BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueConstructedValue factory function for _BACnetPriorityValueConstructedValue
func NewBACnetPriorityValueConstructedValue(constructedValue BACnetConstructedData, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueConstructedValue {
	_result := &_BACnetPriorityValueConstructedValue{
		ConstructedValue:     constructedValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueConstructedValue(structType any) BACnetPriorityValueConstructedValue {
	if casted, ok := structType.(BACnetPriorityValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueConstructedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueConstructedValue) GetTypeName() string {
	return "BACnetPriorityValueConstructedValue"
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueConstructedValueParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValueConstructedValue, error) {
	return BACnetPriorityValueConstructedValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueConstructedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueConstructedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (constructedValue)
	if pullErr := readBuffer.PullContext("constructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for constructedValue")
	}
	_constructedValue, _constructedValueErr := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), nil)
	if _constructedValueErr != nil {
		return nil, errors.Wrap(_constructedValueErr, "Error parsing 'constructedValue' field of BACnetPriorityValueConstructedValue")
	}
	constructedValue := _constructedValue.(BACnetConstructedData)
	if closeErr := readBuffer.CloseContext("constructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for constructedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueConstructedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueConstructedValue{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		ConstructedValue: constructedValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueConstructedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueConstructedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueConstructedValue")
		}

		// Simple Field (constructedValue)
		if pushErr := writeBuffer.PushContext("constructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for constructedValue")
		}
		_constructedValueErr := writeBuffer.WriteSerializable(ctx, m.GetConstructedValue())
		if popErr := writeBuffer.PopContext("constructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for constructedValue")
		}
		if _constructedValueErr != nil {
			return errors.Wrap(_constructedValueErr, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueConstructedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueConstructedValue) isBACnetPriorityValueConstructedValue() bool {
	return true
}

func (m *_BACnetPriorityValueConstructedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
