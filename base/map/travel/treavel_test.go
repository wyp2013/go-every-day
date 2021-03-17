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


func printMap(m map[string]interface{}) {
	for key, val := range m {
		fmt.Println(key, val)
	}

	fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
}

func TestMapValueInterface(t *testing.T) {
	tMap := make(map[string]interface{}, 0)
	tMap["xxxx"] = 1.5
	tMap["yyyy"] = 6

	tMap["struct"] = struct {
		X int
		Y string
	}{X: 1, Y: "test"}
	tMap["ysdf"] = 7
	tMap["zzzzz"] = 8

	printMap(tMap)
	printMap(tMap)
}

func TestMapDelete(t *testing.T) {
	fmt.Println("xxxxxx")

	tMap := make(map[string]interface{}, 0)
	tMap["xxxx"] = 1.5
	tMap["yyyy"] = 6
	delete(tMap, "xxxx")
}

func TestMapTrap(t *testing.T) {
	fmt.Println("xxxxxx")

	tMap := make(map[string]int, 0)
	tMap["xxxx"] = 1
	tMap["yyyy"] = 6
	tMap["zzzz"]  = 10

	ttMap := make(map[string]*int, 0)
	for key, val := range tMap {
		ttMap[key] = &val
	}

	for key, val := range ttMap {
		fmt.Println(key, *val)
	}

}
