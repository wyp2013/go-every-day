package mergesort

import "sync"

func MergeSort(data []int) {
	mergeSort(0, len(data)-1, data)
}


func mergeSort(l, r int, data [] int) {
	mid := (l + r)/2

	wait := &sync.WaitGroup{}
	if mid > l {
		go func() {
			wait.Add(1)
			mergeSort(l, mid, data)
			wait.Done()
		}()
	}

	if mid + 1 < r {
		go func() {
			wait.Add(1)
			mergeSort(mid+1, r, data)
			wait.Done()
		}()
	}

	wait.Wait()
	merge(l, mid, r, data)
}

func merge(l, mid, r int, data []int) {
	for i, k := l, mid + 1 ; i <= mid && k <= r ; {
	}
}