package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/types"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if length := len(c.Errors); length > 0 {
			err := c.Errors[length-1].Err
			var Err *types.AppError
			switch err {
			case types.ErrParam:
				if e, ok := err.(*types.AppError); ok {
					Err = e
				}
				c.JSON(http.StatusBadRequest, types.NewAppResponse(Err.Code, Err.Message))
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}
}
