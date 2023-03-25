package comment

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/pack"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userpc"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"

	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	var comments []types.Comment
	fromUserID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errno.ErrorHandle(errno.AuthErrCode, "token解析失败: %+v", err.Error())
	}
	listResp, err := l.svcCtx.CommentRpc.CommentList(l.ctx, &commentrpc.CommentListReq{
		VideoID: req.VideoID,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.CommentListErrCode, "获取评论列表失败: %+v", err.Error())
	}
	for _, comment := range listResp.CommentList {
		c := pack.CommentConvHTTP(comment)
		infoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userpc.UserInfoReq{
			UserID:     c.ID,
			FromUserID: fromUserID,
		})
		if err != nil {
			return nil, errno.ErrorHandle(errno.CommentListErrCode, "获取用户信息失败: %+v", err.Error())
		}
		c.User = pack.UserConvHTTP(infoResp.User)
		comments = append(comments, c)
	}
	return &types.CommentListResp{
		StatusCode:  errno.SuccessCode,
		CommentList: comments,
	}, nil
}
