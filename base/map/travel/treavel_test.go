package travel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapType(t *testing.T) {
	atc := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(atc), reflect.TypeOf(atc))

	stc := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(stc))

	st2 := make([]int, 0)
	fmt.Println(reflect.TypeOf(st2))


	st := make(map[int]int, 0)
	fmt.Println(reflect.TypeOf(st))



	fmt.Println("xxxxx")
}

func TestMapTravel(t *testing.T) {
	atc := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(atc), reflect.TypeOf(atc))

	stc := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(stc))

	st2 := make([]int, 0)
	fmt.Println(reflect.TypeOf(st2))


	st := make(map[int]int, 0)
	fmt.Println(reflect.TypeOf(st))



	fmt.Println("xxxxx")
}

func TestMapDelete(t *testing.T) {
	fmt.Println("xxxxxx")
}
