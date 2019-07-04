package _select

import "fmt"

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
