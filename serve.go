package ctg_medsenger_bot

import (
	"fmt"
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/TikhonP/ctg-medsenger-bot/handler"
	"github.com/TikhonP/ctg-medsenger-bot/util"
	"github.com/TikhonP/maigo"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type handlers struct {
	root     *handler.RootHandler
	init     *handler.InitHandler
	status   *handler.StatusHandler
	remove   *handler.RemoveHandler
	settings *handler.SettingsHandler
}

func createHandlers(MAIClient *maigo.Client) *handlers {
	return &handlers{
		init: &handler.InitHandler{MAIClient: MAIClient},
	}
}

func Serve(cfg *appconfig.Server) {
	MAIClient := maigo.Init(cfg.MedsengerAgentKey)
	handlers := createHandlers(MAIClient)

	e := echo.New()
	e.Debug = cfg.Debug
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if !cfg.Debug {
		e.Use(sentryecho.New(sentryecho.Options{}))
	}
	e.Validator = util.NewDefaultValidator()

	e.GET("/", handlers.root.Handle)
	e.POST("/init", handlers.init.Handle, util.ApiKeyJSON(cfg))
	e.POST("/status", handlers.status.Handle, util.ApiKeyJSON(cfg))
	e.POST("/remove", handlers.remove.Handle, util.ApiKeyJSON(cfg))
	e.GET("/settings", handlers.settings.Handle, util.ApiKeyGetParam(cfg))

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	err := e.Start(addr)
	if err != nil {
		panic(err)
	}
}
