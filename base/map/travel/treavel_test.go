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


func TestMapSlice(t *testing.T) {
	tMapSlice := make(map[string][]string, 0)

	tMapSlice["key"] = append(tMapSlice["key"], "yyy")

	fmt.Println(tMapSlice)

	var tM map[string]string

	for k,v := range tM {
		fmt.Println(k, v)
	}
}

func TestReflect(t *testing.T) {
	// string type
	var1 := "hello world"

	// integer
	var2 := 10

	// float
	var3 := 1.55

	// boolean
	var4 := true

	// shorthand string array declaration
	var5 := []interface{}{"foo", "bar", "baz"}

	// map is reference datatype
	var6 := map[int]interface{}{100: "Ana", 101: "Lisa", 102: "Rob"}

	// complex64 and complex128
	// is basic datatype
	var7 := complex(9, 15)

	var8 :=[3]int{1,2,3}

	// using %T format specifier to
	// determine the datatype of the variables

	fmt.Println("Using Percent T with Printf")
	fmt.Println()

	fmt.Printf("var1 = %T\n", var1)
	fmt.Printf("var2 = %T\n", var2)
	fmt.Printf("var3 = %T\n", var3)
	fmt.Printf("var4 = %T\n", var4)
	fmt.Printf("var5 = %T\n", var5)
	fmt.Printf("var6 = %T\n", var6)
	fmt.Printf("var7 = %T\n", var7)
	fmt.Printf("var8=%T\n", var8)

	// using TypeOf() method of reflect package
	// to determine the datatype of the variables
	fmt.Println()
	fmt.Println("Using reflect.TypeOf Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.TypeOf(var1))
	fmt.Println("var2 = ", reflect.TypeOf(var2))
	fmt.Println("var3 = ", reflect.TypeOf(var3))
	fmt.Println("var4 = ", reflect.TypeOf(var4))
	fmt.Println("var5 = ", reflect.TypeOf(var5))
	fmt.Println("var6 = ", reflect.TypeOf(var6))
	fmt.Println("var7 = ", reflect.TypeOf(var7))
	fmt.Println("var8 = ", reflect.TypeOf(var8))

	// using ValueOf() method of reflect package
	// to determine the value of the variable
	// Kind() method returns the datatype of the
	// value fetched by the ValueOf() method
	fmt.Println()
	fmt.Println("Using reflect.ValueOf.Kind() Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.ValueOf(var1).Kind())
	fmt.Println("var2 = ", reflect.ValueOf(var2).Kind())
	fmt.Println("var3 = ", reflect.ValueOf(var3).Kind())
	fmt.Println("var4 = ", reflect.ValueOf(var4).Kind())
	fmt.Println("var5 = ", reflect.ValueOf(var5).Kind())
	fmt.Println("var6 = ", reflect.ValueOf(var6).Kind())
	fmt.Println("var7 = ", reflect.ValueOf(var7).Kind())
	fmt.Println("var7 = ", reflect.ValueOf(var8).Kind())

	v := reflect.ValueOf(var6).Kind()
	fmt.Println(v.String())
}


