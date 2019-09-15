#!/usr/bin/env bash

# Exit with an error if any part of the script fails
set -e

# Install gcloud
curl https://sdk.cloud.google.com | bash > /dev/null

# Promote gcloud to PATH top priority (prevent using old version from Travis)
source $HOME/google-cloud-sdk/path.bash.inc

# Make sure kubectl is updated to latest version
gcloud components update kubectl

gcloud auth activate-service-account --key-file gcloud-service-account-secret.json
gcloud container clusters get-credentials gke-cluster --zone us-central1-a --project www-miketrout-dev
kubectl apply -f skills-service-deployment.yaml \
  -f skills-service-service-service.yaml

# https://github.com/kubernetes/kubernetes/issues/27081#issuecomment-238078103
kubectl patch deployment skills-service-deployment \
  -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"
