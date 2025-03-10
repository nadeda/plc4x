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

// DeviceGroup is an enum
type DeviceGroup uint8

type IDeviceGroup interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DeviceGroup_PG_OR_PC DeviceGroup = 0x01
	DeviceGroup_OS       DeviceGroup = 0x02
	DeviceGroup_OTHERS   DeviceGroup = 0x03
)

var DeviceGroupValues []DeviceGroup

func init() {
	_ = errors.New
	DeviceGroupValues = []DeviceGroup{
		DeviceGroup_PG_OR_PC,
		DeviceGroup_OS,
		DeviceGroup_OTHERS,
	}
}

func DeviceGroupByValue(value uint8) (enum DeviceGroup, ok bool) {
	switch value {
	case 0x01:
		return DeviceGroup_PG_OR_PC, true
	case 0x02:
		return DeviceGroup_OS, true
	case 0x03:
		return DeviceGroup_OTHERS, true
	}
	return 0, false
}

func DeviceGroupByName(value string) (enum DeviceGroup, ok bool) {
	switch value {
	case "PG_OR_PC":
		return DeviceGroup_PG_OR_PC, true
	case "OS":
		return DeviceGroup_OS, true
	case "OTHERS":
		return DeviceGroup_OTHERS, true
	}
	return 0, false
}

func DeviceGroupKnows(value uint8) bool {
	for _, typeValue := range DeviceGroupValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDeviceGroup(structType any) DeviceGroup {
	castFunc := func(typ any) DeviceGroup {
		if sDeviceGroup, ok := typ.(DeviceGroup); ok {
			return sDeviceGroup
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceGroup) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m DeviceGroup) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceGroupParse(ctx context.Context, theBytes []byte) (DeviceGroup, error) {
	return DeviceGroupParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceGroupParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceGroup, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("DeviceGroup", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DeviceGroup")
	}
	if enum, ok := DeviceGroupByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DeviceGroup")
		return DeviceGroup(val), nil
	} else {
		return enum, nil
	}
}

func (e DeviceGroup) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DeviceGroup) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("DeviceGroup", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DeviceGroup) PLC4XEnumName() string {
	switch e {
	case DeviceGroup_PG_OR_PC:
		return "PG_OR_PC"
	case DeviceGroup_OS:
		return "OS"
	case DeviceGroup_OTHERS:
		return "OTHERS"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e DeviceGroup) String() string {
	return e.PLC4XEnumName()
}
