// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ToolsRemover is an autogenerated mock type for the ToolsRemover type
type ToolsRemover struct {
	mock.Mock
}

// RemoveTool provides a mock function with given fields: ctx, id
func (_m *ToolsRemover) RemoveTool(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}