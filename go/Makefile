DEPS = main.go
PRODUCT_NAME = libcatweb_parser
GO_BUILD = CGO_ENABLED=1 go build -ldflags "-w -s"
GO_SHARED = -buildmode=c-shared
GO_ARCHIVE = -buildmode=c-archive

.PHONY: build-android-arm64-v8a
build-android-arm64-v8a: GOARCH=arm64
build-android-arm64-v8a: ARCH=arm64-v8a
build-android-arm64-v8a: CC=aarch64-linux-android21-clang
build-android-arm64-v8a: build-android-arch

.PHONY: build-android-armeabi-v7a
build-android-armeabi-v7a: GOARCH=arm
build-android-armeabi-v7a: ARCH=armeabi-v7a
build-android-armeabi-v7a: CC=armv7a-linux-androideabi19-clang
build-android-armeabi-v7a: build-android-arch

.PHONY: build-android-x86_64
build-android-x86_64: GOARCH=amd64
build-android-x86_64: ARCH=x86_64
build-android-x86_64: CC=x86_64-linux-android24-clang
build-android-x86_64: build-android-arch


build-android-arch:
	@echo "Building for Android $(ARCH)"
	mkdir -p build$(SEP)$(ARCH)
	env GOOS=android GOARCH=$(GOARCH) $(GO_BUILD) $(GO_SHARED) -o build/$(ARCH)/$(PRODUCT_NAME).so $(DEPS)

build-windows-x86_64:
	@echo "Building for Windows"
	mkdir -p build/windows
	env GOOS=windows GOARCH=amd64 $(GO_BUILD) $(GO_SHARED) -o build/windows/$(PRODUCT_NAME).dll $(DEPS)

build-ios:
	@echo "Building for iOS"
	- mkdir -p build/ios
	env GOOS=ios GOARCH=arm64 CC=$(GOROOT)/misc/ios/clangwrap.sh $(GO_BUILD) $(GO_ARCHIVE) -o build/ios/$(PRODUCT_NAME).a $(DEPS)

build-macos:
	@echo "Building for macOS"
	- mkdir -p  build/macos
	env GOOS=darwin GOARCH=amd64 $(GO_BUILD) $(GO_ARCHIVE) -o build/macos/$(PRODUCT_NAME).a $(DEPS)

clean:
	rm -rf build