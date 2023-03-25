package model

import "time"

type Comment struct {
	ID         uint64 `gorm:"primarykey"`
	UserID     uint64
	VideoID    uint64
	Content    string
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (c Comment) TableName() string {
	return "comment"
}
