package config 

import (
	"github.com/gorilla/mux"
	"storm/app/controllers"
	"net/http"
)

var ApiCtrl   controllers.ApiController
var IndexCtrl controllers.IndexController

func Routes() *mux.Router {
  
	router := mux.NewRouter()

	// web
	router.HandleFunc("/",      IndexCtrl.Index).Methods("GET")
	router.HandleFunc("/check", IndexCtrl.CheckRequest)

	// api
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", ApiCtrl.APICheck)

	// public 
	// router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Show all register routes (in app/helpers/routes)
	// err := router.Walk(helpers.GetRoutes)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return router
}




