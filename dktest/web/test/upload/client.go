package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	// file1
	fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
	file1, err := os.Open("/Users/bitmain/gowork/src/go-every-day/dktest/web/test/upload/file1.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file1.Close()
	io.Copy(fileWriter1, file1)

	// file2
	fileWriter2, _ := bodyWriter.CreateFormFile("files", "file2.txt")
	file2, err := os.Open("/Users/bitmain/gowork/src/go-every-day/dktest/web/test/upload/file2.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file2.Close()
	io.Copy(fileWriter2, file2)


	bodyWriter.WriteField("test", "lalalalallalal")
	bodyWriter.WriteField("hahaha", "xxxxxxxxx")
	bodyWriter.WriteField("subject", "测试发附件")
	bodyWriter.WriteField("text", "测试")
	bodyWriter.WriteField("to", "smtl_2013@163.com")


	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post("http://localhost:30000/test/file/upload", contentType, bodyBuffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}