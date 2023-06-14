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

// LevelInformation is the corresponding interface of LevelInformation
type LevelInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetRaw returns Raw (property field)
	GetRaw() uint16
	// GetNibble1 returns Nibble1 (virtual field)
	GetNibble1() uint8
	// GetNibble2 returns Nibble2 (virtual field)
	GetNibble2() uint8
	// GetNibble3 returns Nibble3 (virtual field)
	GetNibble3() uint8
	// GetNibble4 returns Nibble4 (virtual field)
	GetNibble4() uint8
	// GetIsAbsent returns IsAbsent (virtual field)
	GetIsAbsent() bool
	// GetIsCorruptedByNoise returns IsCorruptedByNoise (virtual field)
	GetIsCorruptedByNoise() bool
	// GetIsCorruptedByNoiseOrLevelsDiffer returns IsCorruptedByNoiseOrLevelsDiffer (virtual field)
	GetIsCorruptedByNoiseOrLevelsDiffer() bool
	// GetIsCorrupted returns IsCorrupted (virtual field)
	GetIsCorrupted() bool
}

// LevelInformationExactly can be used when we want exactly this type and not a type which fulfills LevelInformation.
// This is useful for switch cases.
type LevelInformationExactly interface {
	LevelInformation
	isLevelInformation() bool
}

// _LevelInformation is the data-structure of this message
type _LevelInformation struct {
	_LevelInformationChildRequirements
	Raw uint16
}

type _LevelInformationChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetIsAbsent() bool
	GetIsCorrupted() bool
}

type LevelInformationParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LevelInformation, serializeChildFunction func() error) error
	GetTypeName() string
}

type LevelInformationChild interface {
	utils.Serializable
	InitializeParent(parent LevelInformation, raw uint16)
	GetParent() *LevelInformation

	GetTypeName() string
	LevelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LevelInformation) GetRaw() uint16 {
	return m.Raw
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_LevelInformation) GetNibble1() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0xF000) >> uint8(12))
}

func (m *_LevelInformation) GetNibble2() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x0F00) >> uint8(8))
}

func (m *_LevelInformation) GetNibble3() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x00F0) >> uint8(4))
}

func (m *_LevelInformation) GetNibble4() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x000F) >> uint8(0))
}

func (m *_LevelInformation) GetIsAbsent() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(bool(bool((m.GetNibble1()) == (0x0))) && bool(bool((m.GetNibble2()) == (0x0)))) && bool(bool((m.GetNibble3()) == (0x0)))) && bool(bool((m.GetNibble4()) == (0x0))))
}

func (m *_LevelInformation) GetIsCorruptedByNoise() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(!(m.GetIsAbsent())) && bool((bool(bool(bool((bool(bool((bool((m.GetNibble1()) < (0x5)))) || bool((bool((m.GetNibble1()) == (0x8))))) || bool((bool((m.GetNibble1()) == (0xC)))))) || bool((bool(bool((bool((m.GetNibble2()) < (0x5)))) || bool((bool((m.GetNibble2()) == (0x8))))) || bool((bool((m.GetNibble2()) == (0xC))))))) || bool((bool(bool((bool((m.GetNibble3()) < (0x5)))) || bool((bool((m.GetNibble3()) == (0x8))))) || bool((bool((m.GetNibble3()) == (0xC))))))) || bool((bool(bool((bool((m.GetNibble4()) < (0x5)))) || bool((bool((m.GetNibble4()) == (0x8))))) || bool((bool((m.GetNibble4()) == (0xC)))))))))
}

func (m *_LevelInformation) GetIsCorruptedByNoiseOrLevelsDiffer() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(!(m.GetIsAbsent())) && bool((bool(bool(bool((bool(bool((bool((m.GetNibble1()) == (0x7)))) || bool((bool((m.GetNibble1()) == (0xB))))) || bool((bool((m.GetNibble1()) > (0xC)))))) || bool((bool(bool((bool((m.GetNibble2()) == (0x7)))) || bool((bool((m.GetNibble2()) == (0xB))))) || bool((bool((m.GetNibble2()) > (0xC))))))) || bool((bool(bool((bool((m.GetNibble3()) == (0x7)))) || bool((bool((m.GetNibble3()) == (0xB))))) || bool((bool((m.GetNibble3()) > (0xC))))))) || bool((bool(bool((bool((m.GetNibble4()) == (0x7)))) || bool((bool((m.GetNibble4()) == (0xB))))) || bool((bool((m.GetNibble4()) > (0xC)))))))))
}

