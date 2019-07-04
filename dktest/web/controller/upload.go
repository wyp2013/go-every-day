package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"go-every-day/dktest/web/common"
	"go-every-day/dktest/web/model"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func UploadFile(c echo.Context) error {
	// idStr := c.QueryParam("id")
	cc := c.(*common.CustomContext)
	reader, err := cc.Request().MultipartReader()
	if err != nil {
		return cc.SendRspErrMsg(1, err.Error())
	}

	keyValMap := make(map[string]string, )
	fileNameList := make([]string, 0)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
		if part.FormName() == "files" {
			dst, err := os.Create("./" + part.FileName())
			if err != nil {
				fmt.Println(fmt.Sprintf("create file failed, filename: %s", part.FileName))
				continue
			}
			defer dst.Close()
			io.Copy(dst, part)

			fileInfo, err := dst.Stat()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println("file ", fileInfo.Name())
			fileNameList = append(fileNameList, fileInfo.Name())

		} else {
			fmt.Println("from ", part.FormName())
			data, err := ioutil.ReadAll(part)
			if err != nil {
				fmt.Println("part.FormName()")
			}
			fmt.Printf("key=[%s] value=[%s]\n", part.FormName(), string(data))
			keyValMap[part.FormName()] = string(data)
		}
	}

	fmt.Println(common.ObjToStr(keyValMap))

	//if _, ok := keyValMap["from"]; !ok {
	//	return cc.SendRspErrMsg(1, "from参数不存在")
	//}

	if _, ok := keyValMap["subject"]; !ok {
		return cc.SendRspErrMsg(1, "subject参数不存在")
	}

	if _, ok := keyValMap["text"]; !ok {
		return cc.SendRspErrMsg(1, "text参数不存在")
	}

	to, ok := keyValMap["to"]
	if !ok {
		return cc.SendRspErrMsg(1, "text参数不存在")
	}

	toList := strings.Split(to, ",")
	if len(toList) == 0 {
		return cc.SendRspErrMsg(1, "to参数错误")
	}

	fmt.Println(fileNameList)
	mailgun := model.NewMaigunModel("sandboxbd75d75a2de54f9c9c2a49c4226ed396.mailgun.org", "08ae36004944c7e466090aa158c14d6b-2b778fc3-a2840547")
	id, err := mailgun.SendMessage(
		"Excited User <mailgun@sandboxbd75d75a2de54f9c9c2a49c4226ed396.mailgun.org>",
		keyValMap["subject"],
		keyValMap["text"],
		toList,
		fileNameList)
	if err != nil {
		fmt.Println(err.Error())
		cc.SendRspErrMsg(1, err.Error())
	} else {
		fmt.Println(id)
	}

	return  cc.SendRspOK("ok")
}
