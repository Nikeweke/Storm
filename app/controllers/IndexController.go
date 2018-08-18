package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"../helpers"
)

type IndexController struct {}

/**
* Home page
*/
func (this IndexController) Index(c echo.Context) error {
	type StringArray helpers.StringArray

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

  return helpers.Render("home", viewArgs, c.Response())
}

/**
* Check Request
*/
func (this IndexController) CheckRequest(c echo.Context) error {
  var req *http.Request = c.Request()
  message := helpers.CheckRequest(req)
  return c.JSON(200, message)
}






