package str

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

func TestStruct(t *testing.T) {
	s := struct {
	}{}

	fmt.Println(unsafe.Sizeof(s))
}

func TestStr(t *testing.T) {
	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)
}

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}

	var buff bytes.Buffer
	s := 0
	e := 1
	for ;; {
		if e >= len(S) {
			str := fmt.Sprintf("%c%d", S[s], e-s)
			buff.WriteString(str)
			break
		}

		if S[s] == S[e] {
			e++
		} else {
			str := fmt.Sprintf("%c%d", S[s], e-s)
			buff.WriteString(str)
			s = e
			e++
		}
	}

	res := buff.String()
	if len(res) < len(S) {
		return res
	}

	return S
}

func TestCompressString(t *testing.T) {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))

}


func TestSwitch(t *testing.T) {
	x := 1
	switch x {
	case 1:
	case 2:
		fmt.Print("xxxyyy")
	case 10:
	case 11:
	default:
		fmt.Print("xxxxx")
	}
}


func TestSwitchb(t *testing.T) {
	var x *int
	y := 5
	x = &y

	fmt.Print(x, &y, reflect.TypeOf(&x))
}
