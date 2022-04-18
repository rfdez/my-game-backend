// Code generated by mockery v2.10.4. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mygame "github.com/rfdez/my-game-backend/internal"
	mock "github.com/stretchr/testify/mock"
)

// EventRepository is an autogenerated mock type for the EventRepository type
type EventRepository struct {
	mock.Mock
}

// SearchAll provides a mock function with given fields: ctx
func (_m *EventRepository) SearchAll(ctx context.Context) ([]mygame.Event, error) {
	ret := _m.Called(ctx)

	var r0 []mygame.Event
	if rf, ok := ret.Get(0).(func(context.Context) []mygame.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mygame.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
