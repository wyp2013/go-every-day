package model

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func InsertHashrate(db *sql.DB, data *Hashrate) error {
	insert, err := db.Prepare("INSERT INTO `hashrate` (`pool_name`, `pool_id`, `hashrate`, `update_at`, `time`) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Error(err)
		return err
	}
	defer insert.Close()

	result, err := insert.Exec(data.PoolName, data.PoolId, data.Hashrate, time.Now().UTC(), data.Time)
	if err != nil {
		log.Error(err)
		return err
	} else {
		id, _ := result.LastInsertId()
		log.Info(fmt.Sprintf("insert success, PoolName:%s, hashrateid:%f, id: %d", data.PoolName, data.Hashrate, id))
	}

	return nil
}

func GetPoolHashrate(db *sql.DB, poolId int, start, end time.Time) ([]*Hashrate, error) {
	querySql := fmt.Sprintf("select * from hashrate where pool_id=%d and time > \"%s\" and time < \"%s\"", poolId, start.String(), end.String())
	log.Info("hashSql: ", querySql)
	rows, err := db.Query(querySql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}


	var datas []*Hashrate
	for rows.Next() {
		var data Hashrate
		err := rows.Scan(
			&data.Time,
			&data.PoolName,
			&data.PoolId,
			&data.Hashrate,
			&data.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		datas = append(datas, &data)
	}

	return datas, nil
}
