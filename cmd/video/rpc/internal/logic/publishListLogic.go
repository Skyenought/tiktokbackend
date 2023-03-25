package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	videoSrv *service.VideoService
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		videoSrv: service.NewVideoService(ctx, svcCtx.DB),
		Logger:   logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *videorpc.PublishlistReq) (*videorpc.PublishlistResp, error) {
	videos, err := l.videoSrv.PublishVideos(uint64(in.UserID))
	if err != nil {
		return nil, errno.ErrorHandle(errno.PublishListErrCode, "获取视频列表失败! err: %+v", err)
	}
	return &videorpc.PublishlistResp{
		StatusCode: errno.SuccessCode,
		VideoList:  pack.VideosConvProto(videos),
	}, nil
}
