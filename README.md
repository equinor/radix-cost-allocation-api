# radix-cost-allocation-api
API for Radix Cost Allocation
The Radix Cost Allocation API is an HTTP server for accessing functionality on the cost allocation for [Radix](https://www.radix.equinor.com) platform. This document is for Radix developers, or anyone interested in poking around.

## Security

Authentication and authorisation are performed through an HTTP bearer token, which is (in most cases) relayed to the Kubernetes API. The Kubernetes AAD integration then performs its authentication and resource authorisation checks, and the result is relayed to the the user.

## Developing

You need Go installed. Make sure `GOPATH` and `GOROOT` are properly set up.

Also needed:

- [`go-swagger`](https://github.com/go-swagger/go-swagger) (on a Mac, you can install it with Homebrew: `brew install go-swagger`)
- [`statik`](https://github.com/rakyll/statik) (install with `go get github.com/rakyll/statik`)

Clone the repo into your `GOPATH` and run `go mod download`.

### Dependencies - go modules

Go modules are used for dependency management. See [link](https://blog.golang.org/using-go-modules) for information how to add, upgrade and remove dependencies. E.g. To update `radix-operator` dependency:

- list versions: `go list -m -versions github.com/equinor/radix-operator`
- update: `go get github.com/equinor/radix-operator@v1.3.1`

### Running locally

The following env vars are needed. Useful default values in brackets.

- `RADIX_CONTAINER_REGISTRY` - (`radixdev.azurecr.io`)

#### Common errors running locally

- **Problem**: `panic: statik/fs: no zip data registered`

  **Solution**: `make swagger`

### Manual redeployment on existing cluster

#### Prerequisites

1. Install draft (https://draft.sh/)
2. `draft init` from project directory (inside `radix-cost-allocation-api`)
3. `draft config set registry radixdev.azurecr.io`
4. `az acr login --name radixdev`

#### Process

1. Update version in header of swagger version in `main.go` so that you can see that the version in the environment corresponds with what you wanted
2. Execute `draft up` to install to dev environment of radix-api
3. Wait for pods to start
4. Go to `https://server-radix-cost-allocation-api-dev.<cluster name>.dev.radix.equinor.com/swaggerui/` to see if the version in the swagger corresponds with the version you set in the header.

## Deployment

Radix Cost Allocation API follows the [standard procedure](https://github.com/equinor/radix-private/blob/master/docs/how-we-work/development-practices.md#standard-radix-applications) defined in _how we work_.

Radix Cost Allocation API is installed as a Radix application in [script](https://github.com/equinor/radix-platform/blob/master/scripts/install_base_components.sh) when setting up a cluster. It will setup API environment with [aliases](https://github.com/equinor/radix-platform/blob/master/scripts/create_alias.sh), and a Webhook so that changes to this repository will be reflected in Radix platform. 
```
If radix-operator is updated to a new tag, `go.mod` should be updated as follows: 
   
    github.com/equinor/radix-operator <NEW_OPERATOR_TAG>
```