func (m *_LevelInformation) GetIsCorrupted() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(m.GetIsCorruptedByNoise()) || bool(m.GetIsCorruptedByNoiseOrLevelsDiffer()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLevelInformation factory function for _LevelInformation
func NewLevelInformation(raw uint16) *_LevelInformation {
	return &_LevelInformation{Raw: raw}
}

// Deprecated: use the interface for direct cast
func CastLevelInformation(structType any) LevelInformation {
	if casted, ok := structType.(LevelInformation); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformation); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformation) GetTypeName() string {
	return "LevelInformation"
}

func (m *_LevelInformation) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_LevelInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LevelInformationParse(ctx context.Context, theBytes []byte) (LevelInformation, error) {
	return LevelInformationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LevelInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LevelInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (raw)
	currentPos = positionAware.GetPos()
	raw, _err := readBuffer.ReadUint16("raw", 16)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'raw' field of LevelInformation")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_nibble1 := (raw & 0xF000) >> uint8(12)
	nibble1 := uint8(_nibble1)
	_ = nibble1

	// Virtual field
	_nibble2 := (raw & 0x0F00) >> uint8(8)
	nibble2 := uint8(_nibble2)
	_ = nibble2

	// Virtual field
	_nibble3 := (raw & 0x00F0) >> uint8(4)
	nibble3 := uint8(_nibble3)
	_ = nibble3

	// Virtual field
	_nibble4 := (raw & 0x000F) >> uint8(0)
	nibble4 := uint8(_nibble4)
	_ = nibble4

	// Virtual field
	_isAbsent := bool(bool(bool(bool((nibble1) == (0x0))) && bool(bool((nibble2) == (0x0)))) && bool(bool((nibble3) == (0x0)))) && bool(bool((nibble4) == (0x0)))
	isAbsent := bool(_isAbsent)
	_ = isAbsent

	// Virtual field
	_isCorruptedByNoise := bool(!(isAbsent)) && bool((bool(bool(bool((bool(bool((bool((nibble1) < (0x5)))) || bool((bool((nibble1) == (0x8))))) || bool((bool((nibble1) == (0xC)))))) || bool((bool(bool((bool((nibble2) < (0x5)))) || bool((bool((nibble2) == (0x8))))) || bool((bool((nibble2) == (0xC))))))) || bool((bool(bool((bool((nibble3) < (0x5)))) || bool((bool((nibble3) == (0x8))))) || bool((bool((nibble3) == (0xC))))))) || bool((bool(bool((bool((nibble4) < (0x5)))) || bool((bool((nibble4) == (0x8))))) || bool((bool((nibble4) == (0xC))))))))
	isCorruptedByNoise := bool(_isCorruptedByNoise)
	_ = isCorruptedByNoise

	// Virtual field
	_isCorruptedByNoiseOrLevelsDiffer := bool(!(isAbsent)) && bool((bool(bool(bool((bool(bool((bool((nibble1) == (0x7)))) || bool((bool((nibble1) == (0xB))))) || bool((bool((nibble1) > (0xC)))))) || bool((bool(bool((bool((nibble2) == (0x7)))) || bool((bool((nibble2) == (0xB))))) || bool((bool((nibble2) > (0xC))))))) || bool((bool(bool((bool((nibble3) == (0x7)))) || bool((bool((nibble3) == (0xB))))) || bool((bool((nibble3) > (0xC))))))) || bool((bool(bool((bool((nibble4) == (0x7)))) || bool((bool((nibble4) == (0xB))))) || bool((bool((nibble4) > (0xC))))))))
	isCorruptedByNoiseOrLevelsDiffer := bool(_isCorruptedByNoiseOrLevelsDiffer)
	_ = isCorruptedByNoiseOrLevelsDiffer

	// Virtual field
	_isCorrupted := bool(isCorruptedByNoise) || bool(isCorruptedByNoiseOrLevelsDiffer)
	isCorrupted := bool(_isCorrupted)
	_ = isCorrupted

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type LevelInformationChildSerializeRequirement interface {
		LevelInformation
		InitializeParent(LevelInformation, uint16)
		GetParent() LevelInformation
	}
	var _childTemp any
	var _child LevelInformationChildSerializeRequirement
	var typeSwitchError error
	switch {
	case isAbsent == bool(true): // LevelInformationAbsent
		_childTemp, typeSwitchError = LevelInformationAbsentParseWithBuffer(ctx, readBuffer)
	case 0 == 0 && isCorrupted == bool(true): // LevelInformationCorrupted
		_childTemp, typeSwitchError = LevelInformationCorruptedParseWithBuffer(ctx, readBuffer)
	case 0 == 0: // LevelInformationNormal
		_childTemp, typeSwitchError = LevelInformationNormalParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [isAbsent=%v, isCorrupted=%v]", isAbsent, isCorrupted)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of LevelInformation")
	}
	_child = _childTemp.(LevelInformationChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("LevelInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformation")
	}

	// Finish initializing
	_child.InitializeParent(_child, raw)
	return _child, nil
}

