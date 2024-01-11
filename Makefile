lib_install:
	go install -v golang.org/x/mobile/cmd/gomobile@latest
	go install -v golang.org/x/mobile/cmd/gobind@latest


build-api-android: lib_install
	gomobile init
	go mod tidy -v
	go get -d golang.org/x/mobile
	go get -d golang.org/x/mobile/bind
	go get -u github.com/khayyamov/mieru
	gomobile bind -v -androidapi 19 -ldflags='-s -w' ./

build-api-ios: lib_install
	gomobile init
	go mod tidy -v
	go get -d golang.org/x/mobile
	go get -d golang.org/x/mobile/bind
	go get -u github.com/khayyamov/mieru
	gomobile bind -a -ldflags '-s -w -extldflags -lresolv' -target=ios -iosversion=12.0 -o Mieru.xcframework Mieru