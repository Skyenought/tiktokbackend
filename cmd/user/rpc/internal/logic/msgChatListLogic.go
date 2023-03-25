package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMsgChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgChatListLogic {
	return &MsgChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgChatListLogic) MsgChatList(in *userpc.MsgChatListReq) (*userpc.MsgChatListResp, error) {
	// todo: add your logic here and delete this line

	return &userpc.MsgChatListResp{}, nil
}
