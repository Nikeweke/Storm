/*
*  routes / api.go
*
*  API Routes
*/
package routes

import (
	"github.com/labstack/echo"
	"../app/controllers"
	// "./middlewares"
	"github.com/labstack/echo/middleware"
)

var ApiCtrl controllers.ApiController

// api/...
func Api(router *echo.Echo) {
	api := router.Group("/api")

	// # Using middleware for this group (custom)
	// api.Use(middlewares.CheckRequest)

	// # Using middleware from Echo library (needs install jwt-go package)
	// api.Use(middleware.Logger())

	// # Using middleware logger formatted  (needs install jwt-go package)
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	  Format: "{${method}} \"${uri}\" (status=${status})\n",
	}))

	api.Any("", ApiCtrl.CheckApi)   
}
