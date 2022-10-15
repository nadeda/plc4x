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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// AirConditioningDataSetPlantHumidityLevel is the corresponding interface of AirConditioningDataSetPlantHumidityLevel
type AirConditioningDataSetPlantHumidityLevel interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
}

// AirConditioningDataSetPlantHumidityLevelExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetPlantHumidityLevel.
// This is useful for switch cases.
type AirConditioningDataSetPlantHumidityLevelExactly interface {
	AirConditioningDataSetPlantHumidityLevel
	isAirConditioningDataSetPlantHumidityLevel() bool
}

// _AirConditioningDataSetPlantHumidityLevel is the data-structure of this message
type _AirConditioningDataSetPlantHumidityLevel struct {
	*_AirConditioningData
	ZoneGroup            byte
	ZoneList             HVACZoneList
	HumidityModeAndFlags HVACHumidityModeAndFlags
	HumidityType         HVACHumidityType
	Level                HVACHumidity
	RawLevel             HVACRawLevels
	AuxLevel             HVACAuxiliaryLevel
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetPlantHumidityLevel) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetPlantHumidityLevel) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetPlantHumidityLevel factory function for _AirConditioningDataSetPlantHumidityLevel
func NewAirConditioningDataSetPlantHumidityLevel(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType, level HVACHumidity, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetPlantHumidityLevel {
	_result := &_AirConditioningDataSetPlantHumidityLevel{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HumidityModeAndFlags: humidityModeAndFlags,
		HumidityType:         humidityType,
		Level:                level,
		RawLevel:             rawLevel,
		AuxLevel:             auxLevel,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetPlantHumidityLevel(structType interface{}) AirConditioningDataSetPlantHumidityLevel {
	if casted, ok := structType.(AirConditioningDataSetPlantHumidityLevel); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetPlantHumidityLevel); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetTypeName() string {
	return "AirConditioningDataSetPlantHumidityLevel"
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits()

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits()

	// Simple field (humidityType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits()
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits()
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataSetPlantHumidityLevelParse(readBuffer utils.ReadBuffer) (AirConditioningDataSetPlantHumidityLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetPlantHumidityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetPlantHumidityLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetPlantHumidityLevel")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParse(readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataSetPlantHumidityLevel")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (humidityModeAndFlags)
	if pullErr := readBuffer.PullContext("humidityModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityModeAndFlags")
	}
	_humidityModeAndFlags, _humidityModeAndFlagsErr := HVACHumidityModeAndFlagsParse(readBuffer)
	if _humidityModeAndFlagsErr != nil {
		return nil, errors.Wrap(_humidityModeAndFlagsErr, "Error parsing 'humidityModeAndFlags' field of AirConditioningDataSetPlantHumidityLevel")
	}
	humidityModeAndFlags := _humidityModeAndFlags.(HVACHumidityModeAndFlags)
	if closeErr := readBuffer.CloseContext("humidityModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityModeAndFlags")
	}

	// Simple Field (humidityType)
	if pullErr := readBuffer.PullContext("humidityType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityType")
	}
	_humidityType, _humidityTypeErr := HVACHumidityTypeParse(readBuffer)
	if _humidityTypeErr != nil {
		return nil, errors.Wrap(_humidityTypeErr, "Error parsing 'humidityType' field of AirConditioningDataSetPlantHumidityLevel")
	}
	humidityType := _humidityType
	if closeErr := readBuffer.CloseContext("humidityType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityType")
	}

	// Optional Field (level) (Can be skipped, if a given expression evaluates to false)
	var level HVACHumidity = nil
	if humidityModeAndFlags.GetIsLevelHumidity() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("level"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for level")
		}
		_val, _err := HVACHumidityParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'level' field of AirConditioningDataSetPlantHumidityLevel")
		default:
			level = _val.(HVACHumidity)
			if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for level")
			}
		}
	}

	// Optional Field (rawLevel) (Can be skipped, if a given expression evaluates to false)
	var rawLevel HVACRawLevels = nil
	if humidityModeAndFlags.GetIsLevelRaw() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("rawLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for rawLevel")
		}
		_val, _err := HVACRawLevelsParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawLevel' field of AirConditioningDataSetPlantHumidityLevel")
		default:
			rawLevel = _val.(HVACRawLevels)
			if closeErr := readBuffer.CloseContext("rawLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for rawLevel")
			}
		}
	}

	// Optional Field (auxLevel) (Can be skipped, if a given expression evaluates to false)
	var auxLevel HVACAuxiliaryLevel = nil
	if humidityModeAndFlags.GetIsAuxLevelUsed() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("auxLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for auxLevel")
		}
		_val, _err := HVACAuxiliaryLevelParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'auxLevel' field of AirConditioningDataSetPlantHumidityLevel")
		default:
			auxLevel = _val.(HVACAuxiliaryLevel)
			if closeErr := readBuffer.CloseContext("auxLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for auxLevel")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetPlantHumidityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetPlantHumidityLevel")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetPlantHumidityLevel{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HumidityModeAndFlags: humidityModeAndFlags,
		HumidityType:         humidityType,
		Level:                level,
		RawLevel:             rawLevel,
		AuxLevel:             auxLevel,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetPlantHumidityLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetPlantHumidityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetPlantHumidityLevel")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (humidityModeAndFlags)
		if pushErr := writeBuffer.PushContext("humidityModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityModeAndFlags")
		}
		_humidityModeAndFlagsErr := writeBuffer.WriteSerializable(m.GetHumidityModeAndFlags())
		if popErr := writeBuffer.PopContext("humidityModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityModeAndFlags")
		}
		if _humidityModeAndFlagsErr != nil {
			return errors.Wrap(_humidityModeAndFlagsErr, "Error serializing 'humidityModeAndFlags' field")
		}

		// Simple Field (humidityType)
		if pushErr := writeBuffer.PushContext("humidityType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityType")
		}
		_humidityTypeErr := writeBuffer.WriteSerializable(m.GetHumidityType())
		if popErr := writeBuffer.PopContext("humidityType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityType")
		}
		if _humidityTypeErr != nil {
			return errors.Wrap(_humidityTypeErr, "Error serializing 'humidityType' field")
		}

		// Optional Field (level) (Can be skipped, if the value is null)
		var level HVACHumidity = nil
		if m.GetLevel() != nil {
			if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for level")
			}
			level = m.GetLevel()
			_levelErr := writeBuffer.WriteSerializable(level)
			if popErr := writeBuffer.PopContext("level"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for level")
			}
			if _levelErr != nil {
				return errors.Wrap(_levelErr, "Error serializing 'level' field")
			}
		}

		// Optional Field (rawLevel) (Can be skipped, if the value is null)
		var rawLevel HVACRawLevels = nil
		if m.GetRawLevel() != nil {
			if pushErr := writeBuffer.PushContext("rawLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for rawLevel")
			}
			rawLevel = m.GetRawLevel()
			_rawLevelErr := writeBuffer.WriteSerializable(rawLevel)
			if popErr := writeBuffer.PopContext("rawLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for rawLevel")
			}
			if _rawLevelErr != nil {
				return errors.Wrap(_rawLevelErr, "Error serializing 'rawLevel' field")
			}
		}

		// Optional Field (auxLevel) (Can be skipped, if the value is null)
		var auxLevel HVACAuxiliaryLevel = nil
		if m.GetAuxLevel() != nil {
			if pushErr := writeBuffer.PushContext("auxLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for auxLevel")
			}
			auxLevel = m.GetAuxLevel()
			_auxLevelErr := writeBuffer.WriteSerializable(auxLevel)
			if popErr := writeBuffer.PopContext("auxLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for auxLevel")
			}
			if _auxLevelErr != nil {
				return errors.Wrap(_auxLevelErr, "Error serializing 'auxLevel' field")
			}
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetPlantHumidityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetPlantHumidityLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetPlantHumidityLevel) isAirConditioningDataSetPlantHumidityLevel() bool {
	return true
}

func (m *_AirConditioningDataSetPlantHumidityLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
