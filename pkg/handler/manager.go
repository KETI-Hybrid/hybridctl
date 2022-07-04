package handler

import "Hybrid_Cluster/util/clusterManager"

type HttpManager struct {
	HTTPServer_IP   string
	HTTPServer_PORT string
	ClusterManager  *clusterManager.ClusterManager
}
