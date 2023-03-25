package model

import "time"

type User struct {
	ID              uint64 `gorm:"primarykey"`
	Username        string
	Password        string
	Signature       string
	Avatar          string    `gorm:"default:https://s1.hdslb.com/bfs/static/jinkela/international-home/assets/icon_slide_selected.png"`
	BackgroundImage string    `gorm:"default:https://s1.hdslb.com/bfs/static/jinkela/international-home/assets/icon_slide_selected.png"` // 背景图片
	FollowCount     int32     `gorm:"default:0"`                                                                                         // 关注数
	FanCount        int32     `gorm:"default:0"`                                                                                         /// 粉丝数
	WorkCount       int32     `gorm:"default:0"`                                                                                         /// 作品数
	FavoriteCount   int32     `gorm:"default:0"`                                                                                         /// 点赞数
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime:nano"`
}

func (m *User) TableName() string {
	return "user"
}
