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

// SecurityDataPanicActivated is the corresponding interface of SecurityDataPanicActivated
type SecurityDataPanicActivated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataPanicActivatedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataPanicActivated.
// This is useful for switch cases.
type SecurityDataPanicActivatedExactly interface {
	SecurityDataPanicActivated
	isSecurityDataPanicActivated() bool
}

// _SecurityDataPanicActivated is the data-structure of this message
type _SecurityDataPanicActivated struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataPanicActivated) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataPanicActivated) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataPanicActivated factory function for _SecurityDataPanicActivated
func NewSecurityDataPanicActivated(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataPanicActivated {
	_result := &_SecurityDataPanicActivated{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPanicActivated(structType any) SecurityDataPanicActivated {
	if casted, ok := structType.(SecurityDataPanicActivated); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPanicActivated); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPanicActivated) GetTypeName() string {
	return "SecurityDataPanicActivated"
}

func (m *_SecurityDataPanicActivated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataPanicActivated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataPanicActivatedParse(ctx context.Context, theBytes []byte) (SecurityDataPanicActivated, error) {
	return SecurityDataPanicActivatedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataPanicActivatedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataPanicActivated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPanicActivated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPanicActivated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataPanicActivated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPanicActivated")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataPanicActivated{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataPanicActivated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataPanicActivated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPanicActivated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPanicActivated")
		}

		if popErr := writeBuffer.PopContext("SecurityDataPanicActivated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPanicActivated")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataPanicActivated) isSecurityDataPanicActivated() bool {
	return true
}

func (m *_SecurityDataPanicActivated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
