build:
	go build -o ./build/privesc-toolkit-manager

build-all: build-mac build-win build-linux

build-mac: build-mac-arm build-mac-amd

build-win: build-win64 build-win32

build-linux: build-linux-amd build-linux-arm build-linux-386

build-win64:
	env GOOS=windows GOARCH=amd64 go build -o ./build/privesc-toolkit-manager-win64.exe

build-win32:
	env GOOS=windows GOARCH=386 go build -o ./build/privesc-toolkit-manager-win32.exe

build-mac-arm:
	env GOOS=darwin GOARCH=arm64 go build -o ./build/privesc-toolkit-manager-mac-arm64

build-mac-amd:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/privesc-toolkit-manager-mac-amd64

build-linux-amd:
	env GOOS=linux GOARCH=amd64 go build -o ./build/privesc-toolkit-manager-linux-amd64

build-linux-arm:
	env GOOS=linux GOARCH=arm64 go build -o ./build/privesc-toolkit-manager-linux-arm64

build-linux-386:
	env GOOS=linux GOARCH=386 go build -o ./build/privesc-toolkit-manager-linux-386
