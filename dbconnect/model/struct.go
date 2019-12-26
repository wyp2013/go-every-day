package model

import "time"

type Hashrate struct {
	PoolName string    `json:"pool_name"`
	PoolId   int       `json:"pool_id"`
	Hashrate float64   `json:"Hashrate"`
	UpdateAt time.Time `json:"update_at"`
	Time     time.Time `json:"time"`
}

type BlockMothCount struct {
	Id        int       `json:"id"`
	PoolId    int       `json:"pool_id"`
	Year      int       `json:"year"`
	M01       int       `json:"m_01"`
	M02       int       `json:"m_02"`
	M03       int       `json:"m_03"`
	M04       int       `json:"m_04"`
	M05       int       `json:"m_05"`
	M06       int       `json:"m_06"`
	M07       int       `json:"m_07"`
	M08       int       `json:"m_08"`
	M09       int       `json:"m_09"`
	M10       int       `json:"m_10"`
	M11       int       `json:"m_11"`
	M12       int       `json:"m_12"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PoolMonthInfo struct {
	PoolName   string  `json:"pool_name"`
	Hashrate   float64 `json:"hashrate"`
	YM         string  `json:"ym"`
	BlockCount int     `json:"block_count"`
}
