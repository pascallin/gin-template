package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/pascallin/gin-template/types"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept-Language")
		c.Set(types.LangCtxKey, accept)
		c.Next()
	}
}
