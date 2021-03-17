package timeout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func doA() {}
func doB() {}
func doC() {}

// do something within 2 second
func doWithTimeout(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		doA()
		doB()
		doC()
		time.Sleep(3*time.Second)
		done <-struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Print("超时")
		return
	case <-done:
		fmt.Print("success done")
	}
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2 * time.Second)
	defer cancel()

	doWithTimeout(ctx)
}
