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

// CreateChallenge provides a mock function with given fields: ctx, challenge
func (_m *MissionRepository) CreateChallenge(ctx context.Context, mission *entity.Challenge) (*entity.Challenge, error) {
	ret := _m.Called(ctx, mission)

	var r0 *entity.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Challenge) (*entity.Challenge, error)); ok {
		return rf(ctx, mission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Challenge) *entity.Challenge); ok {
		r0 = rf(ctx, mission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Challenge) error); ok {
		r1 = rf(ctx, mission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_CreateMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateChallenge'
type MissionRepository_CreateMission_Call struct {
	*mock.Call
}

// CreateMission is a helper method to define mock.On call
//   - ctx context.Context
//   - challenge *entity.Challenge
func (_e *MissionRepository_Expecter) CreateMission(ctx interface{}, mission interface{}) *MissionRepository_CreateMission_Call {
	return &MissionRepository_CreateMission_Call{Call: _e.mock.On("CreateChallenge", ctx, mission)}
}

func (_c *MissionRepository_CreateMission_Call) Run(run func(ctx context.Context, mission *entity.Challenge)) *MissionRepository_CreateMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Challenge))
	})
	return _c
}

func (_c *MissionRepository_CreateMission_Call) Return(_a0 *entity.Challenge, _a1 error) *MissionRepository_CreateMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_CreateMission_Call) RunAndReturn(run func(context.Context, *entity.Challenge) (*entity.Challenge, error)) *MissionRepository_CreateMission_Call {
	_c.Call.Return(run)
	return _c
}

// GetChallenge provides a mock function with given fields: ctx, missionID
func (_m *MissionRepository) GetChallenge(ctx context.Context, missionID uint) (*entity.Challenge, error) {
	ret := _m.Called(ctx, missionID)

	var r0 *entity.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*entity.Challenge, error)); ok {
		return rf(ctx, missionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Challenge); ok {
		r0 = rf(ctx, missionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, missionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_GetMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChallenge'
type MissionRepository_GetMission_Call struct {
	*mock.Call
}

// GetMission is a helper method to define mock.On call
//   - ctx context.Context
//   - missionID uint
func (_e *MissionRepository_Expecter) GetMission(ctx interface{}, missionID interface{}) *MissionRepository_GetMission_Call {
	return &MissionRepository_GetMission_Call{Call: _e.mock.On("GetChallenge", ctx, missionID)}
}

func (_c *MissionRepository_GetMission_Call) Run(run func(ctx context.Context, missionID uint)) *MissionRepository_GetMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *MissionRepository_GetMission_Call) Return(_a0 *entity.Challenge, _a1 error) *MissionRepository_GetMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_GetMission_Call) RunAndReturn(run func(context.Context, uint) (*entity.Challenge, error)) *MissionRepository_GetMission_Call {
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

// ListChallenges provides a mock function with given fields: ctx, userID
func (_m *MissionRepository) ListChallenges(ctx context.Context, userID entity.BinaryUUID) ([]entity.Challenge, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entity.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.BinaryUUID) ([]entity.Challenge, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.BinaryUUID) []entity.Challenge); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.BinaryUUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_ListMissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListChallenges'
type MissionRepository_ListMissions_Call struct {
	*mock.Call
}

// ListMissions is a helper method to define mock.On call
//   - ctx context.Context
//   - userID entity.BinaryUUID
func (_e *MissionRepository_Expecter) ListMissions(ctx interface{}, userID interface{}) *MissionRepository_ListMissions_Call {
	return &MissionRepository_ListMissions_Call{Call: _e.mock.On("ListChallenges", ctx, userID)}
}

func (_c *MissionRepository_ListMissions_Call) Run(run func(ctx context.Context, userID entity.BinaryUUID)) *MissionRepository_ListMissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.BinaryUUID))
	})
	return _c
}

func (_c *MissionRepository_ListMissions_Call) Return(_a0 []entity.Challenge, _a1 error) *MissionRepository_ListMissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_ListMissions_Call) RunAndReturn(run func(context.Context, entity.BinaryUUID) ([]entity.Challenge, error)) *MissionRepository_ListMissions_Call {
	_c.Call.Return(run)
	return _c
}

// ListMultiModeMissions provides a mock function with given fields: ctx, params
func (_m *MissionRepository) ListMultiModeMissions(ctx context.Context, params entity.ListMultiModeMissionsParams) ([]entity.Challenge, error) {
	ret := _m.Called(ctx, params)

	var r0 []entity.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ListMultiModeMissionsParams) ([]entity.Challenge, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ListMultiModeMissionsParams) []entity.Challenge); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ListMultiModeMissionsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_ListMultiModeMissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMultiModeMissions'
type MissionRepository_ListMultiModeMissions_Call struct {
	*mock.Call
}

// ListMultiModeMissions is a helper method to define mock.On call
//   - ctx context.Context
//   - params entity.ListMultiModeMissionsParams
func (_e *MissionRepository_Expecter) ListMultiModeMissions(ctx interface{}, params interface{}) *MissionRepository_ListMultiModeMissions_Call {
	return &MissionRepository_ListMultiModeMissions_Call{Call: _e.mock.On("ListMultiModeMissions", ctx, params)}
}

func (_c *MissionRepository_ListMultiModeMissions_Call) Run(run func(ctx context.Context, params entity.ListMultiModeMissionsParams)) *MissionRepository_ListMultiModeMissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ListMultiModeMissionsParams))
	})
	return _c
}

func (_c *MissionRepository_ListMultiModeMissions_Call) Return(_a0 []entity.Challenge, _a1 error) *MissionRepository_ListMultiModeMissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_ListMultiModeMissions_Call) RunAndReturn(run func(context.Context, entity.ListMultiModeMissionsParams) ([]entity.Challenge, error)) *MissionRepository_ListMultiModeMissions_Call {
	_c.Call.Return(run)
	return _c
}

// PatchChallenge provides a mock function with given fields: ctx, challenge
func (_m *MissionRepository) PatchChallenge(ctx context.Context, mission *entity.Challenge) (*entity.Challenge, error) {
	ret := _m.Called(ctx, mission)

	var r0 *entity.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Challenge) (*entity.Challenge, error)); ok {
		return rf(ctx, mission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Challenge) *entity.Challenge); ok {
		r0 = rf(ctx, mission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Challenge) error); ok {
		r1 = rf(ctx, mission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MissionRepository_PatchMission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchChallenge'
type MissionRepository_PatchMission_Call struct {
	*mock.Call
}

// PatchMission is a helper method to define mock.On call
//   - ctx context.Context
//   - challenge *entity.Challenge
func (_e *MissionRepository_Expecter) PatchMission(ctx interface{}, mission interface{}) *MissionRepository_PatchMission_Call {
	return &MissionRepository_PatchMission_Call{Call: _e.mock.On("PatchChallenge", ctx, mission)}
}

func (_c *MissionRepository_PatchMission_Call) Run(run func(ctx context.Context, mission *entity.Challenge)) *MissionRepository_PatchMission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Challenge))
	})
	return _c
}

func (_c *MissionRepository_PatchMission_Call) Return(_a0 *entity.Challenge, _a1 error) *MissionRepository_PatchMission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MissionRepository_PatchMission_Call) RunAndReturn(run func(context.Context, *entity.Challenge) (*entity.Challenge, error)) *MissionRepository_PatchMission_Call {
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
