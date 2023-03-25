package user

import (
	"context"
	"time"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/logic"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/pkg/bcryptx"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userservice.RegisterReq{
		Username: req.Username, Password: bcryptx.PasswordEncrypt(req.Password),
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.RegisterErrCode, "注册失败: %s", err.Error())
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.Expire
	token, err := logic.GenToken(l.svcCtx.Config.Auth.Secret, now, accessExpire, rpcResp.UserID)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "生成token失败: %s", err.Error())
	}
	return &types.RegisterResp{
		StatusCode: rpcResp.StatusCode,
		StatusMsg:  "注册成功",
		UserID:     rpcResp.UserID,
		Token:      token,
	}, nil
}
