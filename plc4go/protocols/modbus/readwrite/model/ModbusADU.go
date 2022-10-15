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
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusADU is the corresponding interface of ModbusADU
type ModbusADU interface {
	utils.LengthAware
	utils.Serializable
	// GetDriverType returns DriverType (discriminator field)
	GetDriverType() DriverType
}

// ModbusADUExactly can be used when we want exactly this type and not a type which fulfills ModbusADU.
// This is useful for switch cases.
type ModbusADUExactly interface {
	ModbusADU
	isModbusADU() bool
}

// _ModbusADU is the data-structure of this message
type _ModbusADU struct {
	_ModbusADUChildRequirements

	// Arguments.
	Response bool
}

type _ModbusADUChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetDriverType() DriverType
}

type ModbusADUParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ModbusADU, serializeChildFunction func() error) error
	GetTypeName() string
}

type ModbusADUChild interface {
	utils.Serializable
	InitializeParent(parent ModbusADU)
	GetParent() *ModbusADU

	GetTypeName() string
	ModbusADU
}

// NewModbusADU factory function for _ModbusADU
func NewModbusADU(response bool) *_ModbusADU {
	return &_ModbusADU{Response: response}
}

// Deprecated: use the interface for direct cast
func CastModbusADU(structType interface{}) ModbusADU {
	if casted, ok := structType.(ModbusADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusADU) GetTypeName() string {
	return "ModbusADU"
}

func (m *_ModbusADU) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ModbusADU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusADUParse(readBuffer utils.ReadBuffer, driverType DriverType, response bool) (ModbusADU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ModbusADUChildSerializeRequirement interface {
		ModbusADU
		InitializeParent(ModbusADU)
		GetParent() ModbusADU
	}
	var _childTemp interface{}
	var _child ModbusADUChildSerializeRequirement
	var typeSwitchError error
	switch {
	case driverType == DriverType_MODBUS_TCP: // ModbusTcpADU
		_childTemp, typeSwitchError = ModbusTcpADUParse(readBuffer, driverType, response)
	case driverType == DriverType_MODBUS_RTU: // ModbusRtuADU
		_childTemp, typeSwitchError = ModbusRtuADUParse(readBuffer, driverType, response)
	case driverType == DriverType_MODBUS_ASCII: // ModbusAsciiADU
		_childTemp, typeSwitchError = ModbusAsciiADUParse(readBuffer, driverType, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [driverType=%v]", driverType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ModbusADU")
	}
	_child = _childTemp.(ModbusADUChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ModbusADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusADU")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ModbusADU) SerializeParent(writeBuffer utils.WriteBuffer, child ModbusADU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusADU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusADU")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ModbusADU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusADU")
	}
	return nil
}

////
// Arguments Getter

func (m *_ModbusADU) GetResponse() bool {
	return m.Response
}

//
////

func (m *_ModbusADU) isModbusADU() bool {
	return true
}

func (m *_ModbusADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
