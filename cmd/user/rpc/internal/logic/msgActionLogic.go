package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMsgActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgActionLogic {
	return &MsgActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgActionLogic) MsgAction(in *userpc.MsgActionReq) (*userpc.MsgActionResp, error) {
	// todo: add your logic here and delete this line

	return &userpc.MsgActionResp{}, nil
}
