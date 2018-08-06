package controllers

import (
  // "fmt"
	"net/http"
	// "encoding/json"
	"../helpers"
);

type IndexController struct {}

/**
* Index - Home page
*/
func (this IndexController) Index(res http.ResponseWriter, req *http.Request) {
	items := [...]string{
		"item1", 
		"item2", 
		"item3",
	}

  itemsNested := [...]StringArray{
		{"id": 1, "name": "Krake"},
		{"id": 2, "name": "Ivan"},
		{"id": 3, "name": "Human"},
	} 

	viewArgs := StringArray{
		 "title":    "Storm is greeting you!",
		 "greeting": "Storm",
		 "words":    "Fast and reliable",
		 "items": items,
		 "items2": itemsNested,
	}	

  Render("home", viewArgs, res)
}

/**
* Check Request
*/
func (this IndexController) CheckRequest(res http.ResponseWriter, req *http.Request) {
	requestInfo := helpers.CheckRequest(req)
  Send(requestInfo, res)
}