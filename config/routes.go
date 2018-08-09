package config 

import (
  "github.com/labstack/echo"
  "../app/controllers"
  "../routes"
)

var IndexCtrl controllers.IndexController
var ApiCtrl   controllers.ApiController

func Routes() *echo.Echo {
  router := echo.New()
  
  // web 
  routes.Web(router)
  
  // api 
  routes.Api(router)
  
  // public 
  router.Static("/public", "public")

  return router
}




