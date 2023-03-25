package svc

import (
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
