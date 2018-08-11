package config 

import (
  "github.com/labstack/echo"
  "../routes"
)

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




