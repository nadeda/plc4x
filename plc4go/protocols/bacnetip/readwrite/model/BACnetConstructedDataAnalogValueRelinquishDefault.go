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

// BACnetConstructedDataAnalogValueRelinquishDefault is the corresponding interface of BACnetConstructedDataAnalogValueRelinquishDefault
type BACnetConstructedDataAnalogValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataAnalogValueRelinquishDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogValueRelinquishDefault.
// This is useful for switch cases.
type BACnetConstructedDataAnalogValueRelinquishDefaultExactly interface {
	BACnetConstructedDataAnalogValueRelinquishDefault
	isBACnetConstructedDataAnalogValueRelinquishDefault() bool
}

// _BACnetConstructedDataAnalogValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataAnalogValueRelinquishDefault struct {
	*_BACnetConstructedData
	RelinquishDefault BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagReal {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogValueRelinquishDefault factory function for _BACnetConstructedDataAnalogValueRelinquishDefault
func NewBACnetConstructedDataAnalogValueRelinquishDefault(relinquishDefault BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValueRelinquishDefault {
	_result := &_BACnetConstructedDataAnalogValueRelinquishDefault{
		RelinquishDefault:      relinquishDefault,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValueRelinquishDefault(structType any) BACnetConstructedDataAnalogValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueRelinquishDefault"
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAnalogValueRelinquishDefaultParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogValueRelinquishDefault, error) {
	return BACnetConstructedDataAnalogValueRelinquishDefaultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAnalogValueRelinquishDefaultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogValueRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
	_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field of BACnetConstructedDataAnalogValueRelinquishDefault")
	}
	relinquishDefault := _relinquishDefault.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	// Virtual field
	_actualValue := relinquishDefault
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAnalogValueRelinquishDefault{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RelinquishDefault: relinquishDefault,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueRelinquishDefault")
		}

		// Simple Field (relinquishDefault)
		if pushErr := writeBuffer.PushContext("relinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relinquishDefault")
		}
		_relinquishDefaultErr := writeBuffer.WriteSerializable(ctx, m.GetRelinquishDefault())
		if popErr := writeBuffer.PopContext("relinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relinquishDefault")
		}
		if _relinquishDefaultErr != nil {
			return errors.Wrap(_relinquishDefaultErr, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) isBACnetConstructedDataAnalogValueRelinquishDefault() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
