package middlewares

import (
	"net/http"
	// "encoding/json"
	"../../config/structs"
	"fmt"
	"log"
	"../../app/helpers"
)

type StringArray structs.StringArray

func CheckRequest(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	
    checkedRequest := helpers.CheckRequest(req)
		
		log.Println("------------------------------> Check Request")
		fmt.Println("Method =",       checkedRequest["Method"])
		fmt.Println("Body =",         checkedRequest["Body"])
		fmt.Println("Query =",        checkedRequest["Query"])
		fmt.Println("Content-Type =", checkedRequest["Content-Type"])

	  next.ServeHTTP(res, req)
  })
}
