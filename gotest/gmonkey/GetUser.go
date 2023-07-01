package gmonkey

import (
	"fmt"
	"go-every-day/gotest/gmonkey/user"
	"go-every-day/gotest/gmonkey/util"
)

func GetUser(name string) {
	us := user.UserSer{}

	u, _ := us.GetUserINfo(name)

	fmt.Println(util.Obj2String(u))
}