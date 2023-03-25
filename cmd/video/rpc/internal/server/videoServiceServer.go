// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/logic"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	videorpc.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServiceServer) FavoriteAction(ctx context.Context, in *videorpc.FavoriteActionReq) (*videorpc.FavoriteActionResp, error) {
	l := logic.NewFavoriteActionLogic(ctx, s.svcCtx)
	return l.FavoriteAction(in)
}

func (s *VideoServiceServer) FavoriteVideoList(ctx context.Context, in *videorpc.FavoriteVideoListReq) (*videorpc.FavoriteVideoListResp, error) {
	l := logic.NewFavoriteVideoListLogic(ctx, s.svcCtx)
	return l.FavoriteVideoList(in)
}

func (s *VideoServiceServer) PublishAction(ctx context.Context, in *videorpc.PublishActionReq) (*videorpc.PublishActionResp, error) {
	l := logic.NewPublishActionLogic(ctx, s.svcCtx)
	return l.PublishAction(in)
}

func (s *VideoServiceServer) PublishList(ctx context.Context, in *videorpc.PublishlistReq) (*videorpc.PublishlistResp, error) {
	l := logic.NewPublishListLogic(ctx, s.svcCtx)
	return l.PublishList(in)
}

func (s *VideoServiceServer) Feed(ctx context.Context, in *videorpc.FeedReq) (*videorpc.FeedResp, error) {
	l := logic.NewFeedLogic(ctx, s.svcCtx)
	return l.Feed(in)
}

func (s *VideoServiceServer) UpdateCommentCount(ctx context.Context, in *videorpc.UpdateCommentCountReq) (*videorpc.UpdateCommentCountResp, error) {
	l := logic.NewUpdateCommentCountLogic(ctx, s.svcCtx)
	return l.UpdateCommentCount(in)
}