package logic

import (
	"context"
	"time"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	var videoList []types.Video
	now := time.Now().Unix()
	if req.LatestTime > now {
		req.LatestTime = now
	}
	feedResp, err := l.svcCtx.VideoRpc.Feed(l.ctx, &videorpc.FeedReq{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.FeedErrCode, "获取视频列表失败: %+v", err)
	}
	for _, v := range feedResp.VideoList {
		userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userpc.UserInfoReq{
			UserID: v.UserID,
		})
		if err != nil {
			return nil, errno.ErrorHandle(errno.GetUserInfoErrCode, "获取用户信息失败: %+v", err)
		}
		videoList = append(videoList, pack.VideoConvHTTP(v, userInfoResp.User))
	}
	return &types.FeedResp{
		StatusCode: errno.SuccessCode,
		VideoList:  videoList,
		NextTime:   feedResp.NextTime,
		StatusMsg:  "请求成功",
	}, nil
}
