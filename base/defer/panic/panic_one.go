package panic

import (
	"fmt"
	 "time"
)

/*
panic 影响学习
1. panic 只能在同一个goroutine中被捕捉
2. Go 语言在发生 panic 时只会执行当前协程中的 defer 函数
3.
*/

func Panic() {
	panic("panic test")
}

func taskDo(taskName string) {
	defer func() {
		fmt.Println("defer is run or not")
	}()

	//// 如果这里不捕捉panic 会影响调用他的线程
	defer func() {
		fmt.Println("panic")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if taskName == "task1" {
		Panic() // 与直接 panic("panic test")一样，但是换成 go Panic 就不一样了
		// go Panic() // 程序会直接崩溃
	} else {
		fmt.Println(taskName)
	}
}

func TimeDoTask() {
	timer := time.NewTicker(time.Duration(2) * time.Second)
	tasks := []string{"task1", "task2"}

	timeOut := time.After(time.Duration(10) * time.Second)
	for {
		select {
		case <-timer.C:
			for _, task := range tasks {
				//这里不能扑捉到 taskDo的panic
				defer func() {
					fmt.Println("for defer")
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				go taskDo(task) // 此goroutine中有defer recover
			}

		case <-timeOut:
			fmt.Println("time out")
			goto timeout
		}
	}

	timeout:
		fmt.Println("over")

}

