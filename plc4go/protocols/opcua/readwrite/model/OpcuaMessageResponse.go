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

// OpcuaMessageResponse is the corresponding interface of OpcuaMessageResponse
type OpcuaMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetChunk returns Chunk (property field)
	GetChunk() string
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetSecureTokenId returns SecureTokenId (property field)
	GetSecureTokenId() int32
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() int32
	// GetRequestId returns RequestId (property field)
	GetRequestId() int32
	// GetMessage returns Message (property field)
	GetMessage() []byte
}

// OpcuaMessageResponseExactly can be used when we want exactly this type and not a type which fulfills OpcuaMessageResponse.
// This is useful for switch cases.
type OpcuaMessageResponseExactly interface {
	OpcuaMessageResponse
	isOpcuaMessageResponse() bool
}

// _OpcuaMessageResponse is the data-structure of this message
type _OpcuaMessageResponse struct {
	*_MessagePDU
	Chunk           string
	SecureChannelId int32
	SecureTokenId   int32
	SequenceNumber  int32
	RequestId       int32
	Message         []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageResponse) GetMessageType() string {
	return "MSG"
}

func (m *_OpcuaMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageResponse) InitializeParent(parent MessagePDU) {}

func (m *_OpcuaMessageResponse) GetParent() MessagePDU {
	return m._MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageResponse) GetChunk() string {
	return m.Chunk
}

func (m *_OpcuaMessageResponse) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpcuaMessageResponse) GetSecureTokenId() int32 {
	return m.SecureTokenId
}

func (m *_OpcuaMessageResponse) GetSequenceNumber() int32 {
	return m.SequenceNumber
}

func (m *_OpcuaMessageResponse) GetRequestId() int32 {
	return m.RequestId
}

func (m *_OpcuaMessageResponse) GetMessage() []byte {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaMessageResponse factory function for _OpcuaMessageResponse
func NewOpcuaMessageResponse(chunk string, secureChannelId int32, secureTokenId int32, sequenceNumber int32, requestId int32, message []byte) *_OpcuaMessageResponse {
	_result := &_OpcuaMessageResponse{
		Chunk:           chunk,
		SecureChannelId: secureChannelId,
		SecureTokenId:   secureTokenId,
		SequenceNumber:  sequenceNumber,
		RequestId:       requestId,
		Message:         message,
		_MessagePDU:     NewMessagePDU(),
	}
	_result._MessagePDU._MessagePDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaMessageResponse(structType any) OpcuaMessageResponse {
	if casted, ok := structType.(OpcuaMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageResponse) GetTypeName() string {
	return "OpcuaMessageResponse"
}

func (m *_OpcuaMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (messageSize)
	lengthInBits += 32

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (secureTokenId)
	lengthInBits += 32

	// Simple field (sequenceNumber)
	lengthInBits += 32

	// Simple field (requestId)
	lengthInBits += 32

	// Array field
	if len(m.Message) > 0 {
		lengthInBits += 8 * uint16(len(m.Message))
	}

	return lengthInBits
}

func (m *_OpcuaMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaMessageResponseParse(ctx context.Context, theBytes []byte, response bool) (OpcuaMessageResponse, error) {
	return OpcuaMessageResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpcuaMessageResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpcuaMessageResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (chunk)
	_chunk, _chunkErr := readBuffer.ReadString("chunk", uint32(8), "UTF-8")
	if _chunkErr != nil {
		return nil, errors.Wrap(_chunkErr, "Error parsing 'chunk' field of OpcuaMessageResponse")
	}
	chunk := _chunk

	// Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadInt32("messageSize", 32)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field of OpcuaMessageResponse")
	}

	// Simple Field (secureChannelId)
	_secureChannelId, _secureChannelIdErr := readBuffer.ReadInt32("secureChannelId", 32)
	if _secureChannelIdErr != nil {
		return nil, errors.Wrap(_secureChannelIdErr, "Error parsing 'secureChannelId' field of OpcuaMessageResponse")
	}
	secureChannelId := _secureChannelId

	// Simple Field (secureTokenId)
	_secureTokenId, _secureTokenIdErr := readBuffer.ReadInt32("secureTokenId", 32)
	if _secureTokenIdErr != nil {
		return nil, errors.Wrap(_secureTokenIdErr, "Error parsing 'secureTokenId' field of OpcuaMessageResponse")
	}
	secureTokenId := _secureTokenId

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadInt32("sequenceNumber", 32)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of OpcuaMessageResponse")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (requestId)
	_requestId, _requestIdErr := readBuffer.ReadInt32("requestId", 32)
	if _requestIdErr != nil {
		return nil, errors.Wrap(_requestIdErr, "Error parsing 'requestId' field of OpcuaMessageResponse")
	}
	requestId := _requestId
	// Byte Array field (message)
	numberOfBytesmessage := int(uint16(messageSize) - uint16(uint16(24)))
	message, _readArrayErr := readBuffer.ReadByteArray("message", numberOfBytesmessage)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'message' field of OpcuaMessageResponse")
	}

	if closeErr := readBuffer.CloseContext("OpcuaMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageResponse")
	}

	// Create a partially initialized instance
	_child := &_OpcuaMessageResponse{
		_MessagePDU:     &_MessagePDU{},
		Chunk:           chunk,
		SecureChannelId: secureChannelId,
		SecureTokenId:   secureTokenId,
		SequenceNumber:  sequenceNumber,
		RequestId:       requestId,
		Message:         message,
	}
	_child._MessagePDU._MessagePDUChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageResponse")
		}

		// Simple Field (chunk)
		chunk := string(m.GetChunk())
		_chunkErr := writeBuffer.WriteString("chunk", uint32(8), "UTF-8", (chunk))
		if _chunkErr != nil {
			return errors.Wrap(_chunkErr, "Error serializing 'chunk' field")
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := int32(int32(m.GetLengthInBytes(ctx)))
		_messageSizeErr := writeBuffer.WriteInt32("messageSize", 32, int32((messageSize)))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (secureChannelId)
		secureChannelId := int32(m.GetSecureChannelId())
		_secureChannelIdErr := writeBuffer.WriteInt32("secureChannelId", 32, int32((secureChannelId)))
		if _secureChannelIdErr != nil {
			return errors.Wrap(_secureChannelIdErr, "Error serializing 'secureChannelId' field")
		}

		// Simple Field (secureTokenId)
		secureTokenId := int32(m.GetSecureTokenId())
		_secureTokenIdErr := writeBuffer.WriteInt32("secureTokenId", 32, int32((secureTokenId)))
		if _secureTokenIdErr != nil {
			return errors.Wrap(_secureTokenIdErr, "Error serializing 'secureTokenId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := int32(m.GetSequenceNumber())
		_sequenceNumberErr := writeBuffer.WriteInt32("sequenceNumber", 32, int32((sequenceNumber)))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (requestId)
		requestId := int32(m.GetRequestId())
		_requestIdErr := writeBuffer.WriteInt32("requestId", 32, int32((requestId)))
		if _requestIdErr != nil {
			return errors.Wrap(_requestIdErr, "Error serializing 'requestId' field")
		}

		// Array Field (message)
		// Byte Array field (message)
		if err := writeBuffer.WriteByteArray("message", m.GetMessage()); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaMessageResponse) isOpcuaMessageResponse() bool {
	return true
}

func (m *_OpcuaMessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
