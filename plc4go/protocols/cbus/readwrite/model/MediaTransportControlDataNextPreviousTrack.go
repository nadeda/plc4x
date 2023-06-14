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

// MediaTransportControlDataNextPreviousTrack is the corresponding interface of MediaTransportControlDataNextPreviousTrack
type MediaTransportControlDataNextPreviousTrack interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsSetThePreviousTrack returns IsSetThePreviousTrack (virtual field)
	GetIsSetThePreviousTrack() bool
	// GetIsSetTheNextTrack returns IsSetTheNextTrack (virtual field)
	GetIsSetTheNextTrack() bool
}

// MediaTransportControlDataNextPreviousTrackExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataNextPreviousTrack.
// This is useful for switch cases.
type MediaTransportControlDataNextPreviousTrackExactly interface {
	MediaTransportControlDataNextPreviousTrack
	isMediaTransportControlDataNextPreviousTrack() bool
}

// _MediaTransportControlDataNextPreviousTrack is the data-structure of this message
type _MediaTransportControlDataNextPreviousTrack struct {
	*_MediaTransportControlData
	Operation byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) GetIsSetThePreviousTrack() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetIsSetTheNextTrack() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataNextPreviousTrack factory function for _MediaTransportControlDataNextPreviousTrack
func NewMediaTransportControlDataNextPreviousTrack(operation byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataNextPreviousTrack {
	_result := &_MediaTransportControlDataNextPreviousTrack{
		Operation:                  operation,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataNextPreviousTrack(structType any) MediaTransportControlDataNextPreviousTrack {
	if casted, ok := structType.(MediaTransportControlDataNextPreviousTrack); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataNextPreviousTrack); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetTypeName() string {
	return "MediaTransportControlDataNextPreviousTrack"
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataNextPreviousTrackParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataNextPreviousTrack, error) {
	return MediaTransportControlDataNextPreviousTrackParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataNextPreviousTrackParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataNextPreviousTrack, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataNextPreviousTrack"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataNextPreviousTrack")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operation)
	_operation, _operationErr := readBuffer.ReadByte("operation")
	if _operationErr != nil {
		return nil, errors.Wrap(_operationErr, "Error parsing 'operation' field of MediaTransportControlDataNextPreviousTrack")
	}
	operation := _operation

	// Virtual field
	_isSetThePreviousTrack := bool((operation) == (0x00))
	isSetThePreviousTrack := bool(_isSetThePreviousTrack)
	_ = isSetThePreviousTrack

	// Virtual field
	_isSetTheNextTrack := bool((operation) != (0x00))
	isSetTheNextTrack := bool(_isSetTheNextTrack)
	_ = isSetTheNextTrack

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataNextPreviousTrack"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataNextPreviousTrack")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataNextPreviousTrack{
		_MediaTransportControlData: &_MediaTransportControlData{},
		Operation:                  operation,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataNextPreviousTrack"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataNextPreviousTrack")
		}

		// Simple Field (operation)
		operation := byte(m.GetOperation())
		_operationErr := writeBuffer.WriteByte("operation", (operation))
		if _operationErr != nil {
			return errors.Wrap(_operationErr, "Error serializing 'operation' field")
		}
		// Virtual field
		if _isSetThePreviousTrackErr := writeBuffer.WriteVirtual(ctx, "isSetThePreviousTrack", m.GetIsSetThePreviousTrack()); _isSetThePreviousTrackErr != nil {
			return errors.Wrap(_isSetThePreviousTrackErr, "Error serializing 'isSetThePreviousTrack' field")
		}
		// Virtual field
		if _isSetTheNextTrackErr := writeBuffer.WriteVirtual(ctx, "isSetTheNextTrack", m.GetIsSetTheNextTrack()); _isSetTheNextTrackErr != nil {
			return errors.Wrap(_isSetTheNextTrackErr, "Error serializing 'isSetTheNextTrack' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataNextPreviousTrack"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataNextPreviousTrack")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataNextPreviousTrack) isMediaTransportControlDataNextPreviousTrack() bool {
	return true
}

func (m *_MediaTransportControlDataNextPreviousTrack) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
