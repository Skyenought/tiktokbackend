package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	usrSrv *service.UserService
	subSrv *service.SubscribeService
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		usrSrv: service.NewUserService(svcCtx.DB, ctx),
		subSrv: service.NewSubscribeService(svcCtx.DB, ctx),
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *userpc.UserInfoReq) (*userpc.UserInfoResp, error) {
	var user model.User
	var err error
	user, err = l.usrSrv.FindOneByID(uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.UserNotExistErrCode, "查询用户信息失败: %s", err.Error())
	}
	isSubscribe, err := l.subSrv.IsSubscribe(uint64(in.FromUserID), uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.GetSubscribeRelationErrCode, "获取关注关系失败: %s", err.Error())
	}
	protoUser := pack.UserConvProto(user, isSubscribe)
	return &userpc.UserInfoResp{
		StatusCode: errno.SuccessCode,
		User:       &protoUser,
	}, nil
}
