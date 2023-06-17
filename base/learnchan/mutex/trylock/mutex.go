package trylock

import (
	"context"
	"time"
)

type Mutex struct {
	ch chan struct {}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	<- m.ch
}

func (m *Mutex) TryLock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m *Mutex) TrylockWithTime(t time.Duration) bool {
	timer := time.NewTimer(t)

	select {
	case m.ch <- struct{}{}:
		return true
	case <- timer.C:
		return false
	}
}

func (m *Mutex) TrylockWithTime2(t time.Duration) bool {
	ctx, cancle := context.WithTimeout(context.TODO(), t)
	defer cancle()

	select {
	case m.ch <- struct{}{}:
		return true
	case <- ctx.Done():
		return false
	}
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) > 0
}
