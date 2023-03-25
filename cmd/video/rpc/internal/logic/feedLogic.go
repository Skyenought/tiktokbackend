package logic

import (
	"context"
	"time"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *videorpc.FeedReq) (*videorpc.FeedResp, error) {
	var videos []model.Video
	var latestTime string
	var nextTime int64
	if in.LatestTime == 0 {
		latestTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		latestTime = time.Unix(in.LatestTime, 0).Format("2006-01-02 15:04:05")
	}
	if find := l.svcCtx.DB.WithContext(l.ctx).Order("create_time DESC").Limit(30).Where("create_time <= ?", latestTime).Find(&videos); find.Error != nil {
		return nil, errno.ErrorHandle(errno.FeedErrCode, "获取视频列表失败: %+v", find.Error)
	}
	nextTime = videos[0].CreateTime.Unix()
	proto := pack.VideosConvProto(videos)
	return &videorpc.FeedResp{
		StatusCode: errno.SuccessCode,
		NextTime:   nextTime,
		VideoList:  proto,
	}, nil
}
