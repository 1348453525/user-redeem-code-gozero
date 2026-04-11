package userlogic

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/pkg/helper"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/zeromicro/go-zero/core/logx"
)

// 密码加密选项常量
const (
	PasswordSaltLen    = 16
	PasswordIterations = 10000 // 提高迭代次数增强安全性
	PasswordKeyLen     = 32
)

// getPasswordOptions 获取密码加密选项
func getPasswordOptions() *password.Options {
	return &password.Options{
		SaltLen:      PasswordSaltLen,
		Iterations:   PasswordIterations,
		KeyLen:       PasswordKeyLen,
		HashFunction: sha512.New,
	}
}

// encodePassword 加密密码
func encodePassword(pwd string) string {
	salt, encodedPwd := password.Encode(pwd, getPasswordOptions())
	return fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// verifyPassword 验证密码
func verifyPassword(pwd, encrypted string) bool {
	parts := strings.Split(encrypted, "$")
	if len(parts) != 3 {
		return false
	}
	return password.Verify(pwd, parts[1], parts[2], getPasswordOptions())
}

// parseBirthday 解析生日字符串为时间指针
func parseBirthday(birthdayStr string) *time.Time {
	if birthdayStr == "" {
		return nil
	}
	if t, err := time.Parse("2006-01-02", birthdayStr); err == nil {
		return &t
	}
	return nil
}

// buildUserInfoDvo 构建用户信息
func buildUserInfoDvo(user *model.User) proto.UserInfoResponse {
	return proto.UserInfoResponse{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Birthday: helper.FormatDate(user.Birthday),
	}
}

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(r *proto.RegisterRequest) (*proto.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	// 验证密码是否一致
	if r.Password != r.ConfirmPassword {
		return nil, errorx.ToGrpcError(errorx.ErrPasswordNotMatch)
	}

	// 验证用户名长度
	if len(r.Username) < 3 || len(r.Username) > 20 {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证密码长度
	if len(r.Password) < 6 {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证手机号格式
	if r.Mobile != "" && !helper.IsValidMobile(r.Mobile) {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 查询用户是否存在
	var count int64
	result := l.svcCtx.DB.Model(&model.User{}).Where("username = ?", r.Username).Or("mobile = ?", r.Mobile).Count(&count)
	if result.Error != nil {
		logx.Errorw("查询用户失败", logx.Field("err", result.Error))
		return nil, errorx.ToGrpcError(errorx.ErrInternal)
	}
	if count > 0 {
		logx.Error("用户已存在")
		return nil, errorx.ToGrpcError(errorx.ErrUserExisted)
	}

	// 生成密码
	pwd := encodePassword(r.Password)

	// 处理生日
	birthday := parseBirthday(r.Birthday)

	// 创建用户
	newUser := model.User{
		Username: r.Username,
		Password: pwd,
		Nickname: r.Nickname,
		Mobile:   r.Mobile,
		Gender:   r.Gender,
		Birthday: birthday,
		IsDel:    2,
	}
	result = l.svcCtx.DB.Create(&newUser)
	if result.Error != nil {
		return nil, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}

	// 返回数据
	resp := buildUserInfoDvo(&newUser)
	return &resp, nil
}
