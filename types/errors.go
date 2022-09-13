package types

import "errors"

const (
	SucceedCode     = "SUCCEED"
	ParamErrorCode  = "PARAM_ERROR"
	SystemErrorCode = "SYSTEM_ERROR"
)

var (
	ErrParam  = errors.New(ParamErrorCode)
	ErrSystem = errors.New(SystemErrorCode)
)

type AppError struct {
	StatusCode uint   `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode uint, Code string, msg string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       Code,
		Message:    msg,
	}
}
