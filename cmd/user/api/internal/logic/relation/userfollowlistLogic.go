package relation

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfollowlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfollowlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfollowlistLogic {
	return &UserfollowlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfollowlistLogic) Userfollowlist(req *types.UserfollowlistReq) (resp *types.UserfollowlistResp, err error) {
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取用户ID失败: %s", err.Error())
	}
	followlistResp, err := l.svcCtx.UserRpc.RelationFollowlist(l.ctx, &userpc.RelationFollowlistReq{
		FromUserID: userID,
		UserID:     req.UserID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取用户关注列表失败: %s", err.Error())
	}
	userList := pack.UsersConvHTTP(followlistResp.UserList)
	return &types.UserfollowlistResp{
		StatusCode: followlistResp.StatusCode,
		StatusMsg:  "获取用户关注列表成功",
		UserList:   userList,
	}, nil
}
