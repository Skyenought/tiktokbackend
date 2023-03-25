package relation

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfriendlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfriendlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfriendlistLogic {
	return &UserfriendlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfriendlistLogic) Userfriendlist(req *types.UserfriendListReq) (resp *types.UserfriendListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
