// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "pixelix/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// DeleteUser provides a mock function with given fields: ctx, req
func (_m *UserService) DeleteUser(ctx context.Context, req entity.DeleteUserRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.DeleteUserRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserService_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.DeleteUserRequest
func (_e *UserService_Expecter) DeleteUser(ctx interface{}, req interface{}) *UserService_DeleteUser_Call {
	return &UserService_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, req)}
}

func (_c *UserService_DeleteUser_Call) Run(run func(ctx context.Context, req entity.DeleteUserRequest)) *UserService_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.DeleteUserRequest))
	})
	return _c
}

func (_c *UserService_DeleteUser_Call) Return(_a0 error) *UserService_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_DeleteUser_Call) RunAndReturn(run func(context.Context, entity.DeleteUserRequest) error) *UserService_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// OAuthLoginUser provides a mock function with given fields: ctx, req
func (_m *UserService) OAuthLoginUser(ctx context.Context, req entity.OAuthLoginUserRequest) (*entity.OAuthLoginUserResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *entity.OAuthLoginUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.OAuthLoginUserRequest) (*entity.OAuthLoginUserResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.OAuthLoginUserRequest) *entity.OAuthLoginUserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OAuthLoginUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.OAuthLoginUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_OAuthLoginUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OAuthLoginUser'
type UserService_OAuthLoginUser_Call struct {
	*mock.Call
}

// OAuthLoginUser is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.OAuthLoginUserRequest
func (_e *UserService_Expecter) OAuthLoginUser(ctx interface{}, req interface{}) *UserService_OAuthLoginUser_Call {
	return &UserService_OAuthLoginUser_Call{Call: _e.mock.On("OAuthLoginUser", ctx, req)}
}

func (_c *UserService_OAuthLoginUser_Call) Run(run func(ctx context.Context, req entity.OAuthLoginUserRequest)) *UserService_OAuthLoginUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.OAuthLoginUserRequest))
	})
	return _c
}

func (_c *UserService_OAuthLoginUser_Call) Return(_a0 *entity.OAuthLoginUserResponse, _a1 error) *UserService_OAuthLoginUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_OAuthLoginUser_Call) RunAndReturn(run func(context.Context, entity.OAuthLoginUserRequest) (*entity.OAuthLoginUserResponse, error)) *UserService_OAuthLoginUser_Call {
	_c.Call.Return(run)
	return _c
}

// ReadUser provides a mock function with given fields: ctx, req
func (_m *UserService) ReadUser(ctx context.Context, req entity.ReadUserRequest) (*entity.ReadUserResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *entity.ReadUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ReadUserRequest) (*entity.ReadUserResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ReadUserRequest) *entity.ReadUserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ReadUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ReadUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_ReadUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadUser'
type UserService_ReadUser_Call struct {
	*mock.Call
}

// ReadUser is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.ReadUserRequest
func (_e *UserService_Expecter) ReadUser(ctx interface{}, req interface{}) *UserService_ReadUser_Call {
	return &UserService_ReadUser_Call{Call: _e.mock.On("ReadUser", ctx, req)}
}

func (_c *UserService_ReadUser_Call) Run(run func(ctx context.Context, req entity.ReadUserRequest)) *UserService_ReadUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ReadUserRequest))
	})
	return _c
}

func (_c *UserService_ReadUser_Call) Return(_a0 *entity.ReadUserResponse, _a1 error) *UserService_ReadUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_ReadUser_Call) RunAndReturn(run func(context.Context, entity.ReadUserRequest) (*entity.ReadUserResponse, error)) *UserService_ReadUser_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, req
func (_m *UserService) UpdateUser(ctx context.Context, req entity.UpdateUserRequest) (*entity.UpdateUserResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *entity.UpdateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateUserRequest) (*entity.UpdateUserResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateUserRequest) *entity.UpdateUserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UpdateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.UpdateUserRequest
func (_e *UserService_Expecter) UpdateUser(ctx interface{}, req interface{}) *UserService_UpdateUser_Call {
	return &UserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, req)}
}

func (_c *UserService_UpdateUser_Call) Run(run func(ctx context.Context, req entity.UpdateUserRequest)) *UserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.UpdateUserRequest))
	})
	return _c
}

func (_c *UserService_UpdateUser_Call) Return(_a0 *entity.UpdateUserResponse, _a1 error) *UserService_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_UpdateUser_Call) RunAndReturn(run func(context.Context, entity.UpdateUserRequest) (*entity.UpdateUserResponse, error)) *UserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
