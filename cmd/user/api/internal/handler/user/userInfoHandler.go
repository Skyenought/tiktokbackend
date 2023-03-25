package user

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/logic/user"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
