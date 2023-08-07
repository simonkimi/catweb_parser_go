#!/usr/bin/env bash

export CGO_CFLAGS="-fembed-bitcode"
export CGO_ENABLED=1

cd ./go || exit

export GOARCH=arm64
export GOOS=darwin
go build -ldflags "-w -s" -buildmode=c-archive -o ../src/macos/catweb_parser.dylib main.go


export GOARCH=arm64
export GOOS=ios
export CC=$GOROOT/misc/ios/clangwrap.sh
go build -ldflags "-w -s" -buildmode=c-archive -o ../src/ios/catweb_parser.dylib main.go