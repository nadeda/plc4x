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

// BACnetConstructedDataTimerRunning is the corresponding interface of BACnetConstructedDataTimerRunning
type BACnetConstructedDataTimerRunning interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimerRunning returns TimerRunning (property field)
	GetTimerRunning() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataTimerRunningExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimerRunning.
// This is useful for switch cases.
type BACnetConstructedDataTimerRunningExactly interface {
	BACnetConstructedDataTimerRunning
	isBACnetConstructedDataTimerRunning() bool
}

// _BACnetConstructedDataTimerRunning is the data-structure of this message
type _BACnetConstructedDataTimerRunning struct {
	*_BACnetConstructedData
	TimerRunning BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerRunning) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimerRunning) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIMER_RUNNING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerRunning) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerRunning) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerRunning) GetTimerRunning() BACnetApplicationTagBoolean {
	return m.TimerRunning
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerRunning) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetTimerRunning())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerRunning factory function for _BACnetConstructedDataTimerRunning
func NewBACnetConstructedDataTimerRunning(timerRunning BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerRunning {
	_result := &_BACnetConstructedDataTimerRunning{
		TimerRunning:           timerRunning,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerRunning(structType interface{}) BACnetConstructedDataTimerRunning {
	if casted, ok := structType.(BACnetConstructedDataTimerRunning); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerRunning); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerRunning) GetTypeName() string {
	return "BACnetConstructedDataTimerRunning"
}

func (m *_BACnetConstructedDataTimerRunning) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimerRunning) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timerRunning)
	lengthInBits += m.TimerRunning.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerRunning) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerRunningParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerRunning, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerRunning"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerRunning")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerRunning)
	if pullErr := readBuffer.PullContext("timerRunning"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timerRunning")
	}
	_timerRunning, _timerRunningErr := BACnetApplicationTagParse(readBuffer)
	if _timerRunningErr != nil {
		return nil, errors.Wrap(_timerRunningErr, "Error parsing 'timerRunning' field of BACnetConstructedDataTimerRunning")
	}
	timerRunning := _timerRunning.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("timerRunning"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timerRunning")
	}

	// Virtual field
	_actualValue := timerRunning
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerRunning"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerRunning")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerRunning{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimerRunning: timerRunning,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerRunning) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerRunning"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerRunning")
		}

		// Simple Field (timerRunning)
		if pushErr := writeBuffer.PushContext("timerRunning"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timerRunning")
		}
		_timerRunningErr := writeBuffer.WriteSerializable(m.GetTimerRunning())
		if popErr := writeBuffer.PopContext("timerRunning"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timerRunning")
		}
		if _timerRunningErr != nil {
			return errors.Wrap(_timerRunningErr, "Error serializing 'timerRunning' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerRunning"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerRunning")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerRunning) isBACnetConstructedDataTimerRunning() bool {
	return true
}

func (m *_BACnetConstructedDataTimerRunning) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
