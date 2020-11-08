//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "io"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetServiceAckGetEventInformation struct {
    Parent *BACnetServiceAck
    IBACnetServiceAckGetEventInformation
}

// The corresponding interface
type IBACnetServiceAckGetEventInformation interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckGetEventInformation) ServiceChoice() uint8 {
    return 0x1D
}


func (m *BACnetServiceAckGetEventInformation) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckGetEventInformation() *BACnetServiceAck {
    child := &BACnetServiceAckGetEventInformation{
        Parent: NewBACnetServiceAck(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetServiceAckGetEventInformation(structType interface{}) BACnetServiceAckGetEventInformation {
    castFunc := func(typ interface{}) BACnetServiceAckGetEventInformation {
        if casted, ok := typ.(BACnetServiceAckGetEventInformation); ok {
            return casted
        }
        if casted, ok := typ.(*BACnetServiceAckGetEventInformation); ok {
            return *casted
        }
        if casted, ok := typ.(BACnetServiceAck); ok {
            return CastBACnetServiceAckGetEventInformation(casted.Child)
        }
        if casted, ok := typ.(*BACnetServiceAck); ok {
            return CastBACnetServiceAckGetEventInformation(casted.Child)
        }
        return BACnetServiceAckGetEventInformation{}
    }
    return castFunc(structType)
}

func (m *BACnetServiceAckGetEventInformation) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BACnetServiceAckGetEventInformation) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetServiceAckGetEventInformationParse(io *utils.ReadBuffer) (*BACnetServiceAck, error) {

    // Create a partially initialized instance
    _child := &BACnetServiceAckGetEventInformation{
        Parent: &BACnetServiceAck{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetServiceAckGetEventInformation) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetServiceAckGetEventInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *BACnetServiceAckGetEventInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

