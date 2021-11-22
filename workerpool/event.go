package workerpool

import "errors"

type EventType int
const (
	EventSaveMsg          EventType = 1
	EventDispatchMsg      EventType = 2
	EventTimer            EventType = 3
)

type EventHandle func(interface{}) error

type Event struct {
	EventType    EventType
	Param        interface {}
	HandleFunc   EventHandle
}


func DefaultProcessor(job Job) error {
	event, ok := job.(*Event)
	if !ok {
		return errors.New("job type is not event")
	}

	return event.HandleFunc(event.Param)
}




