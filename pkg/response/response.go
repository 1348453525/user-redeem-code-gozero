package response

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 统一响应体结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 业务数据
}

// 成功响应（固定code=200，message=success）
func Success(w http.ResponseWriter, data interface{}) {
	if data == nil {
		data = struct{}{} // 空数据时返回空对象，避免null
	}
	httpx.OkJson(w, Response{
		Code:    http.StatusOK, // 200
		Message: "success",
		Data:    data,
	})
}

// 失败响应（自定义状态码和提示信息）
func Error(w http.ResponseWriter, code int, message string) {
	httpx.OkJson(w, Response{
		Code:    code,
		Message: message,
		Data:    struct{}{},
	})
}

func Errorx(w http.ResponseWriter, err error) {
	var response Response

	switch e := err.(type) {
	case *errorx.Errorx:
		response.Code = e.Code
		response.Message = e.Message
	default:
		response.Code = 500
		response.Message = e.Error()
	}

	response.Data = struct{}{}

	httpx.OkJson(w, response)
}
