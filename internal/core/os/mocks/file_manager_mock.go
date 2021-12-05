// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FileManager is an autogenerated mock type for the FileManager type
type FileManager struct {
	mock.Mock
}

// CreateDir provides a mock function with given fields: dirName
func (_m *FileManager) CreateDirectory(dirName string) (string, error) {
	ret := _m.Called(dirName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(dirName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r1 = rf(dirName)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

// FileExists provides a mock function with given fields: path
func (_m *FileManager) FileExists(path string) (bool, error) {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
