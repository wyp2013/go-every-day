package visit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFilesAndDir(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			fmt.Println(dirPth + PthSep + fi.Name())
		} else {
			fmt.Println(dirPth + PthSep + fi.Name())
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}


func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			fmt.Println(dirPth + PthSep + fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			fmt.Println(dirPth + PthSep + fi.Name())
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}



type OrderDetail struct {
	Id             int     `json:"id"`
	Pid            int     `json:"pid"`
	FileType       int     `json:"file_type"`
	ProductName    string  `json:"product_name"`
	FilePath       string  `json:"file_path"`
	FileName       string  `json:"file_name"`
	OrderName      string  `json:"order_name"` 
	Content        string  `json:"content"`
}

func ReadFile(filepath string) (string, error) {
	content ,err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	content = []byte{}
	return string(content), nil
}

type Node struct {
	Id             int     `json:"id"`
	Pid            int     `json:"pid"`
	FileType       int     `json:"fileType"`
	ProductName    string  `json:"productName"`
	Path           string  `json:"path"`
	FileName       string  `json:"fileName"`
	OrderName      string  `json:"orderName"`
	Children       []*Node `json:"children"`
}

func NewNode(id, pId, fileType int, productName, path, fileName, orderName string, children []*Node) *Node {
	return &Node{
		Id:          id,
		Pid:         pId,
		FileType:    fileType,
		ProductName: productName,
		Path:        path,
		FileName:    fileName,
		OrderName:   orderName,
		Children:    children,
	}

}

func WalkVisitProduct(productName, dirPth string, parentId *int, ods *[]OrderDetail, parent *Node) error {
	if parent == nil {
		panic("empty")
	}

	fis, err := ioutil.ReadDir(dirPth)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	pId := *parentId
	PthSep := string(os.PathSeparator)
	for _, fi := range fis {
		if fi.IsDir() {
			subPath := dirPth + PthSep + fi.Name()
			fmt.Println(subPath)

			*ods = append(*ods, OrderDetail{
				Id:          0,
				Pid:         pId,
				FileType:    1,
				ProductName: productName,
				FilePath:    subPath,
				FileName:    fi.Name(),
				OrderName:   fi.Name(),
				Content:     "",
			})

			child := NewNode(0, pId, 1, productName, fmt.Sprintf("/%d", 0), fi.Name(), fi.Name(), nil)
			parent.Children = append(parent.Children, child)
			*parentId = *parentId + 1
			WalkVisitProduct(productName, subPath, parentId, ods, child)
		} else {
			ok := strings.HasSuffix(fi.Name(), ".md")
			if ok {
				filePath := dirPth + PthSep + fi.Name()
				fmt.Println(filePath)

				content, er := ReadFile(filePath)
				if er != nil {
					err = er
					//return
				}

				fileName := fi.Name()
				*ods = append(*ods, OrderDetail{
					Id:          0,
					Pid:         pId,
					FileType:    2,
					ProductName: productName,
					FilePath:    filePath,
					FileName:    fileName,
					OrderName:   fileName[0:len(fileName)-3],
					Content:     content,
				})

				child := NewNode(0, pId, 1, productName, fmt.Sprintf("/%d", 0), fi.Name(), fi.Name(), nil)
				parent.Children = append(parent.Children, child)
			}
		}
	}

	return nil
}
