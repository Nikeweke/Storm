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
	"routes": Mode{ "routes", "",    color.FgBlack     },
}


/**
* Start server and out info 
*/
func Server(router *echo.Echo) {
	router.HidePort = true   // disable echo http serve info
	router.HideBanner = true // disable echo banner in console

	var mode Mode
  consoleArgs := os.Args // []string
	
	if len(consoleArgs) > 1 {
		mode = MODES[consoleArgs[1]]
	} else {
		mode = MODES["dev"]
	}

	// output routes and exit 
	if mode.Name == "routes" {
		helpers.ShowRoutes(router)
		os.Exit(1)
	}

  outLaunchInfo(mode)
	router.Start(":" + mode.Port)
}


/**
*  Output colorized string about server starting 
*/
func outLaunchInfo(mode Mode) {
	c := color.New(mode.Color).Add(color.Bold)
	current_time := time.Now().Local().Format("15:04:05 01/02/2006")
	
	c.Print("[" + current_time + "] App ")
	fmt.Print("is running on port: ")
	c.Print(mode.Port + " ")
	fmt.Print("(");
	c.Print(mode.Name + " ") 
	fmt.Print("mode)")
}



