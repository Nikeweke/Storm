package helpers

import (
	"github.com/labstack/echo"
    color "github.com/fatih/color"

	"fmt"
	// "reflect"
	
	// "io/ioutil"
	// "encoding/json"
)

// type Route struct {
// 	Method string `json:"method"`
// 	Path   string `json:"path"`
// 	Name   string `json:"name"`
// } 


func ShowRoutes(router *echo.Echo) {
  var green  *color.Color  = color.New(color.FgGreen)   // GET color
  var yellow *color.Color  = color.New(color.FgYellow) // POST color
  var red    *color.Color  = color.New(color.FgRed)   // DELETE color
  var blue   *color.Color  = color.New(color.FgBlue)   // PUT color
  var magenta *color.Color = color.New(color.FgMagenta)  // for all others

  var selectedColor *color.Color

  for _, route := range router.Routes() {
	switch route.Method {
	  case "GET" :    selectedColor = green
	  case "POST":    selectedColor = yellow
	  case "PUT" :    selectedColor = blue
	  case "DELETE":  selectedColor = red
	  default: selectedColor = magenta
	}

	selectedColor.Print(`{`+ route.Method +`} `)
	fmt.Print(`"`+ route.Path +`"` + "\n")
  }
}



/**
* Make file with list of routes 
*/ 
// func GetRoutes(router *echo.Echo) {

//   var routesPaths  []string
//   var item         Route
//   var routes       []*echo.Route = router.Routes()
//   var outputRoutes []Route

//   // making array of unique routes paths
//   for _, route := range routes {
//     exists, _ := in_array(route.Path, routesPaths)
//     if !exists {
//       routesPaths = append(routesPaths, route.Path)
//     }
//   }

//   // making new array of routes 
//   for _, path := range routesPaths {
//     for _, route := range routes {

//       var isAnyRoute  = (path == route.Path && route.Method == "TRACE") 
//       var isOneMethod =  (path == route.Path && (route.Method == "GET" || route.Method == "POST" || route.Method == "PUT" || route.Method == "DELETE") &&  !isInRoutes(route.Path, outputRoutes))

//       // if route equal to route path and have method "trace"
//       // then its "any" defined route
//       if isAnyRoute {
//         item = Route{"ANY", path, route.Name}
//         outputRoutes = append(outputRoutes, item)

//       } else if isOneMethod {
//         item = Route{route.Method, route.Path, route.Name}
//         outputRoutes = append(outputRoutes, item)
//       }
//     }
//   }

//   fmt.Println("Length ==", len(outputRoutes))
//   for _, route := range outputRoutes {
//     fmt.Println(route)
//   }
// }


// func in_array(val interface{}, array interface{}) (exists bool, index int) {
//   exists = false
//   index = -1

//   switch reflect.TypeOf(array).Kind() {
//   case reflect.Slice:
//       s := reflect.ValueOf(array)

//       for i := 0; i < s.Len(); i++ {
//           if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
//               index = i
//               exists = true
//               return
//           }
//       }
//   }

//   return
// }


// func isInRoutes(routePath string, routes []Route) bool {
//   for _, route := range routes {
//     if routePath == route.Path {
//       return true
//     }
//   }
//   return false
// }

// func writeRoutesIntoFile() {
// 	data, err := json.MarshalIndent(router.Routes(), "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ioutil.WriteFile("routes.json", data, 0644)
// }