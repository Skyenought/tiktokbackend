package favorite

import (
	"context"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"
	
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoListLogic {
	return &FavoriteVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteVideoListLogic) FavoriteVideoList(req *types.FavoriteVideoListReq) (resp *types.FavoriteVideoListResp, err error) {
	var videoList []types.Video
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "token解析失败: %s", err.Error())
	}
	listResp, err := l.svcCtx.VideoRpc.FavoriteVideoList(l.ctx, &videorpc.FavoriteVideoListReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "获取收藏视频列表失败: %s", err.Error())
	}
	for _, video := range listResp.VideoList {
		infoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userpc.UserInfoReq{
			UserID:     video.UserID,
			FromUserID: userID,
		})
		if err != nil {
			return nil, errno.ErrorHandle(errno.AuthErrCode, "获取用户信息失败: %s", err.Error())
		}
		videoConvHTTP := pack.VideoConvHTTP(video, infoResp.User)
		videoList = append(videoList, videoConvHTTP)
	}
	return &types.FavoriteVideoListResp{
		StatusCode: errno.SuccessCode,
		StatusMsg:  "获取收藏视频列表成功",
		VideoList:  videoList,
	}, nil
}
