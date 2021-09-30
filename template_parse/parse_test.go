package template_parse

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"
)

type Session struct {
	UserId    string
	UserName  string
}

func (s *Session) IsValidate() bool {
	if s == nil {
		return false
	}

	return true
}

func TestFunc(t *testing.T) {
	var s *Session = nil
	if !s.IsValidate() {
		fmt.Print("safe")
	}
}

func upper(str string) string {
	return strings.ToUpper(str)
}

func test() {
	tplStr := `{{strupper .}}`
	funcMap := template.FuncMap{
		"strupper": upper,
	}
	t1 := template.New("test1")
	tmpl, err := t1.Funcs(funcMap).Parse(tplStr)
	if err != nil {
		panic(err)
	}
	_ = tmpl.Execute(os.Stdout, "go programming")
}

func TestFuncA(t *testing.T) {
	test()
}

type Person struct {
	Name string
	Age    int
}


func TestPerson(t *testing.T) {
	p := Person{"longshuai", 23}
	tmpl, err := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpl)

	err = tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpl)
}

