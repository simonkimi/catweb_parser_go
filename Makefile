PRODUCT_NAME = libcatweb_parser

build-android-arm64-v8a:
	$(MAKE) -C go build-android-arm64-v8a
	mkdir -p android/libs/arm64-v8a
	cp go/build/arm64-v8a/$(PRODUCT_NAME).so android/libs/arm64-v8a/$(PRODUCT_NAME).so
	$(MAKE) -C go clean

build-windows-x86_64:
	$(MAKE) -C go build-windows-x86_64
	cp go/build/windows/$(PRODUCT_NAME).dll windows/$(PRODUCT_NAME).dll
	$(MAKE) -C go clean