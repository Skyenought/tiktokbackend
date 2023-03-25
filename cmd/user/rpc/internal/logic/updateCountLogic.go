package logic

import (
	"context"
	
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"gorm.io/gorm"
	
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountLogic {
	return &UpdateCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCountLogic) UpdateCount(in *userpc.UpdateCountReq) (*userpc.UpdateCountResp, error) {
	switch in.FavoriteType {
	// 1: 增加点赞数 2: 减少点赞数
	case 1:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("favorite_count").
			Where("id = ?", in.UserID).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新获赞数失败: %+v", err)
		}
	case 2:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("favorite_count").
			Where("id = ?", in.UserID).
			Update("favorite_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新获赞数失败: %+v", err)
		}
	}
	// 1: 增加作品数 2: 减少作品数
	if in.VideoType == 1 {
		if err := l.svcCtx.DB.Model(&model.User{}).Select("work_count").
			Where("id = ?", in.UserID).
			Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新作品数失败: %+v", err)
		}
	}
	// 1: 增加粉丝数 2: 减少粉丝数
	switch in.FanType {
	case 1:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("fan_count").
			Where("id = ?", in.UserID).
			Update("fan_count", gorm.Expr("fan_count + ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新关注数失败: %+v", err)
		}
	case 2:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("fan_count").
			Where("id = ?", in.UserID).
			Update("fan_count", gorm.Expr("fan_count - ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新关注数失败: %+v", err)
		}
	}
	// 1: 增加关注数 2: 减少关注数
	switch in.FollowType {
	case 1:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("follow_count").
			Where("id = ?", in.UserID).
			Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新关注数失败: %+v", err)
		}
	case 2:
		if err := l.svcCtx.DB.Model(&model.User{}).Select("follow_count").
			Where("id = ?", in.UserID).
			Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新关注数失败: %+v", err)
		}
	}
	return &userpc.UpdateCountResp{
		StatusCode: errno.SuccessCode,
	}, nil
}
