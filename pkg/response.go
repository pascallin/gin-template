package pkg

import (
	"github.com/gin-gonic/gin"

	"github.com/pascallin/gin-template/types"
)

func NewSucceedAppResponse(c *gin.Context, data interface{}) types.AppResponse {
	return types.AppResponse{
		Code:    types.SucceedCode,
		Message: GetI18nMessage(c, types.SucceedCode),
	}
}
