// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	chat "github.com/iggyster/lets-go-chat/internal/chat"
	mock "github.com/stretchr/testify/mock"
)

// MessageRepo is an autogenerated mock type for the MessageRepo type
type MessageRepo struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *MessageRepo) FindAll() []chat.Msg {
	ret := _m.Called()

	var r0 []chat.Msg
	if rf, ok := ret.Get(0).(func() []chat.Msg); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chat.Msg)
		}
	}

	return r0
}

// Save provides a mock function with given fields: msg
func (_m *MessageRepo) Save(msg *chat.Msg) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chat.Msg) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMessageRepo creates a new instance of MessageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageRepo {
	mock := &MessageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
