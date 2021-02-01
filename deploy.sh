#!/usr/bin/env bash

# Exit with an error if any part of the script fails
set -e

# Install gcloud
sudo apt-get update -y && sudo apt-get install apt-transport-https ca-certificates gnupg -y
echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
sudo apt-get update -y && sudo apt-get install google-cloud-sdk kubectl -y

gcloud auth activate-service-account --key-file gcloud-service-account-secret.json
gcloud container clusters get-credentials gke-cluster --zone us-central1-a --project www-miketrout-dev
kubectl apply -f skills-service-deployment.yaml -f skills-service-service.yaml

# https://github.com/kubernetes/kubernetes/issues/27081#issuecomment-238078103
kubectl patch deployment skills-service-deployment \
  -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"
