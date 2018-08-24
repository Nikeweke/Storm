package helpers

import (
	"html/template"
	"net/http"
	"path"
	color "github.com/fatih/color"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"io/ioutil"
)

var CompiledAssets map[string]*template.Template


/*
|--------------------------------------------------------------------------
|  Render view
|--------------------------------------------------------------------------
*/
func Render(page string, viewArgs map[string]interface{}, res http.ResponseWriter) error {
	// указываем путь к вьюхе
	pagePath :=  path.Join("views", page +".html")

	// вот так можно объеденить несколько кусков
	// head   :=  path.Join("views/partials", "head.html")
	// foot   :=  path.Join("views/partials", "footer.html")
	// navbar :=  path.Join("views/partials", "navbar.html")
	// parsedPages := template.Must(template.ParseFiles(head, navbar, path_tmpl, foot))

	parsedPages :=  template.Must( template.ParseFiles(pagePath) )
	return parsedPages.Execute(res, viewArgs) 
}


/*
|--------------------------------------------------------------------------
|  Render view from compiled assets (go-assets-builder)
|  Be sure you gather your views via go-assets-builder
|--------------------------------------------------------------------------
*/
func RenderCompiled(page string, viewArgs map[string]interface{}, res http.ResponseWriter) error {
	// fill in Compiled Assets with parsed templates
  if len(CompiledAssets) == 0 {
	  fillCompiledAssets()
  } 

	pagePath :=  path.Join("/", "views", page +".html")
	view     := CompiledAssets[pagePath]

	return view.Execute(res, viewArgs)
}



/*
|--------------------------------------------------------------------------
|  Find file by name and fill array with templates that was parsed by assets file (views.go)
|--------------------------------------------------------------------------
*/
func fillCompiledAssets() {
  var extension string     = ".html"
	var t *template.Template = template.New("")
	CompiledAssets = make(map[string]*template.Template)

	for name, file := range Assets.Files {
		// проверяем директория ли это или есть ли у файла расширение(окончание) .html 
		if file.IsDir() || !strings.HasSuffix(name, extension) {
			continue
		}

		// вычитываем файл 		
		h, err := ioutil.ReadAll(file) // gives type - []uint8
		if err != nil {
			fmt.Println(err)
		}

		t, err = t.New(name).Parse(string(h))
		if err != nil {
			fmt.Println(err)
		}

		CompiledAssets[name] = t
	}
}


/*
|--------------------------------------------------------------------------
|  Typeof and output
|--------------------------------------------------------------------------
*/
func Typeof(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
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