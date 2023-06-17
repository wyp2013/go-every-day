package time

import (
	"context"
	"fmt"
)

// goroutine超时控制
func DoWithContext(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out return")
			return
		case n, ok := <-ch:
			if !ok {
				fmt.Println("chan is closed")
				return
			}

			fmt.Println("receive from chan:", n)
		}
	}
}


