package util

import "encoding/json"

func Obj2String(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
