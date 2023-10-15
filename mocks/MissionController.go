// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MissionController is an autogenerated mock type for the MissionController type
type MissionController struct {
	mock.Mock
}

type MissionController_Expecter struct {
	mock *mock.Mock
}

func (_m *MissionController) EXPECT() *MissionController_Expecter {
	return &MissionController_Expecter{mock: &_m.Mock}
}

// CreateMission provides a mock function with given fields: c
func (_m *MissionController) CreateMission(c *gin.Context) {
	_m.Called(c)
}

// MissionController_CreateMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMission'
type MissionController_CreateMission_Call struct {
	*mock.Call
}

// CreateMission is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MissionController_Expecter) CreateMission(c interface{}) *MissionController_CreateMission_Call {
	return &MissionController_CreateMission_Call{Call: _e.mock.On("CreateMission", c)}
}

func (_c *MissionController_CreateMission_Call) Run(run func(c *gin.Context)) *MissionController_CreateMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MissionController_CreateMission_Call) Return() *MissionController_CreateMission_Call {
	_c.Call.Return()
	return _c
}

func (_c *MissionController_CreateMission_Call) RunAndReturn(run func(*gin.Context)) *MissionController_CreateMission_Call {
	_c.Call.Return(run)
	return _c
}

// NewMissionController creates a new instance of MissionController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMissionController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MissionController {
	mock := &MissionController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
