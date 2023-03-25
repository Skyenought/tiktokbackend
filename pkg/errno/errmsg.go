package errno

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	SuccessCode = 0

	ServiceErrCode          uint32 = 100001
	ParamErrCode                   = 100002
	AuthErrCode                    = 100003
	GenerateTokenErrCode           = 100004
	DBErrCode                      = 100005
	DBResultNotFoundErrCode        = 100006
	DBUpdateErrCode                = 100007
	DBInsertErrCode                = 100008
	DBDeleteErrCode                = 100010

	LoginErrCode            = 200001
	UserNotExistErrCode     = 200002
	UserAlreadyExistErrCode = 200003
	RegisterErrCode         = 200004
	GetUserInfoErrCode      = 200005
	UpdateCountErrCode      = 200006
	UpdateFollowErrCode     = 200007
	UpdateFansErrCode       = 200008
	UpdateLikeErrCode       = 200009
	UpdateDislikeErrCode    = 200010

	UploadErrCode        = 300001
	FeedErrCode          = 300002
	PublishActionErrCode = 300003
	PublishListErrCode   = 300004

	RelationActionErrCode       = 400001
	RelationAlreadyExistErrCode = 400002
	DisRelationErrCode          = 400003
	FollowListErrCode           = 400004
	GetSubscribeRelationErrCode = 400005
	FansListErrCode             = 400006

	FavoriteErrCode = 500001
	DislikeErrCode  = 500002

	CommentActionErrCode      = 600001
	DeleteCommentErrCode      = 600002
	CommentListErrCode        = 600003
	UpdateCommentCountErrCode = 600004

	MsgActionErrCode  = 700001
	GetMsgListErrCode = 700002
)

type ErrNo struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *ErrNo) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *ErrNo) GetErrMsg() string {
	return e.errMsg
}

func (e *ErrNo) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrNo(errCode uint32, errMsg string) *ErrNo {
	return &ErrNo{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *ErrNo {
	return &ErrNo{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *ErrNo {
	return &ErrNo{errCode: 10001, errMsg: errMsg}
}

var Message map[uint32]string

func init() {
	Message = make(map[uint32]string)
	Message[SuccessCode] = "请求成功"
	Message[ServiceErrCode] = "服务器开小差啦,稍后再来试一试"
	Message[ParamErrCode] = "参数错误"
	Message[AuthErrCode] = "token失效，请重新登陆"
	Message[GenerateTokenErrCode] = "生成token失败"

	Message[DBErrCode] = "数据库繁忙,请稍后再试"
	Message[DBResultNotFoundErrCode] = "数据查询结果为空"
	Message[DBUpdateErrCode] = "数据更新失败"
	Message[DBInsertErrCode] = "数据插入失败"
	Message[DBDeleteErrCode] = "数据删除失败"

	Message[LoginErrCode] = "用户名或密码错误"
	Message[UserNotExistErrCode] = "用户不存在"
	Message[UserAlreadyExistErrCode] = "用户已存在"
	Message[RegisterErrCode] = "注册失败"
	Message[GetUserInfoErrCode] = "获取用户信息失败"
	Message[UpdateCountErrCode] = "更新用户信息失败"
	Message[UpdateFollowErrCode] = "更新关注信息失败"
	Message[UpdateFansErrCode] = "更新粉丝信息失败"
	Message[UpdateLikeErrCode] = "更新点赞信息失败"
	Message[UpdateDislikeErrCode] = "更新取消点赞信息失败"

	Message[UploadErrCode] = "上传失败"
	Message[FeedErrCode] = "获取视频流失败"
	Message[PublishActionErrCode] = "投稿失败"
	Message[PublishListErrCode] = "获取作品列表失败"

	Message[RelationActionErrCode] = "关注操作失败"
	Message[RelationAlreadyExistErrCode] = "已经关注过了"
	Message[DisRelationErrCode] = "取消关注失败"
	Message[GetSubscribeRelationErrCode] = "获取关注关系失败"
	Message[FollowListErrCode] = "获取关注列表失败"
	Message[FansListErrCode] = "获取粉丝列表失败"

	Message[FavoriteErrCode] = "点赞失败"
	Message[DislikeErrCode] = "取消点赞失败"

	Message[CommentActionErrCode] = "评论操作失败"
	Message[DeleteCommentErrCode] = "删除评论失败"
	Message[CommentListErrCode] = "获取评论列表失败"
	Message[UpdateCommentCountErrCode] = "更新评论数失败"

	Message[MsgActionErrCode] = "发送消息失败"
	Message[GetMsgListErrCode] = "获取消息列表失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := Message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	// 可以传送的范围
	if _, ok := Message[errcode]; ok && errcode <= DBErrCode && errcode > DBDeleteErrCode {
		return true
	} else {
		return false
	}
}

func ErrorHandle(errcode uint32, format string, args ...interface{}) error {
	return errors.Wrapf(NewErrCode(errcode), format, args...)
}
