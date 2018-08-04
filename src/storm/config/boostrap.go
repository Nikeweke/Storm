package config 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"storm/app/helpers"
	// "strings"
)

func Bootstrap() {
  
	/*
	|--------------------------------------------------------------------------
	| Connect to DB
	|--------------------------------------------------------------------------
	*/
	// database.connect()
  
  
	/*
	|--------------------------------------------------------------------------
	| Set templater
	|--------------------------------------------------------------------------
	|
	*/
	// nunjucks.configure('views', { autoescape: true, express: app })
	// app.set('view engine', 'njk')
  
  
	/*
	|--------------------------------------------------------------------------
	| Static files (СSS, JS)
	|--------------------------------------------------------------------------
	|
	*/
	// app.use(express.static('./public'))
  
  
	/*
	|--------------------------------------------------------------------------
	|  Jobs start
	|--------------------------------------------------------------------------
	|   Запуск здесь 'app / jobs / jobs.js'.
	|
	*/
	// if (global.config.jobs) { require('../app/jobs/jobs.js').jobs() }
  
  
	/*
	|--------------------------------------------------------------------------
	| Sockets start
	|--------------------------------------------------------------------------
	|   Запуск здесь 'app / config / sockets.js', и там же можно изменить порт сокетов
	*/
	// if (global.config.sockets.enabled) { require('./sockets.js').sockets(app) }
  
  
	/*
	|--------------------------------------------------------------------------
	| Routes
	|--------------------------------------------------------------------------
	*/
	router := mux.NewRouter()

	// web
	router.HandleFunc("/",      Index).Methods("GET")
	router.HandleFunc("/check", CheckRequest)

	// api
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", ApiCheckRequest)

	// out all register routes
	err := router.Walk(helpers.GetRoutes)
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", router)  

  
	/*
	|--------------------------------------------------------------------------
	| Reload browser if javascript was changed, not views (for rerendering views use: npm i supervisor -g, npm start hrm)
	|--------------------------------------------------------------------------
	*/
	// if (global.config.reload) { require('reload')(app) }
  
  
	/*
	|--------------------------------------------------------------------------
	| Up the server
	|--------------------------------------------------------------------------
	*/
	Server()
}

// Index
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello its Index")
}

func CheckRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Its Cehck requests")
}

func ApiCheckRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "That is API, Hello from Storm!")
}

