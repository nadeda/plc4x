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

// CommandSpecificDataItem is the corresponding interface of CommandSpecificDataItem
type CommandSpecificDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint16
}

// CommandSpecificDataItemExactly can be used when we want exactly this type and not a type which fulfills CommandSpecificDataItem.
// This is useful for switch cases.
type CommandSpecificDataItemExactly interface {
	CommandSpecificDataItem
	isCommandSpecificDataItem() bool
}

// _CommandSpecificDataItem is the data-structure of this message
type _CommandSpecificDataItem struct {
	_CommandSpecificDataItemChildRequirements
}

type _CommandSpecificDataItemChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetItemType() uint16
}

type CommandSpecificDataItemParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CommandSpecificDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type CommandSpecificDataItemChild interface {
	utils.Serializable
	InitializeParent(parent CommandSpecificDataItem)
	GetParent() *CommandSpecificDataItem

	GetTypeName() string
	CommandSpecificDataItem
}

// NewCommandSpecificDataItem factory function for _CommandSpecificDataItem
func NewCommandSpecificDataItem() *_CommandSpecificDataItem {
	return &_CommandSpecificDataItem{}
}

// Deprecated: use the interface for direct cast
func CastCommandSpecificDataItem(structType any) CommandSpecificDataItem {
	if casted, ok := structType.(CommandSpecificDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*CommandSpecificDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_CommandSpecificDataItem) GetTypeName() string {
	return "CommandSpecificDataItem"
}

func (m *_CommandSpecificDataItem) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CommandSpecificDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CommandSpecificDataItemParse(ctx context.Context, theBytes []byte) (CommandSpecificDataItem, error) {
	return CommandSpecificDataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CommandSpecificDataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CommandSpecificDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CommandSpecificDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CommandSpecificDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint16("itemType", 16)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field of CommandSpecificDataItem")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CommandSpecificDataItemChildSerializeRequirement interface {
		CommandSpecificDataItem
		InitializeParent(CommandSpecificDataItem)
		GetParent() CommandSpecificDataItem
	}
	var _childTemp any
	var _child CommandSpecificDataItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case itemType == 0x000C: // CipIdentity
		_childTemp, typeSwitchError = CipIdentityParseWithBuffer(ctx, readBuffer)
	case itemType == 0x0086: // CipSecurityInformation
		_childTemp, typeSwitchError = CipSecurityInformationParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [itemType=%v]", itemType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CommandSpecificDataItem")
	}
	_child = _childTemp.(CommandSpecificDataItemChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CommandSpecificDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CommandSpecificDataItem")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_CommandSpecificDataItem) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CommandSpecificDataItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CommandSpecificDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CommandSpecificDataItem")
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint16(child.GetItemType())
	_itemTypeErr := writeBuffer.WriteUint16("itemType", 16, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CommandSpecificDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CommandSpecificDataItem")
	}
	return nil
}

func (m *_CommandSpecificDataItem) isCommandSpecificDataItem() bool {
	return true
}

func (m *_CommandSpecificDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
