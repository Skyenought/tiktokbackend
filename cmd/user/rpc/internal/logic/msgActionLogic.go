package logic

import (
	"context"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	msgSrv *service.MsgService
	logx.Logger
}

func NewMsgActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgActionLogic {
	return &MsgActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		msgSrv: service.NewMsgService(ctx, svcCtx.DB),
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgActionLogic) MsgAction(in *userpc.MsgActionReq) (*userpc.MsgActionResp, error) {
	err := l.msgSrv.CreateMsg(uint64(in.UserID), uint64(in.ToUserID), in.Content)
	if err != nil {
		return nil, errno.ErrorHandle(errno.MsgActionErrCode, "发送信息失败: %+v", err.Error())
	}
	return &userpc.MsgActionResp{
		StatusCode: errno.SuccessCode,
	}, nil
}
