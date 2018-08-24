## Storm - brother of Bullz
Fast and reliable app to make problem runs

<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/storm.png?raw=true" width="300" />
</p>

#### Содержание 
* Quick start
* Settings.json
* Commands
* Schema of launch
* Packages you may need
* Pack all views into one binary file (go-assets)

---

### Quick start
Just copy and execute file. At second launch comment packages `go get ...` downloading.
###### windows (runner.bat)
```bat
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM Linux Build settings
REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

go get -u github.com/labstack/echo

REM go-assets-builder views -o views/views.go
REM fresh

go run main.go 
REM go build main.go
REM pause
```

###### linux (runner.sh)
```bash
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

go get -u github.com/labstack/echo

# go-assets-builder views -o views/views.go
# fresh

go run main.go
```

### Settings.json
In `settings.json` you can describe your database connections, sockets, jobs and turn off/on it. **Always keep this file with binary file, in other case it will break down** 

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
go-assets-builder views -o views/views.go
```

Go to `views/views.go` and change package name
```go 
package main 
// to 
package views
```

You can take look at file `views/views.go`. There you can find filenames and bytes of views. And you can use it like:
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

If dont get how it works. You can to `./app/helpers/base.go`. Find **RenderCompiled** and **LoadTemplate**. 
In LoadTemplate you can output files that was bundled by commenting `break` in for and set `Println(name)` 