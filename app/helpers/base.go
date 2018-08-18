package helpers

import (
	"html/template"
	"net/http"
	"path"
	color "github.com/fatih/color"
	"encoding/json"
	"fmt"

	// "../../views"
	// "reflect"
)


/*
|--------------------------------------------------------------------------
|  Render page
|--------------------------------------------------------------------------
*/
func Render(page string, viewArgs map[string]interface{}, res http.ResponseWriter) error {
	pagePath :=  path.Join("views", page +".html")   // указываем путь к вьюхам(темплейтам)

	// head   :=  path.Join("views/partials", "head.html")
	// foot   :=  path.Join("views/partials", "footer.html")
	// navbar :=  path.Join("views/partials", "navbar.html")

	// parsedPages := template.Must(template.ParseFiles(head, navbar, path_tmpl, foot))
	parsedPages :=  template.Must( template.ParseFiles(pagePath) )

	return parsedPages.Execute(res, viewArgs) 
}



/*
|--------------------------------------------------------------------------
|  Render page
|--------------------------------------------------------------------------
*/
// func Render(page string, viewArgs map[string]interface{}, res http.ResponseWriter) error {
// 	pagePath :=  path.Join("views", page +".html")   // указываем путь к вьюхам(темплейтам)

// 	// head      :=  path.Join("views/partials", "head.html")
// 	// foot      :=  path.Join("views/partials", "footer.html")
// 	// navbar      :=  path.Join("views/partials", "navbar.html")

// 	// parsedPages := template.Must(template.ParseFiles(head, navbar, path_tmpl, foot))
// 	parsedPages :=  template.Must( template.ParseFiles(pagePath) )

// 	// fmt.Println(reflect.TypeOf(views.Assets)) // *assets.FileSystem
// 	// fmt.Println(reflect.TypeOf(parsedPages)) // *template.Template
// 	// fmt.Println(pagePath) // string 

// 	file, _ := views.Assets.Open("/views/home.html") // *assets.File
// 	fmt.Println(string(file.Data[:]))

// 	// c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
  
// 	return parsedPages.Execute(res, viewArgs) 
// 	// return parsedPages.ExecuteTemplate(res, page, data)
// }




/*
|--------------------------------------------------------------------------
|  Send Json response
|--------------------------------------------------------------------------
*/
func Send(message interface{}, res http.ResponseWriter) {
  json, _ := json.Marshal(message)
  res.Header().Set("Content-Type", "application/json")
  res.Write(json)
}


/*
|--------------------------------------------------------------------------
|  Redirect
|--------------------------------------------------------------------------
*/
func Redirect(path string, res http.ResponseWriter, req *http.Request) {
  http.Redirect(res, req, "/", 301)
}


/*
|--------------------------------------------------------------------------
|  Error Handler
|--------------------------------------------------------------------------
*/
func ErrorHandler(err error) {
	color.Red("ERROR ====================> ")
	fmt.Println(err)
	color.Red("ERROR ====================> ")
}