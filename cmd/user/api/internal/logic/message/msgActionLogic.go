package message

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgActionLogic {
	return &MsgActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgActionLogic) MsgAction(req *types.MsgActionReq) (resp *types.MsgActionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
