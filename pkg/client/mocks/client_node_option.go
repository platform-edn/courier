// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	client "github.com/platform-edn/courier/pkg/client"
	mock "github.com/stretchr/testify/mock"
)

// ClientNodeOption is an autogenerated mock type for the ClientNodeOption type
type ClientNodeOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: c
func (_m *ClientNodeOption) Execute(c *client.ClientNodeOptions) *client.ClientNodeOptions {
	ret := _m.Called(c)

	var r0 *client.ClientNodeOptions
	if rf, ok := ret.Get(0).(func(*client.ClientNodeOptions) *client.ClientNodeOptions); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientNodeOptions)
		}
	}

	return r0
}
