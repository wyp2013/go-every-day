package str

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStr(t *testing.T) {
	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)
}
