package controllers

import (
  "net/http"
);

type ApiController struct {}

/**
* APICheck - check if api works
*/
func (this ApiController) APICheck(res http.ResponseWriter, req *http.Request) {
  resMsg  := ResponseMessage{1, "That is API, Hello from Storm!"}
  // OR
  // resMsg  := StringArray{ "success": 1, "message": "That is API, Hello from Storm!" } 
  Send(resMsg, res)
}
