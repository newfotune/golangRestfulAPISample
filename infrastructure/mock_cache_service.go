// Code generated by mockery v1.0.0. DO NOT EDIT.

package infrastructure

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockCacheServiceInterface is an autogenerated mock type for the CacheServiceInterface type
type MockCacheServiceInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, key
func (_m *MockCacheServiceInterface) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGet provides a mock function with given fields: ctx, key
func (_m *MockCacheServiceInterface) GetGet(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, value, expiry
func (_m *MockCacheServiceInterface) Set(ctx context.Context, key string, value interface{}, expiry int) error {
	ret := _m.Called(ctx, key, value, expiry)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, int) error); ok {
		r0 = rf(ctx, key, value, expiry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetIfNotExist provides a mock function with given fields: ctx, key, value, expiry
func (_m *MockCacheServiceInterface) SetIfNotExist(ctx context.Context, key string, value interface{}, expiry int) (bool, error) {
	ret := _m.Called(ctx, key, value, expiry)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, int) bool); ok {
		r0 = rf(ctx, key, value, expiry)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, int) error); ok {
		r1 = rf(ctx, key, value, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}