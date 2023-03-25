package service

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserService(db *gorm.DB, ctx context.Context) *UserService {
	return &UserService{db: db, ctx: ctx}
}

func (u *UserService) IsRecordExist(username string) bool {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

func (u *UserService) CreateUser(user *model.User) (int64, error) {
	// 添加默认头像和背景图
	user.Avatar = "https://cdn.apifox.cn/app/project-icon/builtin/13.jpg"
	user.BackgroundImage = "https://i0.hdslb.com/bfs/game/c8a5f9ba45d89f961c1e85f949d6571ac9309a09.png"
	if err := u.db.Create(user).Error; err != nil {
		return 0, err
	}
	return int64(user.ID), nil
}

func (u *UserService) FindOneByUsername(username string) (model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserService) FindOneByID(id uint64) (model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
