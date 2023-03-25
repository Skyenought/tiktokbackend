package svc

import (
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{})
	db.AutoMigrate(&model.Favorite{})
	db.AutoMigrate(&model.Video{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
