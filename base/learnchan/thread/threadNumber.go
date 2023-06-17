package thread

import (
	"fmt"
	"runtime"
	"sync"
)


func taskDo(x int, taskChan chan int) {
	for i := 0; i < 1000000; i++ {
	}

	k := <-taskChan
	fmt.Println(k, x, runtime.NumGoroutine())
}

func ControlMaxThreadNum() {
	cChan := make(chan int, 2)
	wait := sync.WaitGroup{}

	for i := 0; i < 100;  i++ {
		cChan <- i

		go func(x int) {
			wait.Add(1)
			taskDo(x, cChan)
			wait.Done()
		}(i)
	}

	wait.Wait()
}

