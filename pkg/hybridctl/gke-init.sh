#!/bin/bash

if [[ -z "${GKE_PROJECT_ID}" ]]; then 
  GKE_PROJECT_ID=$(gcloud config get-value project)
  echo "export GKE_PROJECT_ID=\"$GKE_PROJECT_ID\"">>~/.bashrc
fi

if [[ -z "${GKE_DEFAULT_ZONE}" ]]; then 
  GKE_DEFAULT_ZONE=$(gcloud config get-value compute/zone)
  echo "export GKE_DEFAULT_ZONE=\"$GKE_DEFAULT_ZONE\"" >> ~/.bashrc
fi 

GKE_DEFAULT_CLUSTER=$(gcloud config get-value container/cluster)
if [[ $GKE_DEFAULT_CLUSTER -eq "(unset)" ]]
then
    if [[ -z "${GKE_DEFAULT_ZONE}" ]]; then 
       echo "export GKE_DEFAULT_CLUSTER=\"$GKE_DEFAULT_CLUSTER\"" >> ~/.bashrc
    fi
fi

source ~/.bashrc