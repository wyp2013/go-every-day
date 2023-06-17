package orchan

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Data int `json:"data"`
}

func NewStruct(data int) TestStruct {
	return TestStruct{Data: data}
}


func TestOrChan(t *testing.T) {
	var chans []chan interface{}
	chans = append(chans, make(chan interface{}))
	chans = append(chans, make(chan interface{}))
	chans = append(chans, make(chan interface{}))

	out := OrChan(chans...)

	go func() {
		chans[0] <- TestStruct{Data:1}
	}()

	go func() {
		chans[1] <- TestStruct{Data:2}
	}()

	go func() {
		chans[2] <- TestStruct{Data:3}
	}()

	select {
	case <-out:
		fmt.Println("over")
	}
}
