// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentservice

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment           = commentrpc.Comment
	CommentActionReq  = commentrpc.CommentActionReq
	CommentActionResp = commentrpc.CommentActionResp
	CommentListReq    = commentrpc.CommentListReq
	CommentListResp   = commentrpc.CommentListResp

	CommentService interface {
		CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	}

	defaultCommentService struct {
		cli zrpc.Client
	}
)

func NewCommentService(cli zrpc.Client) CommentService {
	return &defaultCommentService{
		cli: cli,
	}
}

func (m *defaultCommentService) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	client := commentrpc.NewCommentServiceClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}

func (m *defaultCommentService) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := commentrpc.NewCommentServiceClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}
