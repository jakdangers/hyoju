// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SingleModeController is an autogenerated mock type for the SingleModeController type
type SingleModeController struct {
	mock.Mock
}

type SingleModeController_Expecter struct {
	mock *mock.Mock
}

func (_m *SingleModeController) EXPECT() *SingleModeController_Expecter {
	return &SingleModeController_Expecter{mock: &_m.Mock}
}

// NewSingleModeController creates a new instance of SingleModeController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSingleModeController(t interface {
	mock.TestingT
	Cleanup(func())
}) *SingleModeController {
	mock := &SingleModeController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}