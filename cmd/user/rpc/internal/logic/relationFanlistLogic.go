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

type RelationFanlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	subSrv *service.SubscribeService
	logx.Logger
}

func NewRelationFanlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFanlistLogic {
	return &RelationFanlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		subSrv: service.NewSubscribeService(svcCtx.DB, ctx),
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFanlistLogic) RelationFanlist(in *userpc.RelationFanlistReq) (*userpc.RelationFanlistResp, error) {
	fansList, err := l.subSrv.GetFansList(uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.FansListErrCode, "获取粉丝列表失败: %s", err.Error())
	}
	isSubscribe, err := l.subSrv.IsSubscribe(uint64(in.FromUserID), uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.GetSubscribeRelationErrCode, "获取关注关系: %s", err.Error())
	}
	fansListProto := pack.UsersConvProto(fansList, isSubscribe)
	return &userpc.RelationFanlistResp{
		StatusCode: errno.SuccessCode,
		UserList:   fansListProto,
	}, nil
}
