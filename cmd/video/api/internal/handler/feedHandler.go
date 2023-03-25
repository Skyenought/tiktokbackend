package handler

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/logic"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		result.HttpResult(r, w, resp, err)
	}
}
