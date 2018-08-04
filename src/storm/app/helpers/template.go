package helpers

import (
    "html/template"
    "net/http"
	"path"
)

type Template struct {
}


/*
|--------------------------------------------------------------------------
|  Делаем темлейт страницы
|--------------------------------------------------------------------------
*/
func (this Template) Render(page string, data map[string]interface{}, res http.ResponseWriter) {

	path_tmpl :=  path.Join("views", page +".html")   // указываем путь к вьюхам(темплейтам)

	// head      :=  path.Join("views/partials", "head.html")
	// foot      :=  path.Join("views/partials", "footer.html")
	// navbar      :=  path.Join("views/partials", "navbar.html")

	// t := template.Must(template.ParseFiles(head, navbar, path_tmpl, foot))
    t :=  template.Must(template.ParseFiles(path_tmpl))

	// if err != nil{ // если ошибка не пусто, показать сообщение в консоли
	//   fmt.Println(err)
	// }

	t.ExecuteTemplate(res, page, data)
	// template.Execute(res, data)  // показать вьюху
}