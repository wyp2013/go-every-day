package model

import "time"

type Groups struct {
	Gid       int        `json:"gid"`
	Puid      int        `json:"puid"`
	Name      string     `json:"name"`
	UpdateTime time.Time `json:"updated_at" xorm:"updated"`
	CreateTime time.Time `json:"created_at" xorm:"created"`
}
