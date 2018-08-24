@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

go get -u github.com/labstack/echo
go get github.com/fatih/color
go get github.com/jessevdk/go-assets

REM go get github.com/jessevdk/go-assets-builder
REM go get github.com/dgrijalva/jwt-go
REM go get github.com/pilu/fresh
REM go get github.com/robfig/cron
REM go get github.com/jessevdk/go-assets-builder

REM go get github.com/jinzhu/gorm
REM go get github.com/jinzhu/gorm/dialects/mysql
REM go get github.com/jinzhu/gorm/dialects/sqlite
REM go get github.com/jinzhu/gorm/dialects/mssql
REM go get github.com/jinzhu/gorm/dialects/postgres

REM go get gopkg.in/mgo.v2

REM pause