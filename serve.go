package ctg_medsenger_bot

import (
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/TikhonP/ctg-medsenger-bot/handler"
	"github.com/TikhonP/ctg-medsenger-bot/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

type handlers struct {
	root     *handler.RootHandler
	init     *handler.InitHandler
	status   *handler.StatusHandler
	remove   *handler.RemoveHandler
	settings *handler.SettingsHandler
}

func createHandlers() *handlers {
	return new(handlers)
}

func Serve(cfg *appconfig.Server) {
	handlers := createHandlers()

	e := echo.New()
	e.Debug = cfg.Debug
	e.Use(middleware.Logger())
	e.Validator = util.NewDefaultValidator()

	e.GET("/", handlers.root.Handle)
	e.POST("/init", handlers.init.Handle, util.ApiKeyJSON(cfg))
	e.POST("/status", handlers.status.Handle, util.ApiKeyJSON(cfg))
	e.POST("/remove", handlers.remove.Handle, util.ApiKeyJSON(cfg))
	e.GET("/settings", handlers.settings.Handle, util.ApiKeyGetParam(cfg))

	addr := cfg.Host + ":" + strconv.Itoa(int(cfg.Port))
	err := e.Start(addr)
	if err != nil {
		panic(err)
	}
}
