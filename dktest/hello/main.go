package hello

import (
	"context"
	"fmt"
	"os"
	"time"
)

func write(content string) {
	file, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	file.Write(([]byte)(content))
}

func main() {
	fmt.Println("hello world")
	write("test\n")

	context.WithTimeout(context.Background(), time.Second * 3)
}
