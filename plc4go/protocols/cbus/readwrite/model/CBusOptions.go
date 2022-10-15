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

// CBusOptions is the corresponding interface of CBusOptions
type CBusOptions interface {
	utils.LengthAware
	utils.Serializable
	// GetConnect returns Connect (property field)
	GetConnect() bool
	// GetSmart returns Smart (property field)
	GetSmart() bool
	// GetIdmon returns Idmon (property field)
	GetIdmon() bool
	// GetExstat returns Exstat (property field)
	GetExstat() bool
	// GetMonitor returns Monitor (property field)
	GetMonitor() bool
	// GetMonall returns Monall (property field)
	GetMonall() bool
	// GetPun returns Pun (property field)
	GetPun() bool
	// GetPcn returns Pcn (property field)
	GetPcn() bool
	// GetSrchk returns Srchk (property field)
	GetSrchk() bool
}

// CBusOptionsExactly can be used when we want exactly this type and not a type which fulfills CBusOptions.
// This is useful for switch cases.
type CBusOptionsExactly interface {
	CBusOptions
	isCBusOptions() bool
}

// _CBusOptions is the data-structure of this message
type _CBusOptions struct {
	Connect bool
	Smart   bool
	Idmon   bool
	Exstat  bool
	Monitor bool
	Monall  bool
	Pun     bool
	Pcn     bool
	Srchk   bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusOptions) GetConnect() bool {
	return m.Connect
}

func (m *_CBusOptions) GetSmart() bool {
	return m.Smart
}

func (m *_CBusOptions) GetIdmon() bool {
	return m.Idmon
}

func (m *_CBusOptions) GetExstat() bool {
	return m.Exstat
}

func (m *_CBusOptions) GetMonitor() bool {
	return m.Monitor
}

func (m *_CBusOptions) GetMonall() bool {
	return m.Monall
}

func (m *_CBusOptions) GetPun() bool {
	return m.Pun
}

func (m *_CBusOptions) GetPcn() bool {
	return m.Pcn
}

