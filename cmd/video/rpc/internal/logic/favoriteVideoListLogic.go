package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoListLogic struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	videoSrv *service.VideoService
	logx.Logger
}

func NewFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoListLogic {
	return &FavoriteVideoListLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		videoSrv: service.NewVideoService(ctx, svcCtx.DB),
		Logger:   logx.WithContext(ctx),
	}
}

func (l *FavoriteVideoListLogic) FavoriteVideoList(in *videorpc.FavoriteVideoListReq) (*videorpc.FavoriteVideoListResp, error) {
	var videos []model.Video
	var err error
	videos, err = l.videoSrv.FavoriteVideos(uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.FollowListErrCode, "获取视频列表失败! err: %+v", err)
	}
	return &videorpc.FavoriteVideoListResp{
		StatusCode: errno.SuccessCode,
		VideoList:  pack.VideosConvProto(videos),
	}, nil
}
