package main

import (
	"backend/config"
	"backend/driver/mysql"
	fr "backend/faktoryandroute"
	"backend/migrasi"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := prostgest.InitDB(cfg)
	migrasi.MigrateDB(db)

	e := echo.New()

	fr.FaktoryAndRoute(e, db)
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", cfg.SERVER_PORT)))
}
