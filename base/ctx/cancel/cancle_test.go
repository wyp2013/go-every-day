package cancel

import (
	"context"
	"fmt"
	"testing"
)


func genInt() <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()

	return ch
}

func genWithCancle(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func TestGen(t *testing.T) {
	for n := range genInt() {
		fmt.Println(n)
	}

}

func TestCancel(t *testing.T) {
	ctx := context.Background()
	newctx, cancle := context.WithCancel(ctx)
	defer cancle()


	go func() {
		for n := range genWithCancle(newctx) {
			fmt.Println("gen func is runing ", n)

			if n == 5 {
				cancle()
				break // 没有这个会deadlock
			}
		}

		fmt.Println("gen func exit ")
	}()

	go func() {

	}()


}