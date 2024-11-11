ifeq ($(OS),Windows_NT)
	HAS_ENV = $(shell where env >nul 2>&1 && echo yes || echo no)
	ifeq ($(HAS_ENV),no)
		$(error "Error: 'env' command not found. Please ensure it is installed and available in your PATH.")
	endif
	MKDIR = mkdir
else
	MKDIR = mkdir -p
endif

DEPS = ./go/main.go
PRODUCT_NAME = libcatweb_parser
GO_BUILD = CGO_ENABLED=1 go build -ldflags "-w -s"
GO_SHARED = -buildmode=c-shared
GO_ARCHIVE = -buildmode=c-archive

.PHONY: build-android
build-android: build-android-arm64-v8a build-android-armeabi-v7a build-android-x86_64

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
	- $(MKDIR) android/libs/$(ARCH)
	env GOOS=android GOARCH=$(GOARCH) $(GO_BUILD) $(GO_SHARED) -o android/libs/$(ARCH)/$(PRODUCT_NAME).so $(DEPS)


build-ios:
	- $(MKDIR) ios/Classes
	env GOOS=ios GOARCH=arm64 CC=$(GOROOT)/misc/ios/clangwrap.sh $(GO_BUILD) $(GO_ARCHIVE) -o ios/Classes/$(PRODUCT_NAME).a $(DEPS)

build-macos:
	- $(MKDIR) macos/Classes
	env GOOS=darwin GOARCH=amd64 $(GO_BUILD) $(GO_ARCHIVE) -o macos/Classes/$(PRODUCT_NAME).a $(DEPS)
