package controllers


import (
  "fmt"
  "net/http"
);


func Index(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, its Index ")
}