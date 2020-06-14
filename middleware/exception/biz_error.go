package exception

import "fmt"

type BizError struct {
	error
	Code    int
	Message string
}

func (p *BizError) Error() string {
	return fmt.Sprintf("BizError code: %v, message %s", p.Code, p.Message)
}

func NewBizError(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}
