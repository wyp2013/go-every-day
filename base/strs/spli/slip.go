package spli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func Split() {
	filePath := "/Users/wuyupei/tmp/pods.txt"
	fi, err := os.Open(filePath)
	if err != nil {
		panic("open file failed")
	}

	podId := make(map[string]int)
	r := bufio.NewReader(fi)
	for {
		bt, _,  err:= r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		line := string(bt)

		if strings.Contains(line, "/kubelet/lib/kubelet/pods") {
			subStrs := strings.Split(line, " ")
			sps := strings.Split(subStrs[2], "/")
			pid := sps[5]

			if _, ok := podId[pid]; ok {
				podId[pid]++
			} else {
				podId[pid] = 1
			}
		}
	}

	strPod, _:= json.Marshal(podId)
	fmt.Println(string(strPod))
	total := 0
	for _, v := range podId {
		total = total + v
	}
	fmt.Println(total)
}
