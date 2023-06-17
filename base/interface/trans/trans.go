package main

import (
	"errors"
	"fmt"
)

type Job interface {}

type Param struct {
	X int
	Y int
}


func Process(job Job) error {
	p, ok := job.(*Param)
	if !ok {
		return errors.New("type error")
	}

	fmt.Println(p)
	return nil
}

func main() {
	p := &Param{X:1, Y:2}

	err := Process(p)
	if err != nil {
		fmt.Println(err.Error())
	}
}