#!/bin/bash




# kubectl delete -f deploy/operator/configmap.yaml --context uni-master
kubectl delete -f deploy/operator/operator-uni-cluster.yaml --context uni-master


kubectl delete -f deploy/ --context uni-master

kubectl delete -f deploy/volume/pvc_fedora.yml --context uni-master
kubectl delete -f deploy/volume/pv_fedora.yml --context uni-master


