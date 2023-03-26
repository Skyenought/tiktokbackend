package svc

import (
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	VideoRpc videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Subscribe{})
	db.AutoMigrate(&model.Msg{})
	return &ServiceContext{
		Config:   c,
		DB:       db,
		VideoRpc: videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
