package modify

import (
	"fmt"
	"testing"
)

type duck interface {
	walk()
	modify(str string)
}

type dog struct {
	sound string
}

func (d dog) walk() {
	fmt.Println("dog walk ", d.sound)
}

func (d dog) modify(s string) {
	d.sound = s
}

func TestModify(t *testing.T) {
	d := dog{sound:"wang wang wang"}

	var x duck = &d

	x.modify("xixixi") // y=*x; y.modify("xixixi");
	x.walk() // so x 并没有改变
	d.walk()
}