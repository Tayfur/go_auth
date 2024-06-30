package main

import (
	"go_auth/cmd/db"
	"go_auth/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// db connection
	db, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	// Router
	routes.SetupRoutes(e, db)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
