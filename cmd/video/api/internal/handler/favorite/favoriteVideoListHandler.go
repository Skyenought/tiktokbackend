package favorite

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/logic/favorite"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := favorite.NewFavoriteVideoListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteVideoList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
