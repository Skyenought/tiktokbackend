package service

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubscribeService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewSubscribeService(db *gorm.DB, ctx context.Context) *SubscribeService {
	return &SubscribeService{db: db, ctx: ctx}
}

// Subscribe 订阅
func (s *SubscribeService) Subscribe(fanID, videoAuthorID uint64) error {
	return s.db.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 创建订阅关系
		if err := tx.Create(&model.Subscribe{FanID: fanID, VideoAuthorID: videoAuthorID}).Error; err != nil {
			return err
		}
		// 更新粉丝数
		if err := tx.Model(&model.User{}).Where("id = ?", videoAuthorID).
			Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		// 更新关注数
		if err := tx.Model(&model.User{}).Where("id = ?", fanID).
			Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// UnSubscribe 取消订阅
func (s *SubscribeService) UnSubscribe(fanID, videoAuthorID uint64) error {
	return s.db.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 删除订阅关系
		if err := tx.Where("fan_id = ? AND video_author_id = ?", fanID, videoAuthorID).Delete(&model.Subscribe{}).Error; err != nil {
			return err
		}
		// 更新粉丝数
		if err := tx.Model(&model.User{}).Where("id = ?", videoAuthorID).
			Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}
		// 减少关注数
		if err := tx.Model(&model.User{}).Where("id = ?", fanID).
			Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetSubscribeList 关注列表
func (s *SubscribeService) GetSubscribeList(fanID uint64) ([]model.User, error) {
	var users []model.User
	if err := s.db.WithContext(s.ctx).Model(&model.Subscribe{}).Where("fan_id = ?", fanID).
		Select("user.*").Joins("JOIN user ON user.id = subscribe.video_author_id").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// GetFansList 粉丝列表
func (s *SubscribeService) GetFansList(videoAuthorID uint64) ([]model.User, error) {
	var users []model.User
	if err := s.db.WithContext(s.ctx).Model(&model.Subscribe{}).Where("video_author_id = ?", videoAuthorID).
		Select("user.*").Joins("JOIN user ON user.id = subscribe.fan_id").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// IsSubscribe 是否关注视频作者
func (s *SubscribeService) IsSubscribe(fromUserID, videoAuthorID uint64) (bool, error) {
	var subscribe model.Subscribe
	if err := s.db.WithContext(s.ctx).Where("fan_id = ? AND video_author_id = ?", fromUserID, videoAuthorID).
		First(&subscribe).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
