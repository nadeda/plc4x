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

// ConfirmedPrivateTransferError is the corresponding interface of ConfirmedPrivateTransferError
type ConfirmedPrivateTransferError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetErrorParameters returns ErrorParameters (property field)
	GetErrorParameters() BACnetConstructedData
}

// ConfirmedPrivateTransferErrorExactly can be used when we want exactly this type and not a type which fulfills ConfirmedPrivateTransferError.
// This is useful for switch cases.
type ConfirmedPrivateTransferErrorExactly interface {
	ConfirmedPrivateTransferError
	isConfirmedPrivateTransferError() bool
}

// _ConfirmedPrivateTransferError is the data-structure of this message
type _ConfirmedPrivateTransferError struct {
	*_BACnetError
	ErrorType       ErrorEnclosed
	VendorId        BACnetVendorIdTagged
	ServiceNumber   BACnetContextTagUnsignedInteger
	ErrorParameters BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConfirmedPrivateTransferError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConfirmedPrivateTransferError) InitializeParent(parent BACnetError) {}

func (m *_ConfirmedPrivateTransferError) GetParent() BACnetError {
	return m._BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConfirmedPrivateTransferError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ConfirmedPrivateTransferError) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_ConfirmedPrivateTransferError) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_ConfirmedPrivateTransferError) GetErrorParameters() BACnetConstructedData {
	return m.ErrorParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfirmedPrivateTransferError factory function for _ConfirmedPrivateTransferError
func NewConfirmedPrivateTransferError(errorType ErrorEnclosed, vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, errorParameters BACnetConstructedData) *_ConfirmedPrivateTransferError {
	_result := &_ConfirmedPrivateTransferError{
		ErrorType:       errorType,
		VendorId:        vendorId,
		ServiceNumber:   serviceNumber,
		ErrorParameters: errorParameters,
		_BACnetError:    NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConfirmedPrivateTransferError(structType any) ConfirmedPrivateTransferError {
	if casted, ok := structType.(ConfirmedPrivateTransferError); ok {
		return casted
	}
	if casted, ok := structType.(*ConfirmedPrivateTransferError); ok {
		return *casted
	}
	return nil
}

func (m *_ConfirmedPrivateTransferError) GetTypeName() string {
	return "ConfirmedPrivateTransferError"
}

func (m *_ConfirmedPrivateTransferError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits(ctx)

	// Optional Field (errorParameters)
	if m.ErrorParameters != nil {
		lengthInBits += m.ErrorParameters.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_ConfirmedPrivateTransferError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConfirmedPrivateTransferErrorParse(ctx context.Context, theBytes []byte, errorChoice BACnetConfirmedServiceChoice) (ConfirmedPrivateTransferError, error) {
	return ConfirmedPrivateTransferErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), errorChoice)
}

func ConfirmedPrivateTransferErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (ConfirmedPrivateTransferError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConfirmedPrivateTransferError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConfirmedPrivateTransferError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
	_errorType, _errorTypeErr := ErrorEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field of ConfirmedPrivateTransferError")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of ConfirmedPrivateTransferError")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceNumber")
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field of ConfirmedPrivateTransferError")
	}
	serviceNumber := _serviceNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceNumber")
	}

	// Optional Field (errorParameters) (Can be skipped, if a given expression evaluates to false)
	var errorParameters BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("errorParameters"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for errorParameters")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(3), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'errorParameters' field of ConfirmedPrivateTransferError")
		default:
			errorParameters = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("errorParameters"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for errorParameters")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("ConfirmedPrivateTransferError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConfirmedPrivateTransferError")
	}

	// Create a partially initialized instance
	_child := &_ConfirmedPrivateTransferError{
		_BACnetError:    &_BACnetError{},
		ErrorType:       errorType,
		VendorId:        vendorId,
		ServiceNumber:   serviceNumber,
		ErrorParameters: errorParameters,
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_ConfirmedPrivateTransferError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConfirmedPrivateTransferError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConfirmedPrivateTransferError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConfirmedPrivateTransferError")
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorType")
		}
		_errorTypeErr := writeBuffer.WriteSerializable(ctx, m.GetErrorType())
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorType")
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceNumber")
		}
		_serviceNumberErr := writeBuffer.WriteSerializable(ctx, m.GetServiceNumber())
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceNumber")
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (errorParameters) (Can be skipped, if the value is null)
		var errorParameters BACnetConstructedData = nil
		if m.GetErrorParameters() != nil {
			if pushErr := writeBuffer.PushContext("errorParameters"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for errorParameters")
			}
			errorParameters = m.GetErrorParameters()
			_errorParametersErr := writeBuffer.WriteSerializable(ctx, errorParameters)
			if popErr := writeBuffer.PopContext("errorParameters"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for errorParameters")
			}
			if _errorParametersErr != nil {
				return errors.Wrap(_errorParametersErr, "Error serializing 'errorParameters' field")
			}
		}

		if popErr := writeBuffer.PopContext("ConfirmedPrivateTransferError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConfirmedPrivateTransferError")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConfirmedPrivateTransferError) isConfirmedPrivateTransferError() bool {
	return true
}

func (m *_ConfirmedPrivateTransferError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
