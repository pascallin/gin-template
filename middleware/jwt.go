package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pascallin/gin-template/config"
	"github.com/pascallin/gin-template/model"
	"github.com/sirupsen/logrus"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if len(auth) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := strings.Fields(auth)[1]

		jwtToken, err := jwt.ParseWithClaims(token, &model.CustomerClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(config.GetAppConfig().AppJwtSecret), nil
		})
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if claim, ok := jwtToken.Claims.(*model.CustomerClaims); ok && jwtToken.Valid {
			logrus.Debug(claim)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
