/*
*  routes / web.go
*
*  Web Routes
*/
package routes

import (
	"github.com/gorilla/mux"
	"../app/controllers"
	"./middlewares"
)

var IndexCtrl controllers.IndexController

func Web(router *mux.Router) *mux.Router {

	// home 
	router.HandleFunc("/", IndexCtrl.Index).Methods("GET")

    // check request - gives info about request
	router.HandleFunc("/check", IndexCtrl.CheckRequest)

	// test middleware
	router.Use(middlewares.CheckRequest)

	return router
}

