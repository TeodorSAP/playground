#!/bin/sh
# Creates a kubeconfig file with a static token for the current context

CONTEXT=$(kubectl config current-context)
CLUSTER=$(kubectl config view -o json | jq -r --arg context "${CONTEXT}" '.contexts[] | select(.name == $context) | .context.cluster')
URL=$(kubectl config view -o json | jq -r --arg cluster "${CLUSTER}" '.clusters[] | select(.name == $cluster) | .cluster.server')
SERVICE_ACCOUNT="$(whoami | tr '[:upper:]' '[:lower:]')-admin"
kubectl -n default create serviceaccount "$SERVICE_ACCOUNT"

kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  annotations:
    kubernetes.io/service-account.name: "$SERVICE_ACCOUNT"
  name: "$SERVICE_ACCOUNT"
  namespace: default
type: kubernetes.io/service-account-token
EOF

kubectl create clusterrolebinding "$SERVICE_ACCOUNT" --clusterrole cluster-admin --serviceaccount "default:$SERVICE_ACCOUNT"
CA=$(kubectl get secret "$SERVICE_ACCOUNT" -o jsonpath="{.data.ca\.crt}")
TOKEN=$(kubectl get secret "$SERVICE_ACCOUNT" -o jsonpath="{.data.token}" | base64 --decode)
KUBECONFIG="$SERVICE_ACCOUNT-kubeconfig.yaml"

echo "apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: $CA
    server: $URL
  name: default
contexts:
- context:
    cluster: default
    namespace: default
    user: default
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: default
  user:
    token: $TOKEN" >> "$KUBECONFIG"
