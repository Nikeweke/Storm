package helpers

import (
    "html/template"
    "net/http"
	"path"
	// "fmt"
    "encoding/json"
)


/*
|--------------------------------------------------------------------------
|  Render page
|--------------------------------------------------------------------------
*/
func Render(page string, viewArgs map[string]interface{}, res http.ResponseWriter) error {
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
	return parsedPages.Execute(res, viewArgs) 
}



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