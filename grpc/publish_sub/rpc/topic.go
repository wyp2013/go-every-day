package rpc

import "strings"

func subGo(topic interface{}) bool {
	str, ok := topic.(string)
	if !ok {
		return false
	}

	return strings.Contains(str, "go")
}


func subPython(topic interface{}) bool {
	str, ok := topic.(string)
	if !ok {
		return false
	}

	return strings.Contains(str, "python")
}