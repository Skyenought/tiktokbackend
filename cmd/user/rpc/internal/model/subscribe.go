package model

import "time"

type Subscribe struct {
	ID            uint64    `gorm:"primarykey"`
	FanID         uint64    `gorm:"column:fan_id"`          //	粉丝的 ID
	VideoAuthorID uint64    `gorm:"column:video_author_id"` // 被关注者的 ID, 即博主的 ID
	CreateTime    time.Time `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime"`
	UpdateTime    time.Time `gorm:"column:update_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP;autoUpdateTime:nano"`
}

func (m *Subscribe) TableName() string {
	return "`subscribe`"
}
