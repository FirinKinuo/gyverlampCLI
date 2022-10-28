VERSION := $(file < VERSION)
PROJECT := $(shell cat go.mod | grep module | sed 's|module\sgithub.com/FirinKinuo/||g')
GO_MAIN := cmd/gyverlamp/main.go

clean:
	rm -rf _build/ release/

build:
	go mod download
	CGO_ENABLED=0 go build -v -tags release -o $(PROJECT)-$(VERSION) $(GO_MAIN)

build-all:
	mkdir -p _build
	GOOS=darwin  GOARCH=amd64   CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-darwin-amd64 $(GO_MAIN)
	GOOS=darwin  GOARCH=arm64   CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-darwin-arm64 $(GO_MAIN)
	GOOS=linux   GOARCH=amd64   CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-linux-amd64 $(GO_MAIN)
	GOOS=linux   GOARCH=arm     CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-linux-arm $(GO_MAIN)
	GOOS=linux   GOARCH=arm64   CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-linux-arm64 $(GO_MAIN)
	GOOS=windows GOARCH=amd64   CGO_ENABLED=0 go build -v -tags release -o _build/$(PROJECT)-$(VERSION)-windows-amd64.exe $(GO_MAIN)
	cd _build; sha256sum * > sha256sums.txt

install:
	make build
	mv ./$(PROJECT)-$(VERSION) /usr/bin/$(PROJECT)

release:
	make clean
	make build-all
	mkdir release
	cp _build/* release
	cd release; sha256sum --quiet --check sha256sums.txt && \
	gh release create v$(VERSION) -d -t v$(VERSION) *