package model

import "time"

type Video struct {
	ID            uint64 `gorm:"primarykey"`
	Title         string
	PlayURL       string
	CoverURL      string `gorm:"default:https://s1.hdslb.com/bfs/static/jinkela/international-home/assets/icon_slide_selected.png"`
	UserID        uint64
	FavoriteCount uint32    `gorm:"default:0"`
	CommentCount  uint32    `gorm:"default:0"`
	CreateTime    time.Time `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime"`
}

func (m *Video) TableName() string {
	return "video"
}
