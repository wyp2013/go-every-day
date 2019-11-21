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

func Panic2() {
	panic("panic test2")
}

func taskDo2(taskName string) {
	if taskName == "task1" {
		Panic2() // 与直接 panic("panic test")一样，但是换成 go Panic 就不一样了
	} else {
		fmt.Println(taskName)
	}
}

func TimeDoTask2() {
	defer func() {
		fmt.Println("main defer")
	}()
	timer := time.NewTicker(time.Duration(2) * time.Second)
	tasks := []string{"task1", "task2"}

	timeOut := time.After(time.Duration(10) * time.Second)
	for {
		select {
		case <-timer.C:
			for _, task := range tasks {
				// 这里不能扑捉到 taskDo的panic
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				taskDo2(task) // 此goroutine中有defer recover
			}

		case <-timeOut:
			fmt.Println("time out")
			goto timeout
		}
	}

timeout:
	fmt.Println("over")

}
