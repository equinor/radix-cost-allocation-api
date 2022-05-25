BINS	= radix-cost-allocation-api
.PHONY: build
build: $(BINS)

.PHONY: test
test:
	go test -cover `go list ./...`


.PHONY: generate-radix-api-client
generate-radix-api-client:
	swagger generate client -t ./models/radix_api/generated_client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi

# This make command is only needed for local testing now
# we also do make swagger inside Dockerfile
.PHONY: swagger
swagger:
	rm -f ./swaggerui_src/swagger.json ./swaggerui/statik.go
	swagger generate spec -o ./swagger.json --scan-models --exclude=github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models
	mv swagger.json ./swaggerui_src/swagger.json
	statik -src=./swaggerui_src/ -p swaggerui

.PHONY: mocks
mocks:
	mockgen -source ./repository/repository.go -destination ./repository/mock/repository_mock.go -package mock
	mockgen -source ./models/radix_api/client.go -destination ./api/test/mock/radix_api_client_mock.go -package mock
	mockgen -source ./api/utils/auth/auth_provider.go -destination ./api/test/mock/auth_provider_mock.go -package mock
	mockgen -source ./service/costservice.go -destination ./service/mock/costservice.go -package mock