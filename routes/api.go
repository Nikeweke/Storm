/*
*  routes / api.go
*
*  API Routes
*/
package routes

import (
	"github.com/labstack/echo"
	"../app/controllers"
	"./middlewares"
)

var ApiCtrl controllers.ApiController

// api/...
func Api(router *echo.Echo) {
	api := router.Group("/api")

	api.Use(middlewares.CheckRequest)

	api.GET("", ApiCtrl.CheckApi)   

}
