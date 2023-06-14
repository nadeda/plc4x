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

// Request is the corresponding interface of Request
type Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() RequestType
	// GetStartingCR returns StartingCR (property field)
	GetStartingCR() *RequestType
	// GetResetMode returns ResetMode (property field)
	GetResetMode() *RequestType
	// GetSecondPeek returns SecondPeek (property field)
	GetSecondPeek() RequestType
	// GetTermination returns Termination (property field)
	GetTermination() RequestTermination
	// GetActualPeek returns ActualPeek (virtual field)
	GetActualPeek() RequestType
}

// RequestExactly can be used when we want exactly this type and not a type which fulfills Request.
// This is useful for switch cases.
type RequestExactly interface {
	Request
	isRequest() bool
}

// _Request is the data-structure of this message
type _Request struct {
	_RequestChildRequirements
	PeekedByte  RequestType
	StartingCR  *RequestType
	ResetMode   *RequestType
	SecondPeek  RequestType
	Termination RequestTermination

	// Arguments.
	CBusOptions CBusOptions
}

type _RequestChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetActualPeek() RequestType
}

type RequestParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Request, serializeChildFunction func() error) error
	GetTypeName() string
}

type RequestChild interface {
	utils.Serializable
	InitializeParent(parent Request, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination)
	GetParent() *Request

	GetTypeName() string
	Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Request) GetPeekedByte() RequestType {
	return m.PeekedByte
}

func (m *_Request) GetStartingCR() *RequestType {
	return m.StartingCR
}

func (m *_Request) GetResetMode() *RequestType {
	return m.ResetMode
}

func (m *_Request) GetSecondPeek() RequestType {
	return m.SecondPeek
}

func (m *_Request) GetTermination() RequestTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_Request) GetActualPeek() RequestType {
	ctx := context.Background()
	_ = ctx
	startingCR := m.StartingCR
	_ = startingCR
	resetMode := m.ResetMode
	_ = resetMode
	return CastRequestType(CastRequestType(utils.InlineIf(bool((bool(bool((m.GetStartingCR()) == (nil))) && bool(bool((m.GetResetMode()) == (nil))))) || bool((bool(bool(bool((m.GetStartingCR()) == (nil))) && bool(bool((m.GetResetMode()) != (nil)))) && bool(bool((m.GetSecondPeek()) == (RequestType_EMPTY))))), func() any { return CastRequestType(m.GetPeekedByte()) }, func() any { return CastRequestType(m.GetSecondPeek()) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequest factory function for _Request
func NewRequest(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_Request {
	return &_Request{PeekedByte: peekedByte, StartingCR: startingCR, ResetMode: resetMode, SecondPeek: secondPeek, Termination: termination, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastRequest(structType any) Request {
	if casted, ok := structType.(Request); ok {
		return casted
	}
	if casted, ok := structType.(*Request); ok {
		return *casted
	}
	return nil
}

func (m *_Request) GetTypeName() string {
	return "Request"
}

func (m *_Request) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (startingCR)
	if m.StartingCR != nil {
		lengthInBits += 8
	}

	// Optional Field (resetMode)
	if m.ResetMode != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (Request, error) {
	return RequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func RequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (Request, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedByte)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedByte")
	}
	peekedByte, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekedByte' field of Request")
	}
	if closeErr := readBuffer.CloseContext("peekedByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for peekedByte")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (startingCR) (Can be skipped, if a given expression evaluates to false)
	var startingCR *RequestType = nil
	if bool((peekedByte) == (RequestType_EMPTY)) {
		if pullErr := readBuffer.PullContext("startingCR"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for startingCR")
		}
		_val, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'startingCR' field of Request")
		}
		startingCR = &_val
		if closeErr := readBuffer.CloseContext("startingCR"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for startingCR")
		}
	}

	// Optional Field (resetMode) (Can be skipped, if a given expression evaluates to false)
	var resetMode *RequestType = nil
	if bool((peekedByte) == (RequestType_RESET)) {
		if pullErr := readBuffer.PullContext("resetMode"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for resetMode")
		}
		_val, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'resetMode' field of Request")
		}
		resetMode = &_val
		if closeErr := readBuffer.CloseContext("resetMode"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for resetMode")
		}
	}

	// Peek Field (secondPeek)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("secondPeek"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for secondPeek")
	}
	secondPeek, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'secondPeek' field of Request")
	}
	if closeErr := readBuffer.CloseContext("secondPeek"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for secondPeek")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_actualPeek := CastRequestType(utils.InlineIf(bool((bool(bool((startingCR) == (nil))) && bool(bool((resetMode) == (nil))))) || bool((bool(bool(bool((startingCR) == (nil))) && bool(bool((resetMode) != (nil)))) && bool(bool((secondPeek) == (RequestType_EMPTY))))), func() any { return CastRequestType(peekedByte) }, func() any { return CastRequestType(secondPeek) }))
	actualPeek := RequestType(_actualPeek)
	_ = actualPeek

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type RequestChildSerializeRequirement interface {
		Request
		InitializeParent(Request, RequestType, *RequestType, *RequestType, RequestType, RequestTermination)
		GetParent() Request
	}
	var _childTemp any
	var _child RequestChildSerializeRequirement
	var typeSwitchError error
	switch {
	case actualPeek == RequestType_SMART_CONNECT_SHORTCUT: // RequestSmartConnectShortcut
		_childTemp, typeSwitchError = RequestSmartConnectShortcutParseWithBuffer(ctx, readBuffer, cBusOptions)
	case actualPeek == RequestType_RESET: // RequestReset
		_childTemp, typeSwitchError = RequestResetParseWithBuffer(ctx, readBuffer, cBusOptions)
	case actualPeek == RequestType_DIRECT_COMMAND: // RequestDirectCommandAccess
		_childTemp, typeSwitchError = RequestDirectCommandAccessParseWithBuffer(ctx, readBuffer, cBusOptions)
	case actualPeek == RequestType_REQUEST_COMMAND: // RequestCommand
		_childTemp, typeSwitchError = RequestCommandParseWithBuffer(ctx, readBuffer, cBusOptions)
	case actualPeek == RequestType_NULL: // RequestNull
		_childTemp, typeSwitchError = RequestNullParseWithBuffer(ctx, readBuffer, cBusOptions)
	case actualPeek == RequestType_EMPTY: // RequestEmpty
		_childTemp, typeSwitchError = RequestEmptyParseWithBuffer(ctx, readBuffer, cBusOptions)
	case 0 == 0: // RequestObsolete
		_childTemp, typeSwitchError = RequestObsoleteParseWithBuffer(ctx, readBuffer, cBusOptions)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [actualPeek=%v]", actualPeek)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of Request")
	}
	_child = _childTemp.(RequestChildSerializeRequirement)

	// Simple Field (termination)
	if pullErr := readBuffer.PullContext("termination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for termination")
	}
	_termination, _terminationErr := RequestTerminationParseWithBuffer(ctx, readBuffer)
	if _terminationErr != nil {
		return nil, errors.Wrap(_terminationErr, "Error parsing 'termination' field of Request")
	}
	termination := _termination.(RequestTermination)
	if closeErr := readBuffer.CloseContext("termination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for termination")
	}

	if closeErr := readBuffer.CloseContext("Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Request")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedByte, startingCR, resetMode, secondPeek, termination)
	return _child, nil
}

