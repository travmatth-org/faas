package logger

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

var instance *zerolog.Logger
var destination io.Writer = os.Stderr

func SetLogger(base *zerolog.Logger) {
	instance = base
}

func GetLogger() *zerolog.Logger {
	if instance != nil {
		return instance
	}

	zerolog.TimeFieldFormat = time.RFC1123
	base := zerolog.New(destination).With().
		Timestamp().
		Caller().
		Str("role", "faas").
		Logger()
	SetLogger(&base)
	return instance
}

func Error() *zerolog.Event {
	return GetLogger().Error()
}

func Info() *zerolog.Event {
	return GetLogger().Info()
}

func Warn() *zerolog.Event {
	return GetLogger().Warn()
}

func Debug() *zerolog.Event {
	return GetLogger().Debug()
}

func ErrorReq(r *http.Request) *zerolog.Event {
	return hlog.FromRequest(r).Error()
}

func InfoReq(r *http.Request) *zerolog.Event {
	return hlog.FromRequest(r).Info()
}

func WarnReq(r *http.Request) *zerolog.Event {
	return hlog.FromRequest(r).Warn()
}

func DebugReq(r *http.Request) *zerolog.Event {
	return hlog.FromRequest(r).Debug()
}