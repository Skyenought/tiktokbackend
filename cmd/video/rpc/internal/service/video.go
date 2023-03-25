package service

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type VideoService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewVideoService(ctx context.Context, db *gorm.DB) *VideoService {
	return &VideoService{ctx: ctx, db: db}
}

// FavoriteVideos 获取用户喜欢的视频列表
func (v *VideoService) FavoriteVideos(userID uint64) ([]model.Video, error) {
	var videos []model.Video
	var favorites []model.Favorite
	err := v.db.WithContext(v.ctx).Transaction(func(tx *gorm.DB) error {
		// 查询 user_id 喜欢的 video_id 的列表
		if err := tx.Model(&model.Favorite{}).Select("video_id").
			Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		var videoIDs []uint64
		for _, favorite := range favorites {
			videoIDs = append(videoIDs, favorite.VideoID)
		}
		// 查询 video_id 对应的视频列表
		if err := v.db.WithContext(v.ctx).Find(&videos, videoIDs).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoService) PublishVideos(userID uint64) ([]model.Video, error) {
	var videos []model.Video
	if err := v.db.WithContext(v.ctx).Where("user_id = ?", userID).Find(&videos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return videos, nil
}
