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

// ApduDataRestart is the corresponding interface of ApduDataRestart
type ApduDataRestart interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataRestartExactly can be used when we want exactly this type and not a type which fulfills ApduDataRestart.
// This is useful for switch cases.
type ApduDataRestartExactly interface {
	ApduDataRestart
	isApduDataRestart() bool
}

// _ApduDataRestart is the data-structure of this message
type _ApduDataRestart struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataRestart) GetApciType() uint8 {
	return 0xE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataRestart) InitializeParent(parent ApduData) {}

func (m *_ApduDataRestart) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataRestart factory function for _ApduDataRestart
func NewApduDataRestart(dataLength uint8) *_ApduDataRestart {
	_result := &_ApduDataRestart{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataRestart(structType any) ApduDataRestart {
	if casted, ok := structType.(ApduDataRestart); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataRestart); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataRestart) GetTypeName() string {
	return "ApduDataRestart"
}

func (m *_ApduDataRestart) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataRestart) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataRestartParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduDataRestart, error) {
	return ApduDataRestartParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataRestartParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataRestart, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataRestart"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataRestart")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataRestart"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataRestart")
	}

	// Create a partially initialized instance
	_child := &_ApduDataRestart{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataRestart) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataRestart) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataRestart"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataRestart")
		}

		if popErr := writeBuffer.PopContext("ApduDataRestart"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataRestart")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataRestart) isApduDataRestart() bool {
	return true
}

func (m *_ApduDataRestart) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
