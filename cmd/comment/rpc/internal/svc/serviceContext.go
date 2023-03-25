package svc

import (
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Comment{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
