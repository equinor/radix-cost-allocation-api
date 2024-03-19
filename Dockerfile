FROM docker.io/golang:1.21-alpine3.18 as builder
ENV GO111MODULE=on

RUN apk update && \
    apk add bash jq alpine-sdk sed gawk git ca-certificates curl mc && \
    apk add --no-cache gcc musl-dev

WORKDIR /go/src/github.com/equinor/radix-cost-allocation-api/

# get dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy api code
COPY . .

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
