package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/service"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	commentSrv *service.CommentService
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		commentSrv: service.NewCommentService(ctx, svcCtx.DB),
		Logger:     logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *commentrpc.CommentActionReq) (*commentrpc.CommentActionResp, error) {
	var comment model.Comment
	var err error
	switch in.ActionType {
	case 1:
		comment, err = l.commentSrv.AddComment(uint64(in.VideoID), uint64(in.UserID), in.Content)
		if err != nil {
			return nil, errno.ErrorHandle(errno.CommentActionErrCode, "添加评论失败: %+v", err.Error())
		}
		return &commentrpc.CommentActionResp{
			StatusCode: errno.SuccessCode,
			Comment:    pack.CommentConvProto(comment),
		}, nil
	case 2:
		comment, err = l.commentSrv.DeleteComment(uint64(in.CommentID))
		if err != nil {
			return nil, errno.ErrorHandle(errno.CommentActionErrCode, "删除评论失败: %+v", err.Error())
		}
		return &commentrpc.CommentActionResp{
			StatusCode: errno.SuccessCode,
			Comment:    pack.CommentConvProto(comment),
		}, nil
	}
	return nil, errno.ErrorHandle(errno.CommentActionErrCode, "操作类型错误")
}
