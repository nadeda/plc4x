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

// OpcuaOpenRequest is the corresponding interface of OpcuaOpenRequest
type OpcuaOpenRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetChunk returns Chunk (property field)
	GetChunk() string
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetEndpoint returns Endpoint (property field)
	GetEndpoint() PascalString
	// GetSenderCertificate returns SenderCertificate (property field)
	GetSenderCertificate() PascalByteString
	// GetReceiverCertificateThumbprint returns ReceiverCertificateThumbprint (property field)
	GetReceiverCertificateThumbprint() PascalByteString
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() int32
	// GetRequestId returns RequestId (property field)
	GetRequestId() int32
	// GetMessage returns Message (property field)
	GetMessage() []byte
}

// OpcuaOpenRequestExactly can be used when we want exactly this type and not a type which fulfills OpcuaOpenRequest.
// This is useful for switch cases.
type OpcuaOpenRequestExactly interface {
	OpcuaOpenRequest
	isOpcuaOpenRequest() bool
}

// _OpcuaOpenRequest is the data-structure of this message
type _OpcuaOpenRequest struct {
	*_MessagePDU
	Chunk                         string
	SecureChannelId               int32
	Endpoint                      PascalString
	SenderCertificate             PascalByteString
	ReceiverCertificateThumbprint PascalByteString
	SequenceNumber                int32
	RequestId                     int32
	Message                       []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaOpenRequest) GetMessageType() string {
	return "OPN"
}

func (m *_OpcuaOpenRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaOpenRequest) InitializeParent(parent MessagePDU) {}

func (m *_OpcuaOpenRequest) GetParent() MessagePDU {
	return m._MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaOpenRequest) GetChunk() string {
	return m.Chunk
}

func (m *_OpcuaOpenRequest) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpcuaOpenRequest) GetEndpoint() PascalString {
	return m.Endpoint
}

func (m *_OpcuaOpenRequest) GetSenderCertificate() PascalByteString {
	return m.SenderCertificate
}

func (m *_OpcuaOpenRequest) GetReceiverCertificateThumbprint() PascalByteString {
	return m.ReceiverCertificateThumbprint
}

func (m *_OpcuaOpenRequest) GetSequenceNumber() int32 {
	return m.SequenceNumber
}

func (m *_OpcuaOpenRequest) GetRequestId() int32 {
	return m.RequestId
}

