package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/bcryptx"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *service.UserService
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: service.NewUserService(svcCtx.DB, ctx),
		Logger:  logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userpc.LoginReq) (*userpc.LoginResp, error) {
	var user model.User
	var err error
	user, err = l.service.FindOneByUsername(in.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrorHandle(errno.LoginErrCode, "不存在该用户: %s", err.Error())
		}
	}

	if err := bcryptx.CheckPassword(in.Password, user.Password); err != nil {
		return nil, errno.ErrorHandle(errno.LoginErrCode, "密码错误: %s", err.Error())
	}
	return &userpc.LoginResp{
		StatusCode: errno.SuccessCode,
		UserID:     int64(user.ID),
	}, nil
}
