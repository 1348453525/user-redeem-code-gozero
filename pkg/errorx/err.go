package errorx

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToGrpcError(err error) error {
	// if errors.Is(err, ErrInternal) {
	// 	return status.Error(codes.Internal, err.Error())
	// } else if errors.Is(err, ErrParam) {
	// 	return status.Error(codes.InvalidArgument, err.Error())
	// }
	// return status.Error(codes.Unknown, err.Error())

	// switch {
	// case errors.Is(err, ErrInternal):
	// 	return status.Error(codes.Code(ErrInternal.Code), ErrInternal.Message)
	// case errors.Is(err, ErrParam):
	// 	return status.Error(codes.Code(ErrParam.Code), ErrParam.Message)
	// default:
	// 	return status.Error(codes.Unknown, err.Error())
	// }

	var e *Errorx
	if errors.As(err, &e) {
		return status.Error(codes.Code(e.Code), e.Message)
	}
	return status.Error(codes.Unknown, err.Error())
}

// 系统错误
var (
	ErrInternal        = New(500, "内部错误")
	ErrParam           = New(400, "参数错误")
	ErrOperationFailed = New(500, "操作失败")
)

// 用户错误
var (
	ErrPasswordNotMatch = New(1401, "密码不匹配")
	ErrUserExisted      = New(1402, "用户已存在")
	ErrUserDisabled     = New(1403, "账户已停用")
	ErrPasswordError    = New(1404, "密码错误")
	ErrUserNotExisted   = New(1405, "用户不存在")
)

// 兑换码错误
var (
	ErrRedeemCodeUsedUp  = New(1401, "兑换码已用完")
	ErrRedeemCodeExpired = New(1402, "兑换码已过期")
	ErrRedeemCodeInvalid = New(1403, "兑换码无效")
)
