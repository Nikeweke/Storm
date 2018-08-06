/*
*  routes / api.go
*
*  API Routes
*/
package routes

import (
	"github.com/gorilla/mux"
	"../app/controllers"
)

var ApiCtrl controllers.ApiController

func Api(router *mux.Router) *mux.Router {
	api := router.PathPrefix("/api").Subrouter()

	// greeting from api
	api.HandleFunc("/", ApiCtrl.APICheck)

	return router
}
