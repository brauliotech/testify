// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Printer is an autogenerated mock type for the Printer type
type Printer struct {
	mock.Mock
}

// Print provides a mock function with given fields: paperSize, content
func (_m *Printer) Print(paperSize string, content string) error {
	ret := _m.Called(paperSize, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(paperSize, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPrinter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPrinter creates a new instance of Printer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPrinter(t mockConstructorTestingTNewPrinter) *Printer {
	mock := &Printer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
