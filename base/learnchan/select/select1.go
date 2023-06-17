package _select

import (
	"context"
	"fmt"
	//"sync"
	"time"
)

func writeChan(in chan int) {
	for i :=0; ; i++ {
		in <- i
	}
}

func Select_default() {
	c1 := make(chan int,)
	c2 := make(chan string,)

	go writeChan(c1)
	for {
		select {
		case x, ok := <-c1:
			if !ok {
				fmt.Println("c1 is clone")
				goto final
			}

			fmt.Println(x)

		case s, ok := <-c2:
			if !ok {
				fmt.Println("c2 is clone")
				goto final
			}

			fmt.Println(s)
		default:
			fmt.Println("default")
		}
	}

	final:
		fmt.Println("final")

	fmt.Println("over")
}

func doSomething() chan struct{} {
	done := make(chan struct{})
	for i:=0; i<10; i++ {
		fmt.Println("test")
	}

	return done
}

func selectFun() {
	done := make(chan struct{})
	ctx, _:= context.WithTimeout(context.TODO(), time.Duration(3) * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exec time out return ")
				return

			case <-done:
				// todo

				fmt.Println("do something done")
				return
			}
		}
	}()

	select {
	case <-done:
		fmt.Println("testing  test")
		mp := make(map[string]int)
		mp["test"] = 1

	case <-ctx.Done():

 	}
}


func testFunc() {
	//  todo
	fmt.Println("test do something")
}

func getSourcePath(volume  string) string {
	return volume
}

func getBindPath(volume string) string {
	return volume
}

func bindVolume(sourcePath, mountPath string) error {
	return nil
}

func bindAll(allVolumes []string) {
	// wg := &sync.WaitGroup{}
	for _, volume := range allVolumes {
		sourcePath := getSourcePath(volume)
		bindPath := getBindPath(volume)

		newBindVolume(sourcePath, bindPath, 1)
	}
}

func newBindVolume(source string, target string, timeout int)  {
	done :=  make(chan struct{})
	ctx, _:= context.WithTimeout(context.TODO(), time.Duration(1))

	go func() {
		bindVolume(source, target)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		return
	}
}


