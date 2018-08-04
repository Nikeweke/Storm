package main

import (
	"fmt"
	// "log"
	"net/http"
	"github.com/gorilla/mux"
	color "github.com/fatih/color"
);

func main () {
	// fmt.Println("asdasdasdasd")
	// fmt.Scanln()

	// Gorilla router
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	http.Handle("/", r)  // init routes by gorilla router

	
	// START SERVER
	var port string = "8000";
	color.Yellow("App is running on port: " + port)
	http.ListenAndServe(":"+port, nil)
}


func Index(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, its Index ")
}