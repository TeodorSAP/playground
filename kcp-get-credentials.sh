#!/bin/zsh -i
#@description: - get credentials for kcp landscape
#@arg:landscape - BTP Landscape [options:!dev!|stage|prod]

# TODO: SET SHUTL_LANDSCAPE to dev/stage/prod BEFORE EXECUTING THIS SCRIPT

if ! gcloud auth print-access-token >/dev/null 2>&1 ; then
    gcloud auth login
fi
mps=~/$SHUTL_LANDSCAPE/kcp
mkdir -p $mps
touch $mps/kubeconfig.yaml
KUBECONFIG=$mps/kubeconfig.yaml gcloud container clusters get-credentials kyma-mps-${SHUTL_LANDSCAPE} --region europe-west1 --project sap-ti-dx-kyma-mps-${SHUTL_LANDSCAPE} || exit 1

cd $mps