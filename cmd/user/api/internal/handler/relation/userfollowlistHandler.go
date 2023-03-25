package relation

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/logic/relation"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserfollowlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserfollowlistReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := relation.NewUserfollowlistLogic(r.Context(), svcCtx)
		resp, err := l.Userfollowlist(&req)
		result.HttpResult(r, w, resp, err)
	}
}
