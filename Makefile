.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: build
build:
	go build ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run-basic-auth
run-basic-auth:
	cd examples/basic_auth && go run main.go

.PHONY: run-cookie-auth
run-cookie-auth:
	cd examples/cookie_auth && go run main.go

.PHONY: clean
clean:
	rm -f examples/*/cookies.json
	go clean ./...
