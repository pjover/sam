// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// OsService is an autogenerated mock type for the OsService type
type OsService struct {
	mock.Mock
}

// CopyFile provides a mock function with given fields: sourceFilePath, destinationFilePath
func (_m *OsService) CopyFile(sourceFilePath string, destinationFilePath string) error {
	ret := _m.Called(sourceFilePath, destinationFilePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(sourceFilePath, destinationFilePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDirectory provides a mock function with given fields: dirPath
func (_m *OsService) CreateDirectory(dirPath string) error {
	ret := _m.Called(dirPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dirPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateZipFile provides a mock function with given fields: zipFilePath, files
func (_m *OsService) CreateZipFile(zipFilePath string, files []string) error {
	ret := _m.Called(zipFilePath, files)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(zipFilePath, files)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTempDirectory provides a mock function with given fields:
func (_m *OsService) GetTempDirectory() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemExists provides a mock function with given fields: itemPath
func (_m *OsService) ItemExists(itemPath string) (bool, error) {
	ret := _m.Called(itemPath)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(itemPath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(itemPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFiles provides a mock function with given fields: dir, ext
func (_m *OsService) ListFiles(dir string, ext string) ([]string, error) {
	ret := _m.Called(dir, ext)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(dir, ext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dir, ext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Now provides a mock function with given fields:
func (_m *OsService) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// OpenUrlInBrowser provides a mock function with given fields: url
func (_m *OsService) OpenUrlInBrowser(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadFile provides a mock function with given fields: filePath
func (_m *OsService) ReadFile(filePath string) ([]byte, error) {
	ret := _m.Called(filePath)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunCommand provides a mock function with given fields: command, args
func (_m *OsService) RunCommand(command string, args ...string) error {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteFile provides a mock function with given fields: dirPath, filename, content
func (_m *OsService) WriteFile(dirPath string, filename string, content []byte) (string, error) {
	ret := _m.Called(dirPath, filename, content)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, []byte) string); ok {
		r0 = rf(dirPath, filename, content)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []byte) error); ok {
		r1 = rf(dirPath, filename, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}