func (m *_CBusOptions) GetSrchk() bool {
	return m.Srchk
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusOptions factory function for _CBusOptions
func NewCBusOptions(connect bool, smart bool, idmon bool, exstat bool, monitor bool, monall bool, pun bool, pcn bool, srchk bool) *_CBusOptions {
	return &_CBusOptions{Connect: connect, Smart: smart, Idmon: idmon, Exstat: exstat, Monitor: monitor, Monall: monall, Pun: pun, Pcn: pcn, Srchk: srchk}
}

// Deprecated: use the interface for direct cast
func CastCBusOptions(structType interface{}) CBusOptions {
	if casted, ok := structType.(CBusOptions); ok {
		return casted
	}
	if casted, ok := structType.(*CBusOptions); ok {
		return *casted
	}
	return nil
}

func (m *_CBusOptions) GetTypeName() string {
	return "CBusOptions"
}

func (m *_CBusOptions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusOptions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (connect)
	lengthInBits += 1

	// Simple field (smart)
	lengthInBits += 1

	// Simple field (idmon)
	lengthInBits += 1

	// Simple field (exstat)
	lengthInBits += 1

	// Simple field (monitor)
	lengthInBits += 1

	// Simple field (monall)
	lengthInBits += 1

	// Simple field (pun)
	lengthInBits += 1

	// Simple field (pcn)
	lengthInBits += 1

	// Simple field (srchk)
	lengthInBits += 1

	return lengthInBits
}

func (m *_CBusOptions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusOptionsParse(readBuffer utils.ReadBuffer) (CBusOptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (connect)
	_connect, _connectErr := readBuffer.ReadBit("connect")
	if _connectErr != nil {
		return nil, errors.Wrap(_connectErr, "Error parsing 'connect' field of CBusOptions")
	}
	connect := _connect

	// Simple Field (smart)
	_smart, _smartErr := readBuffer.ReadBit("smart")
	if _smartErr != nil {
		return nil, errors.Wrap(_smartErr, "Error parsing 'smart' field of CBusOptions")
	}
	smart := _smart

	// Simple Field (idmon)
	_idmon, _idmonErr := readBuffer.ReadBit("idmon")
	if _idmonErr != nil {
		return nil, errors.Wrap(_idmonErr, "Error parsing 'idmon' field of CBusOptions")
	}
	idmon := _idmon

	// Simple Field (exstat)
	_exstat, _exstatErr := readBuffer.ReadBit("exstat")
	if _exstatErr != nil {
		return nil, errors.Wrap(_exstatErr, "Error parsing 'exstat' field of CBusOptions")
	}
	exstat := _exstat

	// Simple Field (monitor)
	_monitor, _monitorErr := readBuffer.ReadBit("monitor")
	if _monitorErr != nil {
		return nil, errors.Wrap(_monitorErr, "Error parsing 'monitor' field of CBusOptions")
	}
	monitor := _monitor

	// Simple Field (monall)
	_monall, _monallErr := readBuffer.ReadBit("monall")
	if _monallErr != nil {
		return nil, errors.Wrap(_monallErr, "Error parsing 'monall' field of CBusOptions")
	}
	monall := _monall

	// Simple Field (pun)
	_pun, _punErr := readBuffer.ReadBit("pun")
	if _punErr != nil {
		return nil, errors.Wrap(_punErr, "Error parsing 'pun' field of CBusOptions")
	}
	pun := _pun

	// Simple Field (pcn)
	_pcn, _pcnErr := readBuffer.ReadBit("pcn")
	if _pcnErr != nil {
		return nil, errors.Wrap(_pcnErr, "Error parsing 'pcn' field of CBusOptions")
	}
	pcn := _pcn

	// Simple Field (srchk)
	_srchk, _srchkErr := readBuffer.ReadBit("srchk")
	if _srchkErr != nil {
		return nil, errors.Wrap(_srchkErr, "Error parsing 'srchk' field of CBusOptions")
	}
	srchk := _srchk

	if closeErr := readBuffer.CloseContext("CBusOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusOptions")
	}

	// Create the instance
	return &_CBusOptions{
		Connect: connect,
		Smart:   smart,
		Idmon:   idmon,
		Exstat:  exstat,
		Monitor: monitor,
		Monall:  monall,
		Pun:     pun,
		Pcn:     pcn,
		Srchk:   srchk,
	}, nil
}

func (m *_CBusOptions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusOptions")
	}

	// Simple Field (connect)
	connect := bool(m.GetConnect())
	_connectErr := writeBuffer.WriteBit("connect", (connect))
	if _connectErr != nil {
		return errors.Wrap(_connectErr, "Error serializing 'connect' field")
	}

	// Simple Field (smart)
	smart := bool(m.GetSmart())
	_smartErr := writeBuffer.WriteBit("smart", (smart))
	if _smartErr != nil {
		return errors.Wrap(_smartErr, "Error serializing 'smart' field")
	}

	// Simple Field (idmon)
	idmon := bool(m.GetIdmon())
	_idmonErr := writeBuffer.WriteBit("idmon", (idmon))
	if _idmonErr != nil {
		return errors.Wrap(_idmonErr, "Error serializing 'idmon' field")
	}

	// Simple Field (exstat)
	exstat := bool(m.GetExstat())
	_exstatErr := writeBuffer.WriteBit("exstat", (exstat))
	if _exstatErr != nil {
		return errors.Wrap(_exstatErr, "Error serializing 'exstat' field")
	}

	// Simple Field (monitor)
	monitor := bool(m.GetMonitor())
	_monitorErr := writeBuffer.WriteBit("monitor", (monitor))
	if _monitorErr != nil {
		return errors.Wrap(_monitorErr, "Error serializing 'monitor' field")
	}

	// Simple Field (monall)
	monall := bool(m.GetMonall())
	_monallErr := writeBuffer.WriteBit("monall", (monall))
	if _monallErr != nil {
		return errors.Wrap(_monallErr, "Error serializing 'monall' field")
	}

	// Simple Field (pun)
	pun := bool(m.GetPun())
	_punErr := writeBuffer.WriteBit("pun", (pun))
	if _punErr != nil {
		return errors.Wrap(_punErr, "Error serializing 'pun' field")
	}

	// Simple Field (pcn)
	pcn := bool(m.GetPcn())
	_pcnErr := writeBuffer.WriteBit("pcn", (pcn))
	if _pcnErr != nil {
		return errors.Wrap(_pcnErr, "Error serializing 'pcn' field")
	}

	// Simple Field (srchk)
	srchk := bool(m.GetSrchk())
	_srchkErr := writeBuffer.WriteBit("srchk", (srchk))
	if _srchkErr != nil {
		return errors.Wrap(_srchkErr, "Error serializing 'srchk' field")
	}

	if popErr := writeBuffer.PopContext("CBusOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusOptions")
	}
	return nil
}

func (m *_CBusOptions) isCBusOptions() bool {
	return true
}

func (m *_CBusOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
