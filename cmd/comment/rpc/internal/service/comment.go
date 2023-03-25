package service

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/model"
	"gorm.io/gorm"
)

type CommentService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewCommentService(ctx context.Context, db *gorm.DB) *CommentService {
	return &CommentService{
		db:  db,
		ctx: ctx,
	}
}

func (c *CommentService) CommentList() {
}

// AddComment 添加评论
func (c *CommentService) AddComment(videoID uint64, userID uint64, content string) (model.Comment, error) {
	comment := model.Comment{
		VideoID: videoID,
		UserID:  userID,
		Content: content,
	}
	if err := c.db.WithContext(c.ctx).Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

// DeleteComment 删除评论
func (c *CommentService) DeleteComment(commentID uint64) (model.Comment, error) {
	comment := model.Comment{ID: commentID}
	if err := c.db.WithContext(c.ctx).First(&comment).Error; err != nil {
		return comment, err
	}
	if err := c.db.WithContext(c.ctx).Delete(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}
