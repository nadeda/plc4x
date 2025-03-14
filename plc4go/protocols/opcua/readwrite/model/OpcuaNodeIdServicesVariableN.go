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

// OpcuaNodeIdServicesVariableN is an enum
type OpcuaNodeIdServicesVariableN int32

type IOpcuaNodeIdServicesVariableN interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Definition       OpcuaNodeIdServicesVariableN = 12069
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_ValuePrecision   OpcuaNodeIdServicesVariableN = 12070
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_InstrumentRange  OpcuaNodeIdServicesVariableN = 12071
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EURange          OpcuaNodeIdServicesVariableN = 12072
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EngineeringUnits OpcuaNodeIdServicesVariableN = 12073
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Title            OpcuaNodeIdServicesVariableN = 12074
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisScaleType    OpcuaNodeIdServicesVariableN = 12075
	OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisDefinition   OpcuaNodeIdServicesVariableN = 12076
)

var OpcuaNodeIdServicesVariableNValues []OpcuaNodeIdServicesVariableN

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNValues = []OpcuaNodeIdServicesVariableN{
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Definition,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_ValuePrecision,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_InstrumentRange,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EURange,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EngineeringUnits,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Title,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisScaleType,
		OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisDefinition,
	}
}

func OpcuaNodeIdServicesVariableNByValue(value int32) (enum OpcuaNodeIdServicesVariableN, ok bool) {
	switch value {
	case 12069:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Definition, true
	case 12070:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_ValuePrecision, true
	case 12071:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_InstrumentRange, true
	case 12072:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EURange, true
	case 12073:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EngineeringUnits, true
	case 12074:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Title, true
	case 12075:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisScaleType, true
	case 12076:
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNByName(value string) (enum OpcuaNodeIdServicesVariableN, ok bool) {
	switch value {
	case "NDimensionArrayItemType_Definition":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Definition, true
	case "NDimensionArrayItemType_ValuePrecision":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_ValuePrecision, true
	case "NDimensionArrayItemType_InstrumentRange":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_InstrumentRange, true
	case "NDimensionArrayItemType_EURange":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EURange, true
	case "NDimensionArrayItemType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EngineeringUnits, true
	case "NDimensionArrayItemType_Title":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Title, true
	case "NDimensionArrayItemType_AxisScaleType":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisScaleType, true
	case "NDimensionArrayItemType_AxisDefinition":
		return OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableN(structType any) OpcuaNodeIdServicesVariableN {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableN {
		if sOpcuaNodeIdServicesVariableN, ok := typ.(OpcuaNodeIdServicesVariableN); ok {
			return sOpcuaNodeIdServicesVariableN
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableN) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableN) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableN, error) {
	return OpcuaNodeIdServicesVariableNParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableN, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableN", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableN")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableN")
		return OpcuaNodeIdServicesVariableN(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableN) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableN) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableN", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableN) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Definition:
		return "NDimensionArrayItemType_Definition"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_ValuePrecision:
		return "NDimensionArrayItemType_ValuePrecision"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_InstrumentRange:
		return "NDimensionArrayItemType_InstrumentRange"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EURange:
		return "NDimensionArrayItemType_EURange"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_EngineeringUnits:
		return "NDimensionArrayItemType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_Title:
		return "NDimensionArrayItemType_Title"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisScaleType:
		return "NDimensionArrayItemType_AxisScaleType"
	case OpcuaNodeIdServicesVariableN_NDimensionArrayItemType_AxisDefinition:
		return "NDimensionArrayItemType_AxisDefinition"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableN) String() string {
	return e.PLC4XEnumName()
}
