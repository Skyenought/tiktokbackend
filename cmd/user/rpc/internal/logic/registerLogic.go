package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *service.UserService
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: service.NewUserService(svcCtx.DB, ctx),
		Logger:  logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userpc.RegisterReq) (*userpc.RegisterResp, error) {
	var user model.User
	if !l.service.IsRecordExist(in.Username) {
		user.Username = in.Username
		user.Password = in.Password
		id, err := l.service.CreateUser(&user)
		if err != nil {
			return nil, errno.ErrorHandle(errno.DBInsertErrCode, "创建用户失败: %s", err.Error())
		}
		return &userpc.RegisterResp{
			StatusCode: errno.SuccessCode,
			UserID:     id,
		}, nil
	}
	return nil, errno.ErrorHandle(errno.UserAlreadyExistErrCode, "用户已存在")
}
