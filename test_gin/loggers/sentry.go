package logger

import (
	"test_gin/config"
	"time"

	"github.com/getsentry/sentry-go"
)

var initialized bool

func InitSentry() error {
	if initialized {
		return nil
	}

	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 3
	err := sentry.Init(sentry.ClientOptions{
		Dsn:                config.CFG.V.GetString("DSN"),
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
		Environment:        config.CFG.V.GetString("SENTRY_ENVIRONMENT"),
		Transport:          sentrySyncTransport,
	})
	if err != nil {
		return err
	}

	initialized = true

	defer sentry.Flush(2 * time.Second)
	return nil
}
