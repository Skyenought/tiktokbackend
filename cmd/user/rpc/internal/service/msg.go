package service

import (
	"context"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MsgService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewMsgService(ctx context.Context, db *gorm.DB) *MsgService {
	return &MsgService{ctx: ctx, db: db}
}

func (m *MsgService) MsgList(fromUserID, toUserID uint64) (msgs []model.Msg, err error) {
	err = m.db.WithContext(m.ctx).Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Find(&msgs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

func (m *MsgService) CreateMsg(fromUserID, toUserID uint64, content string) (err error) {
	return m.db.WithContext(m.ctx).Create(&model.Msg{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
	}).Error
}
