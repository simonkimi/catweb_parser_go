ifeq ($(OS),Windows_NT)
	SEP := \\
	ENV := set
else
	SEP := /
	ENV := export
endif

ANDROID_OUT := android$(SEP)libs

DEPS := main.go
PRODUCT_NAME := libcatweb_parser
GO_BUILD := $(ENV) CGO_ENABLED=1 && go build -ldflags "-w -s"
C_SHARED := -buildmode=c-shared
C_ARCHIVE := -buildmode=c-archive

.PHONY: android ios macos
android: android-arm64-v8a android-armeabi-v7a android-x86_64

android-arm64-v8a: android-build-arm64-v8a
android-armeabi-v7a: android-build-armeabi-v7a
android-x86_64: android-build-x86_64

android-build-%:
	$(eval GO_BUILD_ENV := )
	$(eval ARCH := $*)
	$(eval OUT_DIR := $(ANDROID_OUT)$(SEP)$(ARCH))
	$(call build_android, $(ARCH))


define build_android
	@echo "Building for Android $(ARCH)"
	$(eval GO_BUILD_ENV += $(ENV) GOOS=android &&)
	$(if $(filter $(1), arm64-v8a),  \
		$(eval GO_BUILD_ENV += $(ENV) GOARCH=arm64 &&) \
		$(eval GO_BUILD_ENV += $(ENV) CC=aarch64-linux-android21-clang &&) \
	)
	$(if $(filter $(1), armeabi-v7a),  \
		$(eval GO_BUILD_ENV += $(ENV) GOARCH=arm &&) \
		$(eval GO_BUILD_ENV += $(ENV) CC=armv7a-linux-androideabi19-clang &&) \
	)
	$(if $(filter $(1), x86_64),  \
		$(eval GO_BUILD_ENV += $(ENV) GOARCH=amd64 &&) \
		$(eval GO_BUILD_ENV += $(ENV) CC=x86_64-linux-android24-clang &&) \
	)
	$(GO_BUILD_ENV) $(GO_BUILD) $(C_SHARED) -o $(OUT_DIR)$(SEP)lib$(PRODUCT_NAME).so $(DEPS)
endef
