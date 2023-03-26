package logic

import (
	"context"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"gorm.io/gorm"
	
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentCountLogic {
	return &UpdateCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentCountLogic) UpdateCommentCount(in *videorpc.UpdateCommentCountReq) (*videorpc.UpdateCommentCountResp, error) {
	switch in.ActionType {
	case 1:
		// 增加 comment_count
		if err := l.svcCtx.DB.WithContext(l.ctx).
			Model(&model.Video{}).
			Select("comment_count").
			Where("id = ?", in.VideoID).
			Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCommentCountErrCode, "更新评论数失败: %+v", err.Error())
		}
	case 2:
		// 减少 comment_count
		if err := l.svcCtx.DB.WithContext(l.ctx).
			Model(&model.Video{}).
			Select("comment_count").
			Where("id = ?", in.VideoID).
			Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return nil, errno.ErrorHandle(errno.UpdateCommentCountErrCode, "更新评论数失败: %+v", err.Error())
		}
	}
	return &videorpc.UpdateCommentCountResp{
		StatusCode: errno.SuccessCode,
	}, nil
}
