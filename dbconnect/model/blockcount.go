package model

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func GetPoolBlockMonthCount(db *sql.DB, poolId int, year int) (*BlockMothCount, error) {
	querySql := fmt.Sprintf("select * from pool_block_count_months where pool_id=%d and year=%d ", poolId, year)
	log.Info("blockSql: ", querySql)
	row := db.QueryRow(querySql)

	var data BlockMothCount
	err := row.Scan(
		&data.Id,
		&data.PoolId,
		&data.Year,
		&data.M01,
		&data.M02,
		&data.M03,
		&data.M04,
		&data.M05,
		&data.M06,
		&data.M07,
		&data.M08,
		&data.M09,
		&data.M10,
		&data.M11,
		&data.M12,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
