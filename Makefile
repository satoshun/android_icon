build: build-darwin build-linux

build-darwin:
	GOARCH=amd64 GOOS=darwin go build -v -o pkg/darwin/converticon

build-linux:
	GOARCH=amd64 GOOS=linux go build -v -o pkg/linux/converticon