func (m *_OpcuaOpenRequest) GetMessage() []byte {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaOpenRequest factory function for _OpcuaOpenRequest
func NewOpcuaOpenRequest(chunk string, secureChannelId int32, endpoint PascalString, senderCertificate PascalByteString, receiverCertificateThumbprint PascalByteString, sequenceNumber int32, requestId int32, message []byte) *_OpcuaOpenRequest {
	_result := &_OpcuaOpenRequest{
		Chunk:                         chunk,
		SecureChannelId:               secureChannelId,
		Endpoint:                      endpoint,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
		SequenceNumber:                sequenceNumber,
		RequestId:                     requestId,
		Message:                       message,
		_MessagePDU:                   NewMessagePDU(),
	}
	_result._MessagePDU._MessagePDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaOpenRequest(structType any) OpcuaOpenRequest {
	if casted, ok := structType.(OpcuaOpenRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaOpenRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaOpenRequest) GetTypeName() string {
	return "OpcuaOpenRequest"
}

func (m *_OpcuaOpenRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (messageSize)
	lengthInBits += 32

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (endpoint)
	lengthInBits += m.Endpoint.GetLengthInBits(ctx)

	// Simple field (senderCertificate)
	lengthInBits += m.SenderCertificate.GetLengthInBits(ctx)

	// Simple field (receiverCertificateThumbprint)
	lengthInBits += m.ReceiverCertificateThumbprint.GetLengthInBits(ctx)

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

func (m *_OpcuaOpenRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaOpenRequestParse(ctx context.Context, theBytes []byte, response bool) (OpcuaOpenRequest, error) {
	return OpcuaOpenRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpcuaOpenRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpcuaOpenRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaOpenRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaOpenRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (chunk)
	_chunk, _chunkErr := readBuffer.ReadString("chunk", uint32(8), "UTF-8")
	if _chunkErr != nil {
		return nil, errors.Wrap(_chunkErr, "Error parsing 'chunk' field of OpcuaOpenRequest")
	}
	chunk := _chunk

	// Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadInt32("messageSize", 32)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field of OpcuaOpenRequest")
	}

	// Simple Field (secureChannelId)
	_secureChannelId, _secureChannelIdErr := readBuffer.ReadInt32("secureChannelId", 32)
	if _secureChannelIdErr != nil {
		return nil, errors.Wrap(_secureChannelIdErr, "Error parsing 'secureChannelId' field of OpcuaOpenRequest")
	}
	secureChannelId := _secureChannelId

	// Simple Field (endpoint)
	if pullErr := readBuffer.PullContext("endpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endpoint")
	}
	_endpoint, _endpointErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _endpointErr != nil {
		return nil, errors.Wrap(_endpointErr, "Error parsing 'endpoint' field of OpcuaOpenRequest")
	}
	endpoint := _endpoint.(PascalString)
	if closeErr := readBuffer.CloseContext("endpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endpoint")
	}

	// Simple Field (senderCertificate)
	if pullErr := readBuffer.PullContext("senderCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for senderCertificate")
	}
	_senderCertificate, _senderCertificateErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _senderCertificateErr != nil {
		return nil, errors.Wrap(_senderCertificateErr, "Error parsing 'senderCertificate' field of OpcuaOpenRequest")
	}
	senderCertificate := _senderCertificate.(PascalByteString)
	if closeErr := readBuffer.CloseContext("senderCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for senderCertificate")
	}

	// Simple Field (receiverCertificateThumbprint)
	if pullErr := readBuffer.PullContext("receiverCertificateThumbprint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for receiverCertificateThumbprint")
	}
	_receiverCertificateThumbprint, _receiverCertificateThumbprintErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _receiverCertificateThumbprintErr != nil {
		return nil, errors.Wrap(_receiverCertificateThumbprintErr, "Error parsing 'receiverCertificateThumbprint' field of OpcuaOpenRequest")
	}
	receiverCertificateThumbprint := _receiverCertificateThumbprint.(PascalByteString)
	if closeErr := readBuffer.CloseContext("receiverCertificateThumbprint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for receiverCertificateThumbprint")
	}

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadInt32("sequenceNumber", 32)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of OpcuaOpenRequest")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (requestId)
	_requestId, _requestIdErr := readBuffer.ReadInt32("requestId", 32)
	if _requestIdErr != nil {
		return nil, errors.Wrap(_requestIdErr, "Error parsing 'requestId' field of OpcuaOpenRequest")
	}
	requestId := _requestId
	// Byte Array field (message)
	numberOfBytesmessage := int(uint16(uint16(uint16(uint16(messageSize)-uint16((utils.InlineIf(bool((endpoint.GetStringLength()) == (-(1))), func() any { return uint16(uint16(0)) }, func() any { return uint16(endpoint.GetStringLength()) }).(uint16))))-uint16((utils.InlineIf(bool((senderCertificate.GetStringLength()) == (-(1))), func() any { return uint16(uint16(0)) }, func() any { return uint16(senderCertificate.GetStringLength()) }).(uint16))))-uint16((utils.InlineIf(bool((receiverCertificateThumbprint.GetStringLength()) == (-(1))), func() any { return uint16(uint16(0)) }, func() any { return uint16(receiverCertificateThumbprint.GetStringLength()) }).(uint16)))) - uint16(uint16(32)))
	message, _readArrayErr := readBuffer.ReadByteArray("message", numberOfBytesmessage)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'message' field of OpcuaOpenRequest")
	}

	if closeErr := readBuffer.CloseContext("OpcuaOpenRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaOpenRequest")
	}

	// Create a partially initialized instance
	_child := &_OpcuaOpenRequest{
		_MessagePDU:                   &_MessagePDU{},
		Chunk:                         chunk,
		SecureChannelId:               secureChannelId,
		Endpoint:                      endpoint,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
		SequenceNumber:                sequenceNumber,
		RequestId:                     requestId,
		Message:                       message,
	}
	_child._MessagePDU._MessagePDUChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaOpenRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaOpenRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaOpenRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaOpenRequest")
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

		// Simple Field (endpoint)
		if pushErr := writeBuffer.PushContext("endpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpoint")
		}
		_endpointErr := writeBuffer.WriteSerializable(ctx, m.GetEndpoint())
		if popErr := writeBuffer.PopContext("endpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpoint")
		}
		if _endpointErr != nil {
			return errors.Wrap(_endpointErr, "Error serializing 'endpoint' field")
		}

		// Simple Field (senderCertificate)
		if pushErr := writeBuffer.PushContext("senderCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for senderCertificate")
		}
		_senderCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetSenderCertificate())
		if popErr := writeBuffer.PopContext("senderCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for senderCertificate")
		}
		if _senderCertificateErr != nil {
			return errors.Wrap(_senderCertificateErr, "Error serializing 'senderCertificate' field")
		}

		// Simple Field (receiverCertificateThumbprint)
		if pushErr := writeBuffer.PushContext("receiverCertificateThumbprint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for receiverCertificateThumbprint")
		}
		_receiverCertificateThumbprintErr := writeBuffer.WriteSerializable(ctx, m.GetReceiverCertificateThumbprint())
		if popErr := writeBuffer.PopContext("receiverCertificateThumbprint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for receiverCertificateThumbprint")
		}
		if _receiverCertificateThumbprintErr != nil {
			return errors.Wrap(_receiverCertificateThumbprintErr, "Error serializing 'receiverCertificateThumbprint' field")
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

		if popErr := writeBuffer.PopContext("OpcuaOpenRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaOpenRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaOpenRequest) isOpcuaOpenRequest() bool {
	return true
}

func (m *_OpcuaOpenRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
