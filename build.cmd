@echo off

cd ./go
set CGO_ENABLED=1
go build -ldflags "-w -s" -buildmode=c-shared -o ../windows/catweb_parser.dll main.go


set CD=%cd%
set ANDROID_PATH=%cd%\..\android\libs
set ANDROID_NDK_HOME=F:\sdk\Android\ndk\25.2.9519653

md %ANDROID_PATH%\armeabi-v7a
md %ANDROID_PATH%\arm64-v8a
md %ANDROID_PATH%\x86_64
md %ANDROID_PATH%\x86

del /f /s /q %ANDROID_PATH%\armeabi-v7a\libcatweb_parser.so
del /f /s /q %ANDROID_PATH%\arm64-v8a\libcatweb_parser.so
del /f /s /q %ANDROID_PATH%\x86_64\libcatweb_parser.so
del /f /s /q %ANDROID_PATH%\x86\libcatweb_parser.so

set GOARCH=arm64
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\arm64-v8a\libcatweb_parser.so main.go

set GOARCH=arm
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi19-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\armeabi-v7a\libcatweb_parser.so main.go


set GOARCH=amd64
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\x86_64-linux-android24-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\x86_64\libcatweb_parser.so main.go

set GOARCH=386
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\i686-linux-android24-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\x86\libcatweb_parser.so main.go
