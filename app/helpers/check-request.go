package helpers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../../config/structs"
);
  
type StringArray structs.StringArray

func CheckRequest(req *http.Request) StringArray {

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
	
	return StringArray{
		"Method": req.Method,
		"Body"  : body,
		"Query" : req.URL.Query(),   
		"Content-Type": contentType, 
	}
}