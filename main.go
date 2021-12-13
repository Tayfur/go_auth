package main

import (
	"go_auth/pkg/db"
	"go_auth/pkg/redis"
	"go_auth/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// db connection
	db, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}
	//redis connection
	redis := redis.ConnectRedis()
	e := echo.New()
	// Router
	routes.SetupRoutes(e, db, redis)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
