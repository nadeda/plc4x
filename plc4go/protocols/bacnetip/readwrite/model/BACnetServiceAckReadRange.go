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

// BACnetServiceAckReadRange is the corresponding interface of BACnetServiceAckReadRange
type BACnetServiceAckReadRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetResultFlags returns ResultFlags (property field)
	GetResultFlags() BACnetResultFlagsTagged
	// GetItemCount returns ItemCount (property field)
	GetItemCount() BACnetContextTagUnsignedInteger
	// GetItemData returns ItemData (property field)
	GetItemData() BACnetConstructedData
	// GetFirstSequenceNumber returns FirstSequenceNumber (property field)
	GetFirstSequenceNumber() BACnetContextTagUnsignedInteger
}

// BACnetServiceAckReadRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckReadRange.
// This is useful for switch cases.
type BACnetServiceAckReadRangeExactly interface {
	BACnetServiceAckReadRange
	isBACnetServiceAckReadRange() bool
}

// _BACnetServiceAckReadRange is the data-structure of this message
type _BACnetServiceAckReadRange struct {
	*_BACnetServiceAck
	ObjectIdentifier    BACnetContextTagObjectIdentifier
	PropertyIdentifier  BACnetPropertyIdentifierTagged
	PropertyArrayIndex  BACnetContextTagUnsignedInteger
	ResultFlags         BACnetResultFlagsTagged
	ItemCount           BACnetContextTagUnsignedInteger
	ItemData            BACnetConstructedData
	FirstSequenceNumber BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadRange) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckReadRange) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetServiceAckReadRange) GetResultFlags() BACnetResultFlagsTagged {
	return m.ResultFlags
}

func (m *_BACnetServiceAckReadRange) GetItemCount() BACnetContextTagUnsignedInteger {
	return m.ItemCount
}

func (m *_BACnetServiceAckReadRange) GetItemData() BACnetConstructedData {
	return m.ItemData
}

