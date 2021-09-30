package fileVisit

func GetFilesAndDir(dirPth string) []string {
	return nil
}

func isDir(fi string) bool {
	return true
}

// 根据配置文件路径，读取路径下的每个产品
func fileVisit(dirPth string) {
	// 获取目录下面所有产品目录 easydata、bml、ais
	dirs := GetFilesAndDir(dirPth)

	parent := 0
    for _, fi := range dirs {
    	if isDir(fi) {
			walkVisitProduct(fi, &parent)
		}
	}
}

// 数据库中新增一行菜单记录
func insertDb(path, fileName string, fileType int, parentId int) {
}

// 读取入库
func walkVisitProduct(dirPath string, parentId *int) {
	fis := GetFilesAndDir(dirPath)

	pId := *parentId
	for _, fi := range fis {
		if isDir(fi) {
			// 入库
			insertDb(dirPath+fi, fi, 2, pId)
			// 继续遍历
			*parentId = *parentId + 1
			walkVisitProduct(dirPath+fi, parentId)
		} else {
			// 入库
			insertDb(dirPath+fi, fi, 1, pId)
		}
	}
}


// 读取入库
func WalkVisitProduct(dirPath string, parentId *int) {
	fis := GetFilesAndDir(dirPath)

	pId := *parentId
	for _, fi := range fis {
		if isDir(fi) {
			// 入库
			insertDb(dirPath+fi, fi, 2, pId)
			// 继续遍历
			*parentId = *parentId + 1
			WalkVisitProduct(dirPath+fi, parentId)
		} else {
			// 入库
			insertDb(dirPath+fi, fi, 1, pId)
		}
	}
}