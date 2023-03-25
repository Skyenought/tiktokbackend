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

type RelationFollowlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	subSrv *service.SubscribeService
	logx.Logger
}

func NewRelationFollowlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowlistLogic {
	return &RelationFollowlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		subSrv: service.NewSubscribeService(svcCtx.DB, ctx),
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFollowlistLogic) RelationFollowlist(in *userpc.RelationFollowlistReq) (*userpc.RelationFollowlistResp, error) {
	users, err := l.subSrv.GetSubscribeList(uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.FollowListErrCode, "获取关注列表失败: %s", err.Error())
	}
	isSubscribe, err := l.subSrv.IsSubscribe(uint64(in.FromUserID), uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.GetSubscribeRelationErrCode, "获取关注关系: %s", err.Error())
	}
	userList := pack.UsersConvProto(users, isSubscribe)
	return &userpc.RelationFollowlistResp{
		StatusCode: errno.SuccessCode,
		UserList:   userList,
	}, nil
}
