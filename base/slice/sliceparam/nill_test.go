package sliceparam

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNilSlice(t *testing.T) {
	var emtptySlice []int

	if emtptySlice == nil {
		fmt.Println("empty slice")
	}

	emtptySlice = []int{}
	fmt.Println(reflect.TypeOf(emtptySlice).String())

	if emtptySlice == nil {
		fmt.Println("empty slice")
	}

}
