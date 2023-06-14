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

// BACnetPropertyStatesExtendedValue is the corresponding interface of BACnetPropertyStatesExtendedValue
type BACnetPropertyStatesExtendedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetExtendedValue returns ExtendedValue (property field)
	GetExtendedValue() BACnetContextTagUnsignedInteger
}

// BACnetPropertyStatesExtendedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesExtendedValue.
// This is useful for switch cases.
type BACnetPropertyStatesExtendedValueExactly interface {
	BACnetPropertyStatesExtendedValue
	isBACnetPropertyStatesExtendedValue() bool
}

// _BACnetPropertyStatesExtendedValue is the data-structure of this message
type _BACnetPropertyStatesExtendedValue struct {
	*_BACnetPropertyStates
	ExtendedValue BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesExtendedValue) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesExtendedValue) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesExtendedValue) GetExtendedValue() BACnetContextTagUnsignedInteger {
	return m.ExtendedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesExtendedValue factory function for _BACnetPropertyStatesExtendedValue
func NewBACnetPropertyStatesExtendedValue(extendedValue BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesExtendedValue {
	_result := &_BACnetPropertyStatesExtendedValue{
		ExtendedValue:         extendedValue,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesExtendedValue(structType any) BACnetPropertyStatesExtendedValue {
	if casted, ok := structType.(BACnetPropertyStatesExtendedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesExtendedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesExtendedValue) GetTypeName() string {
	return "BACnetPropertyStatesExtendedValue"
}

func (m *_BACnetPropertyStatesExtendedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (extendedValue)
	lengthInBits += m.ExtendedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesExtendedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesExtendedValueParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesExtendedValue, error) {
	return BACnetPropertyStatesExtendedValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesExtendedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesExtendedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesExtendedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesExtendedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (extendedValue)
	if pullErr := readBuffer.PullContext("extendedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedValue")
	}
	_extendedValue, _extendedValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _extendedValueErr != nil {
		return nil, errors.Wrap(_extendedValueErr, "Error parsing 'extendedValue' field of BACnetPropertyStatesExtendedValue")
	}
	extendedValue := _extendedValue.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("extendedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesExtendedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesExtendedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesExtendedValue{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		ExtendedValue:         extendedValue,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesExtendedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesExtendedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesExtendedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesExtendedValue")
		}

		// Simple Field (extendedValue)
		if pushErr := writeBuffer.PushContext("extendedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedValue")
		}
		_extendedValueErr := writeBuffer.WriteSerializable(ctx, m.GetExtendedValue())
		if popErr := writeBuffer.PopContext("extendedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedValue")
		}
		if _extendedValueErr != nil {
			return errors.Wrap(_extendedValueErr, "Error serializing 'extendedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesExtendedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesExtendedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesExtendedValue) isBACnetPropertyStatesExtendedValue() bool {
	return true
}

func (m *_BACnetPropertyStatesExtendedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
