package main

import (
	"context"
	ctg "github.com/TikhonP/ctg-medsenger-bot"
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/getsentry/sentry-go"
)

func setupSentry(dsn string) {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/local/app_config.pkl")
	if err != nil {
		panic(err)
	}
	if !cfg.Server.Debug {
		setupSentry(cfg.SentryDSN)
	}
	db.Connect(cfg.Db)
	ctg.Serve(cfg.Server)
}
