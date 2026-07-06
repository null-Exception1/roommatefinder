package init

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Logging() {
	format := os.Getenv("LOG_FORMAT")

	switch format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
	}

	level := os.Getenv("LOG_LEVEL")

	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "none":
		log.SetOutput(io.Discard)
	default:
		log.SetLevel(log.InfoLevel)
	}

	logrus.WithFields(logrus.Fields{
		"package": "init",
	}).Info("initiated logrus")

}
