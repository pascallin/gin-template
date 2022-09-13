package types

type AppResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAppResponse(code string, msg string) AppResponse {
	return AppResponse{
		Code:    code,
		Message: msg,
	}
}
