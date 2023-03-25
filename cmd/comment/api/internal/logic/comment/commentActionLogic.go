package comment

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"
	"google.golang.org/grpc/status"

	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if userID != req.UserID {
		l.Logger.Errorf("用户信息有误")
		return nil, status.Error(errno.AuthErrCode, "用户信息有误")
	}
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "token解析失败: %+v", err.Error())
	}
	actionResp, err := l.svcCtx.CommentRpc.CommentAction(l.ctx, &commentrpc.CommentActionReq{
		UserID:     req.UserID,
		VideoID:    req.VideoID,
		ActionType: req.ActionType,
		Content:    req.CommentText,
		CommentID:  req.CommentID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.CommentActionErrCode, "评论失败: %+v", err.Error())
	}
	_, err = l.svcCtx.VideoRpc.UpdateCommentCount(l.ctx, &videorpc.UpdateCommentCountReq{
		VideoID:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.UpdateCommentCountErrCode, "更新评论失败: %+v", err.Error())
	}
	return &types.CommentActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  "操作成功",
	}, nil
}
