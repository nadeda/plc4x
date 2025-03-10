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

// ReadValueId is the corresponding interface of ReadValueId
type ReadValueId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetDataEncoding returns DataEncoding (property field)
	GetDataEncoding() QualifiedName
}

// ReadValueIdExactly can be used when we want exactly this type and not a type which fulfills ReadValueId.
// This is useful for switch cases.
type ReadValueIdExactly interface {
	ReadValueId
	isReadValueId() bool
}

// _ReadValueId is the data-structure of this message
type _ReadValueId struct {
	*_ExtensionObjectDefinition
	NodeId       NodeId
	AttributeId  uint32
	IndexRange   PascalString
	DataEncoding QualifiedName
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadValueId) GetIdentifier() string {
	return "628"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadValueId) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ReadValueId) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadValueId) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_ReadValueId) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_ReadValueId) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_ReadValueId) GetDataEncoding() QualifiedName {
	return m.DataEncoding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReadValueId factory function for _ReadValueId
func NewReadValueId(nodeId NodeId, attributeId uint32, indexRange PascalString, dataEncoding QualifiedName) *_ReadValueId {
	_result := &_ReadValueId{
		NodeId:                     nodeId,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
		DataEncoding:               dataEncoding,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReadValueId(structType any) ReadValueId {
	if casted, ok := structType.(ReadValueId); ok {
		return casted
	}
	if casted, ok := structType.(*ReadValueId); ok {
		return *casted
	}
	return nil
}

func (m *_ReadValueId) GetTypeName() string {
	return "ReadValueId"
}

func (m *_ReadValueId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (dataEncoding)
	lengthInBits += m.DataEncoding.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReadValueId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReadValueIdParse(ctx context.Context, theBytes []byte, identifier string) (ReadValueId, error) {
	return ReadValueIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ReadValueIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ReadValueId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReadValueId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadValueId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of ReadValueId")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (attributeId)
	_attributeId, _attributeIdErr := readBuffer.ReadUint32("attributeId", 32)
	if _attributeIdErr != nil {
		return nil, errors.Wrap(_attributeIdErr, "Error parsing 'attributeId' field of ReadValueId")
	}
	attributeId := _attributeId

	// Simple Field (indexRange)
	if pullErr := readBuffer.PullContext("indexRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for indexRange")
	}
	_indexRange, _indexRangeErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _indexRangeErr != nil {
		return nil, errors.Wrap(_indexRangeErr, "Error parsing 'indexRange' field of ReadValueId")
	}
	indexRange := _indexRange.(PascalString)
	if closeErr := readBuffer.CloseContext("indexRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for indexRange")
	}

	// Simple Field (dataEncoding)
	if pullErr := readBuffer.PullContext("dataEncoding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataEncoding")
	}
	_dataEncoding, _dataEncodingErr := QualifiedNameParseWithBuffer(ctx, readBuffer)
	if _dataEncodingErr != nil {
		return nil, errors.Wrap(_dataEncodingErr, "Error parsing 'dataEncoding' field of ReadValueId")
	}
	dataEncoding := _dataEncoding.(QualifiedName)
	if closeErr := readBuffer.CloseContext("dataEncoding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataEncoding")
	}

	if closeErr := readBuffer.CloseContext("ReadValueId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadValueId")
	}

	// Create a partially initialized instance
	_child := &_ReadValueId{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
		DataEncoding:               dataEncoding,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ReadValueId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadValueId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadValueId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadValueId")
		}

		// Simple Field (nodeId)
		if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeId")
		}
		_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
		if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeId")
		}
		if _nodeIdErr != nil {
			return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
		}

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (indexRange)
		if pushErr := writeBuffer.PushContext("indexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for indexRange")
		}
		_indexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetIndexRange())
		if popErr := writeBuffer.PopContext("indexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for indexRange")
		}
		if _indexRangeErr != nil {
			return errors.Wrap(_indexRangeErr, "Error serializing 'indexRange' field")
		}

		// Simple Field (dataEncoding)
		if pushErr := writeBuffer.PushContext("dataEncoding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataEncoding")
		}
		_dataEncodingErr := writeBuffer.WriteSerializable(ctx, m.GetDataEncoding())
		if popErr := writeBuffer.PopContext("dataEncoding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataEncoding")
		}
		if _dataEncodingErr != nil {
			return errors.Wrap(_dataEncodingErr, "Error serializing 'dataEncoding' field")
		}

		if popErr := writeBuffer.PopContext("ReadValueId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadValueId")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadValueId) isReadValueId() bool {
	return true
}

func (m *_ReadValueId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
