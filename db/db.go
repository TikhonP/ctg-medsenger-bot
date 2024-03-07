package db

import (
	"fmt"
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// db is a global database.
var db *sqlx.DB

func dataSourceName(cfg *appconfig.Database) string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", cfg.User, cfg.Dbname, cfg.Password, cfg.Host)
}

// Connect creates a new in-memory SQLite database and initializes it with the schema.
func Connect(cfg *appconfig.Database) {
	db = sqlx.MustConnect("postgres", dataSourceName(cfg))
	db.MustExec(schema)
}
