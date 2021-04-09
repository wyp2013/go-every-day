package inter

import (
	"fmt"
	"testing"
)

type Cat struct {}

func (c Cat) Walk() {
	fmt.Println("cat walking")
}


type Dog struct {}
func (c *Dog) Walk() {
	fmt.Println("Dog walking")
}

func TestWalk(t *testing.T) {
	var duck Duck = Cat{}
	duck.Walk()
}

func TestWalkPointer(t *testing.T) {
	var duck Duck = &Cat{}
	duck.Walk()
}


// 结构体不能接受，指针定义的方法
func TestDogWalk(t *testing.T) {
	var duck Duck = Dog{}
	duck.Walk()
}

func TestDogWalkPointer(t *testing.T) {
	var duck Duck = &Dog{}
	duck.Walk()
}