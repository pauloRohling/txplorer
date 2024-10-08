// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocktransaction

import (
	context "context"
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// RunTransaction provides a mock function with given fields: ctx, fn
func (_m *MockManager) RunTransaction(ctx context.Context, fn func(context.Context) error) error {
	ret := _m.Called(ctx, fn)

	if len(ret) == 0 {
		panic("no return value specified for RunTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_RunTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunTransaction'
type MockManager_RunTransaction_Call struct {
	*mock.Call
}

// RunTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - fn func(context.Context) error
func (_e *MockManager_Expecter) RunTransaction(ctx interface{}, fn interface{}) *MockManager_RunTransaction_Call {
	return &MockManager_RunTransaction_Call{Call: _e.mock.On("RunTransaction", ctx, fn)}
}

func (_c *MockManager_RunTransaction_Call) Run(run func(ctx context.Context, fn func(context.Context) error)) *MockManager_RunTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *MockManager_RunTransaction_Call) Return(_a0 error) *MockManager_RunTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_RunTransaction_Call) RunAndReturn(run func(context.Context, func(context.Context) error) error) *MockManager_RunTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// RunTransactionWithOptions provides a mock function with given fields: ctx, fn, options
func (_m *MockManager) RunTransactionWithOptions(ctx context.Context, fn func(context.Context) error, options *sql.TxOptions) error {
	ret := _m.Called(ctx, fn, options)

	if len(ret) == 0 {
		panic("no return value specified for RunTransactionWithOptions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error, *sql.TxOptions) error); ok {
		r0 = rf(ctx, fn, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_RunTransactionWithOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunTransactionWithOptions'
type MockManager_RunTransactionWithOptions_Call struct {
	*mock.Call
}

// RunTransactionWithOptions is a helper method to define mock.On call
//   - ctx context.Context
//   - fn func(context.Context) error
//   - options *sql.TxOptions
func (_e *MockManager_Expecter) RunTransactionWithOptions(ctx interface{}, fn interface{}, options interface{}) *MockManager_RunTransactionWithOptions_Call {
	return &MockManager_RunTransactionWithOptions_Call{Call: _e.mock.On("RunTransactionWithOptions", ctx, fn, options)}
}

func (_c *MockManager_RunTransactionWithOptions_Call) Run(run func(ctx context.Context, fn func(context.Context) error, options *sql.TxOptions)) *MockManager_RunTransactionWithOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error), args[2].(*sql.TxOptions))
	})
	return _c
}

func (_c *MockManager_RunTransactionWithOptions_Call) Return(_a0 error) *MockManager_RunTransactionWithOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_RunTransactionWithOptions_Call) RunAndReturn(run func(context.Context, func(context.Context) error, *sql.TxOptions) error) *MockManager_RunTransactionWithOptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
