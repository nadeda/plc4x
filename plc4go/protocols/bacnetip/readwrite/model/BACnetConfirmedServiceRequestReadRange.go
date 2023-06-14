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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReadRange is the corresponding interface of BACnetConfirmedServiceRequestReadRange
type BACnetConfirmedServiceRequestReadRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetReadRange returns ReadRange (property field)
	GetReadRange() BACnetConfirmedServiceRequestReadRangeRange
}

// BACnetConfirmedServiceRequestReadRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadRange.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadRangeExactly interface {
	BACnetConfirmedServiceRequestReadRange
	isBACnetConfirmedServiceRequestReadRange() bool
}

// _BACnetConfirmedServiceRequestReadRange is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRange struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	PropertyArrayIndex BACnetContextTagUnsignedInteger
	ReadRange          BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetReadRange() BACnetConfirmedServiceRequestReadRangeRange {
	return m.ReadRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRange factory function for _BACnetConfirmedServiceRequestReadRange
func NewBACnetConfirmedServiceRequestReadRange(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, readRange BACnetConfirmedServiceRequestReadRangeRange, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadRange {
	_result := &_BACnetConfirmedServiceRequestReadRange{
		ObjectIdentifier:               objectIdentifier,
		PropertyIdentifier:             propertyIdentifier,
		PropertyArrayIndex:             propertyArrayIndex,
		ReadRange:                      readRange,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRange(structType any) BACnetConfirmedServiceRequestReadRange {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRange"
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (readRange)
	if m.ReadRange != nil {
		lengthInBits += m.ReadRange.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestReadRangeParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReadRange, error) {
	return BACnetConfirmedServiceRequestReadRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestReadRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReadRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetConfirmedServiceRequestReadRange")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetConfirmedServiceRequestReadRange")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if a given expression evaluates to false)
	var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyArrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyArrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field of BACnetConfirmedServiceRequestReadRange")
		default:
			propertyArrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyArrayIndex")
			}
		}
	}

	// Optional Field (readRange) (Can be skipped, if a given expression evaluates to false)
	var readRange BACnetConfirmedServiceRequestReadRangeRange = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("readRange"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for readRange")
		}
		_val, _err := BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'readRange' field of BACnetConfirmedServiceRequestReadRange")
		default:
			readRange = _val.(BACnetConfirmedServiceRequestReadRangeRange)
			if closeErr := readBuffer.CloseContext("readRange"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for readRange")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadRange{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		PropertyArrayIndex: propertyArrayIndex,
		ReadRange:          readRange,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRange")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
		}
		_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyIdentifier")
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
		var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
		if m.GetPropertyArrayIndex() != nil {
			if pushErr := writeBuffer.PushContext("propertyArrayIndex"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for propertyArrayIndex")
			}
			propertyArrayIndex = m.GetPropertyArrayIndex()
			_propertyArrayIndexErr := writeBuffer.WriteSerializable(ctx, propertyArrayIndex)
			if popErr := writeBuffer.PopContext("propertyArrayIndex"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for propertyArrayIndex")
			}
			if _propertyArrayIndexErr != nil {
				return errors.Wrap(_propertyArrayIndexErr, "Error serializing 'propertyArrayIndex' field")
			}
		}

		// Optional Field (readRange) (Can be skipped, if the value is null)
		var readRange BACnetConfirmedServiceRequestReadRangeRange = nil
		if m.GetReadRange() != nil {
			if pushErr := writeBuffer.PushContext("readRange"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for readRange")
			}
			readRange = m.GetReadRange()
			_readRangeErr := writeBuffer.WriteSerializable(ctx, readRange)
			if popErr := writeBuffer.PopContext("readRange"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for readRange")
			}
			if _readRangeErr != nil {
				return errors.Wrap(_readRangeErr, "Error serializing 'readRange' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRange) isBACnetConfirmedServiceRequestReadRange() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
