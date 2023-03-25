package user

import (
	"context"
	"time"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/logic"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userservice.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.LoginErrCode, "登录失败: %s", err.Error())
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.Expire
	token, err := logic.GenToken(l.svcCtx.Config.Auth.Secret, now, accessExpire, rpcResp.UserID)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "生成token失败: %s", err.Error())
	}
	return &types.LoginResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg:  "登录成功",
		UserID:     rpcResp.UserID,
		Token:      token,
	}, nil
}
