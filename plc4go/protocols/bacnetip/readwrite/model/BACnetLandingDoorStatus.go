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

// BACnetLandingDoorStatus is the corresponding interface of BACnetLandingDoorStatus
type BACnetLandingDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLandingDoors returns LandingDoors (property field)
	GetLandingDoors() BACnetLandingDoorStatusLandingDoorsList
}

// BACnetLandingDoorStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetLandingDoorStatus.
// This is useful for switch cases.
type BACnetLandingDoorStatusExactly interface {
	BACnetLandingDoorStatus
	isBACnetLandingDoorStatus() bool
}

// _BACnetLandingDoorStatus is the data-structure of this message
type _BACnetLandingDoorStatus struct {
	LandingDoors BACnetLandingDoorStatusLandingDoorsList
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingDoorStatus) GetLandingDoors() BACnetLandingDoorStatusLandingDoorsList {
	return m.LandingDoors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingDoorStatus factory function for _BACnetLandingDoorStatus
func NewBACnetLandingDoorStatus(landingDoors BACnetLandingDoorStatusLandingDoorsList) *_BACnetLandingDoorStatus {
	return &_BACnetLandingDoorStatus{LandingDoors: landingDoors}
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingDoorStatus(structType any) BACnetLandingDoorStatus {
	if casted, ok := structType.(BACnetLandingDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingDoorStatus) GetTypeName() string {
	return "BACnetLandingDoorStatus"
}

func (m *_BACnetLandingDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (landingDoors)
	lengthInBits += m.LandingDoors.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLandingDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLandingDoorStatusParse(ctx context.Context, theBytes []byte) (BACnetLandingDoorStatus, error) {
	return BACnetLandingDoorStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLandingDoorStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingDoorStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (landingDoors)
	if pullErr := readBuffer.PullContext("landingDoors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingDoors")
	}
	_landingDoors, _landingDoorsErr := BACnetLandingDoorStatusLandingDoorsListParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _landingDoorsErr != nil {
		return nil, errors.Wrap(_landingDoorsErr, "Error parsing 'landingDoors' field of BACnetLandingDoorStatus")
	}
	landingDoors := _landingDoors.(BACnetLandingDoorStatusLandingDoorsList)
	if closeErr := readBuffer.CloseContext("landingDoors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingDoors")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingDoorStatus")
	}

	// Create the instance
	return &_BACnetLandingDoorStatus{
		LandingDoors: landingDoors,
	}, nil
}

func (m *_BACnetLandingDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLandingDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingDoorStatus")
	}

	// Simple Field (landingDoors)
	if pushErr := writeBuffer.PushContext("landingDoors"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for landingDoors")
	}
	_landingDoorsErr := writeBuffer.WriteSerializable(ctx, m.GetLandingDoors())
	if popErr := writeBuffer.PopContext("landingDoors"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for landingDoors")
	}
	if _landingDoorsErr != nil {
		return errors.Wrap(_landingDoorsErr, "Error serializing 'landingDoors' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingDoorStatus")
	}
	return nil
}

func (m *_BACnetLandingDoorStatus) isBACnetLandingDoorStatus() bool {
	return true
}

func (m *_BACnetLandingDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
