package controllers

import (
  "fmt"
	"net/http"
	// "../helpers"
	// "../../config/structs"
);

type IndexController struct {}

// Index - Home page
func (this IndexController) Index(res http.ResponseWriter, req *http.Request) {

	items := [...]string{"item1", "item2", "item3"}
	
	data := StringArray{
		 "title":    "Storm is greeting you!",
		 "greeting": "Storm",
		 "words":    "Fast and reliable",
		 "items": items,
	}	

  render("home", data, res)
}

// CheckRequest
func (this IndexController) CheckRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Its Cehck requests")
}