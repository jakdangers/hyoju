// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

type TaskRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskRepository) EXPECT() *TaskRepository_Expecter {
	return &TaskRepository_Expecter{mock: &_m.Mock}
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}