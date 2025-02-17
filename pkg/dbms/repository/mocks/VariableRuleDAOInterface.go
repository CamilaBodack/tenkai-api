// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/softplan/tenkai-api/pkg/dbms/model"
	mock "github.com/stretchr/testify/mock"
)

// VariableRuleDAOInterface is an autogenerated mock type for the VariableRuleDAOInterface type
type VariableRuleDAOInterface struct {
	mock.Mock
}

// CreateVariableRule provides a mock function with given fields: e
func (_m *VariableRuleDAOInterface) CreateVariableRule(e model.VariableRule) (int, error) {
	ret := _m.Called(e)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.VariableRule) int); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.VariableRule) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVariableRule provides a mock function with given fields: id
func (_m *VariableRuleDAOInterface) DeleteVariableRule(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditVariableRule provides a mock function with given fields: e
func (_m *VariableRuleDAOInterface) EditVariableRule(e model.VariableRule) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.VariableRule) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListVariableRules provides a mock function with given fields:
func (_m *VariableRuleDAOInterface) ListVariableRules() ([]model.VariableRule, error) {
	ret := _m.Called()

	var r0 []model.VariableRule
	if rf, ok := ret.Get(0).(func() []model.VariableRule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.VariableRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
