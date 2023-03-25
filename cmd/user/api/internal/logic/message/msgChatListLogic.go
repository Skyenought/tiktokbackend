package message

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgChatListLogic {
	return &MsgChatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgChatListLogic) MsgChatList(req *types.MsgChatListReq) (resp *types.MsgChatListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
