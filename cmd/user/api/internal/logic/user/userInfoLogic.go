package user

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

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "token解析失败: %s", err.Error())
	}
	infoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userpc.UserInfoReq{
		UserID:     req.UserID,
		FromUserID: userID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取用户信息失败: %s", err.Error())
	}
	return &types.UserInfoResp{
		StatusCode: infoResp.StatusCode,
		StatusMsg:  "获取用户信息成功",
		User:       pack.UserConvHTTP(infoResp.User),
	}, nil
}
