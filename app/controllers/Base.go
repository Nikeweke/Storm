package controllers

import (
    "html/template"
    "net/http"
	"path"
	// "fmt"
	"../../config/structs"
    "encoding/json"
)

// Иниц. все структуры используемые в пространстве "controllers"
type StringArray structs.StringArray
type ResponseMessage structs.ResponseMessage


/*
|--------------------------------------------------------------------------
|  Render page
|--------------------------------------------------------------------------
*/
func render(page string, data StringArray, res http.ResponseWriter) {
	pagePath :=  path.Join("views", page +".html")   // указываем путь к вьюхам(темплейтам)

	// head      :=  path.Join("views/partials", "head.html")
	// foot      :=  path.Join("views/partials", "footer.html")
	// navbar      :=  path.Join("views/partials", "navbar.html")

	// parsedPages := template.Must(template.ParseFiles(head, navbar, path_tmpl, foot))
	parsedPages :=  template.Must( template.ParseFiles(pagePath) )
	// if err != nil{ // если ошибка не пусто, показать сообщение в консоли
	//   fmt.Println(err)
	// }

	// parsedPages.ExecuteTemplate(res, page, data)
	parsedPages.Execute(res, data) 
}


/*
|--------------------------------------------------------------------------
|  Send Json response
|--------------------------------------------------------------------------
*/
func send(message interface{}, res http.ResponseWriter) {
  json, _ := json.Marshal(message)
  res.Header().Set("Content-Type", "application/json")
  res.Write(json)
}


/*
|--------------------------------------------------------------------------
|  Redirect
|--------------------------------------------------------------------------
*/
func redirect(path string, res http.ResponseWriter, req *http.Request) {
  http.Redirect(res, req, "/", 301)
}