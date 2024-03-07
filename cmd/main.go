package main

import (
	"context"
	ctg "github.com/TikhonP/ctg-medsenger-bot"
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/TikhonP/ctg-medsenger-bot/db"
)

func main() {
	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/local/app_config.pkl")
	if err != nil {
		panic(err)
	}
	db.Connect(cfg.Db)
	ctg.Serve(cfg.Server)
}
