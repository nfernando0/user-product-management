package main

import (
	"fmt"
	"products/db"
	"products/migration"
	"products/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db.DatabaseInit()
	migration.RunMigration()

	routes.RouteInit(e.Group("/api"))

	fmt.Println("Server running on port 5000")
	e.Logger.Fatal(e.Start(":5000"))
}