func (m *_BACnetServiceAckReadRange) GetFirstSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.FirstSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadRange factory function for _BACnetServiceAckReadRange
func NewBACnetServiceAckReadRange(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, resultFlags BACnetResultFlagsTagged, itemCount BACnetContextTagUnsignedInteger, itemData BACnetConstructedData, firstSequenceNumber BACnetContextTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckReadRange {
	_result := &_BACnetServiceAckReadRange{
		ObjectIdentifier:    objectIdentifier,
		PropertyIdentifier:  propertyIdentifier,
		PropertyArrayIndex:  propertyArrayIndex,
		ResultFlags:         resultFlags,
		ItemCount:           itemCount,
		ItemData:            itemData,
		FirstSequenceNumber: firstSequenceNumber,
		_BACnetServiceAck:   NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadRange(structType any) BACnetServiceAckReadRange {
	if casted, ok := structType.(BACnetServiceAckReadRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadRange) GetTypeName() string {
	return "BACnetServiceAckReadRange"
}

func (m *_BACnetServiceAckReadRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (resultFlags)
	lengthInBits += m.ResultFlags.GetLengthInBits(ctx)

	// Simple field (itemCount)
	lengthInBits += m.ItemCount.GetLengthInBits(ctx)

	// Optional Field (itemData)
	if m.ItemData != nil {
		lengthInBits += m.ItemData.GetLengthInBits(ctx)
	}

	// Optional Field (firstSequenceNumber)
	if m.FirstSequenceNumber != nil {
		lengthInBits += m.FirstSequenceNumber.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckReadRangeParse(ctx context.Context, theBytes []byte, serviceAckLength uint32) (BACnetServiceAckReadRange, error) {
	return BACnetServiceAckReadRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckReadRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckReadRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetServiceAckReadRange")
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
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetServiceAckReadRange")
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
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field of BACnetServiceAckReadRange")
		default:
			propertyArrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyArrayIndex")
			}
		}
	}

	// Simple Field (resultFlags)
	if pullErr := readBuffer.PullContext("resultFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for resultFlags")
	}
	_resultFlags, _resultFlagsErr := BACnetResultFlagsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _resultFlagsErr != nil {
		return nil, errors.Wrap(_resultFlagsErr, "Error parsing 'resultFlags' field of BACnetServiceAckReadRange")
	}
	resultFlags := _resultFlags.(BACnetResultFlagsTagged)
	if closeErr := readBuffer.CloseContext("resultFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for resultFlags")
	}

	// Simple Field (itemCount)
	if pullErr := readBuffer.PullContext("itemCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for itemCount")
	}
	_itemCount, _itemCountErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _itemCountErr != nil {
		return nil, errors.Wrap(_itemCountErr, "Error parsing 'itemCount' field of BACnetServiceAckReadRange")
	}
	itemCount := _itemCount.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("itemCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for itemCount")
	}

	// Optional Field (itemData) (Can be skipped, if a given expression evaluates to false)
	var itemData BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("itemData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for itemData")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(5), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'itemData' field of BACnetServiceAckReadRange")
		default:
			itemData = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("itemData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for itemData")
			}
		}
	}

	// Optional Field (firstSequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var firstSequenceNumber BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("firstSequenceNumber"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for firstSequenceNumber")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(6), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'firstSequenceNumber' field of BACnetServiceAckReadRange")
		default:
			firstSequenceNumber = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("firstSequenceNumber"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for firstSequenceNumber")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckReadRange{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier:    objectIdentifier,
		PropertyIdentifier:  propertyIdentifier,
		PropertyArrayIndex:  propertyArrayIndex,
		ResultFlags:         resultFlags,
		ItemCount:           itemCount,
		ItemData:            itemData,
		FirstSequenceNumber: firstSequenceNumber,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckReadRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckReadRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadRange")
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

		// Simple Field (resultFlags)
		if pushErr := writeBuffer.PushContext("resultFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for resultFlags")
		}
		_resultFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetResultFlags())
		if popErr := writeBuffer.PopContext("resultFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for resultFlags")
		}
		if _resultFlagsErr != nil {
			return errors.Wrap(_resultFlagsErr, "Error serializing 'resultFlags' field")
		}

		// Simple Field (itemCount)
		if pushErr := writeBuffer.PushContext("itemCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for itemCount")
		}
		_itemCountErr := writeBuffer.WriteSerializable(ctx, m.GetItemCount())
		if popErr := writeBuffer.PopContext("itemCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for itemCount")
		}
		if _itemCountErr != nil {
			return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
		}

		// Optional Field (itemData) (Can be skipped, if the value is null)
		var itemData BACnetConstructedData = nil
		if m.GetItemData() != nil {
			if pushErr := writeBuffer.PushContext("itemData"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for itemData")
			}
			itemData = m.GetItemData()
			_itemDataErr := writeBuffer.WriteSerializable(ctx, itemData)
			if popErr := writeBuffer.PopContext("itemData"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for itemData")
			}
			if _itemDataErr != nil {
				return errors.Wrap(_itemDataErr, "Error serializing 'itemData' field")
			}
		}

		// Optional Field (firstSequenceNumber) (Can be skipped, if the value is null)
		var firstSequenceNumber BACnetContextTagUnsignedInteger = nil
		if m.GetFirstSequenceNumber() != nil {
			if pushErr := writeBuffer.PushContext("firstSequenceNumber"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for firstSequenceNumber")
			}
			firstSequenceNumber = m.GetFirstSequenceNumber()
			_firstSequenceNumberErr := writeBuffer.WriteSerializable(ctx, firstSequenceNumber)
			if popErr := writeBuffer.PopContext("firstSequenceNumber"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for firstSequenceNumber")
			}
			if _firstSequenceNumberErr != nil {
				return errors.Wrap(_firstSequenceNumberErr, "Error serializing 'firstSequenceNumber' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckReadRange) isBACnetServiceAckReadRange() bool {
	return true
}

func (m *_BACnetServiceAckReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
