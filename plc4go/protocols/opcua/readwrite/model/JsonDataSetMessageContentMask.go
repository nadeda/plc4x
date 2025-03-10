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

// JsonDataSetMessageContentMask is an enum
type JsonDataSetMessageContentMask uint32

type IJsonDataSetMessageContentMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskNone                    JsonDataSetMessageContentMask = 0
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterId         JsonDataSetMessageContentMask = 1
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMetaDataVersion         JsonDataSetMessageContentMask = 2
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskSequenceNumber          JsonDataSetMessageContentMask = 4
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskTimestamp               JsonDataSetMessageContentMask = 8
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskStatus                  JsonDataSetMessageContentMask = 16
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMessageType             JsonDataSetMessageContentMask = 32
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterName       JsonDataSetMessageContentMask = 64
	JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskReversibleFieldEncoding JsonDataSetMessageContentMask = 128
)

var JsonDataSetMessageContentMaskValues []JsonDataSetMessageContentMask

func init() {
	_ = errors.New
	JsonDataSetMessageContentMaskValues = []JsonDataSetMessageContentMask{
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskNone,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterId,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMetaDataVersion,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskSequenceNumber,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskTimestamp,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskStatus,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMessageType,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterName,
		JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskReversibleFieldEncoding,
	}
}

func JsonDataSetMessageContentMaskByValue(value uint32) (enum JsonDataSetMessageContentMask, ok bool) {
	switch value {
	case 0:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskNone, true
	case 1:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterId, true
	case 128:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskReversibleFieldEncoding, true
	case 16:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskStatus, true
	case 2:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMetaDataVersion, true
	case 32:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMessageType, true
	case 4:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskSequenceNumber, true
	case 64:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterName, true
	case 8:
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskTimestamp, true
	}
	return 0, false
}

func JsonDataSetMessageContentMaskByName(value string) (enum JsonDataSetMessageContentMask, ok bool) {
	switch value {
	case "jsonDataSetMessageContentMaskNone":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskNone, true
	case "jsonDataSetMessageContentMaskDataSetWriterId":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterId, true
	case "jsonDataSetMessageContentMaskReversibleFieldEncoding":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskReversibleFieldEncoding, true
	case "jsonDataSetMessageContentMaskStatus":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskStatus, true
	case "jsonDataSetMessageContentMaskMetaDataVersion":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMetaDataVersion, true
	case "jsonDataSetMessageContentMaskMessageType":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMessageType, true
	case "jsonDataSetMessageContentMaskSequenceNumber":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskSequenceNumber, true
	case "jsonDataSetMessageContentMaskDataSetWriterName":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterName, true
	case "jsonDataSetMessageContentMaskTimestamp":
		return JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskTimestamp, true
	}
	return 0, false
}

func JsonDataSetMessageContentMaskKnows(value uint32) bool {
	for _, typeValue := range JsonDataSetMessageContentMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastJsonDataSetMessageContentMask(structType any) JsonDataSetMessageContentMask {
	castFunc := func(typ any) JsonDataSetMessageContentMask {
		if sJsonDataSetMessageContentMask, ok := typ.(JsonDataSetMessageContentMask); ok {
			return sJsonDataSetMessageContentMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m JsonDataSetMessageContentMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m JsonDataSetMessageContentMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func JsonDataSetMessageContentMaskParse(ctx context.Context, theBytes []byte) (JsonDataSetMessageContentMask, error) {
	return JsonDataSetMessageContentMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func JsonDataSetMessageContentMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (JsonDataSetMessageContentMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("JsonDataSetMessageContentMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading JsonDataSetMessageContentMask")
	}
	if enum, ok := JsonDataSetMessageContentMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for JsonDataSetMessageContentMask")
		return JsonDataSetMessageContentMask(val), nil
	} else {
		return enum, nil
	}
}

func (e JsonDataSetMessageContentMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e JsonDataSetMessageContentMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("JsonDataSetMessageContentMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e JsonDataSetMessageContentMask) PLC4XEnumName() string {
	switch e {
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskNone:
		return "jsonDataSetMessageContentMaskNone"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterId:
		return "jsonDataSetMessageContentMaskDataSetWriterId"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskReversibleFieldEncoding:
		return "jsonDataSetMessageContentMaskReversibleFieldEncoding"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskStatus:
		return "jsonDataSetMessageContentMaskStatus"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMetaDataVersion:
		return "jsonDataSetMessageContentMaskMetaDataVersion"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskMessageType:
		return "jsonDataSetMessageContentMaskMessageType"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskSequenceNumber:
		return "jsonDataSetMessageContentMaskSequenceNumber"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskDataSetWriterName:
		return "jsonDataSetMessageContentMaskDataSetWriterName"
	case JsonDataSetMessageContentMask_jsonDataSetMessageContentMaskTimestamp:
		return "jsonDataSetMessageContentMaskTimestamp"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e JsonDataSetMessageContentMask) String() string {
	return e.PLC4XEnumName()
}
