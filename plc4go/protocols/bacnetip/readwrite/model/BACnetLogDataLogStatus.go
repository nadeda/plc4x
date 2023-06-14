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

// BACnetLogDataLogStatus is the corresponding interface of BACnetLogDataLogStatus
type BACnetLogDataLogStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogData
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
}

// BACnetLogDataLogStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogStatus.
// This is useful for switch cases.
type BACnetLogDataLogStatusExactly interface {
	BACnetLogDataLogStatus
	isBACnetLogDataLogStatus() bool
}

// _BACnetLogDataLogStatus is the data-structure of this message
type _BACnetLogDataLogStatus struct {
	*_BACnetLogData
	LogStatus BACnetLogStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogStatus) InitializeParent(parent BACnetLogData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogDataLogStatus) GetParent() BACnetLogData {
	return m._BACnetLogData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogStatus factory function for _BACnetLogDataLogStatus
func NewBACnetLogDataLogStatus(logStatus BACnetLogStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogDataLogStatus {
	_result := &_BACnetLogDataLogStatus{
		LogStatus:      logStatus,
		_BACnetLogData: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogData._BACnetLogDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogStatus(structType any) BACnetLogDataLogStatus {
	if casted, ok := structType.(BACnetLogDataLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogStatus) GetTypeName() string {
	return "BACnetLogDataLogStatus"
}

func (m *_BACnetLogDataLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogStatusParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLogDataLogStatus, error) {
	return BACnetLogDataLogStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogDataLogStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogDataLogStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logStatus)
	if pullErr := readBuffer.PullContext("logStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logStatus")
	}
	_logStatus, _logStatusErr := BACnetLogStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _logStatusErr != nil {
		return nil, errors.Wrap(_logStatusErr, "Error parsing 'logStatus' field of BACnetLogDataLogStatus")
	}
	logStatus := _logStatus.(BACnetLogStatusTagged)
	if closeErr := readBuffer.CloseContext("logStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogStatus{
		_BACnetLogData: &_BACnetLogData{
			TagNumber: tagNumber,
		},
		LogStatus: logStatus,
	}
	_child._BACnetLogData._BACnetLogDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogStatus")
		}

		// Simple Field (logStatus)
		if pushErr := writeBuffer.PushContext("logStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logStatus")
		}
		_logStatusErr := writeBuffer.WriteSerializable(ctx, m.GetLogStatus())
		if popErr := writeBuffer.PopContext("logStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logStatus")
		}
		if _logStatusErr != nil {
			return errors.Wrap(_logStatusErr, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogStatus) isBACnetLogDataLogStatus() bool {
	return true
}

func (m *_BACnetLogDataLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
