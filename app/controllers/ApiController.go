package controllers

import (
  "net/http"
  "github.com/labstack/echo"
)

type ApiController struct {}

/**
* CheckApi - check if api works
*/
func (this ApiController) CheckApi(c echo.Context) error {
  type StringArray map[string]interface{}
  var message = StringArray{
    "success": 1,
    "message": "Hello this is Storm gretting you!",
  }
  return c.JSON(http.StatusOK, message)
}
