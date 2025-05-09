// Code generated by mockery v2.53.0. DO NOT EDIT.

package cmocks

import mock "github.com/stretchr/testify/mock"

// Component is an autogenerated mock type for the Component type
type Component struct {
	mock.Mock
}

// IntervalElapsed provides a mock function with no fields
func (_m *Component) IntervalElapsed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IntervalElapsed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Name provides a mock function with no fields
func (_m *Component) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Paused provides a mock function with no fields
func (_m *Component) Paused() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Paused")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Reload provides a mock function with no fields
func (_m *Component) Reload() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Reload")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLastRan provides a mock function with no fields
func (_m *Component) UpdateLastRan() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UpdateLastRan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewComponent creates a new instance of Component. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *Component {
	mock := &Component{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
