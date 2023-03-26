package logic

import (
	"context"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	msgSrv *service.MsgService
	logx.Logger
}

func NewMsgChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgChatListLogic {
	return &MsgChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		msgSrv: service.NewMsgService(ctx, svcCtx.DB),
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgChatListLogic) MsgChatList(in *userpc.MsgChatListReq) (*userpc.MsgChatListResp, error) {
	msgs, err := l.msgSrv.MsgList(uint64(in.UserID), uint64(in.ToUserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.GetMsgListErrCode, "获取信息列表失败: %+v", err.Error())
	}
	msgList := pack.MsgsConvProto(msgs)
	return &userpc.MsgChatListResp{
		StatusCode: errno.SuccessCode,
		MsgList:    msgList,
	}, nil
}
