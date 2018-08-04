@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM go get github.com/pilu/fresh
REM go get -u github.com/gorilla/mux
REM go get github.com/fatih/color


REM go get -u github.com/logrusorgru/aurora
