#!/bin/bash -i

CURRENT=$(echo $GKE_PROJECT_ID)
NEW_GKE_PROJECT_ID=$(gcloud config get-value project)
echo $NEW_GKE_PROJECT_ID
if [ "${CURRENT}" != "${NEW_GKE_PROJECT_ID}" ]; then
  sed -i -e '/^export GKE_PROJECT_ID/d' ~/.bashrc
  unset GKE_PROJECT_ID
  if [ "${NEW_GKE_PROJECT_ID}" != "" ]; then
    echo "export GKE_PROJECT_ID=\"$NEW_GKE_PROJECT_ID\"">>~/.bashrc
  fi
fi


CURRENT=$(echo $GKE_DEFAULT_ZONE)
NEW_GKE_DEFAULT_ZONE=$(gcloud config get-value compute/zone)
echo $NEW_GKE_DEFAULT_ZONE
if [ "${CURRENT}" != "${NEW_GKE_DEFAULT_ZONE}" ]; then
  sed -i -e '/^export GKE_DEFAULT_ZONE/d' ~/.bashrc
  unset GKE_DEFAULT_ZONE
  if [ "${NEW_GKE_DEFAULT_ZONE}" != "" ]; then
    echo "export GKE_DEFAULT_ZONE=\"$NEW_GKE_DEFAULT_ZONE\"">>~/.bashrc
  fi
fi

CURRENT=$(echo $GKE_DEFAULT_CLUSTER)
NEW_GKE_DEFAULT_CLUSTER=$(gcloud config get-value container/cluster)
echo $NEW_GKE_DEFAULT_CLUSTER
if [ "${CURRENT}" != "${NEW_GKE_DEFAULT_CLUSTER}" ]; then
  sed -i -e '/^export GKE_DEFAULT_CLUSTER/d' ~/.bashrc
  unset GKE_DEFAULT_CLUSTER
  if [ "${NEW_GKE_DEFAULT_CLUSTER}" != "" ]; then
    echo "export GKE_DEFAULT_CLUSTER=\"$NEW_GKE_DEFAULT_CLUSTER\"" >> ~/.bashrc   
  fi
fi

