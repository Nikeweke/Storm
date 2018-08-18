@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM Linux Build settings
REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM go get -u github.com/labstack/echo

REM go-assets-builder views -o views/views.go
REM fresh

go run main.go 
REM go build main.go
REM pause