package pack

import (
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
)

func VideoConvProto(m *model.Video) *videoservice.Video {
	return &videoservice.Video{
		ID:            int64(m.ID),
		UserID:        int64(m.UserID),
		PlayURL:       m.PlayURL,
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: int32(m.FavoriteCount),
		CommentCount:  int32(m.CommentCount),
		Title:         m.Title,
	}
}

func VideosConvProto(ms []model.Video) []*videoservice.Video {
	var res []*videoservice.Video
	for _, m := range ms {
		res = append(res, VideoConvProto(&m))
	}
	return res
}
