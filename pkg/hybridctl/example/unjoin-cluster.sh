#!/bin/bash
kubectl delete ns "kube-federation-system" --context $1;
kubectl delete sa $1-hcp -n kube-federation-system --context $1;
kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:$1 --context $1 ;
kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:$1-hcp --context $1;
kubectl delete kubefedclusters $1 -n kube-federation-system --context master