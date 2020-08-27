#!/bin/bash
# this script 
# - creates a certificate
# - set certificate to valid in kubernetes
# - creates a secret in "title" namespaces containing cert
# - register an admission controller using public key from cert and endpoint "/admissioncontrollers/applications" 
# last step should be moved somewhere else, so that it can be used to register admission controller for 
# different crds as RadixRegistration, RadixApplication and RadixDeployment
title="radix-cost-allocation-api-prod"
kubectl create namespace ${title} || true

csrName=${title}.${title}
tmpdir=$(mktemp -d)
echo "creating certs in tmpdir ${tmpdir} "

cat <<EOF >> ${tmpdir}/csr.conf
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = server.${title}.svc
EOF

openssl genrsa -out ${tmpdir}/server-key.pem 2048
openssl req -new -key ${tmpdir}/server-key.pem -subj "/CN=${title}.${title}.svc" -out ${tmpdir}/server.csr -config ${tmpdir}/csr.conf

# clean-up any previously created CSR for our service. Ignore errors if not present.
kubectl delete csr ${csrName} 2>/dev/null || true

# create  server cert/key CSR and  send to k8s API
cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${csrName}
spec:
  groups:
  - system:authenticated
  request: $(cat ${tmpdir}/server.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - server auth
EOF

# verify CSR has been created
while true; do
    kubectl get csr ${csrName}
    if [ "$?" -eq 0 ]; then
        break
    fi
done

# approve and fetch the signed certificate
kubectl certificate approve ${csrName}
# verify certificate has been signed
for x in $(seq 10); do
    serverCert=$(kubectl get csr ${csrName} -o jsonpath='{.status.certificate}')
    if [[ ${serverCert} != '' ]]; then
        break
    fi
    sleep 1
done
if [[ ${serverCert} == '' ]]; then
    echo "ERROR: After approving csr ${csrName}, the signed certificate did not appear on the resource. Giving up after 10 attempts." >&2
    exit 1
fi
echo ${serverCert} | openssl base64 -d -A -out ${tmpdir}/server-cert.pem


# create the secret with CA cert and server cert/key
kubectl create secret generic ${title}-certs \
        --from-file=key.pem=${tmpdir}/server-key.pem \
        --from-file=cert.pem=${tmpdir}/server-cert.pem \
        --dry-run -o yaml |
    kubectl -n ${title} apply -f -

# update admission controller with new certificate public key
cat <<EOF | kubectl apply -f -
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: ra-admission-controller
webhooks:
- name: admission-policy.applications.io
  failurePolicy: Fail
  rules:
  - apiGroups:
    - radix.equinor.com
    apiVersions:
    - v1
    operations:
    - CREATE
    - UPDATE
    resources:
    - radixapplications
  clientConfig:
    caBundle: $(cat ${tmpdir}/server-cert.pem | base64)
    service:
       name: server
       namespace: radix-cost-allocation-api-prod
       path: "/admissioncontrollers/applications"
EOF