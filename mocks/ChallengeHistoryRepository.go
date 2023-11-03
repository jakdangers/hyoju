// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "pixelix/entity"

	mock "github.com/stretchr/testify/mock"
)

// ChallengeHistoryRepository is an autogenerated mock type for the ChallengeHistoryRepository type
type ChallengeHistoryRepository struct {
	mock.Mock
}

type ChallengeHistoryRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ChallengeHistoryRepository) EXPECT() *ChallengeHistoryRepository_Expecter {
	return &ChallengeHistoryRepository_Expecter{mock: &_m.Mock}
}

// CreateChallengeHistory provides a mock function with given fields: ctx, missionHistory
func (_m *ChallengeHistoryRepository) CreateChallengeHistory(ctx context.Context, missionHistory *entity.ChallengeHistory) (*entity.ChallengeHistory, error) {
	ret := _m.Called(ctx, missionHistory)

	var r0 *entity.ChallengeHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.ChallengeHistory) (*entity.ChallengeHistory, error)); ok {
		return rf(ctx, missionHistory)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.ChallengeHistory) *entity.ChallengeHistory); ok {
		r0 = rf(ctx, missionHistory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ChallengeHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.ChallengeHistory) error); ok {
		r1 = rf(ctx, missionHistory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChallengeHistoryRepository_CreateChallengeHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateChallengeHistory'
type ChallengeHistoryRepository_CreateChallengeHistory_Call struct {
	*mock.Call
}

// CreateChallengeHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - missionHistory *entity.ChallengeHistory
func (_e *ChallengeHistoryRepository_Expecter) CreateChallengeHistory(ctx interface{}, missionHistory interface{}) *ChallengeHistoryRepository_CreateChallengeHistory_Call {
	return &ChallengeHistoryRepository_CreateChallengeHistory_Call{Call: _e.mock.On("CreateChallengeHistory", ctx, missionHistory)}
}

func (_c *ChallengeHistoryRepository_CreateChallengeHistory_Call) Run(run func(ctx context.Context, missionHistory *entity.ChallengeHistory)) *ChallengeHistoryRepository_CreateChallengeHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.ChallengeHistory))
	})
	return _c
}

func (_c *ChallengeHistoryRepository_CreateChallengeHistory_Call) Return(_a0 *entity.ChallengeHistory, _a1 error) *ChallengeHistoryRepository_CreateChallengeHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChallengeHistoryRepository_CreateChallengeHistory_Call) RunAndReturn(run func(context.Context, *entity.ChallengeHistory) (*entity.ChallengeHistory, error)) *ChallengeHistoryRepository_CreateChallengeHistory_Call {
	_c.Call.Return(run)
	return _c
}

// ListGroupChallengeHistories provides a mock function with given fields: ctx, params
func (_m *ChallengeHistoryRepository) ListGroupChallengeHistories(ctx context.Context, params entity.ListGroupChallengeHistoriesParams) ([]entity.ChallengeHistory, error) {
	ret := _m.Called(ctx, params)

	var r0 []entity.ChallengeHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ListGroupChallengeHistoriesParams) ([]entity.ChallengeHistory, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ListGroupChallengeHistoriesParams) []entity.ChallengeHistory); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ChallengeHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ListGroupChallengeHistoriesParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChallengeHistoryRepository_ListGroupChallengeHistories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGroupChallengeHistories'
type ChallengeHistoryRepository_ListGroupChallengeHistories_Call struct {
	*mock.Call
}

// ListGroupChallengeHistories is a helper method to define mock.On call
//   - ctx context.Context
//   - params entity.ListGroupChallengeHistoriesParams
func (_e *ChallengeHistoryRepository_Expecter) ListGroupChallengeHistories(ctx interface{}, params interface{}) *ChallengeHistoryRepository_ListGroupChallengeHistories_Call {
	return &ChallengeHistoryRepository_ListGroupChallengeHistories_Call{Call: _e.mock.On("ListGroupChallengeHistories", ctx, params)}
}

func (_c *ChallengeHistoryRepository_ListGroupChallengeHistories_Call) Run(run func(ctx context.Context, params entity.ListGroupChallengeHistoriesParams)) *ChallengeHistoryRepository_ListGroupChallengeHistories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ListGroupChallengeHistoriesParams))
	})
	return _c
}

func (_c *ChallengeHistoryRepository_ListGroupChallengeHistories_Call) Return(_a0 []entity.ChallengeHistory, _a1 error) *ChallengeHistoryRepository_ListGroupChallengeHistories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChallengeHistoryRepository_ListGroupChallengeHistories_Call) RunAndReturn(run func(context.Context, entity.ListGroupChallengeHistoriesParams) ([]entity.ChallengeHistory, error)) *ChallengeHistoryRepository_ListGroupChallengeHistories_Call {
	_c.Call.Return(run)
	return _c
}

// NewChallengeHistoryRepository creates a new instance of ChallengeHistoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChallengeHistoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChallengeHistoryRepository {
	mock := &ChallengeHistoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
