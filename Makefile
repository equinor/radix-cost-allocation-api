BINS	= radix-cost-allocation-api
.PHONY: build
build: $(BINS)

.PHONY: test
test:
	PRETTY_PRINT=yes go test -cover ./...

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0

.PHONY: radixapiclient
radixapiclient: bootstrap
	swagger generate client -t ./models/radix_api/generated_client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi


.PHONY: radixconfigs
radixconfigs:
	# radix-id-vulnerability-scan-reader-<env>
	AZURE_CLIENT_ID=b8fd30d4-61d0-4842-b6c1-e91ceb58db8c SQL_SERVER=sql-radix-cost-allocation-dev.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.dev.yaml
	AZURE_CLIENT_ID=bb6d92a0-2f6d-421e-80e6-1b2174953d21 SQL_SERVER=sql-radix-cost-allocation-c2.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.c2.yaml
	AZURE_CLIENT_ID=a0bc7c53-d168-4f80-8b7d-dbbf85d6ed73 SQL_SERVER=sql-radix-cost-allocation-c3.database.windows.net envsubst < radixconfig.tpl.yaml > radixconfig.c3.yaml
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

.PHONY: generate
generate: radixconfigs mocks swagger

.PHONY: verify-generate
verify-generate: generate
	git diff --exit-code

HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_MOCKGEN       := $(shell command -v mockgen;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0
endif
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.2
endif
ifndef HAS_MOCKGEN
	go install go.uber.org/mock/mockgen@v0.6.0
endif
