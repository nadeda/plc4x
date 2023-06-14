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

// AlarmMessageAckResponseType is the corresponding interface of AlarmMessageAckResponseType
type AlarmMessageAckResponseType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []uint8
}

// AlarmMessageAckResponseTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageAckResponseType.
// This is useful for switch cases.
type AlarmMessageAckResponseTypeExactly interface {
	AlarmMessageAckResponseType
	isAlarmMessageAckResponseType() bool
}

// _AlarmMessageAckResponseType is the data-structure of this message
type _AlarmMessageAckResponseType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckResponseType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageAckResponseType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageAckResponseType) GetMessageObjects() []uint8 {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageAckResponseType factory function for _AlarmMessageAckResponseType
func NewAlarmMessageAckResponseType(functionId uint8, numberOfObjects uint8, messageObjects []uint8) *_AlarmMessageAckResponseType {
	return &_AlarmMessageAckResponseType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckResponseType(structType any) AlarmMessageAckResponseType {
	if casted, ok := structType.(AlarmMessageAckResponseType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckResponseType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckResponseType) GetTypeName() string {
	return "AlarmMessageAckResponseType"
}

func (m *_AlarmMessageAckResponseType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		lengthInBits += 8 * uint16(len(m.MessageObjects))
	}

	return lengthInBits
}

func (m *_AlarmMessageAckResponseType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageAckResponseTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageAckResponseType, error) {
	return AlarmMessageAckResponseTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageAckResponseTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckResponseType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckResponseType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckResponseType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of AlarmMessageAckResponseType")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of AlarmMessageAckResponseType")
	}
	numberOfObjects := _numberOfObjects

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]uint8, numberOfObjects)
	// This happens when the size is set conditional to 0
	if len(messageObjects) == 0 {
		messageObjects = nil
	}
	{
		_numItems := uint16(numberOfObjects)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field of AlarmMessageAckResponseType")
			}
			messageObjects[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckResponseType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckResponseType")
	}

	// Create the instance
	return &_AlarmMessageAckResponseType{
		FunctionId:      functionId,
		NumberOfObjects: numberOfObjects,
		MessageObjects:  messageObjects,
	}, nil
}

func (m *_AlarmMessageAckResponseType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageAckResponseType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AlarmMessageAckResponseType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckResponseType")
	}

	// Simple Field (functionId)
	functionId := uint8(m.GetFunctionId())
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.GetNumberOfObjects())
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for messageObjects")
	}
	for _curItem, _element := range m.GetMessageObjects() {
		_ = _curItem
		_elementErr := writeBuffer.WriteUint8("", 8, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
		}
	}
	if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for messageObjects")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckResponseType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckResponseType")
	}
	return nil
}

func (m *_AlarmMessageAckResponseType) isAlarmMessageAckResponseType() bool {
	return true
}

func (m *_AlarmMessageAckResponseType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
