// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	registry "github.com/platform-edn/courier/pkg/registry"
	mock "github.com/stretchr/testify/mock"
)

// NodeMapper is an autogenerated mock type for the NodeMapper type
type NodeMapper struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: _a0
func (_m *NodeMapper) AddNode(_a0 registry.Node) {
	_m.Called(_a0)
}

// RemoveNode provides a mock function with given fields: _a0
func (_m *NodeMapper) RemoveNode(_a0 string) {
	_m.Called(_a0)
}
