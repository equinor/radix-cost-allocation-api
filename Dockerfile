# Build stage
FROM docker.io/golang:1.25-alpine3.23 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags "-s -w" -o /build/radix-cost-allocation-api

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /build/radix-cost-allocation-api .
USER 1000
EXPOSE 3003
ENTRYPOINT ["/app/radix-cost-allocation-api"]
