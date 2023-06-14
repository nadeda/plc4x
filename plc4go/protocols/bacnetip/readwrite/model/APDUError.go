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

// APDUError is the corresponding interface of APDUError
type APDUError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetErrorChoice returns ErrorChoice (property field)
	GetErrorChoice() BACnetConfirmedServiceChoice
	// GetError returns Error (property field)
	GetError() BACnetError
}

// APDUErrorExactly can be used when we want exactly this type and not a type which fulfills APDUError.
// This is useful for switch cases.
type APDUErrorExactly interface {
	APDUError
	isAPDUError() bool
}

// _APDUError is the data-structure of this message
type _APDUError struct {
	*_APDU
	OriginalInvokeId uint8
	ErrorChoice      BACnetConfirmedServiceChoice
	Error            BACnetError
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUError) GetApduType() ApduType {
	return ApduType_ERROR_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUError) InitializeParent(parent APDU) {}

func (m *_APDUError) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUError) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return m.ErrorChoice
}

func (m *_APDUError) GetError() BACnetError {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUError factory function for _APDUError
func NewAPDUError(originalInvokeId uint8, errorChoice BACnetConfirmedServiceChoice, error BACnetError, apduLength uint16) *_APDUError {
	_result := &_APDUError{
		OriginalInvokeId: originalInvokeId,
		ErrorChoice:      errorChoice,
		Error:            error,
		_APDU:            NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUError(structType any) APDUError {
	if casted, ok := structType.(APDUError); ok {
		return casted
	}
	if casted, ok := structType.(*APDUError); ok {
		return *casted
	}
	return nil
}

func (m *_APDUError) GetTypeName() string {
	return "APDUError"
}

func (m *_APDUError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (errorChoice)
	lengthInBits += 8

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUErrorParse(ctx context.Context, theBytes []byte, apduLength uint16) (APDUError, error) {
	return APDUErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of APDUError")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field of APDUError")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (errorChoice)
	if pullErr := readBuffer.PullContext("errorChoice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorChoice")
	}
	_errorChoice, _errorChoiceErr := BACnetConfirmedServiceChoiceParseWithBuffer(ctx, readBuffer)
	if _errorChoiceErr != nil {
		return nil, errors.Wrap(_errorChoiceErr, "Error parsing 'errorChoice' field of APDUError")
	}
	errorChoice := _errorChoice
	if closeErr := readBuffer.CloseContext("errorChoice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorChoice")
	}

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for error")
	}
	_error, _errorErr := BACnetErrorParseWithBuffer(ctx, readBuffer, BACnetConfirmedServiceChoice(errorChoice))
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of APDUError")
	}
	error := _error.(BACnetError)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for error")
	}

	if closeErr := readBuffer.CloseContext("APDUError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUError")
	}

	// Create a partially initialized instance
	_child := &_APDUError{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		OriginalInvokeId: originalInvokeId,
		ErrorChoice:      errorChoice,
		Error:            error,
		reservedField0:   reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUError")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 4, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (errorChoice)
		if pushErr := writeBuffer.PushContext("errorChoice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorChoice")
		}
		_errorChoiceErr := writeBuffer.WriteSerializable(ctx, m.GetErrorChoice())
		if popErr := writeBuffer.PopContext("errorChoice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorChoice")
		}
		if _errorChoiceErr != nil {
			return errors.Wrap(_errorChoiceErr, "Error serializing 'errorChoice' field")
		}

		// Simple Field (error)
		if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for error")
		}
		_errorErr := writeBuffer.WriteSerializable(ctx, m.GetError())
		if popErr := writeBuffer.PopContext("error"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for error")
		}
		if _errorErr != nil {
			return errors.Wrap(_errorErr, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("APDUError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUError")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUError) isAPDUError() bool {
	return true
}

func (m *_APDUError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
