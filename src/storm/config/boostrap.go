package config 

func Bootstrap() {
	/*
	|--------------------------------------------------------------------------
	| Routes
	|--------------------------------------------------------------------------
	*/
	router := Routes()
	// if use this, then dont pass router to: http.ListenAndServe(":" + port, router)
	// http.Handle("/", router)  

	/*
	|--------------------------------------------------------------------------
	| Up the server
	|--------------------------------------------------------------------------
	*/
	Server(router)
}



