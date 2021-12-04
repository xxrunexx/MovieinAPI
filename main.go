package main

import (
	"movie-api/driver"
	_middleware "movie-api/middleware"
	"movie-api/routes"
)

func main() {
	driver.InitDB()
	e := routes.New()

	// Log Middleware
	_middleware.LogMiddlewareInit(e)

	// Starting the server
	e.Start(":8000")
}
