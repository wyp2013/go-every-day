package trylock

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type TestMutex struct {
	mu   *Mutex
	Data int
}

func Write(t *TestMutex) {
	for ;; {
		t.mu.Lock()
		fmt.Println("befor write", t.Data)
		t.Data = rand.Intn(1000)
		fmt.Println("after write", t.Data)
		t.mu.Unlock()
	}
}


func Read(t *TestMutex) {
	for ;; {
		if t.mu.TryLock() {
			fmt.Println("read ", t.Data)
			t.mu.Unlock()
		}
	}
}

func ReadTimeOut(t *TestMutex, ti time.Duration) {
	for ;; {
		if t.mu.TrylockWithTime(ti) {
			fmt.Println("read ", t.Data)
			t.mu.Unlock()
		}
	}
}

func TestMutex_Lock(t *testing.T) {
	test := &TestMutex{
		mu : NewMutex(),
	}

	go Write(test)
	go Read(test)

	select {

	}
}

func TestMutex_TrylockWithTime(t *testing.T) {
	test := &TestMutex{
		mu : NewMutex(),
	}

	go Write(test)
	go ReadTimeOut(test, 100 * time.Millisecond)

	select {

	}
}
