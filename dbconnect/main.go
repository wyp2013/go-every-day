package main

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-every-day/dbconnect/model"
	"go-every-day/dbconnect/utils"
	"os"
)

func Init() {
	configName := flag.String("config", "/Users/bitmain/gowork/src/go-every-day/dbconnect/test.yaml", "--config test.yaml")
	flag.Parse()

	cfg, err := utils.InitConfig(*configName)
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

func GetPoolInfo(poolIds []int, year int) map[int][]*model.PoolMonthInfo {
	poolMap := make(map[int][]*model.PoolMonthInfo)

	for _, pooId := range poolIds {
		monBlockCnt, err := GetPoolMothBlockNum(pooId, year)
		if err != nil {
			log.WithFields(log.Fields{"poolId":pooId, "yead":year}).Error("GetPoolMothBlockNum")
		}

		poolName, monHashCnt, err := GetPoolMothHashrate(pooId, year)
		if err != nil {
			panic(err)
		}

		for mon := 1; mon <= 12; mon++ {
			blockCnt, bOk := monBlockCnt[GetYM(year, mon)]
			if !bOk {
				blockCnt = 0
			}

			hashrate, hOk := monHashCnt[GetYM(year, mon)]
			if !hOk {
				hashrate = 0.0
			}

			poolInfo := &model.PoolMonthInfo{
				PoolName:   poolName,
				YM:         GetYM(year, mon),
				Hashrate:   hashrate,
				BlockCount: blockCnt,
			}

			if _, ok := poolMap[pooId]; !ok {
				poolMap[pooId] = make([]*model.PoolMonthInfo, 0)
			}

			poolMap[pooId] = append(poolMap[pooId], poolInfo)
		}
	}

	return poolMap
}

func GetYM(year, mon int) string {
	return fmt.Sprintf("%04d%02d", year, mon)
}

func GetPoolMothBlockNum(poolId int, year int) (map[string]int, error) {
	data, err := model.GetPoolBlockMonthCount(utils.GetDb(), poolId, year)
	if err != nil {
		return nil, err
	}

	result := make(map[string]int)
	result[GetYM(year, 1)] = data.M01
	result[GetYM(year, 2)] = data.M02
	result[GetYM(year, 3)] = data.M03
	result[GetYM(year, 4)] = data.M04
	result[GetYM(year, 5)] = data.M05
	result[GetYM(year, 6)] = data.M06
	result[GetYM(year, 7)] = data.M07
	result[GetYM(year, 8)] = data.M08
	result[GetYM(year, 9)] = data.M09
	result[GetYM(year, 10)] = data.M10
	result[GetYM(year, 11)] = data.M11
	result[GetYM(year, 12)] = data.M12

	return result, nil
}

func GetPoolMothHashrate(poolId int, year int) (string, map[string]float64, error) {
	poolName := ""
	hashMap := make(map[string]float64)

	for mon := 1; mon <= 12; mon++ {
		start, err := utils.MakeTime("2006-01-02 15:04:05", year, mon, 1, 0, 0, 0)
		if err != nil {
			return "", nil, err
		}

		end := start.AddDate(0, 1, 0)
		datas, err := model.GetPoolHashrate(utils.GetDb(), poolId, start, end)
		if err != nil {
			log.WithField("poolId", poolId).WithField("mon", GetYM(year, mon)).Error("GetPoolHashrate failed")
		}

		cnt := 0
		sumHash := float64(0)
		for _, data := range datas {
			if data.Hashrate > 0 {
				cnt++
				sumHash += data.Hashrate
			}

			if len(poolName) == 0 {
				poolName = data.PoolName
			}
		}

		if cnt > 0 {
			sumHash = sumHash / float64(cnt)
		}

		hashMap[GetYM(year, mon)] = sumHash
	}

	if len(poolName) == 0 && poolId == 41 {
		poolName = "BitFury"
	}

	return poolName, hashMap, nil
}

func GenerateCSV(poolMap map[int][]*model.PoolMonthInfo, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err.Error())
	}

	data := "pool,month,hashrate(P),blocks\r\n"
	_, err = file.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	for _, pooInfos := range poolMap {
		for _, info := range pooInfos {
			data := fmt.Sprintf("%s,%s,%.2f,%d\r\n", info.PoolName, info.YM, info.Hashrate, info.BlockCount)
			_, err = file.Write([]byte(data))
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	Init()

	poolMap := GetPoolInfo([]int{7, 22, 29, 41, 54, 55, 57, 59, 79, 81, 87, 91, 95, 96, 97, 98, 100}, 2019)
	data, _ := json.Marshal(poolMap)
	fmt.Println(string(data))

	GenerateCSV(poolMap, utils.GetConfig().PoolConfig)
}
