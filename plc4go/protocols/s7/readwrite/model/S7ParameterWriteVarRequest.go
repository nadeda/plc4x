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

// S7ParameterWriteVarRequest is the corresponding interface of S7ParameterWriteVarRequest
type S7ParameterWriteVarRequest interface {
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7VarRequestParameterItem
}

// S7ParameterWriteVarRequestExactly can be used when we want exactly this type and not a type which fulfills S7ParameterWriteVarRequest.
// This is useful for switch cases.
type S7ParameterWriteVarRequestExactly interface {
	S7ParameterWriteVarRequest
	isS7ParameterWriteVarRequest() bool
}

// _S7ParameterWriteVarRequest is the data-structure of this message
type _S7ParameterWriteVarRequest struct {
	*_S7Parameter
	Items []S7VarRequestParameterItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetParameterType() uint8 {
	return 0x05
}

func (m *_S7ParameterWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterWriteVarRequest) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterWriteVarRequest) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetItems() []S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterWriteVarRequest factory function for _S7ParameterWriteVarRequest
func NewS7ParameterWriteVarRequest(items []S7VarRequestParameterItem) *_S7ParameterWriteVarRequest {
	_result := &_S7ParameterWriteVarRequest{
		Items:        items,
		_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterWriteVarRequest(structType interface{}) S7ParameterWriteVarRequest {
	if casted, ok := structType.(S7ParameterWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterWriteVarRequestParse(readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterWriteVarRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterWriteVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (numItems) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field of S7ParameterWriteVarRequest")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]S7VarRequestParameterItem, numItems)
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(numItems); curItem++ {
			_item, _err := S7VarRequestParameterItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of S7ParameterWriteVarRequest")
			}
			items[curItem] = _item.(S7VarRequestParameterItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterWriteVarRequest")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterWriteVarRequest{
		_S7Parameter: &_S7Parameter{},
		Items:        items,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterWriteVarRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterWriteVarRequest")
		}

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.GetItems())))
		_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _, _element := range m.GetItems() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterWriteVarRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7ParameterWriteVarRequest) isS7ParameterWriteVarRequest() bool {
	return true
}

func (m *_S7ParameterWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
