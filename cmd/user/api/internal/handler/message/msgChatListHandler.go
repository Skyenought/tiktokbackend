package message

import (
	"net/http"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/logic/message"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"
	"github.com/Skyenought/tiktokbackend/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MsgChatListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgChatListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := message.NewMsgChatListLogic(r.Context(), svcCtx)
		resp, err := l.MsgChatList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
