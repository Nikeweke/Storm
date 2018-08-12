/*
*  routes / web.go
*
*  Web Routes
*/
package routes

import (
	"github.com/labstack/echo"
	"../app/controllers"
	// "./middlewares"
)

var IndexCtrl controllers.IndexController

func Web(router *echo.Echo) {

	router.POST("/",      IndexCtrl.Index)
	router.DELETE("/check", IndexCtrl.CheckRequest)
	router.PUT("/check2", IndexCtrl.CheckRequest)

}

