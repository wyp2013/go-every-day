package _select

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSelect_default(t *testing.T) {
	Select_default()
}

func TestSelect(t *testing.T) {
	ctx, _:= context.WithTimeout(context.Background(), time.Duration(10) *time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("time is coming")
			return
		default:
			time.Sleep(time.Duration(1) *time.Second)
			fmt.Println("default is call")
			time.Sleep(time.Duration(20) *time.Second)
		}
	}

}

func doBadThing(done chan bool) {
	time.Sleep(time.Duration(1) *  time.Second)
	done <- true
}

func timeoutFunc(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)

	tick := time.NewTicker(time.Duration(1) * time.Millisecond)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-tick.C:
		return fmt.Errorf("timeout")
	}
}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()

	for  i:=0; i < 1000; i++ {
		timeoutFunc(f)
	}

	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestDoBadThing(t *testing.T) {
	test(t, doBadThing)
}

func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool , 1)
	go f(done)

	tick := time.NewTicker(time.Duration(1) * time.Millisecond)
	select {
	case <-done:
		fmt.Println("done ")
		return nil
	case <-tick.C:
		 fmt.Println("timeout ")
		 return fmt.Errorf("timeout")
	}
}


func TestTimeoutWithBuffer(t *testing.T) {
	 for i:=0; i < 1000; i++ {
	 	timeoutWithBuffer(doBadThing)
	 }
}

func doGoodThing(chan bool) {

}


func TestChannel(t *testing.T) {
	 dChan := make(chan int)

	 go func() {
	 	for {
	 		var data int
			select {
	 		case data = <- dChan:
	 			fmt.Println("go routine one receive data from dChan: ", data)
			}
		}
	 }()

	 go func() {
	 	for {
	 		var data int
			select {
	 		case data = <-dChan:
	 			fmt.Println("go routine two receive data from dChen:  ", data)
			}
		}
	 }()

	 for  i:=0; i<100; i++ {
	 	data := i+1000
	 	dChan <- data
	 	time.Sleep(time.Duration(1) * time.Second)
	 }
}

func TestBufferChannel(t *testing.T) {
	 bufChan := make(chan int, 10)

	 go func() {
	 	for {
	 		select {
	 		case data := <-bufChan:
	 			fmt.Println("go routine one receive data from bufChan: ", data)
			}
		}
	 }()

	 go func() {
	 	for {
			select {
			case data := <-bufChan:
				fmt.Println("go routine two receive data from bufChan: ", data)
			}
		}
	 }()

	 for i:=0; i<100; i++ {
	 	bufChan <- i
	 }

}

