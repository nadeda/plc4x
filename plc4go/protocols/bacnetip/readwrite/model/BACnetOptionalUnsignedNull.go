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

// BACnetOptionalUnsignedNull is the corresponding interface of BACnetOptionalUnsignedNull
type BACnetOptionalUnsignedNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalUnsigned
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetOptionalUnsignedNullExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalUnsignedNull.
// This is useful for switch cases.
type BACnetOptionalUnsignedNullExactly interface {
	BACnetOptionalUnsignedNull
	isBACnetOptionalUnsignedNull() bool
}

// _BACnetOptionalUnsignedNull is the data-structure of this message
type _BACnetOptionalUnsignedNull struct {
	*_BACnetOptionalUnsigned
	NullValue BACnetApplicationTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalUnsignedNull) InitializeParent(parent BACnetOptionalUnsigned, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalUnsignedNull) GetParent() BACnetOptionalUnsigned {
	return m._BACnetOptionalUnsigned
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsignedNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalUnsignedNull factory function for _BACnetOptionalUnsignedNull
func NewBACnetOptionalUnsignedNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetOptionalUnsignedNull {
	_result := &_BACnetOptionalUnsignedNull{
		NullValue:               nullValue,
		_BACnetOptionalUnsigned: NewBACnetOptionalUnsigned(peekedTagHeader),
	}
	_result._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsignedNull(structType any) BACnetOptionalUnsignedNull {
	if casted, ok := structType.(BACnetOptionalUnsignedNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsignedNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsignedNull) GetTypeName() string {
	return "BACnetOptionalUnsignedNull"
}

func (m *_BACnetOptionalUnsignedNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalUnsignedNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOptionalUnsignedNullParse(ctx context.Context, theBytes []byte) (BACnetOptionalUnsignedNull, error) {
	return BACnetOptionalUnsignedNullParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalUnsignedNullParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalUnsignedNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsignedNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsignedNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field of BACnetOptionalUnsignedNull")
	}
	nullValue := _nullValue.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsignedNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsignedNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalUnsignedNull{
		_BACnetOptionalUnsigned: &_BACnetOptionalUnsigned{},
		NullValue:               nullValue,
	}
	_child._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalUnsignedNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalUnsignedNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalUnsignedNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsignedNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := writeBuffer.WriteSerializable(ctx, m.GetNullValue())
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalUnsignedNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsignedNull")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalUnsignedNull) isBACnetOptionalUnsignedNull() bool {
	return true
}

func (m *_BACnetOptionalUnsignedNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
