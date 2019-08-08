package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := make([]int, 0)

	for i := 0; i < 50; i++ {
		data = append(data, rand.Intn(100))
	}

	fmt.Println(data)

	QuickSort(data)

	fmt.Println()
	fmt.Println(data)
}
