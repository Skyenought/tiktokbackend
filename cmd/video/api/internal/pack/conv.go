package pack

import (
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
)

// VideoConvHTTP converts a video model to a video http response.
func VideoConvHTTP(v *videorpc.Video, user *userservice.User) types.Video {
	return types.Video{
		ID: v.ID,
		Author: types.User{
			ID:              user.ID,
			Name:            user.Name,
			IsFollow:        user.IsFollow,
			FollowerCount:   user.FollowerCount,
			FollowCount:     user.FollowCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		},
		CoverUrl:      v.CoverURL,
		PlayUrl:       v.PlayURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    false,
		Title:         v.Title,
	}
}

// VideosConv converts a video model list to a video http response list.
func VideosConv(vs []*videorpc.Video, users []*userservice.User) []types.Video {
	var res []types.Video
	for i, v := range vs {
		res = append(res, VideoConvHTTP(v, users[i]))
	}
	return res
}
