#!/bin/bash


kubectl create ns hcp --context uni-master

kubectl create -f deploy/volume/pv_fedora.yml --context uni-master
kubectl create -f deploy/volume/pvc_fedora.yml --context uni-master


kubectl create -f deploy/role_binding.yaml --context uni-master
kubectl create -f deploy/service_account_before.yaml --context uni-master
kubectl create -f deploy/service.yaml --context uni-master
# kubectl create -f deploy/secret.yaml --context uni-master

