package service

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"gorm.io/gorm"
)

type FavoriteService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewFavoriteService(ctx context.Context, db *gorm.DB) *FavoriteService {
	return &FavoriteService{ctx: ctx, db: db}
}

func (s *FavoriteService) AddFavorite(userID uint64, videoID uint64) error {
	var video model.Video
	return s.db.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 更新视频点赞数
		if err := tx.Model(&video).Select("favorite_count").
			Where("id = ?", videoID).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
			Error; err != nil {
			return err
		}
		// 2. 添加点赞记录
		if err := tx.Create(&model.Favorite{UserID: userID, VideoID: videoID}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *FavoriteService) MinusFavorite(userID uint64, videoID uint64) error {
	return s.db.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		var video model.Video
		// 1. 更新视频点赞数
		if err := tx.Model(&video).Select("favorite_count").
			Where("id = ?", videoID).
			Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
			Error; err != nil {
			return err
		}
		// 2. 删除点赞记录
		if err := tx.Delete(&model.Favorite{UserID: userID, VideoID: videoID}).Error; err != nil {
			return err
		}
		return nil
	})
}
