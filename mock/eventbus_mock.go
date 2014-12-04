package mock

import "fmt"

type eventbusMock struct {
	RegisterHandlerFunc   func(fn interface{}) error
	UnregisterHandlerFunc func(fn interface{}) error
	PublishFunc           func(fn interface{}) error
}

func New() *eventbusMock {
	e := new(eventbusMock)
	return e
}

func (e *eventbusMock) RegisterHandler(fn interface{}) error {
	if e.RegisterHandlerFunc == nil {
		return fmt.Errorf("RegisterHandlerFunc not defined")
	}
	return e.RegisterHandlerFunc(fn)
}

func (e *eventbusMock) UnregisterHandler(fn interface{}) error {
	if e.UnregisterHandlerFunc == nil {
		return fmt.Errorf("UnregisterHandlerFunc not defined")
	}
	return e.UnregisterHandlerFunc(fn)
}

func (e *eventbusMock) Publish(event interface{}) error {
	if e.PublishFunc == nil {
		return fmt.Errorf("PublishFunc not defined")
	}
	return e.PublishFunc(event)
}
