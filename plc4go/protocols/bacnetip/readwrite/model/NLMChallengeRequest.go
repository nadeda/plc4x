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

// NLMChallengeRequest is the corresponding interface of NLMChallengeRequest
type NLMChallengeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetMessageChallenge returns MessageChallenge (property field)
	GetMessageChallenge() byte
	// GetOriginalMessageId returns OriginalMessageId (property field)
	GetOriginalMessageId() uint32
	// GetOriginalTimestamp returns OriginalTimestamp (property field)
	GetOriginalTimestamp() uint32
}

// NLMChallengeRequestExactly can be used when we want exactly this type and not a type which fulfills NLMChallengeRequest.
// This is useful for switch cases.
type NLMChallengeRequestExactly interface {
	NLMChallengeRequest
	isNLMChallengeRequest() bool
}

// _NLMChallengeRequest is the data-structure of this message
type _NLMChallengeRequest struct {
	*_NLM
	MessageChallenge  byte
	OriginalMessageId uint32
	OriginalTimestamp uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMChallengeRequest) GetMessageType() uint8 {
	return 0x0A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMChallengeRequest) InitializeParent(parent NLM) {}

func (m *_NLMChallengeRequest) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMChallengeRequest) GetMessageChallenge() byte {
	return m.MessageChallenge
}

func (m *_NLMChallengeRequest) GetOriginalMessageId() uint32 {
	return m.OriginalMessageId
}

func (m *_NLMChallengeRequest) GetOriginalTimestamp() uint32 {
	return m.OriginalTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMChallengeRequest factory function for _NLMChallengeRequest
func NewNLMChallengeRequest(messageChallenge byte, originalMessageId uint32, originalTimestamp uint32, apduLength uint16) *_NLMChallengeRequest {
	_result := &_NLMChallengeRequest{
		MessageChallenge:  messageChallenge,
		OriginalMessageId: originalMessageId,
		OriginalTimestamp: originalTimestamp,
		_NLM:              NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMChallengeRequest(structType any) NLMChallengeRequest {
	if casted, ok := structType.(NLMChallengeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*NLMChallengeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_NLMChallengeRequest) GetTypeName() string {
	return "NLMChallengeRequest"
}

func (m *_NLMChallengeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (messageChallenge)
	lengthInBits += 8

	// Simple field (originalMessageId)
	lengthInBits += 32

	// Simple field (originalTimestamp)
	lengthInBits += 32

	return lengthInBits
}

func (m *_NLMChallengeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMChallengeRequestParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMChallengeRequest, error) {
	return NLMChallengeRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMChallengeRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMChallengeRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMChallengeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMChallengeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (messageChallenge)
	_messageChallenge, _messageChallengeErr := readBuffer.ReadByte("messageChallenge")
	if _messageChallengeErr != nil {
		return nil, errors.Wrap(_messageChallengeErr, "Error parsing 'messageChallenge' field of NLMChallengeRequest")
	}
	messageChallenge := _messageChallenge

	// Simple Field (originalMessageId)
	_originalMessageId, _originalMessageIdErr := readBuffer.ReadUint32("originalMessageId", 32)
	if _originalMessageIdErr != nil {
		return nil, errors.Wrap(_originalMessageIdErr, "Error parsing 'originalMessageId' field of NLMChallengeRequest")
	}
	originalMessageId := _originalMessageId

	// Simple Field (originalTimestamp)
	_originalTimestamp, _originalTimestampErr := readBuffer.ReadUint32("originalTimestamp", 32)
	if _originalTimestampErr != nil {
		return nil, errors.Wrap(_originalTimestampErr, "Error parsing 'originalTimestamp' field of NLMChallengeRequest")
	}
	originalTimestamp := _originalTimestamp

	if closeErr := readBuffer.CloseContext("NLMChallengeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMChallengeRequest")
	}

	// Create a partially initialized instance
	_child := &_NLMChallengeRequest{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		MessageChallenge:  messageChallenge,
		OriginalMessageId: originalMessageId,
		OriginalTimestamp: originalTimestamp,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMChallengeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMChallengeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMChallengeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMChallengeRequest")
		}

		// Simple Field (messageChallenge)
		messageChallenge := byte(m.GetMessageChallenge())
		_messageChallengeErr := writeBuffer.WriteByte("messageChallenge", (messageChallenge))
		if _messageChallengeErr != nil {
			return errors.Wrap(_messageChallengeErr, "Error serializing 'messageChallenge' field")
		}

		// Simple Field (originalMessageId)
		originalMessageId := uint32(m.GetOriginalMessageId())
		_originalMessageIdErr := writeBuffer.WriteUint32("originalMessageId", 32, (originalMessageId))
		if _originalMessageIdErr != nil {
			return errors.Wrap(_originalMessageIdErr, "Error serializing 'originalMessageId' field")
		}

		// Simple Field (originalTimestamp)
		originalTimestamp := uint32(m.GetOriginalTimestamp())
		_originalTimestampErr := writeBuffer.WriteUint32("originalTimestamp", 32, (originalTimestamp))
		if _originalTimestampErr != nil {
			return errors.Wrap(_originalTimestampErr, "Error serializing 'originalTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("NLMChallengeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMChallengeRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMChallengeRequest) isNLMChallengeRequest() bool {
	return true
}

func (m *_NLMChallengeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
