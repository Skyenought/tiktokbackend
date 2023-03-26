package model

import "time"

type Msg struct {
	ID         uint64    `gorm:"primary_key;auto_increment;not null" json:"id"`
	ToUserID   uint64    `gorm:"not null" json:"to_user_id"`
	FromUserID uint64    `gorm:"not null" json:"from_user_id"`
	Content    string    `gorm:"not null" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (m *Msg) TableName() string {
	return "message"
}
