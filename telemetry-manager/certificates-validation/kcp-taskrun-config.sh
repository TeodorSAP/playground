#!/bin/bash

ENV=$1

export KCP_OIDC_ISSUER_URL="https://kymatest.accounts400.ondemand.com"
export KCP_OIDC_CLIENT_ID="9bd05ed7-a930-44e6-8c79-e6defeb7dec9"
export KCP_KEB_API_URL="https://kyma-env-broker.cp.$ENV.kyma.cloud.sap"
export KCP_MOTHERSHIP_API_URL="https://mothership-reconciler.cp.$ENV.kyma.cloud.sap/v1"
export KCP_KUBECONFIG_API_URL="https://kubeconfig-service.cp.$ENV.kyma.cloud.sap"
export KCPCONFIG="$HOME/.kcp/config-$ENV.yaml"
export KCP_GARDENER_NAMESPACE="garden-kyma-$ENV"
export KCP_GARDENER_KUBECONFIG="$HOME/.kcp/gardener-kubeconfig/kubeconfig-garden-kyma-$ENV.yaml" # Edit it and delete the line (- '--oidc-use-pkce').
