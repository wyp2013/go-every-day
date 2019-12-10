package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"go-every-day/dktest/web/common"
	"go-every-day/dktest/web/model"
	"strconv"
)

func GetGroups(c echo.Context) error {
	fmt.Println("get group")

	idStr := c.QueryParam("id")
	cc := c.(*common.CustomContext)

	if len(idStr) == 0 {
		return cc.SendRspErrMsg(common.ERR_PARAM, "参数id为空")
	}

	gid, err := strconv.Atoi(idStr)
	if err != nil {
		return cc.SendRspErrMsg(common.ERR_PARAM, fmt.Sprintf("参数id转int错误：%s", idStr))
	}

	groupsModel := model.NewGroupsModel(cc)
	g, err := groupsModel.GetGroups(gid)
	if err != nil {
		return cc.SendRspErrMsg(common.ERR_DB, err.Error())
	}

	return cc.SendRspOK(g)
}

