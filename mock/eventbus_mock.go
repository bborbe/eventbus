package mock

type eventbusMock struct {
}

func New() *eventbusMock {
	e := new(eventbusMock)
	return e
}

func (e *eventbusMock) RegisterHandler(fn interface{}) error {
	return nil
}

func (e *eventbusMock) UnregisterHandler(fn interface{}) error {
	return nil
}

func (e *eventbusMock) Publish(event interface{}) error {
	return nil
}
