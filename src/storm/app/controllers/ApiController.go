package controllers

import (
  // "fmt"
  "net/http"
  "encoding/json"
  "storm/config/structs"
);

type ApiController struct {}
type ResponseMessage structs.ResponseMessage

// APICheck - checl if api works
func (this ApiController) APICheck(res http.ResponseWriter, req *http.Request) {

  resMsg  := ResponseMessage{1, "That is API, Hello from Storm!"}
  json, _ := json.Marshal(resMsg)

  res.Header().Set("Content-Type", "application/json")
  res.Write(json)
}
