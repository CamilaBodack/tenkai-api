// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	model "github.com/softplan/tenkai-api/pkg/dbms/model"
	mock "github.com/stretchr/testify/mock"
)

// UserEnvironmentRoleDAOInterface is an autogenerated mock type for the UserEnvironmentRoleDAOInterface type
type UserEnvironmentRoleDAOInterface struct {
	mock.Mock
}

// CreateOrUpdate provides a mock function with given fields: so
func (_m *UserEnvironmentRoleDAOInterface) CreateOrUpdate(so model.UserEnvironmentRole) error {
	ret := _m.Called(so)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.UserEnvironmentRole) error); ok {
		r0 = rf(so)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRoleByUserAndEnvironment provides a mock function with given fields: user, envID
func (_m *UserEnvironmentRoleDAOInterface) GetRoleByUserAndEnvironment(user model.User, envID uint) (*model.SecurityOperation, error) {
	ret := _m.Called(user, envID)

	var r0 *model.SecurityOperation
	if rf, ok := ret.Get(0).(func(model.User, uint) *model.SecurityOperation); ok {
		r0 = rf(user, envID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SecurityOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User, uint) error); ok {
		r1 = rf(user, envID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersAndRoleByEnv provides a mock function with given fields: id
func (_m *UserEnvironmentRoleDAOInterface) GetUsersAndRoleByEnv(id int) ([]model.UserEnvRole, error) {
	ret := _m.Called(id)

	var r0 []model.UserEnvRole
	if rf, ok := ret.Get(0).(func(int) []model.UserEnvRole); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.UserEnvRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
