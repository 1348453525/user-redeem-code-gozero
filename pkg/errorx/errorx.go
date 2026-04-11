package errorx

import "fmt"

type Errorx struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) *Errorx {
	return &Errorx{
		Code:    code,
		Message: message,
	}
}

func (e Errorx) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
