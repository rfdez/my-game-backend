// Code generated by mockery v2.10.4. DO NOT EDIT.

package loggermocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: _a0
func (_m *Logger) Debug(_a0 string) {
	_m.Called(_a0)
}

// Error provides a mock function with given fields: _a0
func (_m *Logger) Error(_a0 string) {
	_m.Called(_a0)
}

// Fatal provides a mock function with given fields: _a0
func (_m *Logger) Fatal(_a0 string) {
	_m.Called(_a0)
}

// Info provides a mock function with given fields: _a0
func (_m *Logger) Info(_a0 string) {
	_m.Called(_a0)
}
