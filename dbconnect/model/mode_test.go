package model

import (
	"encoding/json"
	"fmt"
	"go-every-day/dbconnect/utils"
	"testing"
)

func init() {
	cfg, err := utils.InitConfig("/Users/bitmain/gowork/src/go-every-day/dbconnect/test.yaml")
	if err != nil {
		panic(err.Error())
	} else {
		r, _ := json.Marshal(cfg)
		fmt.Println(string(r))
	}

	if err = utils.InitLogger(cfg.Log.Level, cfg.Log.Path, cfg.Log.ShowSource); err != nil {
		panic(err)
	}

	_, err = utils.InitMysql(cfg.GetPoolHashrateDBConn(), cfg.HashrateDB.MaxOpenConns, cfg.HashrateDB.MaxIdleConns)
	if err != nil {
		panic(err)
	}
}
func TestGetPoolHashrate(t *testing.T) {
	start, err := utils.MakeTime("2006-01-02 15:04:05", 2019, 12, 20, 12, 13, 13)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	end, err := utils.MakeTime("2006-01-02 15:04:05", 2019, 12, 20, 13, 13, 13)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	datas, err := GetPoolHashrate(utils.GetDb(), 55, start, end)
	if err != nil {
		t.Fatal(err.Error())
	}else {
		bytes, _ := json.Marshal(datas)
		fmt.Println(string(bytes))
	}

}

func TestGetPoolBlockMonthCount(t *testing.T) {
	info, err := GetPoolBlockMonthCount(utils.GetDb(), 55, 2019)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		bytes, _ := json.Marshal(info)
		fmt.Println(string(bytes))
	}
}
