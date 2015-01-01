package mock

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/eventbus"
)

func TestImplements(t *testing.T) {
	eventBus := New()
	var i *eventbus.EventBus
	err := AssertThat(eventBus, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
