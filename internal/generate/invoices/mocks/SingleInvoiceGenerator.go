// Id generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SingleInvoiceGenerator is an autogenerated mock type for the SingleInvoiceGenerator type
type SingleInvoiceGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields: invoiceId
func (_m *SingleInvoiceGenerator) Generate(invoiceId string) (string, error) {
	ret := _m.Called(invoiceId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(invoiceId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
