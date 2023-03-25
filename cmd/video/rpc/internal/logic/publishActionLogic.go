package logic

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *videorpc.PublishActionReq) (*videorpc.PublishActionResp, error) {
	// 查询是否存在重复视频
	var video model.Video
	var err error

	// minioplayURL := MinioSprintfPlayURL(l.svcCtx.Config.Minio.EndPoint, l.svcCtx.Config.Minio.UploadPath, in.Title)
	tencentOssPlayURL := TencentSprintfPlayURL(l.svcCtx.Config.Oss.BucketURL, l.svcCtx.Config.Oss.UpdatePath, in.Title)
	video = model.Video{PlayURL: tencentOssPlayURL, UserID: uint64(in.UserID), Title: in.Title}
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("play_url = ? AND title = ?", tencentOssPlayURL, in.Title).
		First(&video).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 存入视频信息
			if err := l.svcCtx.DB.WithContext(l.ctx).Create(&video).Error; err != nil {
				return nil, errno.ErrorHandle(errno.PublishActionErrCode, "发布失败: %+v", err)
			}
		}
	}

	return &videorpc.PublishActionResp{StatusCode: errno.SuccessCode}, nil
}

// MinioSprintfPlayURL mino 格式的 URL
func MinioSprintfPlayURL(s1, s2, s3 string) string {
	return fmt.Sprintf("http://%v%v/%v.mp4", s1, s2, s3)
}

// TencentSprintfPlayURL tencent oss 格式的 URL
func TencentSprintfPlayURL(s1, s2, s3 string) string {
	return fmt.Sprintf("%s%s/%v.mp4", s1, s2, s3)
}
