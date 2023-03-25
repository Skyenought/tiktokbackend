package svc

import (
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentservice"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userservice.UserService
	CommentRpc commentservice.CommentService
	VideoRpc   videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
		CommentRpc: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentRpcConf)),
		VideoRpc:   videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
