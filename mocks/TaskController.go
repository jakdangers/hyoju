// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// TaskController is an autogenerated mock type for the TaskController type
type TaskController struct {
	mock.Mock
}

type TaskController_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskController) EXPECT() *TaskController_Expecter {
	return &TaskController_Expecter{mock: &_m.Mock}
}

// CreateTask provides a mock function with given fields: c
func (_m *TaskController) CreateMission(c *gin.Context) {
	_m.Called(c)
}

// TaskController_CreateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMission'
type TaskController_CreateTask_Call struct {
	*mock.Call
}

// CreateTask is a helper method to define mock.On call
//   - c *gin.Context
func (_e *TaskController_Expecter) CreateTask(c interface{}) *TaskController_CreateTask_Call {
	return &TaskController_CreateTask_Call{Call: _e.mock.On("CreateMission", c)}
}

func (_c *TaskController_CreateTask_Call) Run(run func(c *gin.Context)) *TaskController_CreateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *TaskController_CreateTask_Call) Return() *TaskController_CreateTask_Call {
	_c.Call.Return()
	return _c
}

func (_c *TaskController_CreateTask_Call) RunAndReturn(run func(*gin.Context)) *TaskController_CreateTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskController creates a new instance of TaskController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskController(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskController {
	mock := &TaskController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
