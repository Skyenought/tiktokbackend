package relation

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FollowAction 关注或取消关注
func (l *FollowActionLogic) FollowAction(req *types.FollowActionReq) (resp *types.FollowActionResp, err error) {
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)

	actionResp, err := l.svcCtx.UserRpc.RelationAction(l.ctx, &userpc.RelationActionReq{
		ActionType: req.ActionType,
		FromUserID: userID,       // 关注者
		UserID:     req.ToUserID, // 被关注者
	})

	// 增加 or 减少 关注者的粉丝数
	_, err = l.svcCtx.UserRpc.UpdateCount(l.ctx, &userpc.UpdateCountReq{
		UserID:  req.ToUserID,
		FanType: req.ActionType,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.UpdateFansErrCode, "更新粉丝数失败: %s", err.Error())
	}
	// 增加 or 减少 被关注者的关注数
	_, err = l.svcCtx.UserRpc.UpdateCount(l.ctx, &userpc.UpdateCountReq{
		UserID:     userID,
		FollowType: req.ActionType,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.UpdateFansErrCode, "更新关注数失败: %s", err.Error())
	}

	return &types.FollowActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  "关注或取消关注成功",
	}, nil
}
