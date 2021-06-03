// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// BlacklistAdd provides a mock function with given fields: ctx, subnet
func (_m *Repository) BlacklistAdd(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlacklistRemove provides a mock function with given fields: ctx, subnet
func (_m *Repository) BlacklistRemove(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhitelistAdd provides a mock function with given fields: ctx, subnet
func (_m *Repository) WhitelistAdd(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhitelistRemove provides a mock function with given fields: ctx, subnet
func (_m *Repository) WhitelistRemove(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}