// Code generated by mockery v2.32.4. DO NOT EDIT.

package cluster

import (
	context "context"

	internalcluster "github.com/tochemey/goakt/internal/cluster"
	internalpb "github.com/tochemey/goakt/internal/internalpb"

	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

type Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *Interface) EXPECT() *Interface_Expecter {
	return &Interface_Expecter{mock: &_m.Mock}
}

// AdvertisedAddress provides a mock function with given fields:
func (_m *Interface) AdvertisedAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Interface_AdvertisedAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AdvertisedAddress'
type Interface_AdvertisedAddress_Call struct {
	*mock.Call
}

// AdvertisedAddress is a helper method to define mock.On call
func (_e *Interface_Expecter) AdvertisedAddress() *Interface_AdvertisedAddress_Call {
	return &Interface_AdvertisedAddress_Call{Call: _e.mock.On("AdvertisedAddress")}
}

func (_c *Interface_AdvertisedAddress_Call) Run(run func()) *Interface_AdvertisedAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Interface_AdvertisedAddress_Call) Return(_a0 string) *Interface_AdvertisedAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_AdvertisedAddress_Call) RunAndReturn(run func() string) *Interface_AdvertisedAddress_Call {
	_c.Call.Return(run)
	return _c
}

// Events provides a mock function with given fields:
func (_m *Interface) Events() <-chan *internalcluster.Event {
	ret := _m.Called()

	var r0 <-chan *internalcluster.Event
	if rf, ok := ret.Get(0).(func() <-chan *internalcluster.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *internalcluster.Event)
		}
	}

	return r0
}

// Interface_Events_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Events'
type Interface_Events_Call struct {
	*mock.Call
}

// Events is a helper method to define mock.On call
func (_e *Interface_Expecter) Events() *Interface_Events_Call {
	return &Interface_Events_Call{Call: _e.mock.On("Events")}
}

func (_c *Interface_Events_Call) Run(run func()) *Interface_Events_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Interface_Events_Call) Return(_a0 <-chan *internalcluster.Event) *Interface_Events_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_Events_Call) RunAndReturn(run func() <-chan *internalcluster.Event) *Interface_Events_Call {
	_c.Call.Return(run)
	return _c
}

// GetActor provides a mock function with given fields: ctx, actorName
func (_m *Interface) GetActor(ctx context.Context, actorName string) (*internalpb.WireActor, error) {
	ret := _m.Called(ctx, actorName)

	var r0 *internalpb.WireActor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*internalpb.WireActor, error)); ok {
		return rf(ctx, actorName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *internalpb.WireActor); ok {
		r0 = rf(ctx, actorName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.WireActor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, actorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_GetActor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActor'
type Interface_GetActor_Call struct {
	*mock.Call
}

// GetActor is a helper method to define mock.On call
//   - ctx context.Context
//   - actorName string
func (_e *Interface_Expecter) GetActor(ctx interface{}, actorName interface{}) *Interface_GetActor_Call {
	return &Interface_GetActor_Call{Call: _e.mock.On("GetActor", ctx, actorName)}
}

func (_c *Interface_GetActor_Call) Run(run func(ctx context.Context, actorName string)) *Interface_GetActor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Interface_GetActor_Call) Return(_a0 *internalpb.WireActor, _a1 error) *Interface_GetActor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_GetActor_Call) RunAndReturn(run func(context.Context, string) (*internalpb.WireActor, error)) *Interface_GetActor_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartition provides a mock function with given fields: actorName
func (_m *Interface) GetPartition(actorName string) int {
	ret := _m.Called(actorName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(actorName)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Interface_GetPartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartition'
type Interface_GetPartition_Call struct {
	*mock.Call
}

// GetPartition is a helper method to define mock.On call
//   - actorName string
func (_e *Interface_Expecter) GetPartition(actorName interface{}) *Interface_GetPartition_Call {
	return &Interface_GetPartition_Call{Call: _e.mock.On("GetPartition", actorName)}
}

func (_c *Interface_GetPartition_Call) Run(run func(actorName string)) *Interface_GetPartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Interface_GetPartition_Call) Return(_a0 int) *Interface_GetPartition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_GetPartition_Call) RunAndReturn(run func(string) int) *Interface_GetPartition_Call {
	_c.Call.Return(run)
	return _c
}

// KeyExists provides a mock function with given fields: ctx, key
func (_m *Interface) KeyExists(ctx context.Context, key string) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_KeyExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeyExists'
type Interface_KeyExists_Call struct {
	*mock.Call
}

// KeyExists is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Interface_Expecter) KeyExists(ctx interface{}, key interface{}) *Interface_KeyExists_Call {
	return &Interface_KeyExists_Call{Call: _e.mock.On("KeyExists", ctx, key)}
}

func (_c *Interface_KeyExists_Call) Run(run func(ctx context.Context, key string)) *Interface_KeyExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Interface_KeyExists_Call) Return(_a0 bool, _a1 error) *Interface_KeyExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_KeyExists_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *Interface_KeyExists_Call {
	_c.Call.Return(run)
	return _c
}

