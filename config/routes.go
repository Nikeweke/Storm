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


  // get Routes
//   data, err := json.MarshalIndent(e.Routes(), "", "  ")
// if err != nil {
// 	return err
// }
// ioutil.WriteFile("routes.json", data, 0644)

  return router
}




