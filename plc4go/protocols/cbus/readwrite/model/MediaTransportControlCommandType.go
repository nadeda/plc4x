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

// MediaTransportControlCommandType is an enum
type MediaTransportControlCommandType uint8

type IMediaTransportControlCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	MediaTransportControlCommandType_STOP                                   MediaTransportControlCommandType = 0x00
	MediaTransportControlCommandType_PLAY                                   MediaTransportControlCommandType = 0x01
	MediaTransportControlCommandType_PAUSE_RESUME                           MediaTransportControlCommandType = 0x02
	MediaTransportControlCommandType_SELECT_CATEGORY                        MediaTransportControlCommandType = 0x03
	MediaTransportControlCommandType_SELECT_SELECTION                       MediaTransportControlCommandType = 0x04
	MediaTransportControlCommandType_SELECT_TRACK                           MediaTransportControlCommandType = 0x05
	MediaTransportControlCommandType_SHUFFLE_ON_OFF                         MediaTransportControlCommandType = 0x06
	MediaTransportControlCommandType_REPEAT_ON_OFF                          MediaTransportControlCommandType = 0x07
	MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY                 MediaTransportControlCommandType = 0x08
	MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION                MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_NEXT_PREVIOUS_TRACK                    MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_FAST_FORWARD                           MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_REWIND                                 MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_SOURCE_POWER_CONTROL                   MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_TOTAL_TRACKS                           MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_STATUS_REQUEST                         MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_ENUMERATE_CATEGORIES_SELECTIONS_TRACKS MediaTransportControlCommandType = 0x09
	MediaTransportControlCommandType_ENUMERATION_SIZE                       MediaTransportControlCommandType = 0x0A
	MediaTransportControlCommandType_TRACK_NAME                             MediaTransportControlCommandType = 0x0B
	MediaTransportControlCommandType_SELECTION_NAME                         MediaTransportControlCommandType = 0x0C
	MediaTransportControlCommandType_CATEGORY_NAME                          MediaTransportControlCommandType = 0x0D
)

var MediaTransportControlCommandTypeValues []MediaTransportControlCommandType

func init() {
	_ = errors.New
	MediaTransportControlCommandTypeValues = []MediaTransportControlCommandType{
		MediaTransportControlCommandType_STOP,
		MediaTransportControlCommandType_PLAY,
		MediaTransportControlCommandType_PAUSE_RESUME,
		MediaTransportControlCommandType_SELECT_CATEGORY,
		MediaTransportControlCommandType_SELECT_SELECTION,
		MediaTransportControlCommandType_SELECT_TRACK,
		MediaTransportControlCommandType_SHUFFLE_ON_OFF,
		MediaTransportControlCommandType_REPEAT_ON_OFF,
		MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY,
		MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION,
		MediaTransportControlCommandType_NEXT_PREVIOUS_TRACK,
		MediaTransportControlCommandType_FAST_FORWARD,
		MediaTransportControlCommandType_REWIND,
		MediaTransportControlCommandType_SOURCE_POWER_CONTROL,
		MediaTransportControlCommandType_TOTAL_TRACKS,
		MediaTransportControlCommandType_STATUS_REQUEST,
		MediaTransportControlCommandType_ENUMERATE_CATEGORIES_SELECTIONS_TRACKS,
		MediaTransportControlCommandType_ENUMERATION_SIZE,
		MediaTransportControlCommandType_TRACK_NAME,
		MediaTransportControlCommandType_SELECTION_NAME,
		MediaTransportControlCommandType_CATEGORY_NAME,
	}
}

func (e MediaTransportControlCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0
		}
	case 0x01:
		{ /* '0x01' */
			return 0
		}
	case 0x02:
		{ /* '0x02' */
			return 1
		}
	case 0x03:
		{ /* '0x03' */
			return 1
		}
	case 0x04:
		{ /* '0x04' */
			return 2
		}
	case 0x05:
		{ /* '0x05' */
			return 4
		}
	case 0x06:
		{ /* '0x06' */
			return 1
		}
	case 0x07:
		{ /* '0x07' */
			return 1
		}
	case 0x08:
		{ /* '0x08' */
			return 1
		}
	case 0x09:
		{ /* '0x09' */
			return 1
		}
	case 0x0A:
		{ /* '0x0A' */
			return 3
		}
	case 0x0B:
		{ /* '0x0B' */
			return 1
		}
	case 0x0C:
		{ /* '0x0C' */
			return 1
		}
	case 0x0D:
		{ /* '0x0D' */
			return 1
		}
	default:
		{
			return 0
		}
	}
}

func MediaTransportControlCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (MediaTransportControlCommandType, error) {
	for _, sizeValue := range MediaTransportControlCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func MediaTransportControlCommandTypeByValue(value uint8) (enum MediaTransportControlCommandType, ok bool) {
	switch value {
	case 0x00:
		return MediaTransportControlCommandType_STOP, true
	case 0x01:
		return MediaTransportControlCommandType_PLAY, true
	case 0x02:
		return MediaTransportControlCommandType_PAUSE_RESUME, true
	case 0x03:
		return MediaTransportControlCommandType_SELECT_CATEGORY, true
	case 0x04:
		return MediaTransportControlCommandType_SELECT_SELECTION, true
	case 0x05:
		return MediaTransportControlCommandType_SELECT_TRACK, true
	case 0x06:
		return MediaTransportControlCommandType_SHUFFLE_ON_OFF, true
	case 0x07:
		return MediaTransportControlCommandType_REPEAT_ON_OFF, true
	case 0x08:
		return MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY, true
	case 0x09:
		return MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION, true
	case 0x0A:
		return MediaTransportControlCommandType_ENUMERATION_SIZE, true
	case 0x0B:
		return MediaTransportControlCommandType_TRACK_NAME, true
	case 0x0C:
		return MediaTransportControlCommandType_SELECTION_NAME, true
	case 0x0D:
		return MediaTransportControlCommandType_CATEGORY_NAME, true
	}
	return 0, false
}

func MediaTransportControlCommandTypeByName(value string) (enum MediaTransportControlCommandType, ok bool) {
	switch value {
	case "STOP":
		return MediaTransportControlCommandType_STOP, true
	case "PLAY":
		return MediaTransportControlCommandType_PLAY, true
	case "PAUSE_RESUME":
		return MediaTransportControlCommandType_PAUSE_RESUME, true
	case "SELECT_CATEGORY":
		return MediaTransportControlCommandType_SELECT_CATEGORY, true
	case "SELECT_SELECTION":
		return MediaTransportControlCommandType_SELECT_SELECTION, true
	case "SELECT_TRACK":
		return MediaTransportControlCommandType_SELECT_TRACK, true
	case "SHUFFLE_ON_OFF":
		return MediaTransportControlCommandType_SHUFFLE_ON_OFF, true
	case "REPEAT_ON_OFF":
		return MediaTransportControlCommandType_REPEAT_ON_OFF, true
	case "NEXT_PREVIOUS_CATEGORY":
		return MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY, true
	case "NEXT_PREVIOUS_SELECTION":
		return MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION, true
	case "ENUMERATION_SIZE":
		return MediaTransportControlCommandType_ENUMERATION_SIZE, true
	case "TRACK_NAME":
		return MediaTransportControlCommandType_TRACK_NAME, true
	case "SELECTION_NAME":
		return MediaTransportControlCommandType_SELECTION_NAME, true
	case "CATEGORY_NAME":
		return MediaTransportControlCommandType_CATEGORY_NAME, true
	}
	return 0, false
}

func MediaTransportControlCommandTypeKnows(value uint8) bool {
	for _, typeValue := range MediaTransportControlCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMediaTransportControlCommandType(structType any) MediaTransportControlCommandType {
	castFunc := func(typ any) MediaTransportControlCommandType {
		if sMediaTransportControlCommandType, ok := typ.(MediaTransportControlCommandType); ok {
			return sMediaTransportControlCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m MediaTransportControlCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m MediaTransportControlCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlCommandTypeParse(ctx context.Context, theBytes []byte) (MediaTransportControlCommandType, error) {
	return MediaTransportControlCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlCommandType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("MediaTransportControlCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MediaTransportControlCommandType")
	}
	if enum, ok := MediaTransportControlCommandTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for MediaTransportControlCommandType")
		return MediaTransportControlCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e MediaTransportControlCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MediaTransportControlCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("MediaTransportControlCommandType", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MediaTransportControlCommandType) PLC4XEnumName() string {
	switch e {
	case MediaTransportControlCommandType_STOP:
		return "STOP"
	case MediaTransportControlCommandType_PLAY:
		return "PLAY"
	case MediaTransportControlCommandType_PAUSE_RESUME:
		return "PAUSE_RESUME"
	case MediaTransportControlCommandType_SELECT_CATEGORY:
		return "SELECT_CATEGORY"
	case MediaTransportControlCommandType_SELECT_SELECTION:
		return "SELECT_SELECTION"
	case MediaTransportControlCommandType_SELECT_TRACK:
		return "SELECT_TRACK"
	case MediaTransportControlCommandType_SHUFFLE_ON_OFF:
		return "SHUFFLE_ON_OFF"
	case MediaTransportControlCommandType_REPEAT_ON_OFF:
		return "REPEAT_ON_OFF"
	case MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY:
		return "NEXT_PREVIOUS_CATEGORY"
	case MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION:
		return "NEXT_PREVIOUS_SELECTION"
	case MediaTransportControlCommandType_ENUMERATION_SIZE:
		return "ENUMERATION_SIZE"
	case MediaTransportControlCommandType_TRACK_NAME:
		return "TRACK_NAME"
	case MediaTransportControlCommandType_SELECTION_NAME:
		return "SELECTION_NAME"
	case MediaTransportControlCommandType_CATEGORY_NAME:
		return "CATEGORY_NAME"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e MediaTransportControlCommandType) String() string {
	return e.PLC4XEnumName()
}
