package pkg

import (
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/pascallin/gin-template/config"
)

func SetupLogger() {
	if config.GetAppConfig().AppEnv == "prod" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:03:04",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
			funcName := path.Base(frame.Function)
			return funcName, fileName
		},
	})
}
