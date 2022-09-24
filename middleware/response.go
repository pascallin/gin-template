package middleware

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pascallin/gin-template/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./locales/result.en.toml")
	// bundle.MustLoadMessageFile("./locales/result.zh.toml")
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ResponseInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		accept := c.Request.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, accept)

		c.Next()

		statusCode := c.Writer.Status()
		// fmt.Println(statusCode)

		var res types.AppResponse
		if err := json.Unmarshal(blw.body.Bytes(), &res); err != nil {
			logrus.Error(err)
			return
		}

		msg, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: res.Code,
		})

		if err == nil {
			res.Message = msg
		} else {
			logrus.Error(err)
		}

		logrus.Debug(res)

		c.JSON(statusCode, res)
	}
}
