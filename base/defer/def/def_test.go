package def

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)


// 会 panic 两次
func TestDefer(t *testing.T) {
	defer fmt.Println("in main")
	defer func() {
		fmt.Println("xxxx")
		panic("panic again")
	}()

	panic("panic once")
}


func TestGetRecord(t *testing.T) {
	getRecordFunc := func(filePath string) map[string]int {
		recordMap := make(map[string]int, 0)
		file, err := os.Open(filePath)
		if err != nil {
			return recordMap
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			// line = strings.TrimSpace(line)
			subs := strings.Split(line, " ")
			if len(subs) > 2 {
				recordMap[subs[1]]++
			}
		}

		return recordMap
	}

	recordMap := getRecordFunc("test.txt")


	for key, val := range recordMap {
		fmt.Println(key, val)
	}
}

// 关闭一个空channel会panic
func TestCloseChannel(t *testing.T) {
	var stopChan chan struct{}
	close(stopChan)
}

func TestSendChannel(t *testing.T) {
	sendCeaned :=  make(chan int, 1)

	go func() {
		for {
			select {
			case <-time.After(time.Duration(1) * time.Second):
				fmt.Println("timer 1 is coming")

			case <-sendCeaned:
				fmt.Println("timer 1 is closed")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Duration(3) * time.Second):
				fmt.Println("timer 2 is coming")

			case <-sendCeaned:
				fmt.Println("timer 2 is closed")
				return
			}
		}
	}()

	time.Sleep(time.Duration(10) * time.Second)
	close(sendCeaned)

	select {
	case <-time.After(time.Duration(10) *  time.Second):
		return
	}
}

func readWitheList(filePath string)  []string {
	whiteList := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		return whiteList
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.Replace(line, "\n", "", -1)
		line = strings.Replace(line, " ", "", -1)
		fmt.Print("xxxx: ", line)

		whiteList = append(whiteList, line)
	}

	fmt.Println(whiteList)
	return whiteList
}


func TestRead(t *testing.T) {
	readWitheList("test.txt")


}