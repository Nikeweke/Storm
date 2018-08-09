package config 

import (
	"github.com/labstack/echo"
	//  color "github.com/fatih/color"
	// "github.com/gorilla/mux"
);

func Server(router *echo.Echo) {
	// port := "8000";
	// color.Yellow("App is running on port: " + port)
	// http.ListenAndServe(":" + port, router)

	router.HideBanner = true // disable echo banner in console
	router.Logger.Fatal(router.Start(":8000"))
}