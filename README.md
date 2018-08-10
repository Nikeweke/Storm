## Storm - brother of Bullz
Fast and reliable app to make problem runs. Using `echo` framework for HTTP handling

<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/storm.png?raw=true" width="300" />
</p>

#### Содержание 
* Quick start

---

### Quick start
###### windows (runner.bat)
```batch
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
