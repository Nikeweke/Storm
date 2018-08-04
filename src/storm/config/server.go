package config 

import (
	"net/http"
   	color "github.com/fatih/color"
);

func Server() {
    	// START SERVER
	var port string = "8000";
	color.Yellow("App is running on port: " + port)
	http.ListenAndServe(":"+port, nil)
}