func (pm *_LevelInformation) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LevelInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("LevelInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LevelInformation")
	}
	// Virtual field
	if _nibble1Err := writeBuffer.WriteVirtual(ctx, "nibble1", m.GetNibble1()); _nibble1Err != nil {
		return errors.Wrap(_nibble1Err, "Error serializing 'nibble1' field")
	}
	// Virtual field
	if _nibble2Err := writeBuffer.WriteVirtual(ctx, "nibble2", m.GetNibble2()); _nibble2Err != nil {
		return errors.Wrap(_nibble2Err, "Error serializing 'nibble2' field")
	}
	// Virtual field
	if _nibble3Err := writeBuffer.WriteVirtual(ctx, "nibble3", m.GetNibble3()); _nibble3Err != nil {
		return errors.Wrap(_nibble3Err, "Error serializing 'nibble3' field")
	}
	// Virtual field
	if _nibble4Err := writeBuffer.WriteVirtual(ctx, "nibble4", m.GetNibble4()); _nibble4Err != nil {
		return errors.Wrap(_nibble4Err, "Error serializing 'nibble4' field")
	}
	// Virtual field
	if _isAbsentErr := writeBuffer.WriteVirtual(ctx, "isAbsent", m.GetIsAbsent()); _isAbsentErr != nil {
		return errors.Wrap(_isAbsentErr, "Error serializing 'isAbsent' field")
	}
	// Virtual field
	if _isCorruptedByNoiseErr := writeBuffer.WriteVirtual(ctx, "isCorruptedByNoise", m.GetIsCorruptedByNoise()); _isCorruptedByNoiseErr != nil {
		return errors.Wrap(_isCorruptedByNoiseErr, "Error serializing 'isCorruptedByNoise' field")
	}
	// Virtual field
	if _isCorruptedByNoiseOrLevelsDifferErr := writeBuffer.WriteVirtual(ctx, "isCorruptedByNoiseOrLevelsDiffer", m.GetIsCorruptedByNoiseOrLevelsDiffer()); _isCorruptedByNoiseOrLevelsDifferErr != nil {
		return errors.Wrap(_isCorruptedByNoiseOrLevelsDifferErr, "Error serializing 'isCorruptedByNoiseOrLevelsDiffer' field")
	}
	// Virtual field
	if _isCorruptedErr := writeBuffer.WriteVirtual(ctx, "isCorrupted", m.GetIsCorrupted()); _isCorruptedErr != nil {
		return errors.Wrap(_isCorruptedErr, "Error serializing 'isCorrupted' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LevelInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LevelInformation")
	}
	return nil
}

func (m *_LevelInformation) isLevelInformation() bool {
	return true
}

func (m *_LevelInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
