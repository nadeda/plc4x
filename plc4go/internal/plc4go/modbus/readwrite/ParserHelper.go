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
    "errors"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/modbus/readwrite/model"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

type ModbusParserHelper struct {
}

func (m ModbusParserHelper) Parse(typeName string, arguments []string, io *utils.ReadBuffer) (spi.Message, error) {
    switch typeName {
    case "ModbusPDUWriteFileRecordRequestItem":
        return model.ModbusPDUWriteFileRecordRequestItemParse(io)
    case "ModbusPDUReadFileRecordResponseItem":
        return model.ModbusPDUReadFileRecordResponseItemParse(io)
    case "ModbusConstants":
        return model.ModbusConstantsParse(io)
    case "ModbusTcpADU":
        response, err := utils.StrToBool(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.ModbusTcpADUParse(io, response)
    case "ModbusPDUWriteFileRecordResponseItem":
        return model.ModbusPDUWriteFileRecordResponseItemParse(io)
    case "ModbusPDU":
        response, err := utils.StrToBool(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.ModbusPDUParse(io, response)
    case "ModbusPDUReadFileRecordRequestItem":
        return model.ModbusPDUReadFileRecordRequestItemParse(io)
    case "ModbusSerialADU":
        response, err := utils.StrToBool(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.ModbusSerialADUParse(io, response)
    }
    return nil, errors.New("Unsupported type " + typeName)
}
