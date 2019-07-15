package model

import (
	"bufio"
	"errors"
	"fmt"
	"go-every-day/base/files/filecompare/struct"
	"go-every-day/base/files/filecompare/util"
	"io"
	"os"
	"strconv"
	"strings"
)

type TransModel struct {
}

func (t *TransModel) ConstructFromFile(filePath string, separator string) (map[string]*_struct.Transcation, []string, error) {
	exist, err := util.Exists(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}
	if !exist {
		return nil, nil, errors.New("file not exist")
	}

	if len(separator) == 0 {
		separator = " " // 默认空格
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	reader := bufio.NewReader(file)
	transMap := make(map[string]*_struct.Transcation)
	keys := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		str := string(line)
		seps := strings.Split(str, separator)
		if len(seps) == 8 {
			// total size
			size, err := strconv.Atoi(seps[1])
			if err != nil {
				fmt.Println(str, err.Error())
				continue
			}

			transMap[seps[0]] = &_struct.Transcation{
				Txhash: seps[0],
				TotalSize: size,
				ConfirmedName: seps[5],
			}
			keys = append(keys, seps[0])
		} else if len(seps) > 8 {
			size := 0
			for i := 2; i < len(seps); i++ {
				size, err = strconv.Atoi(seps[i])
				if err != nil {
					continue
				} else {
					break
				}
			}

			if size > 0 {
				transMap[seps[0]] = &_struct.Transcation{
					Txhash: seps[0],
					TotalSize: size,
					ConfirmedName: seps[1],
				}
				keys = append(keys, seps[0])
			} else {
				fmt.Println(str)
			}
		} else {

		}
	}

	fmt.Println(util.ObjToStr(transMap))

	return transMap, keys, nil
}
