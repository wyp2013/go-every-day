package gmonkey

import (
	"bou.ke/monkey"
	"go-every-day/gotest/gmonkey/util"
	"reflect"
	"testing"

	"go-every-day/gotest/gmonkey/user"
)

func TestGetUser(t *testing.T) {

	// patch obj method GetUserINfo
	us := &user.UserSer{}
	monkey.PatchInstanceMethod(reflect.TypeOf(us), "GetUserINfo", func(*user.UserSer, string) (*user.User, error) {
		return &user.User{Name: "hhhhhhhhh", Age: 20}, nil
	})

	//monkey.Patch(util.Obj2String, func(obj interface{}) string {
	//	return "path obj2string"
	//})

	GetUser("test")
}

func TestGetUser2(t *testing.T) {
	monkey.Patch(util.Obj2String, func(obj interface{}) string {
		return "path obj2string"
	})

	GetUser("test")
}
