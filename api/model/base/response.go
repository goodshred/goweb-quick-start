package base

import "github.com/pkg/errors"

// ServerResponse common response
type ServerResponse struct {
	Code  int         `json:"code"`  // 0 为成功，前端需要区分 401，403 等特殊 http code
	Error string      `json:"error"` // status 不为 true 的时候填入原因
	Data  interface{} `json:"data"`  // status 为 true 的时候填入结果
}

// ResponseData 成功
func ResponseData(data interface{}) (dgr ServerResponse) {
	dgr.Data = data
	return
}

// ResponseErr 失败
func ResponseErr(err error) (dgr ServerResponse) {
	dgr.Error = err.Error()
	dgr.Code = 500
	return
}

// ResponseErrWrap 失败包
func ResponseErrWrap(err error, extra string) (dgr ServerResponse) {
	dgr.Error = errors.Wrap(err, extra).Error()
	dgr.Code = 400
	return
}

// ResponseErrUnAuth 鉴权失败
func ResponseErrUnAuth() (dgr ServerResponse) {
	dgr.Error = "401 Unauthorized"
	dgr.Code = 401
	return
}

// ResponseErrForbidden 鉴权成功但是权限不足
func ResponseErrForbidden() (dgr ServerResponse) {
	dgr.Error = "403 Forbidden"
	dgr.Code = 403
	return
}

// Response 成功
func ResponseSuccess() (dgr ServerResponse) {
	dgr.Data = "success"
	return
}
