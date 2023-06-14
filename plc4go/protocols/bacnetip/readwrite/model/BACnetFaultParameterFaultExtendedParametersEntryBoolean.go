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

// BACnetFaultParameterFaultExtendedParametersEntryBoolean is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryBoolean
type BACnetFaultParameterFaultExtendedParametersEntryBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
}

// BACnetFaultParameterFaultExtendedParametersEntryBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryBoolean.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryBooleanExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryBoolean
	isBACnetFaultParameterFaultExtendedParametersEntryBoolean() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryBoolean is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryBoolean struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	BooleanValue BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryBoolean factory function for _BACnetFaultParameterFaultExtendedParametersEntryBoolean
func NewBACnetFaultParameterFaultExtendedParametersEntryBoolean(booleanValue BACnetApplicationTagBoolean, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryBoolean {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryBoolean{
		BooleanValue: booleanValue,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryBoolean(structType any) BACnetFaultParameterFaultExtendedParametersEntryBoolean {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryBoolean"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryBooleanParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultExtendedParametersEntryBoolean, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field of BACnetFaultParameterFaultExtendedParametersEntryBoolean")
	}
	booleanValue := _booleanValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryBoolean{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		BooleanValue: booleanValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(ctx, m.GetBooleanValue())
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) isBACnetFaultParameterFaultExtendedParametersEntryBoolean() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
