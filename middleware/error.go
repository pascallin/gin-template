package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"

	"github.com/pascallin/gin-template/pkg"
	"github.com/pascallin/gin-template/types"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if length := len(c.Errors); length > 0 {
			err := c.Errors[length-1].Err
			localizer := pkg.GetLocalizerFromCtx(c)

			logrus.WithField("err", err).Debug("error middleware get error")

			if e, ok := err.(*types.AppError); ok {
				var msg string
				msg, _ = localizer.Localize(&i18n.LocalizeConfig{
					MessageID: e.Code,
				})
				c.JSON(e.StatusCode, types.NewAppResponse(e.Code, msg))
				return
			}

			c.JSON(http.StatusInternalServerError, types.NewAppResponse("INTERNAL_ERROR", "unknown error"))

		}
	}
}
