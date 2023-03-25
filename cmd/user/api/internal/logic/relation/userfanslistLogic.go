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

type UserfanslistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfanslistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfanslistLogic {
	return &UserfanslistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfanslistLogic) Userfanslist(req *types.UserfanslistReq) (resp *types.UserfanslistResp, err error) {
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取用户ID失败: %s", err.Error())
	}
	fanlistResp, err := l.svcCtx.UserRpc.RelationFanlist(l.ctx, &userpc.RelationFanlistReq{
		UserID:     req.UserID,
		FromUserID: userID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取粉丝列表失败: %s", err.Error())
	}
	userList := pack.UsersConvHTTP(fanlistResp.UserList)
	return &types.UserfanslistResp{
		StatusCode: fanlistResp.StatusCode,
		StatusMsg:  "获取粉丝列表成功",
		UserList:   userList,
	}, nil
}
