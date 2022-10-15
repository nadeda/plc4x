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

// ApduDataExtFileStreamInfoReport is the corresponding interface of ApduDataExtFileStreamInfoReport
type ApduDataExtFileStreamInfoReport interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtFileStreamInfoReportExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtFileStreamInfoReport.
// This is useful for switch cases.
type ApduDataExtFileStreamInfoReportExactly interface {
	ApduDataExtFileStreamInfoReport
	isApduDataExtFileStreamInfoReport() bool
}

// _ApduDataExtFileStreamInfoReport is the data-structure of this message
type _ApduDataExtFileStreamInfoReport struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtFileStreamInfoReport) GetExtApciType() uint8 {
	return 0x30
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtFileStreamInfoReport) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtFileStreamInfoReport) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtFileStreamInfoReport factory function for _ApduDataExtFileStreamInfoReport
func NewApduDataExtFileStreamInfoReport(length uint8) *_ApduDataExtFileStreamInfoReport {
	_result := &_ApduDataExtFileStreamInfoReport{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtFileStreamInfoReport(structType interface{}) ApduDataExtFileStreamInfoReport {
	if casted, ok := structType.(ApduDataExtFileStreamInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtFileStreamInfoReport); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtFileStreamInfoReport) GetTypeName() string {
	return "ApduDataExtFileStreamInfoReport"
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtFileStreamInfoReportParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtFileStreamInfoReport, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtFileStreamInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtFileStreamInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtFileStreamInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtFileStreamInfoReport")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtFileStreamInfoReport{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtFileStreamInfoReport) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtFileStreamInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtFileStreamInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtFileStreamInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtFileStreamInfoReport")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtFileStreamInfoReport) isApduDataExtFileStreamInfoReport() bool {
	return true
}

func (m *_ApduDataExtFileStreamInfoReport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
