build-mac:
	if [[ $$(uname -p) = "arm" ]]; then \
		GOOS=darwin GOARCH=arm64 go build -o ~/go/bin/numid ./cmd/numid/main.go; \
	else \
		GOOS=darwin GOARCH=amd64 go build -o ~/go/bin/numid ./cmd/numid/main.go; \
	fi

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./deploy/numid ./cmd/numid/main.go

do-checksum:
	cd deploy/upload && sha256sum numid > numid-checksum

build-linux-with-checksum: build-linux do-checksum
