package pack

import (
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/types"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
)

// CommentConvHTTP conv to http
func CommentConvHTTP(m *commentrpc.Comment) types.Comment {
	return types.Comment{
		ID:         m.ID,
		User:       types.User{ID: m.UserID},
		Content:    m.Content,
		CreateDate: m.CreateDate,
	}
}

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
