package helpers

import (
	"github.com/labstack/echo"
    color "github.com/fatih/color"

	"fmt"
	"reflect"
	
	"io/ioutil"
	"encoding/json"
)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Name   string `json:"name"`
} 

/**
* Show routes and write into file without superfluos(лишней) data
*/ 
func ShowRoutes(router *echo.Echo) {

  var routesPaths  []string
  var routes       []*echo.Route = router.Routes()
  var outputRoutes []Route

  // making array of unique routes paths
  for _, route := range routes {
    exists, _ := in_array(route.Path, routesPaths)
    if !exists {
      routesPaths = append(routesPaths, route.Path)
    }
  }

  // making new array of routes 
  var methods    int
  var methodName string
  var fileName   string
  for _, path := range routesPaths {
    methods = 0
    methodName = ""

    for _, route := range routes {
      if path == route.Path {
        methods++
        methodName = route.Method
        fileName   = route.Name
      }
    }

    if !isInRoutes(path, outputRoutes) && methods == 1 {
      outputRoutes = append(outputRoutes, Route{methodName, path, fileName})

    } else if !isInRoutes(path, outputRoutes) && methods > 2 {
      outputRoutes = append(outputRoutes, Route{"FEW", path, fileName})
    } 
  }

  // fmt.Println("Length ==", len(outputRoutes))
  writeRoutesIntoFile(outputRoutes)
  for _, route := range outputRoutes {
    // fmt.Println(route)
    color := getColorForMethod(route.Method)

    color.Print("{" + route.Method + "} ")
    fmt.Print(route.Path + "\n")
  }
}


/**
*  Check if value in array
*/
func in_array(val interface{}, array interface{}) (exists bool, index int) {
  exists = false
  index = -1

  switch reflect.TypeOf(array).Kind() {
  case reflect.Slice:
      s := reflect.ValueOf(array)

      for i := 0; i < s.Len(); i++ {
          if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
              index = i
              exists = true
              return
          }
      }
  }

  return
}


/**
*  Check if route in array of routes for output
*/
func isInRoutes(routePath string, routes []Route) bool {
  for _, route := range routes {
    if routePath == route.Path {
      return true
    }
  }
  return false
}


/**
*  Get color by method
*/
func getColorForMethod(method string) *color.Color {
	switch method {
    case "GET" :    return color.New(color.FgGreen)
    case "POST":    return color.New(color.FgYellow)
    case "PUT" :    return color.New(color.FgBlue)
    case "DELETE":  return color.New(color.FgRed)
    default:        return color.New(color.FgMagenta)
  }
}


/**
*  Write routes into file
*/
func writeRoutesIntoFile(routes interface{}) {
	data, err := json.MarshalIndent(routes, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)
}



// func ShowRoutes(router *echo.Echo) {
//   var green  *color.Color  = color.New(color.FgGreen)   // GET color
//   var yellow *color.Color  = color.New(color.FgYellow) // POST color
//   var red    *color.Color  = color.New(color.FgRed)   // DELETE color
//   var blue   *color.Color  = color.New(color.FgBlue)   // PUT color
//   var magenta *color.Color = color.New(color.FgMagenta)  // for all others

//   var selectedColor *color.Color

//   for _, route := range router.Routes() {
// 	switch route.Method {
// 	  case "GET" :    selectedColor = green
// 	  case "POST":    selectedColor = yellow
// 	  case "PUT" :    selectedColor = blue
// 	  case "DELETE":  selectedColor = red
// 	  default: selectedColor = magenta
// 	}

// 	selectedColor.Print(`{`+ route.Method +`} `)
// 	fmt.Print(`"`+ route.Path +`"` + "\n")
//   }
// }
