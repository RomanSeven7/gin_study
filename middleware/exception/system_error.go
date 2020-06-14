package exception

import "fmt"

type SystemError struct {
	error
	Code    int
	Message string
}

func (p *SystemError) Error() string {
	return fmt.Sprintf("SystemError code: %v, message %s", p.Code, p.Message)
}

func NewSystemError(code int, message string) *SystemError {
	return &SystemError{
		Code:    code,
		Message: message,
	}
}
