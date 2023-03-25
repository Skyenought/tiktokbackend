package favorite

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionReq) (resp *types.FavoriteActionResp, err error) {
	// 进行点赞或者取消点赞
	actionResp, err := l.svcCtx.VideoRpc.FavoriteAction(l.ctx, &videoservice.FavoriteActionReq{
		UserID:     req.UserID,
		VideoID:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.FavoriteErrCode, "关注失败: %+v", err)
	}
	// 更新获赞数
	_, err = l.svcCtx.UserRpc.UpdateCount(l.ctx, &userservice.UpdateCountReq{
		UserID:       req.UserID,
		FavoriteType: req.ActionType,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.FavoriteErrCode, "更新用户信息失败: %+v", err)
	}
	return &types.FavoriteActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  "操作成功",
	}, nil
}
