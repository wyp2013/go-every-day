package repl

import (
	"fmt"
	"regexp"
	"testing"
)

// 处理key没有加引号的情况
func TestRlace(t *testing.T) {
	str := "[{y:6912,formattedValue:'6,912',dt:'Thursday,July30,2015',reward:'34,762968'}, {y:6912,formattedValue:'6,912',dt:'Thursday,July30,2015',reward:'34,762968'}]"

	reg := regexp.MustCompile("(\\w+):")
	str = reg.ReplaceAllString(str, "\"$1\":")
	fmt.Println(str)
}
