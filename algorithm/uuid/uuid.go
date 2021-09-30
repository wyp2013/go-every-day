package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {

	tF := ""
	tx := ""
	ah := make(map[string]bool)
	for i := 0;  i < 5000000; i++ {
		tF = uuid.New().String()
		ah[tF] = true

		if i == 3000000 {
			tx = tF
		}
	}

	fmt.Println(len("bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"))

	now := time.Now().UnixNano()
	if _, ok := ah[tx]; ok {
		fmt.Println("find")
	} else {

	}

	fmt.Println(float64(time.Now().UnixNano()-now)/1000000.0)
}


