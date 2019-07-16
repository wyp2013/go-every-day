package util

import "encoding/json"

func ObjToStr(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func StrToObj(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	return err
}