package model

import "time"

type Favorite struct {
	ID         uint64 `gorm:"primarykey"`
	UserID     uint64
	VideoID    uint64
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime"`
}

func (m *Favorite) TableName() string {
	return "favorite"
}
