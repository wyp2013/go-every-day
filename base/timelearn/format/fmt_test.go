package format

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	str := time.Unix(1438608792, 0).Format("2006-01-02 15:04:05")
	str1 := time.Unix(1438270158, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str1, str)
}
