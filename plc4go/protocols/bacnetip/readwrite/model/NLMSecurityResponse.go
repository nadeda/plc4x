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

// NLMSecurityResponse is the corresponding interface of NLMSecurityResponse
type NLMSecurityResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetResponseCode returns ResponseCode (property field)
	GetResponseCode() SecurityResponseCode
	// GetOriginalMessageId returns OriginalMessageId (property field)
	GetOriginalMessageId() uint32
	// GetOriginalTimestamp returns OriginalTimestamp (property field)
	GetOriginalTimestamp() uint32
	// GetVariableParameters returns VariableParameters (property field)
	GetVariableParameters() []byte
}

// NLMSecurityResponseExactly can be used when we want exactly this type and not a type which fulfills NLMSecurityResponse.
// This is useful for switch cases.
type NLMSecurityResponseExactly interface {
	NLMSecurityResponse
	isNLMSecurityResponse() bool
}

// _NLMSecurityResponse is the data-structure of this message
type _NLMSecurityResponse struct {
	*_NLM
	ResponseCode       SecurityResponseCode
	OriginalMessageId  uint32
	OriginalTimestamp  uint32
	VariableParameters []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMSecurityResponse) GetMessageType() uint8 {
	return 0x0C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMSecurityResponse) InitializeParent(parent NLM) {}

func (m *_NLMSecurityResponse) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMSecurityResponse) GetResponseCode() SecurityResponseCode {
	return m.ResponseCode
}

func (m *_NLMSecurityResponse) GetOriginalMessageId() uint32 {
	return m.OriginalMessageId
}

func (m *_NLMSecurityResponse) GetOriginalTimestamp() uint32 {
	return m.OriginalTimestamp
}

func (m *_NLMSecurityResponse) GetVariableParameters() []byte {
	return m.VariableParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMSecurityResponse factory function for _NLMSecurityResponse
func NewNLMSecurityResponse(responseCode SecurityResponseCode, originalMessageId uint32, originalTimestamp uint32, variableParameters []byte, apduLength uint16) *_NLMSecurityResponse {
	_result := &_NLMSecurityResponse{
		ResponseCode:       responseCode,
		OriginalMessageId:  originalMessageId,
		OriginalTimestamp:  originalTimestamp,
		VariableParameters: variableParameters,
		_NLM:               NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMSecurityResponse(structType any) NLMSecurityResponse {
	if casted, ok := structType.(NLMSecurityResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NLMSecurityResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NLMSecurityResponse) GetTypeName() string {
	return "NLMSecurityResponse"
}

func (m *_NLMSecurityResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseCode)
	lengthInBits += 8

	// Simple field (originalMessageId)
	lengthInBits += 32

	// Simple field (originalTimestamp)
	lengthInBits += 32

	// Array field
	if len(m.VariableParameters) > 0 {
		lengthInBits += 8 * uint16(len(m.VariableParameters))
	}

	return lengthInBits
}

func (m *_NLMSecurityResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMSecurityResponseParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMSecurityResponse, error) {
	return NLMSecurityResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMSecurityResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMSecurityResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMSecurityResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMSecurityResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseCode)
	if pullErr := readBuffer.PullContext("responseCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseCode")
	}
	_responseCode, _responseCodeErr := SecurityResponseCodeParseWithBuffer(ctx, readBuffer)
	if _responseCodeErr != nil {
		return nil, errors.Wrap(_responseCodeErr, "Error parsing 'responseCode' field of NLMSecurityResponse")
	}
	responseCode := _responseCode
	if closeErr := readBuffer.CloseContext("responseCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseCode")
	}

	// Simple Field (originalMessageId)
	_originalMessageId, _originalMessageIdErr := readBuffer.ReadUint32("originalMessageId", 32)
	if _originalMessageIdErr != nil {
		return nil, errors.Wrap(_originalMessageIdErr, "Error parsing 'originalMessageId' field of NLMSecurityResponse")
	}
	originalMessageId := _originalMessageId

	// Simple Field (originalTimestamp)
	_originalTimestamp, _originalTimestampErr := readBuffer.ReadUint32("originalTimestamp", 32)
	if _originalTimestampErr != nil {
		return nil, errors.Wrap(_originalTimestampErr, "Error parsing 'originalTimestamp' field of NLMSecurityResponse")
	}
	originalTimestamp := _originalTimestamp
	// Byte Array field (variableParameters)
	numberOfBytesvariableParameters := int(uint16(apduLength) - uint16((uint16(uint16(uint16(uint16(1))+uint16(uint16(1)))+uint16(uint16(4))) + uint16(uint16(4)))))
	variableParameters, _readArrayErr := readBuffer.ReadByteArray("variableParameters", numberOfBytesvariableParameters)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'variableParameters' field of NLMSecurityResponse")
	}

	if closeErr := readBuffer.CloseContext("NLMSecurityResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMSecurityResponse")
	}

	// Create a partially initialized instance
	_child := &_NLMSecurityResponse{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		ResponseCode:       responseCode,
		OriginalMessageId:  originalMessageId,
		OriginalTimestamp:  originalTimestamp,
		VariableParameters: variableParameters,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMSecurityResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMSecurityResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMSecurityResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMSecurityResponse")
		}

		// Simple Field (responseCode)
		if pushErr := writeBuffer.PushContext("responseCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseCode")
		}
		_responseCodeErr := writeBuffer.WriteSerializable(ctx, m.GetResponseCode())
		if popErr := writeBuffer.PopContext("responseCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseCode")
		}
		if _responseCodeErr != nil {
			return errors.Wrap(_responseCodeErr, "Error serializing 'responseCode' field")
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

		// Array Field (variableParameters)
		// Byte Array field (variableParameters)
		if err := writeBuffer.WriteByteArray("variableParameters", m.GetVariableParameters()); err != nil {
			return errors.Wrap(err, "Error serializing 'variableParameters' field")
		}

		if popErr := writeBuffer.PopContext("NLMSecurityResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMSecurityResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMSecurityResponse) isNLMSecurityResponse() bool {
	return true
}

func (m *_NLMSecurityResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
