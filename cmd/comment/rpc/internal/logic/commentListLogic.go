package logic

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/pack"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *commentrpc.CommentListReq) (*commentrpc.CommentListResp, error) {
	var comments []model.Comment
	if err := l.svcCtx.DB.WithContext(l.ctx).Where("video_id = ?", in.VideoID).Find(&comments).Error; err != nil {
		return nil, errno.ErrorHandle(errno.CommentListErrCode, "获取评论列表失败: %+v", err.Error())
	}
	commentsProto := pack.CommentsConvProto(comments)
	return &commentrpc.CommentListResp{
		StatusCode:  errno.SuccessCode,
		CommentList: commentsProto,
	}, nil
}
