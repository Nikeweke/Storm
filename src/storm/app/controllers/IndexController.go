package controllers

import (
  "fmt"
  "net/http"
);

type IndexController struct {
}

// Index - Home page
func (this IndexController) Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello its Index")
}

// CheckRequest - Checking request
func (this IndexController) CheckRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Its Cehck requests")
}