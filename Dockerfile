FROM golang:1.21-alpine3.18 as builder
ENV GO111MODULE=on

RUN apk update && \
    apk add bash jq alpine-sdk sed gawk git ca-certificates curl mc && \
    apk add --no-cache gcc musl-dev
RUN go install honnef.co/go/tools/cmd/staticcheck@2023.1.3 && \
    go install github.com/rakyll/statik@v0.1.7
    # Install go-swagger - 57786786=v0.29.0 - get release id from https://api.github.com/repos/go-swagger/go-swagger/releases
RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/57786786 | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    curl -o /usr/local/bin/swagger -L'#' "$download_url" && \
    chmod +x /usr/local/bin/swagger

WORKDIR /go/src/github.com/equinor/radix-cost-allocation-api/

# get dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy api code
COPY . .

# Generate swagger
RUN swagger generate spec -o ./swaggerui_src/swagger.json --scan-models --exclude-deps --exclude=github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models && \
    swagger validate ./swaggerui_src/swagger.json && \
    statik -src=./swaggerui_src/ -p swaggerui

# lint and unit tests
RUN staticcheck ./... && \
    go vet ./... && \
    CGO_ENABLED=0 GOOS=linux go test ./...

# Build radix cost allocation API go project
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/radix-cost-allocation-api

RUN addgroup -S -g 1000 radix-cost
RUN adduser -S -u 1000 -G radix-cost radix-cost

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/radix-cost-allocation-api /usr/local/bin/radix-cost-allocation-api
COPY --from=builder /etc/passwd /etc/passwd
USER 1000
EXPOSE 3003
ENTRYPOINT ["/usr/local/bin/radix-cost-allocation-api"]
