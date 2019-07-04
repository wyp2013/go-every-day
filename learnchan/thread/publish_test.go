package thread

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type Event struct {
	body string
	id   int
}

func TestPublisher_PublishMessage(t *testing.T) {
	publish := NewPublisher(2 * time.Second, 3)

	subGo := publish.SubscribeTopic(func(msg interface{}) bool {
		str, ok := msg.(string)
		if !ok {
			return false
		}

		return strings.Contains(str, "go")
	})

	go func() {
		for msg := range subGo {
			fmt.Println(msg)
		}
	}()

	subEvent := publish.SubscribeTopic(func(event interface{}) bool {
		e, ok := event.(Event)
		if !ok {
			return false
		}

		return e.id == 1
	})

	go func() {
		for e := range subEvent {
			fmt.Println(e)
		}
	}()


	publish.PublishMessage("hello, go world")
	publish.PublishMessage("hello, world")
	publish.PublishMessage(Event{body:"hahahaha", id: 1})
	publish.PublishMessage(Event{body:"hahahaha", id: 2})
}
