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

// MeteringDataMeasureOil is the corresponding interface of MeteringDataMeasureOil
type MeteringDataMeasureOil interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
}

// MeteringDataMeasureOilExactly can be used when we want exactly this type and not a type which fulfills MeteringDataMeasureOil.
// This is useful for switch cases.
type MeteringDataMeasureOilExactly interface {
	MeteringDataMeasureOil
	isMeteringDataMeasureOil() bool
}

// _MeteringDataMeasureOil is the data-structure of this message
type _MeteringDataMeasureOil struct {
	*_MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataMeasureOil) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataMeasureOil) GetParent() MeteringData {
	return m._MeteringData
}

// NewMeteringDataMeasureOil factory function for _MeteringDataMeasureOil
func NewMeteringDataMeasureOil(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureOil {
	_result := &_MeteringDataMeasureOil{
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureOil(structType any) MeteringDataMeasureOil {
	if casted, ok := structType.(MeteringDataMeasureOil); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureOil); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureOil) GetTypeName() string {
	return "MeteringDataMeasureOil"
}

func (m *_MeteringDataMeasureOil) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureOil) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataMeasureOilParse(ctx context.Context, theBytes []byte) (MeteringDataMeasureOil, error) {
	return MeteringDataMeasureOilParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataMeasureOilParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureOil, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureOil"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureOil")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureOil"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureOil")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataMeasureOil{
		_MeteringData: &_MeteringData{},
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataMeasureOil) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureOil) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureOil"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureOil")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureOil"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureOil")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureOil) isMeteringDataMeasureOil() bool {
	return true
}

func (m *_MeteringDataMeasureOil) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
