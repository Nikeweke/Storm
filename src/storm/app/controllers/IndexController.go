package controllers

import (
  "fmt"
	"net/http"
	"storm/app/helpers"
	"storm/config/structs"
);

type IndexController struct {}
type StringArray structs.StringArray

var Template helpers.Template


// Index - Home page
func (this IndexController) Index(res http.ResponseWriter, req *http.Request) {
	data := StringArray{"title": "Главная страница"}	
	
  fmt.Println(data)

  // Template.Render("home", data, res)
	// fmt.Fprintf(res, "Hello its Index")
}

// CheckRequest - Checking request
func (this IndexController) CheckRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Its Cehck requests")
}