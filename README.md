## Storm - brother of Bullz
Fast and reliable app to make problem runs

<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/storm.png?raw=true" width="300" />
</p>

#### Содержание 
* Quick start
* Settings.json
* What is optional?
* Commands
* Schema of launch
* Packages you may need
* Pack all views into one binary file (go-assets)
* [My Sockets examples](https://github.com/Nikeweke/GOLANG-WS)

---

### Quick start
```bash
# install packages
.\packages.bat

# run it
.\runner.bat
```

### Settings.json
In `settings.json` you can describe your database connections, sockets, jobs and turn off/on it. **Always keep this file with binary file, in other case it will break down** 


### What is optional?
* runner.bat
* runner_linux.bat
* packages.bat
* packages_linux.bat
* test.sqlite - its database just for testing `gorm`

### Commands
* **main.exe dev** - start `dev` mode - 8000 port
* **main.exe prod** - start `prod` mode - 80 port
* **main.exe test** - start `test` mode - 3000 port
* **main.exe routes** - output in console registered routes
* **main.exe routes -w** - output routes and write into file

### Schema of launch
<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/schema_storm.png?raw=true" width="600" />
</p>


### Packages you may need
in `packages.bat` you can find ready commands

* **go get -u github.com/gin-gonic/gin** - `echo` like http framework
* **go get github.com/pilu/fresh** - watcher за кодом
* **go get github.com/fatih/color** - цвета для консоли 
* **go get github.com/gorilla/websocket** - sockets
* **go get github.com/robfig/cron** - jobs
* **go get -u github.com/jinzhu/gorm** - driver for different databases
* **go get github.com/go-sql-driver/mysql** - mysql driver
* **go get github.com/jessevdk/go-assets-builder** - transform assets into go files and compile it

### Pack all views into one binary file (go-assets)
Install `go-assets-builder` for bundling your `views` or assets
```bash
go get github.com/jessevdk/go-assets-builder
```

Execute command for a bundle 
```bash
# target folder is "views" where placed all views,
# and output file is views.go
go-assets-builder views -o app/helpers/views.go
```

Go to `./app/helpers/views.go` and change package name
```go 
package main 
// to 
package helpers
```

You can take look at file `./app/helpers/views.go`. There you can find filenames and bytes of views. And you can use it like:
```go
// Example for possible controller
import (
	"net/http"
	"github.com/labstack/echo"
	"../helpers"  // here you can find "RenderCompiled" method (./app/helpers/base.go)
)

func (this someController) someMethod(c echo.Context) error {
  ...
  // this method takes page from generated file of assets "views.go"
  return helpers.RenderCompiled("home", viewArgs, c.Response())
}
```

If you dont get how it works. You can go to `./app/helpers/base.go`. Find **RenderCompiled** and **LoadTemplate**. 
In LoadTemplate you can output files that was bundled by commenting `break` in for and set `Println(name)` 
