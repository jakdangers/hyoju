// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "pixelix/entity"

	mock "github.com/stretchr/testify/mock"
)

// MissionRepository is an autogenerated mock type for the MissionRepository type
type MissionRepository struct {
	mock.Mock
}

type MissionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MissionRepository) EXPECT() *MissionRepository_Expecter {
	return &MissionRepository_Expecter{mock: &_m.Mock}
}

// CreateMission provides a mock function with given fields: ctx, mission
func (_m *MissionRepository) CreateMission(ctx context.Context, mission *entity.Mission) (*entity.Mission, error) {
	ret := _m.Called(ctx, mission)

	var r0 *entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Mission) (*entity.Mission, error)); ok {
		return rf(ctx, mission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Mission) *entity.Mission); ok {
		r0 = rf(ctx, mission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Mission) error); ok {
		r1 = rf(ctx, mission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_CreateMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMission'
type MissionRepository_CreateMission_Call struct {
	*mock.Call
}

// CreateMission is a helper method to define mock.On call
//   - ctx context.Context
//   - mission *entity.Mission
func (_e *MissionRepository_Expecter) CreateMission(ctx interface{}, mission interface{}) *MissionRepository_CreateMission_Call {
	return &MissionRepository_CreateMission_Call{Call: _e.mock.On("CreateMission", ctx, mission)}
}

func (_c *MissionRepository_CreateMission_Call) Run(run func(ctx context.Context, mission *entity.Mission)) *MissionRepository_CreateMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Mission))
	})
	return _c
}

func (_c *MissionRepository_CreateMission_Call) Return(_a0 *entity.Mission, _a1 error) *MissionRepository_CreateMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_CreateMission_Call) RunAndReturn(run func(context.Context, *entity.Mission) (*entity.Mission, error)) *MissionRepository_CreateMission_Call {
	_c.Call.Return(run)
	return _c
}

// GetMission provides a mock function with given fields: ctx, missionID
func (_m *MissionRepository) GetMission(ctx context.Context, missionID uint) (*entity.Mission, error) {
	ret := _m.Called(ctx, missionID)

	var r0 *entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*entity.Mission, error)); ok {
		return rf(ctx, missionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Mission); ok {
		r0 = rf(ctx, missionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, missionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_GetMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMission'
type MissionRepository_GetMission_Call struct {
	*mock.Call
}

// GetMission is a helper method to define mock.On call
//   - ctx context.Context
//   - missionID uint
func (_e *MissionRepository_Expecter) GetMission(ctx interface{}, missionID interface{}) *MissionRepository_GetMission_Call {
	return &MissionRepository_GetMission_Call{Call: _e.mock.On("GetMission", ctx, missionID)}
}

func (_c *MissionRepository_GetMission_Call) Run(run func(ctx context.Context, missionID uint)) *MissionRepository_GetMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *MissionRepository_GetMission_Call) Return(_a0 *entity.Mission, _a1 error) *MissionRepository_GetMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_GetMission_Call) RunAndReturn(run func(context.Context, uint) (*entity.Mission, error)) *MissionRepository_GetMission_Call {
	_c.Call.Return(run)
	return _c
}

// ListActiveSingleMissionIDs provides a mock function with given fields: ctx
func (_m *MissionRepository) ListActiveSingleMissionIDs(ctx context.Context) ([]uint, error) {
	ret := _m.Called(ctx)

	var r0 []uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uint, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uint); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_ListActiveSingleMissionIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListActiveSingleMissionIDs'
type MissionRepository_ListActiveSingleMissionIDs_Call struct {
	*mock.Call
}

// ListActiveSingleMissionIDs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MissionRepository_Expecter) ListActiveSingleMissionIDs(ctx interface{}) *MissionRepository_ListActiveSingleMissionIDs_Call {
	return &MissionRepository_ListActiveSingleMissionIDs_Call{Call: _e.mock.On("ListActiveSingleMissionIDs", ctx)}
}

func (_c *MissionRepository_ListActiveSingleMissionIDs_Call) Run(run func(ctx context.Context)) *MissionRepository_ListActiveSingleMissionIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MissionRepository_ListActiveSingleMissionIDs_Call) Return(_a0 []uint, _a1 error) *MissionRepository_ListActiveSingleMissionIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_ListActiveSingleMissionIDs_Call) RunAndReturn(run func(context.Context) ([]uint, error)) *MissionRepository_ListActiveSingleMissionIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ListMissions provides a mock function with given fields: ctx, userID
func (_m *MissionRepository) ListMissions(ctx context.Context, userID entity.BinaryUUID) ([]entity.Mission, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.BinaryUUID) ([]entity.Mission, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.BinaryUUID) []entity.Mission); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.BinaryUUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_ListMissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMissions'
type MissionRepository_ListMissions_Call struct {
	*mock.Call
}

// ListMissions is a helper method to define mock.On call
//   - ctx context.Context
//   - userID entity.BinaryUUID
func (_e *MissionRepository_Expecter) ListMissions(ctx interface{}, userID interface{}) *MissionRepository_ListMissions_Call {
	return &MissionRepository_ListMissions_Call{Call: _e.mock.On("ListMissions", ctx, userID)}
}

func (_c *MissionRepository_ListMissions_Call) Run(run func(ctx context.Context, userID entity.BinaryUUID)) *MissionRepository_ListMissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.BinaryUUID))
	})
	return _c
}

func (_c *MissionRepository_ListMissions_Call) Return(_a0 []entity.Mission, _a1 error) *MissionRepository_ListMissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_ListMissions_Call) RunAndReturn(run func(context.Context, entity.BinaryUUID) ([]entity.Mission, error)) *MissionRepository_ListMissions_Call {
	_c.Call.Return(run)
	return _c
}

// PatchMission provides a mock function with given fields: ctx, mission
func (_m *MissionRepository) PatchMission(ctx context.Context, mission *entity.Mission) (*entity.Mission, error) {
	ret := _m.Called(ctx, mission)

	var r0 *entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Mission) (*entity.Mission, error)); ok {
		return rf(ctx, mission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Mission) *entity.Mission); ok {
		r0 = rf(ctx, mission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Mission) error); ok {
		r1 = rf(ctx, mission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_PatchMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchMission'
type MissionRepository_PatchMission_Call struct {
	*mock.Call
}

// PatchMission is a helper method to define mock.On call
//   - ctx context.Context
//   - mission *entity.Mission
func (_e *MissionRepository_Expecter) PatchMission(ctx interface{}, mission interface{}) *MissionRepository_PatchMission_Call {
	return &MissionRepository_PatchMission_Call{Call: _e.mock.On("PatchMission", ctx, mission)}
}

func (_c *MissionRepository_PatchMission_Call) Run(run func(ctx context.Context, mission *entity.Mission)) *MissionRepository_PatchMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Mission))
	})
	return _c
}

func (_c *MissionRepository_PatchMission_Call) Return(_a0 *entity.Mission, _a1 error) *MissionRepository_PatchMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_PatchMission_Call) RunAndReturn(run func(context.Context, *entity.Mission) (*entity.Mission, error)) *MissionRepository_PatchMission_Call {
	_c.Call.Return(run)
	return _c
}

// NewMissionRepository creates a new instance of MissionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMissionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MissionRepository {
	mock := &MissionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
