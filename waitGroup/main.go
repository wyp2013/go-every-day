package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Second * 4)
			fmt.Println(num)
		}(i)
	}

	if WaitTimeout(&wg, time.Second*3) {
		fmt.Println("timeout exit")
	}

	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	tmp := make(chan int, )

	go func() {
		wg.Wait()
		close(tmp)
	}()

	select {
	case i, ok := <-tmp:
		if !ok {
			fmt.Println("channel tmp is close")
		}

		fmt.Println(i)
		return false
	case  <-time.After(timeout):
		return true
	}

	context.WithTimeout(context.TODO(), time.Second *10)

	return false
}


