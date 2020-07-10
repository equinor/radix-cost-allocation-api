module github.com/equinor/radix-cost-allocation-api

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/equinor/radix-operator v1.5.15
	github.com/golang/gddo v0.0.0-20200611223618-a4829ef13274
	github.com/gorilla/mux v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/rakyll/statik v0.1.6
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.0.0-20191020214737-6c8691705fc5
	k8s.io/client-go v12.0.0+incompatible
)

replace (
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20190818123050-43acd0e2e93f
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
