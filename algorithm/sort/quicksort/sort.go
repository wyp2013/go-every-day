package quicksort

import "sync"

func QuickSort(data []int) {
	quickSort(0, len(data) - 1, data)
}

func quickSort(l, r int, data []int) {
	pos := findPos(l, r, data)

	wait := &sync.WaitGroup{}
	if pos - 1 > l {
		go func() {
			wait.Add(1)
			quickSort(l, pos-1, data)
			wait.Done()
		}()
	}

	if pos + 1 < r {
		go func() {
			wait.Add(1)
			quickSort(pos+1, r, data)
			wait.Done()
		}()
	}

	wait.Wait()
}

func findPos(l, r int, data []int) int {
	target := data[l]

	for ; l < r ; {
		for ; target < data[r]  && r > l ; r-- {}
		data[l] = data[r]

		for ; target >= data[l] && l < r ; l++ {}
		data[r] = data[l]
	}

	data[l] = target
	return l
}
