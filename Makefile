.PHONY: generate
generate: fetch-spec bundle-spec
	@echo "Generating Go code from OpenAPI spec..."
	go generate ./gen/...
	go fmt ./gen/...

.PHONY: fetch-spec
fetch-spec:
	@echo "Fetching OpenAPI specification..."
	bash scripts/fetch-spec.sh main

.PHONY: bundle-spec
bundle-spec:
	@echo "Bundling OpenAPI specification..."
	npx @redocly/cli bundle api/openapi.yaml -o api/openapi.bundled.yaml --ext yaml --remove-unused-components --force
	@echo "Fixing duplicate type names..."
	@sed -i.bak \
		-e 's|#/components/responses/NotificationV2Response|#/components/responses/NotificationV2ApiResponse|g' \
		-e '/^  responses:/,/^  schemas:/ s/^    NotificationV2Response:/    NotificationV2ApiResponse:/' \
		api/openapi.bundled.yaml && rm -f api/openapi.bundled.yaml.bak

.PHONY: install-tools
install-tools:
	@echo "Installing oapi-codegen..."
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: clean
clean:
	rm -rf gen/*.gen.go api/openapi.bundled.yaml
