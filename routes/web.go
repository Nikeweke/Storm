/*
*  routes / web.go
*
*  Web Routes
*/
package routes

import (
	"github.com/labstack/echo"
	"../app/controllers"
)

var IndexCtrl controllers.IndexController

func Web(router *echo.Echo) {

	router.GET("/",      IndexCtrl.Index)
	router.Any("/check", IndexCtrl.CheckRequest)

}

