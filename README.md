## Storm - brother of Bullz
Fast and reliable app to make problem runs

<p align="center" style="text-align:center;">
  <img src="https://github.com/Nikeweke/Storm/blob/master/public/assets/storm.png?raw=true" width="300" />
</p>

#### Содержание 
* Quick start

---

### Quick start
###### windows
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

go get -u github.com/labstack/echo

go run main.go
```

###### linux
```bash
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

go get -u github.com/labstack/echo

go run main.go
```
