package transmit

import (
	"fmt"
	"testing"
)

func TestCreateCancel(t *testing.T) {
	cancelMap := CreateCancel()

	fmt.Println(cancelMap)
}
