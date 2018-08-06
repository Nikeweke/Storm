package middlewares

import (
	"net/http"
	"encoding/json"
	"../../config/structs"
	"fmt"
	"log"
)

type StringArray structs.StringArray

func CheckRequest(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var contentType string = req.Header.Get("Content-type")
		var body interface{}

		if contentType == "application/x-www-form-urlencoded" {
			req.ParseForm()
			body = req.PostForm
		// req.Form - has and body and query together
		// req.Form.Get("someKey")
		// req.FormValue("someKey")

		} else if contentType == "application/json" {
			body = &StringArray{}
			if err := json.NewDecoder(req.Body).Decode(body); err != nil {
				fmt.Println("Error decoding body: %s", err)
				return
			}
			// body.someKey
		} 

		// req.URL.Query()["someKey"]
		
		log.Println("------------------------------> Check Request")
		fmt.Println("Method =", req.Method)
		fmt.Println("Body =", body)
		fmt.Println("Query =", req.URL.Query())
		fmt.Println("Content-Type =", contentType)

	  next.ServeHTTP(res, req)
  })
}
