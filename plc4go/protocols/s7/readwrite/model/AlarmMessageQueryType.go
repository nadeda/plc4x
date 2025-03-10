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

// Constant values.
const AlarmMessageQueryType_DATALENGTH uint16 = 0xFFFF

// AlarmMessageQueryType is the corresponding interface of AlarmMessageQueryType
type AlarmMessageQueryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectQueryType
}

// AlarmMessageQueryTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageQueryType.
// This is useful for switch cases.
type AlarmMessageQueryTypeExactly interface {
	AlarmMessageQueryType
	isAlarmMessageQueryType() bool
}

// _AlarmMessageQueryType is the data-structure of this message
type _AlarmMessageQueryType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	ReturnCode      DataTransportErrorCode
	TransportSize   DataTransportSize
	MessageObjects  []AlarmMessageObjectQueryType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageQueryType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageQueryType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageQueryType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_AlarmMessageQueryType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_AlarmMessageQueryType) GetMessageObjects() []AlarmMessageObjectQueryType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageQueryType) GetDataLength() uint16 {
	return AlarmMessageQueryType_DATALENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageQueryType factory function for _AlarmMessageQueryType
func NewAlarmMessageQueryType(functionId uint8, numberOfObjects uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, messageObjects []AlarmMessageObjectQueryType) *_AlarmMessageQueryType {
	return &_AlarmMessageQueryType{FunctionId: functionId, NumberOfObjects: numberOfObjects, ReturnCode: returnCode, TransportSize: transportSize, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageQueryType(structType any) AlarmMessageQueryType {
	if casted, ok := structType.(AlarmMessageQueryType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageQueryType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageQueryType) GetTypeName() string {
	return "AlarmMessageQueryType"
}

func (m *_AlarmMessageQueryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Const Field (DataLength)
	lengthInBits += 16

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageQueryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageQueryTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageQueryType, error) {
	return AlarmMessageQueryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageQueryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageQueryType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AlarmMessageQueryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageQueryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of AlarmMessageQueryType")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of AlarmMessageQueryType")
	}
	numberOfObjects := _numberOfObjects

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParseWithBuffer(ctx, readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of AlarmMessageQueryType")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	_transportSize, _transportSizeErr := DataTransportSizeParseWithBuffer(ctx, readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of AlarmMessageQueryType")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Const Field (DataLength)
	DataLength, _DataLengthErr := readBuffer.ReadUint16("DataLength", 16)
	if _DataLengthErr != nil {
		return nil, errors.Wrap(_DataLengthErr, "Error parsing 'DataLength' field of AlarmMessageQueryType")
	}
	if DataLength != AlarmMessageQueryType_DATALENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageQueryType_DATALENGTH) + " but got " + fmt.Sprintf("%d", DataLength))
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]AlarmMessageObjectQueryType, utils.Max(numberOfObjects, 0))
	// This happens when the size is set conditional to 0
	if len(messageObjects) == 0 {
		messageObjects = nil
	}
	{
		_numItems := uint16(utils.Max(numberOfObjects, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := AlarmMessageObjectQueryTypeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field of AlarmMessageQueryType")
			}
			messageObjects[_curItem] = _item.(AlarmMessageObjectQueryType)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageQueryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageQueryType")
	}

	// Create the instance
	return &_AlarmMessageQueryType{
		FunctionId:      functionId,
		NumberOfObjects: numberOfObjects,
		ReturnCode:      returnCode,
		TransportSize:   transportSize,
		MessageObjects:  messageObjects,
	}, nil
}

func (m *_AlarmMessageQueryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageQueryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageQueryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageQueryType")
	}

	// Simple Field (functionId)
	functionId := uint8(m.GetFunctionId())
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, uint8((functionId)))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.GetNumberOfObjects())
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, uint8((numberOfObjects)))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Const Field (DataLength)
	_DataLengthErr := writeBuffer.WriteUint16("DataLength", 16, uint16(0xFFFF))
	if _DataLengthErr != nil {
		return errors.Wrap(_DataLengthErr, "Error serializing 'DataLength' field")
	}

	// Array Field (messageObjects)
	if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for messageObjects")
	}
	for _curItem, _element := range m.GetMessageObjects() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMessageObjects()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
		}
	}
	if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for messageObjects")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageQueryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageQueryType")
	}
	return nil
}

func (m *_AlarmMessageQueryType) isAlarmMessageQueryType() bool {
	return true
}

func (m *_AlarmMessageQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
