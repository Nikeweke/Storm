## Storm - brother of Bullz
Fast and reliable app to make problem runs. Using `echo` framework for HTTP handling

<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/storm.png?raw=true" width="300" />
</p>

#### Содержание 
* Quick start
* Commands
* Schema of launch
* Useful packages

---

### Quick start
Just copy and execute file
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

go run main.go
```

###### linux (runner.sh)
```bash
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

go get -u github.com/labstack/echo

go run main.go
```

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


### Useful packages
* **go get github.com/pilu/fresh** - watcher за кодом
* **go get github.com/fatih/color** - цвета для консоли 
* **go get -u github.com/gin-gonic/gin** - `echo` like http framework
* **go get github.com/gorilla/websocket** - sockets
* **go get github.com/robfig/cron** - jobs
* **go get -u github.com/jinzhu/gorm** - driver for different databases
* **go get github.com/go-sql-driver/mysql** - mysql driver
* **go get github.com/jessevdk/go-assets** - transform assets into go files and compile it