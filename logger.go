package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NewLogger(conf *LoggerConfig) *log.Logger {

	logger := log.New()

	switch {
	case strings.ToUpper(conf.Level) == "TRACE":
		logger.SetLevel(log.TraceLevel)
	case strings.ToUpper(conf.Level) == "DEBUG":
		logger.SetLevel(log.DebugLevel)
	case strings.ToUpper(conf.Level) == "INFO":
		logger.SetLevel(log.InfoLevel)
	case strings.ToUpper(conf.Level) == "WARN":
		logger.SetLevel(log.WarnLevel)
	case strings.ToUpper(conf.Level) == "ERROR":
		logger.SetLevel(log.ErrorLevel)
	case strings.ToUpper(conf.Level) == "FATAL":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.Info("setting default logging to INFO")
		logger.SetLevel(log.InfoLevel)
	}

	switch {
	case strings.ToUpper(conf.Format) == "TEXT":
		logger.SetFormatter(&log.TextFormatter{
			TimestampFormat: conf.TimestampFormat,
			FullTimestamp:   conf.FullTimestamp,
			DisableColors:   conf.DisableColor,
		})
	case strings.ToUpper(conf.Format) == "JSON":
		logger.SetFormatter(&log.JSONFormatter{
			TimestampFormat: conf.TimestampFormat,
			PrettyPrint:     conf.PrettyPrint,
		})
	default:
		logger.WithField("format", conf.Format).Warn("setting default format to TEXT as no valid logging format provided")
	}

	switch {
	case strings.ToUpper(conf.Output) == "STDOUT":
		logger.SetOutput(os.Stdout)
	case strings.ToUpper(conf.Output) == "STDERR":
		logger.SetOutput(os.Stderr)
	}

	switch {
	case conf.ReportCaller:
		logger.SetReportCaller(true)
	default:
		logger.SetReportCaller(false)
	}

	return logger
}
