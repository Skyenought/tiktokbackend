package publish

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/logic/publish"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishlistReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := publish.NewPublishlistLogic(r.Context(), svcCtx)
		resp, err := l.Publishlist(&req)
		result.HttpResult(r, w, resp, err)
	}
}
