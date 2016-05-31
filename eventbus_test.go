package eventbus

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplements(t *testing.T) {
	eventBus := New()
	var i *EventBus
	err := AssertThat(eventBus, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

type TestEventA struct{}
type TestEventB struct{}

func TestPublishHandlerIsCalledIfTypeIsMatching(t *testing.T) {
	counter := 0
	eventBus := New()
	eventBus.RegisterHandler(func(event TestEventA) {
		counter++
	})
	eventBus.Publish(TestEventA{})
	eventBus.Publish(TestEventB{})
	err := AssertThat(counter, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishHandlerIsCalledIfTypeIsMatchingRef(t *testing.T) {
	counter := 0
	eventBus := New()
	eventBus.RegisterHandler(func(event *TestEventA) {
		counter++
	})
	eventBus.Publish(&TestEventA{})
	eventBus.Publish(&TestEventB{})
	err := AssertThat(counter, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishHandlerIsCalledIfTypeIsMatchingReturn(t *testing.T) {
	counter := 0
	eventBus := New()
	eventBus.RegisterHandler(func(event TestEventA) error {
		counter++
		return nil
	})
	eventBus.Publish(TestEventA{})
	eventBus.Publish(TestEventB{})
	err := AssertThat(counter, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublisMultiHandlerAreCalled(t *testing.T) {
	counter := 0
	eventBus := New()
	eventBus.RegisterHandler(func(event TestEventA) {
		counter++
	})
	eventBus.RegisterHandler(func(event TestEventA) {
		counter++
	})
	eventBus.Publish(TestEventA{})
	err := AssertThat(counter, Is(2))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishHandlerIsNotCalledIfTypeIsNotMatching(t *testing.T) {
	called := false
	eventBus := New()
	eventBus.RegisterHandler(func(event TestEventA) {
		called = true
	})
	eventBus.Publish(TestEventB{})
	err := AssertThat(called, Is(false))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterHandlerReturnErrorFuncWithoutArgs(t *testing.T) {
	eventBus := New()
	err := eventBus.RegisterHandler(func() {})
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterHandlerReturnErrorFuncWithTwoArgs(t *testing.T) {
	eventBus := New()
	err := eventBus.RegisterHandler(func(a TestEventA, b TestEventA) {})
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterHandlerReturnNoErrorFuncWithOneArg(t *testing.T) {
	eventBus := New()
	err := eventBus.RegisterHandler(func(event TestEventA) {})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnregisterIllegalFunction(t *testing.T) {
	eventBus := New()
	err := eventBus.UnregisterHandler(func() {})
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnregisterHandler(t *testing.T) {
	counter := 0
	eventBus := New()
	fn := func(event TestEventA) {
		counter++
	}
	eventBus.RegisterHandler(fn)
	eventBus.Publish(TestEventA{})
	eventBus.UnregisterHandler(fn)
	eventBus.Publish(TestEventA{})
	err := AssertThat(counter, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnregisterNotRegisteredFunctionHasNoEffect(t *testing.T) {
	counter := 0
	eventBus := New()
	fnA := func(event TestEventA) {
		counter++
	}
	fnB := func(event TestEventA) {
		counter++
	}
	eventBus.RegisterHandler(fnA)
	eventBus.Publish(TestEventA{})
	eventBus.UnregisterHandler(fnB)
	eventBus.Publish(TestEventA{})
	err := AssertThat(counter, Is(2))
	if err != nil {
		t.Fatal(err)
	}
}
