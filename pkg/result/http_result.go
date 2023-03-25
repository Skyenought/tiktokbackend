package result

import (
	"fmt"
	"net/http"

	"github.com/Skyenought/tiktokbackend/pkg/errno"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type RespSuccess struct {
	StatusCode uint32      `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"data"`
}

func Success(data interface{}) *RespSuccess {
	return &RespSuccess{200, "OK", data}
}

type RespError struct {
	StatusCode uint32 `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Error(errcode uint32, errmsg string) *RespError {
	return &RespError{errcode, errmsg}
}

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		errcode := errno.ServiceErrCode
		errmsg := "服务器开小差啦，稍后再来试一试"
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errno.ErrNo); ok {
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(gstatus.Code())
				if errno.IsCodeErr(grpcCode) {
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[API-ERR]: %+v", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	} else {
		httpx.WriteJson(w, http.StatusOK, resp)
	}
}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", errno.MapErrMsg(errno.ParamErrCode), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(errno.ParamErrCode, errMsg))
}
