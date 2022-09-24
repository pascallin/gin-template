package types

const (
	SucceedCode     = "SUCCEED"
	ParamErrorCode  = "PARAM_ERROR"
	SystemErrorCode = "SYSTEM_ERROR"
)

type AppError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, code string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
	}
}

var (
	ErrParam  = NewAppError(400, ParamErrorCode)
	ErrSystem = NewAppError(500, SystemErrorCode)
)
