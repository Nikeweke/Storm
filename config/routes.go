package config 

import (
  "github.com/gorilla/mux"
  "net/http"
  "../routes"
)

func Routes() *mux.Router {
  router := mux.NewRouter()
  
  router = routes.Web(router)
  router = routes.Api(router)

  // public 
  router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

  // Show all register routes (in app/helpers/routes)
  // err := router.Walk(helpers.GetRoutes)
  // if err != nil {
  // 	fmt.Println(err)
  // }

  return router
}




