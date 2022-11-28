build-mac:
	if [[ $$(uname -p) = "arm" ]]; then GOARCH=arm64; else GOARCH=amd64; fi; \
	GOOS=darwin go build -o ~/go/bin/numid ./cmd/numid/main.go; \

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./deploy/numid ./cmd/numid/main.go

do-checksum:
	cd deploy/upload && sha256sum numid > numid-checksum

build-linux-with-checksum: build-linux do-checksum

deploy-testnet:
	deploy/deploy-testnet.sh

e2e-test:
	deploy/e2e-test.sh

demo:
	@deploy/demo.sh

mock-expected-keepers:
	mockgen -source=x/numi/types/expected_keepers.go -destination=testutil/mock_types/expected_keepers.go 
