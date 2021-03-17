package main


import (
	"fmt"
	"math/rand"
	"go-every-day/algorithm/sort/quicksort"
)

func main() {
	data := make([]int, 0)

	for i := 0; i < 50; i++ {
		data = append(data, rand.Intn(100))
	}

	fmt.Println(data)

	quicksort.QuickSort(data)

	fmt.Println()
	fmt.Println(data)
}
