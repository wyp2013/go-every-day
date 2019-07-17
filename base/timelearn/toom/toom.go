package toom

import (
	"fmt"
	"time"
)

// 创建过多的定时器导致消耗内存：select时每次都会创建time对象
func UseTimer(second int) chan bool {
	ch := make(chan int)
	// 防止goroutine泄漏
	exit := make(chan bool)
	go func() {
		i := 1000

		for {
			// 这里退出goroutine
			select {
			case <-exit:
				close(ch)
				return
			default:
				i++
				ch <- i
			}
		}

	}()

	go func() {
		cnt := 0 // 统计大概创建了多少个timer, 每一次select都会创建一次timer
		for {
			breakFor := false

			select {
			case _, ok := <-ch:
				if !ok {
					breakFor = true
					break
				}

				// fmt.Println(cnt, x)
			case <-time.After(time.Duration(second) * time.Second):
				fmt.Println("time is out")
				breakFor = true
				break
			default:
				cnt++
			}

			if breakFor {
				close(exit)
				fmt.Println("cnt ", cnt)
				break
			}
		}
	}()

	return exit
}


// 创建过多的定时器导致消耗内存：select时每次都会创建time对象
func UseTimer2(second int) chan bool {
	ch := make(chan int)
	// 防止goroutine泄漏
	exit := make(chan bool)
	go func() {
		i := 1000

		for {
			// 这里退出goroutine
			select {
			case <-exit:
				return
			default:
				i++
				ch <- i
			}
		}

	}()

	go func() {
		timer := time.NewTimer(time.Duration(second) * time.Second)
		cnt := 0 // 统计大概创建了多少个timer, 每一次select都会创建一次timer
		for {
			breakFor := false

			select {
			case _, ok := <-ch:
				if !ok {
					breakFor = true
					break
				}

				// fmt.Println(cnt, x)
			case <-timer.C:
				fmt.Println("time is out")
				breakFor = true
				break
			default:
				cnt++
			}

			if breakFor {
				close(exit)
				fmt.Println("cnt ", cnt)
				break
			}
		}
	}()

	return exit
}