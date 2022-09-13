package pkg

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	defaultLang = "en"
	langCtxKey  = "lang"
)

var bundle *i18n.Bundle

func GetBuddle() *i18n.Bundle {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("./locales/result.en.toml")

	return bundle
}

func GetLocalizer(lang string) *i18n.Localizer {
	bundle := GetBuddle()
	tmp := defaultLang
	if lang != "" {
		tmp = lang
	}

	return i18n.NewLocalizer(bundle, tmp)
}

func GetLocalizerFromCtx(c *gin.Context) *i18n.Localizer {
	return GetLocalizer(c.GetString(langCtxKey))
}

func GetI18nMessage(c *gin.Context, key string) string {
	localizer := GetLocalizer(c.GetString(langCtxKey))
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: key,
	})
}
