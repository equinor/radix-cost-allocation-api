# radix-cost-allocation-api
API for Radix Cost Allocation
The Radix Cost Allocation API is an HTTP server for accessing functionality on the cost allocation for [Radix](https://www.radix.equinor.com) platform. This document is for Radix developers, or anyone interested in poking around.

## Security

Authentication and authorisation are performed through an HTTP bearer token, which is (in most cases) relayed to the Kubernetes API. The Kubernetes AAD integration then performs its authentication and resource authorisation checks, and the result is relayed to the the user.

## Developing

You need Go installed. Make sure `GOPATH` and `GOROOT` are properly set up.

Also needed:

- [`go-swagger`](https://github.com/go-swagger/go-swagger) (on a Mac, you can install it with Homebrew: `brew install go-swagger`)
- [`statik`](https://github.com/rakyll/statik) (install with `go install github.com/rakyll/statik@v0.1.7`)
- [`gomock`](https://github.com/golang/mock) (install with `go install github.com/golang/mock/mockgen@v1.6.0`)

Clone the repo into your `GOPATH` and run `go mod download`.

### Generating mocks
We use gomock to generate mocks used in unit test.
You need to regenerate mocks if you make changes to any of the interface types used by the application

`make mocks`

### Dependencies - go modules

Go modules are used for dependency management. See [link](https://blog.golang.org/using-go-modules) for information how to add, upgrade and remove dependencies. E.g. To update `radix-operator` dependency:

- list versions: `go list -m -versions github.com/equinor/radix-operator`
- update: `go get github.com/equinor/radix-operator@v1.3.1`

### Running locally

Run once after cloning of the GitHub repository:

1. `go mod download`
2. `make swagger`
3. `make generate-radix-api-client`

The following env var is needed. Useful default values in brackets.

* `SQL_SERVER` - SQL server name
* `SQL_DATABASE` - SQL database name
* `SQL_USER` - SQL server user name
* `SQL_PASSWORD` - SQL server user password
* `RADIX_ENVIRONMENT` - Radix environment (ex. `qa`)
* `RADIX_CLUSTERNAME` - Radix cluster name (ex. `weekly-33`)
* `RADIX_DNS_ZONE` - Radix DNS zone (ex. `dev.radix.equinor.com`)
* `WHITELIST` - List of applications, not included for cost report `ex. {"whiteList": ["canarycicd-test","canarycicd-test1","canarycicd-test2","canarycicd-test3","canarycicd-test4","radix-api","radix-canary-golang","radix-cost-allocation-api","radix-github-webhook","radix-platform","radix-web-console"]}`)
* `AD_REPORT_READERS` - Azure AD group for user, allowed to get an overall cost report (ex. `{"groups": ["d59ab0b8-2b2c-11eb-adc1-0242ac120002"]}`)
* `TOKEN_ISSUER` - Azure tennant ID (ex. `https://sts.windows.net/f08f9cda-2b2c-11eb-adc1-0242ac120002/`)
* `USE_LOCAL_RADIX_API`
  * `false`, `no` or not set` - connecting to in-cluster `radix-api`
  * `true` or `yes` - connecting to `radix-api`, running on `http://localhost:3002`
* `USE_PROFILER`
  * `false`, `no` or `not set` - do not use profiler
  * `true` or `yes` - use [pprof](https://golang.org/pkg/net/http/pprof/) profiler, running on `http://localhost:7070/debug/pprof`. Use web-UI to profile, when started service:
    ```
        go tool pprof -http=:6070 http://localhost:7070/debug/pprof/heap
    ```

#### Common errors running locally

- **Problem**: `panic: statik/fs: no zip data registered`

  **Solution**: `make swagger`

## Deployment

Radix Cost Allocation API follows the [standard procedure](https://github.com/equinor/radix-private/blob/master/docs/how-we-work/development-practices.md#standard-radix-applications) defined in _how we work_. 

Radix Cost Allocation API is installed as a Radix application in [script](https://github.com/equinor/radix-platform/blob/master/scripts/install_base_components.sh) when setting up a cluster. It will setup API environment with [aliases](https://github.com/equinor/radix-platform/blob/master/scripts/create_alias.sh), and a Webhook so that changes to this repository will be reflected in Radix platform. 
```
If radix-operator is updated to a new tag, `go.mod` should be updated as follows: 
   
    github.com/equinor/radix-operator <NEW_OPERATOR_TAG>
```
To install with `install_base_components.sh`, mentioned above - add RadixRegistration values and application secrets to Azure KeyVault:
1. If using a terminal - login to an Azure and switch to an Azure subscription:
    ```
    az login
    az account set -s "<SUBSCRIPTION-NAME>"
    ```
2. Create a file `radix-cost-allocation-api-radixregistration-values.yaml` with a content:
    ```
    repository: https://github.com/<GIT-USER>/<GIT-REPOSITORY>
    cloneURL: git@github.com:<GIT-USER>@<GIT-REPOSITORY>.git
    adGroups:
      - <AD-GROUP-IF-USED>
    deployKey: |
        -----BEGIN RSA PRIVATE KEY-----
        <YOUR-PRIVATE-KEY>
        -----END RSA PRIVATE KEY-----
    deployKeyPublic: ssh-rsa <YOUR-PUBLIC-KEY>
    sharedSecret: <SOME-RANDOM-TEXT>
    ```
3. Set the secret with following command or with Azure portal (`Key vault`/`Secrets`):
    ```
    az keyvault secret set  \
    -f radix-cost-allocation-api-radixregistration-values.yaml \
    -n radix-cost-allocation-api-radixregistration-values \
    --vault-name "<KEYVAULT>"
    ```
4. To check - run following command and read the created file `...-check.yaml` or with Azure portal (`Key vault`/`Secrets`):
    ```
    az keyvault secret download \
    -f radix-cost-allocation-api-radixregistration-values-check.yaml \
    -n radix-cost-allocation-api-radixregistration-values \
    --vault-name "<KEYVAULT>" 
    ```
5. Create a file `radix-cost-allocation-api-secrets.json` with content:
    ```
    {
      "db": {
        "server":"<SERVER>",
        "database":"<DATABASE>",
        "user":"<USER>",
        "password":"<PASSWORD>"
      },
      "subscriptionCost": {
        "value": "<COST-VALUE>",
        "currency": "<COST-CURRENCY>"
        "whiteList": "{"whiteList": ["APP1", "APP2"]}"
      },
      "auth": {
        "tokenIssuer": "https://sts.windows.net/<TENANT-ID>/",
        "reportReaders": "{"groups": ["<AD-GROUP>"]}"
      }
    }
    ```
6. Set the secret with following command or with Azure portal (`Key vault`/`Secrets`):
    ```
    az keyvault secret set \
    -f radix-cost-allocation-api-secrets.json \
    -n radix-cost-allocation-api-secrets \
    --vault-name "<KEYVAULT>"
    ```
7. To check - run following command and read the created file `...-check.json` or with Azure portal (`Key vault`/`Secrets`):
    ```
    az keyvault secret download \
    -f radix-cost-allocation-api-secrets-check.json \
    -n radix-cost-allocation-api-secrets \
    --vault-name "<KEYVAULT>" 
    ```
 