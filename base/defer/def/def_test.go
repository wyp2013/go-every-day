package def

import (
	"fmt"
	"testing"
)


// 会 panic 两次
func TestDefer(t *testing.T) {
	defer fmt.Println("in main")
	defer func() {
		fmt.Println("xxxx")
		panic("panic again")
	}()

	panic("panic once")
}
