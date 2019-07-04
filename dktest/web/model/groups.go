package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"go-every-day/dktest/web/common"
)

type GroupsModel struct {
	Context *common.CustomContext
}

func NewGroupsModel(cc *common.CustomContext) *GroupsModel {
	return &GroupsModel{
		Context: cc,
	}
}

func (m *GroupsModel) GetGroups(gid int) (*Groups, error) {
	var g Groups

	bFind, err := m.Context.DbEngine.Where("gid=?", gid).Get(&g)
	if err != nil {
		return nil, err
	} else if !bFind {
		return nil, errors.New("Not Find")
	}

	return &g, nil
}