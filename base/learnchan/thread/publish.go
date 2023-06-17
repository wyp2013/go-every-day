package thread

import (
	"fmt"
	"sync"
	"time"
)

type SubscribeType chan interface{}
type TopicType   func(interface{}) bool

type Publish interface {
	PublishMessage(msg interface{})
	SubscribeTopic(topic TopicType) chan interface{}
}

type Publisher struct {
	mutux       sync.Mutex
	timeout     time.Duration
	bufferSize  int
	subscribes  map[SubscribeType]TopicType
}

func NewPublisher(timeout time.Duration, bufferSize int) *Publisher {
	return &Publisher{
		timeout:    timeout,
		bufferSize: bufferSize,
		subscribes: make(map[SubscribeType]TopicType),
	}
}

func (pub *Publisher) PublishMessage(msg interface{}) {
	pub.mutux.Lock()
	defer pub.mutux.Unlock()

	wait := &sync.WaitGroup{}
	for sub, topic := range pub.subscribes {
		wait.Add(1)
		go pub.sendTopic(msg, topic, sub, wait)
	}

	wait.Wait()
}

func (pub *Publisher) SubscribeTopic(topic TopicType) chan interface{} {
	sub := make(chan interface{}, pub.bufferSize)

	pub.mutux.Lock()
	pub.subscribes[sub] = topic
	pub.mutux.Unlock()

	return sub
}

func (pub *Publisher) sendTopic(msg interface{}, topic TopicType, sub SubscribeType, wait *sync.WaitGroup) {
	defer wait.Done()

	if topic != nil && !topic(msg) {
		return
	}

	select {
	case sub <- msg:
		fmt.Println("send msg ", msg)
	case <-time.After(pub.timeout):
		fmt.Println("send timeout")
	}
}

func (pub * Publisher) Close() {
	pub.mutux.Lock()
	defer pub.mutux.Unlock()

	for sub := range pub.subscribes {
		delete(pub.subscribes, sub)
		close(sub)
	}
}

func (pub * Publisher) Retire(sub SubscribeType) bool {
	pub.mutux.Lock()
	defer pub.mutux.Unlock()

	if _, ok := pub.subscribes[sub]; ok {
		delete(pub.subscribes, sub)
		close(sub)
		return true
	}

	return false
}

