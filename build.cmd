@echo off

cd ./go
set CGO_ENABLED=1
set WINDOWS_PATH=%cd%/../windows/libs
go build -ldflags "-w -s" -buildmode=c-shared -o %WINDOWS_PATH%/libgo.dll main.go