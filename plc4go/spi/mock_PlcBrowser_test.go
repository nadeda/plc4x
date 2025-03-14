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

package spi

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcBrowser is an autogenerated mock type for the PlcBrowser type
type MockPlcBrowser struct {
	mock.Mock
}

type MockPlcBrowser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcBrowser) EXPECT() *MockPlcBrowser_Expecter {
	return &MockPlcBrowser_Expecter{mock: &_m.Mock}
}

// Browse provides a mock function with given fields: ctx, browseRequest
func (_m *MockPlcBrowser) Browse(ctx context.Context, browseRequest model.PlcBrowseRequest) <-chan model.PlcBrowseRequestResult {
	ret := _m.Called(ctx, browseRequest)

	var r0 <-chan model.PlcBrowseRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcBrowseRequest) <-chan model.PlcBrowseRequestResult); ok {
		r0 = rf(ctx, browseRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcBrowseRequestResult)
		}
	}

	return r0
}

// MockPlcBrowser_Browse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Browse'
type MockPlcBrowser_Browse_Call struct {
	*mock.Call
}

// Browse is a helper method to define mock.On call
//   - ctx context.Context
//   - browseRequest model.PlcBrowseRequest
func (_e *MockPlcBrowser_Expecter) Browse(ctx interface{}, browseRequest interface{}) *MockPlcBrowser_Browse_Call {
	return &MockPlcBrowser_Browse_Call{Call: _e.mock.On("Browse", ctx, browseRequest)}
}

func (_c *MockPlcBrowser_Browse_Call) Run(run func(ctx context.Context, browseRequest model.PlcBrowseRequest)) *MockPlcBrowser_Browse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcBrowseRequest))
	})
	return _c
}

func (_c *MockPlcBrowser_Browse_Call) Return(_a0 <-chan model.PlcBrowseRequestResult) *MockPlcBrowser_Browse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowser_Browse_Call) RunAndReturn(run func(context.Context, model.PlcBrowseRequest) <-chan model.PlcBrowseRequestResult) *MockPlcBrowser_Browse_Call {
	_c.Call.Return(run)
	return _c
}

// BrowseWithInterceptor provides a mock function with given fields: ctx, browseRequest, interceptor
func (_m *MockPlcBrowser) BrowseWithInterceptor(ctx context.Context, browseRequest model.PlcBrowseRequest, interceptor func(model.PlcBrowseItem) bool) <-chan model.PlcBrowseRequestResult {
	ret := _m.Called(ctx, browseRequest, interceptor)

	var r0 <-chan model.PlcBrowseRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcBrowseRequest, func(model.PlcBrowseItem) bool) <-chan model.PlcBrowseRequestResult); ok {
		r0 = rf(ctx, browseRequest, interceptor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcBrowseRequestResult)
		}
	}

	return r0
}

// MockPlcBrowser_BrowseWithInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BrowseWithInterceptor'
type MockPlcBrowser_BrowseWithInterceptor_Call struct {
	*mock.Call
}

// BrowseWithInterceptor is a helper method to define mock.On call
//   - ctx context.Context
//   - browseRequest model.PlcBrowseRequest
//   - interceptor func(model.PlcBrowseItem) bool
func (_e *MockPlcBrowser_Expecter) BrowseWithInterceptor(ctx interface{}, browseRequest interface{}, interceptor interface{}) *MockPlcBrowser_BrowseWithInterceptor_Call {
	return &MockPlcBrowser_BrowseWithInterceptor_Call{Call: _e.mock.On("BrowseWithInterceptor", ctx, browseRequest, interceptor)}
}

func (_c *MockPlcBrowser_BrowseWithInterceptor_Call) Run(run func(ctx context.Context, browseRequest model.PlcBrowseRequest, interceptor func(model.PlcBrowseItem) bool)) *MockPlcBrowser_BrowseWithInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcBrowseRequest), args[2].(func(model.PlcBrowseItem) bool))
	})
	return _c
}

func (_c *MockPlcBrowser_BrowseWithInterceptor_Call) Return(_a0 <-chan model.PlcBrowseRequestResult) *MockPlcBrowser_BrowseWithInterceptor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowser_BrowseWithInterceptor_Call) RunAndReturn(run func(context.Context, model.PlcBrowseRequest, func(model.PlcBrowseItem) bool) <-chan model.PlcBrowseRequestResult) *MockPlcBrowser_BrowseWithInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcBrowser creates a new instance of MockPlcBrowser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcBrowser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcBrowser {
	mock := &MockPlcBrowser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
