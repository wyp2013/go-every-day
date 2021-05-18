package json_parse

import (
	"fmt"
	parser "github.com/buger/jsonparser"
)

func parser_json() {
	strJson := []byte(`{"optionsxxx":{"id":"3da3b1791aaa","conjunction":"and","children":[{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"sss","id":"2c8d28ea195b"},{"id":"435367d4c3da","conjunction":"and","children":[{"id":"a64cdac5b0b4"}]},{"id":"fa62a37b02cd","left":{"type":"field","field":"shifei"}}]}}`)
	id, _, _, _:= parser.Get(strJson, "optionsxxx", "id")
	fmt.Println(string(id))


	parser.ArrayEach(
		strJson,
		func(value []byte, dataType parser.ValueType, offset int, err error) {

			// 寻找left
			left, _, _, er := parser.Get(value, "left")
			if er == nil {
				// do something, 调用处理left的函数
				fmt.Println(string(left))
			} else {
				fmt.Println(er.Error())
			}

			// 寻找 children
			children, _, _, er := parser.Get(value, "children")
			if er == nil {
				// todo 递归处理children
				fmt.Println(string(children))
			} else {
				// do nothing
			}

			// 其它字段
		}, "optionsxxx", "")
}

func jsonParseArray(data[]byte) {
	parser.ArrayEach(
		data,
		func(value []byte, dataType parser.ValueType, offset int, err error) {

			// 寻找left
			left, _, _, er := parser.Get(value, "left")
			if er == nil {
				jsonParse(left)
			} else {
				fmt.Println(er.Error())
			}

			// 寻找op
			op, _, _, er := parser.Get(value, "op")
			if er == nil {
				// do something, 调用处理left的函数
				fmt.Println(string(op))
			} else {
				fmt.Println(er.Error())
			}

			// 寻找 children
			children, _, _, er := parser.Get(value, "children")
			if er == nil {
				fmt.Println(string(children))
				jsonParseArray(children) // 递归处理
			} else {
				// do nothing
			}
		})
}

func jsonParse(data[]byte) {
	// 寻找 optionsxxx
	optionsxxx, _, _, er := parser.Get(data, "optionsxxx")
	if er == nil {
		jsonParse(optionsxxx)
	} else {
		// do nothing
	}

	left, _, _, er := parser.Get(data, "type")
	if er == nil {
		// do something, 调用处理left的函数
		fmt.Println("type=",string(left))
	} else {
		fmt.Println(er.Error())
	}

	// 寻找 children
	children, _, _, er := parser.Get(data, "children")
	if er == nil {
		fmt.Println("children=", string(children))
		jsonParseArray(children)
	} else {
		// do nothing
	}

}




func main() {
	parser_json()
}
