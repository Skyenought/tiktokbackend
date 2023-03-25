package pack

import (
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
)

// UserConvHTTP UserConv userpc.User conv to types.User
func UserConvHTTP(user *userpc.User) (resp types.User) {
	return types.User{
		ID:              user.ID,
		Name:            user.Name,
		Avatar:          user.Avatar,
		Signature:       user.Signature,
		IsFollow:        user.IsFollow,
		FavoriteCount:   user.FavoriteCount,
		WorkCount:       user.WorkCount,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		TotalFavorited:  user.TotalFavorited,
		BackgroundImage: user.BackgroundImage,
	}
}

func UsersConvHTTP(users []*userpc.User) (resp []types.User) {
	for _, user := range users {
		resp = append(resp, UserConvHTTP(user))
	}
	return
}
