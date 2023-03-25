package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	subSrv *service.SubscribeService
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		subSrv: service.NewSubscribeService(svcCtx.DB, ctx),
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationActionLogic) RelationAction(in *userpc.RelationActionReq) (*userpc.RelationActionResp, error) {
	switch in.ActionType {
	case 1:
		err := l.subSrv.Subscribe(uint64(in.FromUserID), uint64(in.UserID))
		if err != nil {
			return nil, errno.ErrorHandle(errno.RelationActionErrCode, "关注失败: %s", err.Error())
		}
	case 2:
		err := l.subSrv.UnSubscribe(uint64(in.FromUserID), uint64(in.UserID))
		if err != nil {
			return nil, errno.ErrorHandle(errno.DisRelationErrCode, "取消关注失败: %s", err.Error())
		}
	}
	return &userpc.RelationActionResp{
		StatusCode: errno.SuccessCode,
	}, nil
}
