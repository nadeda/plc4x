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

// Code generated by mockery v2.32.4. DO NOT EDIT.

package _default

import (
	context "context"

	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockExpectation is an autogenerated mock type for the Expectation type
type MockExpectation struct {
	mock.Mock
}

type MockExpectation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExpectation) EXPECT() *MockExpectation_Expecter {
	return &MockExpectation_Expecter{mock: &_m.Mock}
}

// GetAcceptsMessage provides a mock function with given fields:
func (_m *MockExpectation) GetAcceptsMessage() spi.AcceptsMessage {
	ret := _m.Called()

	var r0 spi.AcceptsMessage
	if rf, ok := ret.Get(0).(func() spi.AcceptsMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.AcceptsMessage)
		}
	}

	return r0
}

// MockExpectation_GetAcceptsMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAcceptsMessage'
type MockExpectation_GetAcceptsMessage_Call struct {
	*mock.Call
}

// GetAcceptsMessage is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetAcceptsMessage() *MockExpectation_GetAcceptsMessage_Call {
	return &MockExpectation_GetAcceptsMessage_Call{Call: _e.mock.On("GetAcceptsMessage")}
}

func (_c *MockExpectation_GetAcceptsMessage_Call) Run(run func()) *MockExpectation_GetAcceptsMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetAcceptsMessage_Call) Return(_a0 spi.AcceptsMessage) *MockExpectation_GetAcceptsMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetAcceptsMessage_Call) RunAndReturn(run func() spi.AcceptsMessage) *MockExpectation_GetAcceptsMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockExpectation) GetContext() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockExpectation_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockExpectation_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetContext() *MockExpectation_GetContext_Call {
	return &MockExpectation_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockExpectation_GetContext_Call) Run(run func()) *MockExpectation_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetContext_Call) Return(_a0 context.Context) *MockExpectation_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetContext_Call) RunAndReturn(run func() context.Context) *MockExpectation_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetCreationTime provides a mock function with given fields:
func (_m *MockExpectation) GetCreationTime() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// MockExpectation_GetCreationTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreationTime'
type MockExpectation_GetCreationTime_Call struct {
	*mock.Call
}

// GetCreationTime is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetCreationTime() *MockExpectation_GetCreationTime_Call {
	return &MockExpectation_GetCreationTime_Call{Call: _e.mock.On("GetCreationTime")}
}

func (_c *MockExpectation_GetCreationTime_Call) Run(run func()) *MockExpectation_GetCreationTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetCreationTime_Call) Return(_a0 time.Time) *MockExpectation_GetCreationTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetCreationTime_Call) RunAndReturn(run func() time.Time) *MockExpectation_GetCreationTime_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpiration provides a mock function with given fields:
func (_m *MockExpectation) GetExpiration() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// MockExpectation_GetExpiration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpiration'
type MockExpectation_GetExpiration_Call struct {
	*mock.Call
}

// GetExpiration is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetExpiration() *MockExpectation_GetExpiration_Call {
	return &MockExpectation_GetExpiration_Call{Call: _e.mock.On("GetExpiration")}
}

func (_c *MockExpectation_GetExpiration_Call) Run(run func()) *MockExpectation_GetExpiration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetExpiration_Call) Return(_a0 time.Time) *MockExpectation_GetExpiration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetExpiration_Call) RunAndReturn(run func() time.Time) *MockExpectation_GetExpiration_Call {
	_c.Call.Return(run)
	return _c
}

// GetHandleError provides a mock function with given fields:
func (_m *MockExpectation) GetHandleError() spi.HandleError {
	ret := _m.Called()

	var r0 spi.HandleError
	if rf, ok := ret.Get(0).(func() spi.HandleError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.HandleError)
		}
	}

	return r0
}

// MockExpectation_GetHandleError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHandleError'
type MockExpectation_GetHandleError_Call struct {
	*mock.Call
}

// GetHandleError is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetHandleError() *MockExpectation_GetHandleError_Call {
	return &MockExpectation_GetHandleError_Call{Call: _e.mock.On("GetHandleError")}
}

func (_c *MockExpectation_GetHandleError_Call) Run(run func()) *MockExpectation_GetHandleError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetHandleError_Call) Return(_a0 spi.HandleError) *MockExpectation_GetHandleError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetHandleError_Call) RunAndReturn(run func() spi.HandleError) *MockExpectation_GetHandleError_Call {
	_c.Call.Return(run)
	return _c
}

// GetHandleMessage provides a mock function with given fields:
func (_m *MockExpectation) GetHandleMessage() spi.HandleMessage {
	ret := _m.Called()

	var r0 spi.HandleMessage
	if rf, ok := ret.Get(0).(func() spi.HandleMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.HandleMessage)
		}
	}

	return r0
}

// MockExpectation_GetHandleMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHandleMessage'
type MockExpectation_GetHandleMessage_Call struct {
	*mock.Call
}

// GetHandleMessage is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) GetHandleMessage() *MockExpectation_GetHandleMessage_Call {
	return &MockExpectation_GetHandleMessage_Call{Call: _e.mock.On("GetHandleMessage")}
}

func (_c *MockExpectation_GetHandleMessage_Call) Run(run func()) *MockExpectation_GetHandleMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_GetHandleMessage_Call) Return(_a0 spi.HandleMessage) *MockExpectation_GetHandleMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_GetHandleMessage_Call) RunAndReturn(run func() spi.HandleMessage) *MockExpectation_GetHandleMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockExpectation) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockExpectation_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockExpectation_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockExpectation_Expecter) String() *MockExpectation_String_Call {
	return &MockExpectation_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockExpectation_String_Call) Run(run func()) *MockExpectation_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExpectation_String_Call) Return(_a0 string) *MockExpectation_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExpectation_String_Call) RunAndReturn(run func() string) *MockExpectation_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExpectation creates a new instance of MockExpectation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExpectation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExpectation {
	mock := &MockExpectation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
