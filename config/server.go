package config 

import (
	"github.com/labstack/echo"
	color "github.com/fatih/color"
	"fmt"
	"os"
	"time"
	"../app/helpers"
)

type Mode struct {
	Name string 
	Port string 
	Color color.Attribute 
}

var MODES = map[string]Mode{
	"dev":    Mode{ "dev",  "8000",  color.FgYellow,  },
	"test":   Mode{ "test", "3000",  color.FgMagenta, },
	"prod":   Mode{ "prod", "80",    color.FgGreen,   },
}


/**
* Start server and out info 
*/
func Server(router *echo.Echo) {
	router.HidePort = true   // disable echo http serve info
	router.HideBanner = true // disable echo banner in console

	var mode Mode
  consoleArgs := os.Args // []string

	// check if one time command, execute and make exit
	checkOnOneTimeCommand(consoleArgs, router)
	
	if len(consoleArgs) > 1 && MODES[consoleArgs[1]].Name != "" {
		mode = MODES[consoleArgs[1]]
	} else {
		mode = MODES["dev"]
	}

  outLaunchInfo(mode)
	router.Start(":" + mode.Port)
}


/**
* One time commands
*/
func checkOnOneTimeCommand(consoleArgs []string, router *echo.Echo) {
  if len(consoleArgs) < 2 {
    return
	}

	arg1 := consoleArgs[1]

	// output routes
	if arg1 == "routes" {
		helpers.ShowRoutes(router, false)
		os.Exit(1)

  // output routes and write into file		
	} else if arg1 == "routes-file" {
		helpers.ShowRoutes(router, true)
  	os.Exit(1)

	} else {
		return 
	}
}


/**
*  Output colorized string about server starting 
*/
func outLaunchInfo(mode Mode) {
	c := color.New(mode.Color).Add(color.Bold)
	current_time := time.Now().Local().Format("15:04:05 01/02/2006")
	
	c.Print("App ")
	fmt.Print("is running on port: ")
	c.Print(mode.Port + " ")
	fmt.Print("(");
	c.Print(mode.Name + " ") 
	fmt.Print("mode)" + "[" + current_time + "] \n")
}



