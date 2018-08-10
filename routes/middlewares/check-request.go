package middlewares

import (
	"github.com/labstack/echo"
	// "net/http"
	// "encoding/json"
	"fmt"
	"log"
	"../../app/helpers"
)

type StringArray helpers.StringArray

func CheckRequest(next echo.HandlerFunc) echo.HandlerFunc  {
  return func(c echo.Context) error {
	
    checkedRequest := helpers.CheckRequest(c.Request())
		
		log.Println("------------------------------> Check Request")
		fmt.Println("Method =",       checkedRequest["Method"])
		fmt.Println("Body =",         checkedRequest["Body"])
		fmt.Println("Query =",        checkedRequest["Query"])
		fmt.Println("Content-Type =", checkedRequest["Content-Type"])

	  return next(c)
  }
}


