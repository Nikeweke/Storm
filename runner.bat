@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM Linux Build settings
REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM go get github.com/pilu/fresh

REM go get -u github.com/labstack/echo
REM go get github.com/dgrijalva/jwt-go
REM go get github.com/fatih/color

REM go get github.com/jinzhu/gorm
REM go get github.com/jinzhu/gorm/dialects/mysql
REM go get github.com/jinzhu/gorm/dialects/sqlite
REM go get gopkg.in/mgo.v2

REM fresh
go run main.go

REM go build main.go prod
REM pause