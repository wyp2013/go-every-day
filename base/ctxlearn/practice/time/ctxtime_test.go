package time

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// go
func TestDoWithContext(t *testing.T) {
	ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	defer func() {
		fmt.Println("call cxt.cancle()")
		cancle()
	}()

	intCh := make(chan int)
	go DoWithContext(ctx, intCh)

	i := 0
	loopFor:
	for {
		i++

		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done() is return, break for loop")
			// break // break只能跳出select分支，不能跳出for语句
			break loopFor
		default:
			intCh <- i
			time.Sleep(500 * time.Millisecond)
		}
	}
}
