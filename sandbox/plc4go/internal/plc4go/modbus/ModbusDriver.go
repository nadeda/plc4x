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
package modbus

import (
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/pkg/plc4go"
)

type Driver struct {
	fieldHandler spi.PlcFieldHandler
	valueHandler spi.PlcValueHandler
	plc4go.PlcDriver
}

func NewModbusDriver() plc4go.PlcDriver {
	return Driver{
		fieldHandler: NewFieldHandler(),
		valueHandler: NewValueHandler(),
	}
}

func (m Driver) GetProtocolCode() string {
	return "modbus"
}

func (m Driver) GetProtocolName() string {
	return "Modbus"
}

func (m Driver) CheckQuery(query string) error {
	_, err := m.fieldHandler.ParseQuery(query)
	return err
}

func (m Driver) GetConnection(connectionString string) (plc4go.PlcConnection, error) {
	return NewConnection(m.fieldHandler, m.valueHandler), nil
}
