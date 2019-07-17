package test

import (
	"bufio"
	"fmt"
	"go-every-day/base/files/filecompare/model"
	"go-every-day/base/files/filecompare/struct"
	"io"
	"os"
	"strings"
	"testing"
)

func getFindKeys(filePath string) map[string]int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	keyMap := make(map[string]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		keyMap[string(line)] = 1
	}

	return keyMap
}

func getOutputPath() {

}

func TestTransConstructFromFile(t *testing.T) {
	transModel := model.TransModel{}

	computTransMap, _, err := transModel.ConstructFromFile("/Users/bitmain/gowork/src/go-every-day/base/files/output/trans.csv", ",")
	if err != nil {
		t.Failed()
		return
	}

	standTransMap, keys, err := transModel.ConstructFromFile("/Users/bitmain/gowork/src/go-every-day/base/files/output/stand.txt", " ")
	if err != nil {
		t.Failed()
		return
	}

	keyMap := getFindKeys("/Users/bitmain/gowork/src/go-every-day/base/files/output/foundkeys.txt")
	fmt.Println(len(computTransMap), len(standTransMap), len(keyMap))


	rersult := make([]*_struct.Transcation, 0)
	for _, key := range keys {
		ctrans, ok := computTransMap[key]

		if !ok {
			//str, _ := util.ObjToStr();
			st := standTransMap[key]
			rersult = append(rersult, &_struct.Transcation{key, st.TotalSize, "Wallet"})
			continue
		} else {
			size := ctrans.TotalSize
			if size > standTransMap[key].TotalSize {
				size = standTransMap[key].TotalSize
			}
			rersult = append(rersult, &_struct.Transcation{key, size, ctrans.ConfirmedName})
		}
	}

	outputName := "/Users/bitmain/gowork/src/go-every-day/base/files/output/compare.csv"
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Println(err.Error())
		t.Failed()
		return
	}
	defer outFile.Close()

	header := []byte("txhash,total_size,from_pusher\r\n")
	outFile.Write(header)

	for _, res := range rersult {
		str := fmt.Sprintf("%s,%d,%s\r\n", res.Txhash, res.TotalSize, res.ConfirmedName)
		outFile.Write([]byte(str))
	}
}

func TestStringSplit(t *testing.T) {
	seps := strings.Split("79e4cbe557a108440d54591431d43378d57bd71760bd8bc87c87d691a4b52663 btccom             225 469483 *", " ")
	fmt.Println(len(seps))
}
