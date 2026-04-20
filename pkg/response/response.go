package response

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if grpcStatus, ok := status.FromError(e); ok {
			response.Code = grpcCodeToHttpCode(grpcStatus.Code())
			response.Message = grpcStatus.Message()
		} else {
			response.Code = 500
			response.Message = e.Error()
		}
	}

	response.Data = struct{}{}

	httpx.OkJson(w, response)
}

func grpcCodeToHttpCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return 200 // http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return 400 // http.StatusBadRequest
	case codes.Unauthenticated:
		return 401 // http.StatusUnauthorized
	case codes.NotFound:
		return 404 // http.StatusNotFound
	case codes.AlreadyExists, codes.Aborted:
		return 409 // http.StatusConflict
	case codes.PermissionDenied:
		return 403 // http.StatusForbidden
	case codes.ResourceExhausted:
		return 429 // http.StatusTooManyRequests
	case codes.Canceled:
		return 499
	case codes.Unknown, codes.Internal, codes.DataLoss:
		return 500 // http.StatusInternalServerError
	case codes.Unimplemented:
		return 501 // http.StatusNotImplemented
	case codes.Unavailable:
		return 503 // http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return 504 // http.StatusGatewayTimeout
	default:
		// 判断 code 是 err.go 中自定义的错误码，返回自定的错误码
		if code >= 1000 {
			return int(code)
		}
		return 500 // http.StatusInternalServerError
	}
}
