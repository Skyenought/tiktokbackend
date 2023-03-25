package favorite

import (
	"context"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoListLogic {
	return &FavoriteVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteVideoListLogic) FavoriteVideoList(req *types.FavoriteVideoListReq) (resp *types.FavoriteVideoListResp, err error) {
	return
}
