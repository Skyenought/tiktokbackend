package publish

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/Skyenought/tiktokbackend/pkg/errno"
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"
	"github.com/pkg/errors"

	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	httpReq *http.Request
	ctx     context.Context
	svcCtx  *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PublishActionLogic {
	return &PublishActionLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		httpReq: r,
		svcCtx:  svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq) (resp *types.PublishActionResp, err error) {
	// 解析 token
	userID, err := jwtx.GetUserId(req.Token, l.svcCtx.Config.Auth.Secret)
	if err != nil {
		return nil, errors.Wrapf(errno.NewErrCode(errno.AuthErrCode), "解析token失败: %+v", err)
	}
	// 存入 video
	_, fileHeader, err := l.httpReq.FormFile("data")
	if err != nil {
		return nil, errors.Wrapf(errno.NewErrCode(errno.UploadErrCode), "获取文件失败: %+v", err)
	}
	videoContent, err := fileHeader.Open()
	if err != nil {
		return nil, errors.Wrapf(errno.NewErrCode(errno.UploadErrCode), "打开文件失败: %+v", err)
	}

	// minio oss
	//_, err = l.svcCtx.MinioClient.PutObject(
	//	l.svcCtx.Config.Minio.BucketName,
	//	fmt.Sprintf("%v/%v.mp4", l.svcCtx.Config.Minio.UploadPath, req.Title),
	//	videoContent,
	//	fileHeader.Size,
	//	minio.PutObjectOptions{ContentType: "video/mp4"},
	//)

	// tecent oss
	filename := fmt.Sprintf("%v/%v.mp4", l.svcCtx.Config.Minio.UploadPath, req.Title)
	_, err = l.svcCtx.TencentOSS.Object.Put(l.ctx, filename, videoContent, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType:   "video/mp4",
			ContentLength: fileHeader.Size,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(errno.NewErrCode(errno.UploadErrCode), "上传文件失败: %+v", err)
	}
	// 发布作品
	actionResp, err := l.svcCtx.VideoRpc.PublishAction(l.ctx, &videoservice.PublishActionReq{
		UserID: userID,
		Title:  req.Title,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.UploadErrCode, "发布失败: %+v", err)
	}

	// 更新作品数
	_, err = l.svcCtx.UserRpc.UpdateCount(l.ctx, &userservice.UpdateCountReq{
		UserID:    userID,
		VideoType: 1,
	})
	if err != nil {
		return nil, errno.ErrorHandle(errno.UpdateCountErrCode, "更新作品数失败: %+v", err)
	}
	return &types.PublishActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  "上传成功",
	}, nil
}