// NodeHost provides a mock function with given fields:
func (_m *Interface) NodeHost() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Interface_NodeHost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeHost'
type Interface_NodeHost_Call struct {
	*mock.Call
}

// NodeHost is a helper method to define mock.On call
func (_e *Interface_Expecter) NodeHost() *Interface_NodeHost_Call {
	return &Interface_NodeHost_Call{Call: _e.mock.On("NodeHost")}
}

func (_c *Interface_NodeHost_Call) Run(run func()) *Interface_NodeHost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Interface_NodeHost_Call) Return(_a0 string) *Interface_NodeHost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_NodeHost_Call) RunAndReturn(run func() string) *Interface_NodeHost_Call {
	_c.Call.Return(run)
	return _c
}

// NodeRemotingPort provides a mock function with given fields:
func (_m *Interface) NodeRemotingPort() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Interface_NodeRemotingPort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeRemotingPort'
type Interface_NodeRemotingPort_Call struct {
	*mock.Call
}

// NodeRemotingPort is a helper method to define mock.On call
func (_e *Interface_Expecter) NodeRemotingPort() *Interface_NodeRemotingPort_Call {
	return &Interface_NodeRemotingPort_Call{Call: _e.mock.On("NodeRemotingPort")}
}

func (_c *Interface_NodeRemotingPort_Call) Run(run func()) *Interface_NodeRemotingPort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Interface_NodeRemotingPort_Call) Return(_a0 int) *Interface_NodeRemotingPort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_NodeRemotingPort_Call) RunAndReturn(run func() int) *Interface_NodeRemotingPort_Call {
	_c.Call.Return(run)
	return _c
}

// PutActor provides a mock function with given fields: ctx, actor
func (_m *Interface) PutActor(ctx context.Context, actor *internalpb.WireActor) error {
	ret := _m.Called(ctx, actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.WireActor) error); ok {
		r0 = rf(ctx, actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_PutActor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutActor'
type Interface_PutActor_Call struct {
	*mock.Call
}

// PutActor is a helper method to define mock.On call
//   - ctx context.Context
//   - actor *internalpb.WireActor
func (_e *Interface_Expecter) PutActor(ctx interface{}, actor interface{}) *Interface_PutActor_Call {
	return &Interface_PutActor_Call{Call: _e.mock.On("PutActor", ctx, actor)}
}

func (_c *Interface_PutActor_Call) Run(run func(ctx context.Context, actor *internalpb.WireActor)) *Interface_PutActor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.WireActor))
	})
	return _c
}

func (_c *Interface_PutActor_Call) Return(_a0 error) *Interface_PutActor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_PutActor_Call) RunAndReturn(run func(context.Context, *internalpb.WireActor) error) *Interface_PutActor_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveActor provides a mock function with given fields: ctx, actorName
func (_m *Interface) RemoveActor(ctx context.Context, actorName string) error {
	ret := _m.Called(ctx, actorName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, actorName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_RemoveActor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveActor'
type Interface_RemoveActor_Call struct {
	*mock.Call
}

// RemoveActor is a helper method to define mock.On call
//   - ctx context.Context
//   - actorName string
func (_e *Interface_Expecter) RemoveActor(ctx interface{}, actorName interface{}) *Interface_RemoveActor_Call {
	return &Interface_RemoveActor_Call{Call: _e.mock.On("RemoveActor", ctx, actorName)}
}

func (_c *Interface_RemoveActor_Call) Run(run func(ctx context.Context, actorName string)) *Interface_RemoveActor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Interface_RemoveActor_Call) Return(_a0 error) *Interface_RemoveActor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_RemoveActor_Call) RunAndReturn(run func(context.Context, string) error) *Interface_RemoveActor_Call {
	_c.Call.Return(run)
	return _c
}

// SetKey provides a mock function with given fields: ctx, key
func (_m *Interface) SetKey(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_SetKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetKey'
type Interface_SetKey_Call struct {
	*mock.Call
}

// SetKey is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Interface_Expecter) SetKey(ctx interface{}, key interface{}) *Interface_SetKey_Call {
	return &Interface_SetKey_Call{Call: _e.mock.On("SetKey", ctx, key)}
}

func (_c *Interface_SetKey_Call) Run(run func(ctx context.Context, key string)) *Interface_SetKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Interface_SetKey_Call) Return(_a0 error) *Interface_SetKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_SetKey_Call) RunAndReturn(run func(context.Context, string) error) *Interface_SetKey_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Interface) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Interface_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Interface_Expecter) Start(ctx interface{}) *Interface_Start_Call {
	return &Interface_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Interface_Start_Call) Run(run func(ctx context.Context)) *Interface_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Interface_Start_Call) Return(_a0 error) *Interface_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_Start_Call) RunAndReturn(run func(context.Context) error) *Interface_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: ctx
func (_m *Interface) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Interface_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Interface_Expecter) Stop(ctx interface{}) *Interface_Stop_Call {
	return &Interface_Stop_Call{Call: _e.mock.On("Stop", ctx)}
}

func (_c *Interface_Stop_Call) Run(run func(ctx context.Context)) *Interface_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Interface_Stop_Call) Return(_a0 error) *Interface_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_Stop_Call) RunAndReturn(run func(context.Context) error) *Interface_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}