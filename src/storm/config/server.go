package config 

import (
	"net/http"
	 color "github.com/fatih/color"
	"github.com/gorilla/mux"
);

func Server(router *mux.Router) {
	port := "8000";
	color.Yellow("App is running on port: " + port)
	http.ListenAndServe(":" + port, router)
}