func (pm *_Request) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Request, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Request"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Request")
	}

	// Optional Field (startingCR) (Can be skipped, if the value is null)
	var startingCR *RequestType = nil
	if m.GetStartingCR() != nil {
		if pushErr := writeBuffer.PushContext("startingCR"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for startingCR")
		}
		startingCR = m.GetStartingCR()
		_startingCRErr := writeBuffer.WriteSerializable(ctx, startingCR)
		if popErr := writeBuffer.PopContext("startingCR"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for startingCR")
		}
		if _startingCRErr != nil {
			return errors.Wrap(_startingCRErr, "Error serializing 'startingCR' field")
		}
	}

	// Optional Field (resetMode) (Can be skipped, if the value is null)
	var resetMode *RequestType = nil
	if m.GetResetMode() != nil {
		if pushErr := writeBuffer.PushContext("resetMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for resetMode")
		}
		resetMode = m.GetResetMode()
		_resetModeErr := writeBuffer.WriteSerializable(ctx, resetMode)
		if popErr := writeBuffer.PopContext("resetMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for resetMode")
		}
		if _resetModeErr != nil {
			return errors.Wrap(_resetModeErr, "Error serializing 'resetMode' field")
		}
	}
	// Virtual field
	if _actualPeekErr := writeBuffer.WriteVirtual(ctx, "actualPeek", m.GetActualPeek()); _actualPeekErr != nil {
		return errors.Wrap(_actualPeekErr, "Error serializing 'actualPeek' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (termination)
	if pushErr := writeBuffer.PushContext("termination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for termination")
	}
	_terminationErr := writeBuffer.WriteSerializable(ctx, m.GetTermination())
	if popErr := writeBuffer.PopContext("termination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for termination")
	}
	if _terminationErr != nil {
		return errors.Wrap(_terminationErr, "Error serializing 'termination' field")
	}

	if popErr := writeBuffer.PopContext("Request"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Request")
	}
	return nil
}

////
// Arguments Getter

func (m *_Request) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_Request) isRequest() bool {
	return true
}

func (m *_Request) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
