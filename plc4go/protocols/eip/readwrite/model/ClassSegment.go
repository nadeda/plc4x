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

// ClassSegment is the corresponding interface of ClassSegment
type ClassSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPathSegmentType returns PathSegmentType (property field)
	GetPathSegmentType() uint8
	// GetLogicalSegmentType returns LogicalSegmentType (property field)
	GetLogicalSegmentType() uint8
	// GetLogicalSegmentFormat returns LogicalSegmentFormat (property field)
	GetLogicalSegmentFormat() uint8
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() uint8
}

// ClassSegmentExactly can be used when we want exactly this type and not a type which fulfills ClassSegment.
// This is useful for switch cases.
type ClassSegmentExactly interface {
	ClassSegment
	isClassSegment() bool
}

// _ClassSegment is the data-structure of this message
type _ClassSegment struct {
	PathSegmentType      uint8
	LogicalSegmentType   uint8
	LogicalSegmentFormat uint8
	ClassSegment         uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClassSegment) GetPathSegmentType() uint8 {
	return m.PathSegmentType
}

func (m *_ClassSegment) GetLogicalSegmentType() uint8 {
	return m.LogicalSegmentType
}

func (m *_ClassSegment) GetLogicalSegmentFormat() uint8 {
	return m.LogicalSegmentFormat
}

func (m *_ClassSegment) GetClassSegment() uint8 {
	return m.ClassSegment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewClassSegment factory function for _ClassSegment
func NewClassSegment(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, classSegment uint8) *_ClassSegment {
	return &_ClassSegment{PathSegmentType: pathSegmentType, LogicalSegmentType: logicalSegmentType, LogicalSegmentFormat: logicalSegmentFormat, ClassSegment: classSegment}
}

// Deprecated: use the interface for direct cast
func CastClassSegment(structType any) ClassSegment {
	if casted, ok := structType.(ClassSegment); ok {
		return casted
	}
	if casted, ok := structType.(*ClassSegment); ok {
		return *casted
	}
	return nil
}

func (m *_ClassSegment) GetTypeName() string {
	return "ClassSegment"
}

func (m *_ClassSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (pathSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentFormat)
	lengthInBits += 2

	// Simple field (classSegment)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClassSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClassSegmentParse(ctx context.Context, theBytes []byte) (ClassSegment, error) {
	return ClassSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClassSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClassSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClassSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClassSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pathSegmentType)
	_pathSegmentType, _pathSegmentTypeErr := readBuffer.ReadUint8("pathSegmentType", 3)
	if _pathSegmentTypeErr != nil {
		return nil, errors.Wrap(_pathSegmentTypeErr, "Error parsing 'pathSegmentType' field of ClassSegment")
	}
	pathSegmentType := _pathSegmentType

	// Simple Field (logicalSegmentType)
	_logicalSegmentType, _logicalSegmentTypeErr := readBuffer.ReadUint8("logicalSegmentType", 3)
	if _logicalSegmentTypeErr != nil {
		return nil, errors.Wrap(_logicalSegmentTypeErr, "Error parsing 'logicalSegmentType' field of ClassSegment")
	}
	logicalSegmentType := _logicalSegmentType

	// Simple Field (logicalSegmentFormat)
	_logicalSegmentFormat, _logicalSegmentFormatErr := readBuffer.ReadUint8("logicalSegmentFormat", 2)
	if _logicalSegmentFormatErr != nil {
		return nil, errors.Wrap(_logicalSegmentFormatErr, "Error parsing 'logicalSegmentFormat' field of ClassSegment")
	}
	logicalSegmentFormat := _logicalSegmentFormat

	// Simple Field (classSegment)
	_classSegment, _classSegmentErr := readBuffer.ReadUint8("classSegment", 8)
	if _classSegmentErr != nil {
		return nil, errors.Wrap(_classSegmentErr, "Error parsing 'classSegment' field of ClassSegment")
	}
	classSegment := _classSegment

	if closeErr := readBuffer.CloseContext("ClassSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClassSegment")
	}

	// Create the instance
	return &_ClassSegment{
		PathSegmentType:      pathSegmentType,
		LogicalSegmentType:   logicalSegmentType,
		LogicalSegmentFormat: logicalSegmentFormat,
		ClassSegment:         classSegment,
	}, nil
}

func (m *_ClassSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClassSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ClassSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ClassSegment")
	}

	// Simple Field (pathSegmentType)
	pathSegmentType := uint8(m.GetPathSegmentType())
	_pathSegmentTypeErr := writeBuffer.WriteUint8("pathSegmentType", 3, (pathSegmentType))
	if _pathSegmentTypeErr != nil {
		return errors.Wrap(_pathSegmentTypeErr, "Error serializing 'pathSegmentType' field")
	}

	// Simple Field (logicalSegmentType)
	logicalSegmentType := uint8(m.GetLogicalSegmentType())
	_logicalSegmentTypeErr := writeBuffer.WriteUint8("logicalSegmentType", 3, (logicalSegmentType))
	if _logicalSegmentTypeErr != nil {
		return errors.Wrap(_logicalSegmentTypeErr, "Error serializing 'logicalSegmentType' field")
	}

	// Simple Field (logicalSegmentFormat)
	logicalSegmentFormat := uint8(m.GetLogicalSegmentFormat())
	_logicalSegmentFormatErr := writeBuffer.WriteUint8("logicalSegmentFormat", 2, (logicalSegmentFormat))
	if _logicalSegmentFormatErr != nil {
		return errors.Wrap(_logicalSegmentFormatErr, "Error serializing 'logicalSegmentFormat' field")
	}

	// Simple Field (classSegment)
	classSegment := uint8(m.GetClassSegment())
	_classSegmentErr := writeBuffer.WriteUint8("classSegment", 8, (classSegment))
	if _classSegmentErr != nil {
		return errors.Wrap(_classSegmentErr, "Error serializing 'classSegment' field")
	}

	if popErr := writeBuffer.PopContext("ClassSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ClassSegment")
	}
	return nil
}

func (m *_ClassSegment) isClassSegment() bool {
	return true
}

func (m *_ClassSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
