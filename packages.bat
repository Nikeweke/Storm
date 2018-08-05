@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

go get github.com/pilu/fresh
go get -u github.com/gorilla/mux
go get github.com/fatih/color

REM go get github.com/gorilla/websocket
REM go get github.com/robfig/cron
REM go get -u github.com/jinzhu/gorm
REM go get github.com/go-sql-driver/mysql


