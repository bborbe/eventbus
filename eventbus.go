package eventbus

import (
	"reflect"
	"fmt"
	"sync"
)

type EventBus interface {
	RegisterHandler(fn interface {}) error
	UnregisterHandler(fn interface {}) error
	Publish(event interface{}) error
}

type eventBus struct {
	handlers map[reflect.Type][]reflect.Value
	lock     sync.RWMutex
}

func New() *eventBus {
	e := new(eventBus)
	e.handlers = make(map[reflect.Type][]reflect.Value)
	return e
}

func parse(fn interface {}) (reflect.Type, reflect.Value, error) {
	v := reflect.ValueOf(fn)
	def := v.Type()
	if def.NumIn() != 1 {
		return nil, v, fmt.Errorf("Handler must have a single argument")
	}
	argument := def.In(0)
	return argument, v, nil
}

func (e *eventBus) RegisterHandler(fn interface {}) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	argument, v, err := parse(fn)
	if err != nil {
		return err
	}

	e.handlers[argument] = append(e.handlers[argument], v)
	return nil
}

func (e *eventBus) UnregisterHandler(fn interface {}) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	argument, v, err := parse(fn)
	if err != nil {
		return err
	}

	handlers := make([]reflect.Value, 0)
	for _, handler := range e.handlers[argument] {
		if handler != v {
			handlers = append(handlers, handler)
		}
	}
	e.handlers[argument] = handlers
	return nil
}

func (e *eventBus) Publish(event interface{}) error {
	e.lock.RLock()
	defer e.lock.RUnlock()

	t := reflect.TypeOf(event)
	args := [...]reflect.Value{reflect.ValueOf(event)}
	fns := e.handlers[t]
	for _, fn := range fns {
		fn.Call(args[:])
	}
	return nil
}
