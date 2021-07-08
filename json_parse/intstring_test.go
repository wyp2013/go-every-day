package json_parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func Test_IntString(t *testing.T) {
	jsonStr1 := `{"id": "1234567", "score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Print(c1)
}

func Test_intfloat(t *testing.T) {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // float64
	fmt.Println(m2["count"])

	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // int
}

type IntString int

func (i *IntString) UnmarshalJSON(data []byte) error {
	js := string(data)
	var err error
	*(*int)(i), err = strconv.Atoi(js)
	return err
}

type Card2 struct {
	ID    IntString   `json:"id"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func Test_intfloat2(t *testing.T) {
	str := `{"id": 1234567}`
	var m2 map[string]interface{}
	err := json.Unmarshal([]byte(str), &m2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(m2)

	decoder := json.NewDecoder(bytes.NewReader([]byte(str)))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	fmt.Printf("value:%v\n", m2["id"]) // 1
	fmt.Printf("type:%T\n", m2["id"])  // int

}
