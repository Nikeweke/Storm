@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM go get -u github.com/labstack/echo
REM go get github.com/pilu/fresh

REM go get github.com/fatih/color
REM go get -u github.com/gin-gonic/gin
REM go get -u github.com/gorilla/mux
REM go get github.com/gorilla/websocket
REM go get github.com/robfig/cron
REM go get -u github.com/jinzhu/gorm
REM go get github.com/go-sql-driver/mysql


