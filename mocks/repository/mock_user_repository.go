// Code generated by mockery v2.45.0. DO NOT EDIT.

package mockrepository

import (
	context "context"

	model "github.com/pauloRohling/txplorer/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, name, email, password
func (_m *MockUserRepository) Create(ctx context.Context, name string, email string, password string) (*model.User, error) {
	ret := _m.Called(ctx, name, email, password)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*model.User, error)); ok {
		return rf(ctx, name, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *model.User); ok {
		r0 = rf(ctx, name, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - email string
//   - password string
func (_e *MockUserRepository_Expecter) Create(ctx interface{}, name interface{}, email interface{}, password interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", ctx, name, email, password)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(ctx context.Context, name string, email string, password string)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(context.Context, string, string, string) (*model.User, error)) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEmail'
type MockUserRepository_FindByEmail_Call struct {
	*mock.Call
}

// FindByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockUserRepository_Expecter) FindByEmail(ctx interface{}, email interface{}) *MockUserRepository_FindByEmail_Call {
	return &MockUserRepository_FindByEmail_Call{Call: _e.mock.On("FindByEmail", ctx, email)}
}

func (_c *MockUserRepository_FindByEmail_Call) Run(run func(ctx context.Context, email string)) *MockUserRepository_FindByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_FindByEmail_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_FindByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_FindByEmail_Call) RunAndReturn(run func(context.Context, string) (*model.User, error)) *MockUserRepository_FindByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
