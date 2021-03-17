package toom

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 这个程序会一直运行，timer被gc回收，并且还不停的创建timer
func TestUseTimer(t *testing.T) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("before, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	closeCh := UseTimer(1)

	time.Sleep(20 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after 3min, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	select {
	case <-closeCh:
		break
	}

	fmt.Println("over")
}

// 正确的程序
func TestUseTimer2(t *testing.T) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("before, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	closeCh := UseTimer2(1)

	time.Sleep(20 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after 3min, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	select {
	case <-closeCh:
		break
	}

	fmt.Println("over")
}
