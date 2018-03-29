@echo off

set ROOT_DIR=%cd%
set GOPATH=%ROOT_DIR%
set GOBIN=%ROOT_DIR%\bin

go build -o %GOBIN%\dos2unix.exe fender\cmd\dos2unix
go build -o %GOBIN%\unix2dos.exe fender\cmd\unix2dos