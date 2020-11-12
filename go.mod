module github.com/equinor/radix-cost-allocation-api

go 1.13

require (
	github.com/coreos/go-oidc v0.0.0-20180117170138-065b426bd416
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/equinor/radix-operator v1.5.15
	github.com/go-openapi/errors v0.19.7
	github.com/go-openapi/runtime v0.19.16
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.11
	github.com/golang/gddo v0.0.0-20200611223618-a4829ef13274
	github.com/golang/mock v1.4.4
	github.com/gorilla/mux v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/rakyll/statik v0.1.6
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.0.0-20191020214737-6c8691705fc5
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
)

replace (
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20190818123050-43acd0e2e93f
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
