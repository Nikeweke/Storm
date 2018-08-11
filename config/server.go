package config 

import (
	"github.com/labstack/echo"
	color "github.com/fatih/color"
	"log"
	"fmt"
	"os"
  // "../app/helpers"
  // "reflect"
)

type MapNested map[string]map[string]string
type Map map[string]string

var MODES = MapNested{
	"dev":    Map{ "name": "dev",   "port": "8000",  "color": "yellow",  },
	"test":   Map{ "name": "test",  "port": "3000",  "color": "magenta", },
	"prod":   Map{ "name": "prod",  "port": "80",    "color": "green",   },
	"routes": Map{ "name": "routes" },
}

func Server(router *echo.Echo) {

	router.HidePort = true   // disable echo http serve info
	router.HideBanner = true // disable echo banner in console

	consoleArgs := os.Args // []string
  modeArg     := consoleArgs[1]
	var mode Map
  
  if modeArg == "dev" || modeArg == "prod" || modeArg == "test" {
    mode = MODES[modeArg]
  } else {
    mode = MODES["dev"]
  }

  startServerColorInfo(mode)
	router.Start(":8000")
}

func startServerColorInfo(mode Map) {
	c := color.New(color.FgCyan).Add(color.Underline)
	log.Print("")
	c.Print("App ")
	fmt.Print("is running on port: ")
	c.Print(mode["port"] + " ")
	fmt.Print("(");
	c.Print(mode["name"] + " ") 
	fmt.Print("mode)")
}


// что такое process.argv[...] ?
  // это просто параметры которые можно ловить с консоли - node  app  <modeName>
  //                                                        [0]   [1]  [2]
//   let modeName = process.argv[2]

//   // если просмотр маршрутов
//   if (modeName === 'routes') {
//     console.log( getRoutes(app) )    
//     process.exit(1)
//   }

//   // определяем порт (по ум. - 8000). 
//   let mode = modeName ? MODES[modeName] : MODES.dev 

//   // запуск сервера
//   app.listen(mode.port, () => {
//     let coloredMsg = []
//     for (let key of Object.keys(mode)) {
//       coloredMsg.push(colors.bold[mode.color](mode[key]))
//     }
//     coloredMsg[2] = colors.bold[mode.color]('App')

//     util.log(`${coloredMsg[2]} is running on port: ${coloredMsg[1]} (${coloredMsg[0]} mode)`)
// })