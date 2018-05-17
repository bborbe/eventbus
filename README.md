# Eventbus

Package provide eventbus in Go

## Usage

```
type TestEvent struct{}

eventBus := New()
eventBus.RegisterHandler(func(event TestEventA) {
  // handle event
})
eventBus.Publish(TestEventA{})
```
