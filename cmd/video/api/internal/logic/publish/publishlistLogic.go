package publish

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishlistLogic {
	return &PublishlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishlistLogic) Publishlist(req *types.PublishlistReq) (resp *types.PublishlistResp, err error) {
	var videoList []types.Video
	// 解析 token
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.GenerateTokenErrCode, "token 解析失败: %+v", err)
	}
	// 获取视频列表
	actionResp, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &videoservice.PublishlistReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.PublishListErrCode, "获取视频列表失败: %+v", err)
	}
	// 获取用户信息
	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userservice.UserInfoReq{
		UserID:     req.UserID,
		FromUserID: userID,
	})
	// 转换视频信息
	for _, v := range actionResp.VideoList {
		videoList = append(videoList, pack.VideoConvHTTP(v, userInfoResp.User))
	}
	return &types.PublishlistResp{
		StatusCode: errno.SuccessCode,
		StatusMsg:  "获取视频列表成功",
		VideoList:  videoList,
	}, nil
}
