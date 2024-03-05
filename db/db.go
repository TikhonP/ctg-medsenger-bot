package db

import (
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Connect(cfg *appconfig.Database) {
	db = sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(schema)
}
