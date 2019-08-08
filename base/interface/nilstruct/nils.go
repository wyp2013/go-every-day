package main

import (
	"fmt"
	"unsafe"
)

func nonNil(v interface{}) {
	if v == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("non nil")
	}
}

type TestStruct struct {}

type itab struct { // 32 bytes
	inter * int
	_type * int
	hash  uint32 // copy of _type.hash. Used for type switches.
	x     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

func main() {
	var s *TestStruct
	nonNil(s)

	fmt.Println(unsafe.Sizeof(itab{}._type))
}
