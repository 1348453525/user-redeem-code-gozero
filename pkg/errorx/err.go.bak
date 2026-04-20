package errorx

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToGrpcError(err error) error {
	if errors.Is(err, ErrInternal) {
		return status.Error(codes.Internal, err.Error())
	} else if errors.Is(err, ErrParam) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return status.Error(codes.Unknown, err.Error())
}

// 系统错误
var (
	ErrInternal = errors.New("内部错误")

	ErrParam           = errors.New("参数错误")
	ErrOperationFailed = errors.New("操作失败")
)

// 用户错误
var (
	ErrPasswordNotMatch = errors.New("密码不匹配")
	ErrUserExisted      = errors.New("用户已存在")
	ErrUserDisabled     = errors.New("账户已停用")
	ErrPasswordError    = errors.New("密码错误")
	ErrUserNotExisted   = errors.New("用户不存在")
)

// 兑换码错误
var (
	ErrRedeemCodeUsedUp  = errors.New("兑换码已用完")
	ErrRedeemCodeExpired = errors.New("兑换码已过期")
	ErrRedeemCodeInvalid = errors.New("兑换码无效")
)
