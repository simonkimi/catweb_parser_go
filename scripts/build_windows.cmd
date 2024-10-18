cd ../go
set CGO_ENABLED=1
set GOARCH=amd64
set GOOS=windows
go build -ldflags "-w -s" -buildmode=c-shared -o ../windows/catweb_parser.dll main.go