BINS	= radix-cost-allocation-api
.PHONY: build
build: $(BINS)

.PHONY: test
test:
	PRETTY_PRINT=yes go test -cover ./...

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0

.PHONY: generate-radix-api-client
generate-radix-api-client: bootstrap
	swagger generate client -t ./models/radix_api/generated_client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi


.PHONY: generate-radixconfig-envs
generate-radixconfig-envs:
	# radix-id-vulnerability-scan-reader-<env>
	AZURE_CLIENT_ID=b8fd30d4-61d0-4842-b6c1-e91ceb58db8c SQL_SERVER=sql-radix-cost-allocation-dev.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.dev.yaml
	AZURE_CLIENT_ID=bb6d92a0-2f6d-421e-80e6-1b2174953d21 SQL_SERVER=sql-radix-cost-allocation-c2.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.c2.yaml
	AZURE_CLIENT_ID=ed3ef8ee-c9b2-4a89-9b0d-47b40abb2bf1 SQL_SERVER=sql-radix-cost-allocation-platform.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.platform.yaml
	AZURE_CLIENT_ID=bc4b6c73-78c2-4b22-ab08-575706a338ec SQL_SERVER=sql-radix-cost-allocation-playground.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.playground.yaml


# This make command is only needed for local testing now
# we also do make swagger inside Dockerfile
.PHONY: swagger
swagger: bootstrap
	swagger generate spec -o ./swaggerui/html/swagger.json --scan-models --exclude-deps --exclude=github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models
	swagger validate ./swaggerui/html/swagger.json

.PHONY: mocks
mocks: bootstrap
	mockgen -source ./repository/repository.go -destination ./repository/mock/repository_mock.go -package mock
	mockgen -source ./models/radix_api/client.go -destination ./api/test/mock/radix_api_client_mock.go -package mock
	mockgen -source ./api/utils/auth/auth_provider.go -destination ./api/test/mock/auth_provider_mock.go -package mock
	mockgen -source ./service/costservice.go -destination ./service/mock/costservice.go -package mock

HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_MOCKGEN       := $(shell command -v mockgen;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
endif
ifndef HAS_GOLANGCI_LINT
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
endif
ifndef HAS_MOCKGEN
	go install github.com/golang/mock/mockgen@v1.6.0
endif
