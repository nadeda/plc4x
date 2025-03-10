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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEnergyMeterRef is the corresponding interface of BACnetConstructedDataEnergyMeterRef
type BACnetConstructedDataEnergyMeterRef interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEnergyMeterRef returns EnergyMeterRef (property field)
	GetEnergyMeterRef() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataEnergyMeterRefExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEnergyMeterRef.
// This is useful for switch cases.
type BACnetConstructedDataEnergyMeterRefExactly interface {
	BACnetConstructedDataEnergyMeterRef
	isBACnetConstructedDataEnergyMeterRef() bool
}

// _BACnetConstructedDataEnergyMeterRef is the data-structure of this message
type _BACnetConstructedDataEnergyMeterRef struct {
	*_BACnetConstructedData
	EnergyMeterRef BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENERGY_METER_REF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetEnergyMeterRef() BACnetDeviceObjectReference {
	return m.EnergyMeterRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetEnergyMeterRef())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEnergyMeterRef factory function for _BACnetConstructedDataEnergyMeterRef
func NewBACnetConstructedDataEnergyMeterRef(energyMeterRef BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnergyMeterRef {
	_result := &_BACnetConstructedDataEnergyMeterRef{
		EnergyMeterRef:         energyMeterRef,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnergyMeterRef(structType any) BACnetConstructedDataEnergyMeterRef {
	if casted, ok := structType.(BACnetConstructedDataEnergyMeterRef); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnergyMeterRef); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetTypeName() string {
	return "BACnetConstructedDataEnergyMeterRef"
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (energyMeterRef)
	lengthInBits += m.EnergyMeterRef.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataEnergyMeterRefParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEnergyMeterRef, error) {
	return BACnetConstructedDataEnergyMeterRefParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataEnergyMeterRefParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEnergyMeterRef, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnergyMeterRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnergyMeterRef")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (energyMeterRef)
	if pullErr := readBuffer.PullContext("energyMeterRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for energyMeterRef")
	}
	_energyMeterRef, _energyMeterRefErr := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
	if _energyMeterRefErr != nil {
		return nil, errors.Wrap(_energyMeterRefErr, "Error parsing 'energyMeterRef' field of BACnetConstructedDataEnergyMeterRef")
	}
	energyMeterRef := _energyMeterRef.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("energyMeterRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for energyMeterRef")
	}

	// Virtual field
	_actualValue := energyMeterRef
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnergyMeterRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnergyMeterRef")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEnergyMeterRef{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EnergyMeterRef: energyMeterRef,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnergyMeterRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnergyMeterRef")
		}

		// Simple Field (energyMeterRef)
		if pushErr := writeBuffer.PushContext("energyMeterRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for energyMeterRef")
		}
		_energyMeterRefErr := writeBuffer.WriteSerializable(ctx, m.GetEnergyMeterRef())
		if popErr := writeBuffer.PopContext("energyMeterRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for energyMeterRef")
		}
		if _energyMeterRefErr != nil {
			return errors.Wrap(_energyMeterRefErr, "Error serializing 'energyMeterRef' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnergyMeterRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnergyMeterRef")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnergyMeterRef) isBACnetConstructedDataEnergyMeterRef() bool {
	return true
}

func (m *_BACnetConstructedDataEnergyMeterRef) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
