@echo off

cd ./go
set CGO_ENABLED=1
go build -ldflags "-w -s" -buildmode=c-shared -o ../windows/catweb_parser.dll main.go