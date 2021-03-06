// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/platform-edn/courier/pkg/proto"
	mock "github.com/stretchr/testify/mock"
)

// MessageServer is an autogenerated mock type for the MessageServer type
type MessageServer struct {
	mock.Mock
}

// PublishMessage provides a mock function with given fields: _a0, _a1
func (_m *MessageServer) PublishMessage(_a0 context.Context, _a1 *proto.PublishMessageRequest) (*proto.PublishMessageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.PublishMessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.PublishMessageRequest) *proto.PublishMessageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.PublishMessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.PublishMessageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestMessage provides a mock function with given fields: _a0, _a1
func (_m *MessageServer) RequestMessage(_a0 context.Context, _a1 *proto.RequestMessageRequest) (*proto.RequestMessageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.RequestMessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RequestMessageRequest) *proto.RequestMessageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RequestMessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.RequestMessageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResponseMessage provides a mock function with given fields: _a0, _a1
func (_m *MessageServer) ResponseMessage(_a0 context.Context, _a1 *proto.ResponseMessageRequest) (*proto.ResponseMessageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ResponseMessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ResponseMessageRequest) *proto.ResponseMessageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ResponseMessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ResponseMessageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedMessageServer provides a mock function with given fields:
func (_m *MessageServer) mustEmbedUnimplementedMessageServer() {
	_m.Called()
}
