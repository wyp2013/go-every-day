package visit

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllFile(t *testing.T) {
	GetFilesAndDir("/Users/wuyupei/gowork/src/go-every-day/fileVisit")
}

func TestGetFilesAndDirs(t *testing.T) {
	GetFilesAndDirs("/Users/wuyupei/gowork/src/go-every-day/fileVisit")
}

func TestGetFilesAndDirs2(t *testing.T) {
	GetFilesAndDirs("/Users/wuyupei/work/markdown")
}

func TestWalkVisitProduct(t *testing.T) {
	var p int = 0
	var ods []OrderDetail
	root := NewNode(0, 0, 1, "easydata", "", "easydata", "easydata", nil)

	WalkVisitProduct("EasyData3.5.2", "/Users/wuyupei/work/markdown/easydata/3.5.1", &p, &ods, root)

	data, _ := json.Marshal(ods)
	fmt.Println(string(data))

	nodeData, _ := json.Marshal(*root)
	fmt.Println(string(nodeData))
}
