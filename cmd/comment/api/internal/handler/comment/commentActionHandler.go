package comment

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/logic/comment"
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		result.HttpResult(r, w, resp, err)
	}
}
