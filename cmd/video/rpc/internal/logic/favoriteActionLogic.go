package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *service.FavoriteService
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: service.NewFavoriteService(ctx, svcCtx.DB),
		Logger:  logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *videorpc.FavoriteActionReq) (*videorpc.FavoriteActionResp, error) {
	switch in.ActionType {
	// 点赞
	case 1:
		err := l.service.AddFavorite(uint64(in.UserID), uint64(in.VideoID))
		if err != nil {
			return nil, errno.ErrorHandle(errno.FavoriteErrCode, "点赞失败! err: %+v", err)
		}
		return &videorpc.FavoriteActionResp{
			StatusCode: errno.SuccessCode,
		}, nil
	// 取消点赞
	case 2:
		err := l.service.MinusFavorite(uint64(in.UserID), uint64(in.VideoID))
		if err != nil {
			return nil, errno.ErrorHandle(errno.FavoriteErrCode, "取消点赞失败! err: %+v", err)
		}
		return &videorpc.FavoriteActionResp{
			StatusCode: errno.SuccessCode,
		}, nil
	}
	return nil, errno.ErrorHandle(errno.FavoriteErrCode, "未知的操作类型! err: %+v", nil)
}
