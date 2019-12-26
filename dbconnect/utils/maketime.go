package utils

import (
	"fmt"
	"time"
)

func MakeTime(layout string, year, moth, day, hour, minute, second int) (time.Time, error) {
	// todo judge the params
	destTime := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, moth, day, hour, minute, second)
	return time.Parse(layout, destTime)
}
