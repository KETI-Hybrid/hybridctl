package util

import "strconv"

var (
	Option_filetype      string
	Option_context       string
	Option_file          string
	Option_namespace     string
	Option_allnamespace  bool
	Option_ip            string
	Option_clusterName   string
	Option_allcluster    bool
	Option_containerName string
	Option_stdin         bool
	Option_tty           bool
)

var (
	OpenMCPAPIServer = "10.0.3.20:31635" // Initialized by reading config file from root.go

	ResourceMap = map[string]string{

		// WORKLOAD APIS
		"cj":       "cronjobs",
		"cronjob":  "cronjobs",
		"cronjobs": "cronjobs",

		"ds":        "daemonsets",
		"damonset":  "daemonsets",
		"damonsets": "daemonsets",

		"deploy":      "deployments",
		"deployment":  "deployments",
		"deployments": "deployments",

		"job":  "jobs",
		"jobs": "jobs",

		"po":   "pods",
		"pod":  "pods",
		"pods": "pods",

		"rs":          "replicasets",
		"replicaset":  "replicasets",
		"replicasets": "replicasets",

		"rc":                     "replicationcontrollers",
		"replicationcontroller":  "replicationcontrollers",
		"replicationcontrollers": "replicationcontrollers",

		"sts":          "statefulsets",
		"statefulset":  "statefulsets",
		"statefulsets": "statefulsets",

		// SERVICE APIS
		"ep":        "endpoints",
		"endpoint":  "endpoints",
		"endpoints": "endpoints",

		"ing":       "ingresses",
		"ingress":   "ingresses",
		"ingresses": "ingresses",

		"svc":      "services",
		"service":  "services",
		"services": "services",

		// CONFIG AND STORAGE APIS
		"cm":         "configmaps",
		"configmap":  "configmaps",
		"configmaps": "configmaps",

		"secret":  "secrets",
		"secrets": "secrets",

		"pvc":                    "persistentvolumeclaims",
		"persistentvolumeclaim":  "persistentvolumeclaims",
		"persistentvolumeclaims": "persistentvolumeclaims",

		"sc":             "storageclasses",
		"storageclasse":  "storageclasses",
		"storageclasses": "storageclasses",

		// METADATA APIS
		"crd":                       "customresourcedefinitions",
		"crds":                      "customresourcedefinitions",
		"customresourcedefinition":  "customresourcedefinitions",
		"customresourcedefinitions": "customresourcedefinitions",

		"ev":     "events",
		"event":  "events",
		"events": "events",

		"hpa":                      "horizontalpodautoscalers",
		"horizontalpodautoscaler":  "horizontalpodautoscalers",
		"horizontalpodautoscalers": "horizontalpodautoscalers",

		"vpa":                    "verticalpodautoscalers",
		"verticalpodautoscaler":  "verticalpodautoscalers",
		"verticalpodautoscalers": "verticalpodautoscalers",

		"pdb":                  "poddisruptionbudgets",
		"poddisruptionbudget":  "poddisruptionbudgets",
		"poddisruptionbudgets": "poddisruptionbudgets",

		// CLUSTER APIS
		"apiservice":  "apiservices",
		"apiservices": "apiservices",

		"clusterrole":  "clusterroles",
		"clusterroles": "clusterroles",

		"clusterrolebinding":  "clusterrolebindings",
		"clusterrolebindings": "clusterrolebindings",

		"ns":         "namespaces",
		"namespace":  "namespaces",
		"namespaces": "namespaces",

		"no":    "nodes",
		"node":  "nodes",
		"nodes": "nodes",

		"pv":                "persistentvolumes",
		"persistentvolume":  "persistentvolumes",
		"persistentvolumes": "persistentvolumes",

		"quota":          "resourcequotas",
		"resourcequota":  "resourcequotas",
		"resourcequotas": "resourcequotas",

		"role":  "roles",
		"roles": "roles",

		"rolebinding":  "rolebindings",
		"rolebindings": "rolebindings",

		"sa":              "serviceaccounts",
		"serviceaccount":  "serviceaccounts",
		"serviceaccounts": "serviceaccounts",

		// ETC APIS
		"cluster":         "kubefedclusters",
		"kubefedcluster":  "kubefedclusters",
		"kubefedclusters": "kubefedclusters",

		"odeploy":            "openmcpdeployments",
		"openmcpdeployment":  "openmcpdeployments",
		"openmcpdeployments": "openmcpdeployments",

		"osvc":            "openmcpservices",
		"openmcpservice":  "openmcpservices",
		"openmcpservices": "openmcpservices",

		"oing":            "openmcpingresss",
		"openmcpingress":  "openmcpingresss",
		"openmcpingresss": "openmcpingresss",

		"ohas":                     "openmcphybridautoscalers",
		"openmcphybridautoscaler":  "openmcphybridautoscalers",
		"openmcphybridautoscalers": "openmcphybridautoscalers",

		"opol":           "openmcppolicys",
		"openmcppolicy":  "openmcppolicys",
		"openmcppolicys": "openmcppolicys",

		"ocm":               "openmcpconfigmaps",
		"openmcpconfigmap":  "openmcpconfigmaps",
		"openmcpconfigmaps": "openmcpconfigmaps",

		"osec":           "openmcpsecrets",
		"openmcpsecret":  "openmcpsecrets",
		"openmcpsecrets": "openmcpsecrets",

		"ode":                 "openmcpdnsendpoints",
		"openmcpdnsendpoint":  "openmcpdnsendpoints",
		"openmcpdnsendpoints": "openmcpdnsendpoints",
	}
	apiGroup = map[string]string{
		// WORKLOAD APIS
		"cronjobs":               "/apis/batch/v1beta1",
		"daemonsets":             "/apis/apps/v1",
		"deployments":            "/apis/apps/v1",
		"jobs":                   "/apis/batch/v1",
		"pods":                   "/api/v1",
		"replicasets":            "/apis/apps/v1",
		"replicationcontrollers": "/api/v1",
		"statefulsets":           "/apis/apps/v1",

		// SERVICE APIS
		"endpoints": "/api/v1",
		"ingresses": "/apis/networking.k8s.io/v1beta1",
		"services":  "/api/v1",

		// CONFIG AND STORAGE APIS
		"configmaps":             "/api/v1",
		"secrets":                "/api/v1",
		"persistentvolumeclaims": "/api/v1",
		"storageclasses":         "/apis/storage.k8s.io/v1",

		// METADATA APIS
		"customresourcedefinitions": "/apis/apiextensions.k8s.io/v1",
		"events":                    "/api/v1",
		"horizontalpodautoscalers":  "/apis/autoscaling/v1",
		"verticalpodautoscalers":    "/apis/autoscaling.k8s.io/v1beta2",
		"poddisruptionbudgets":      "/apis/policy/v1beta1",

		// CLUSTER APIS
		"apiservices":         "/apis/apiregistration.k8s.io/v1",
		"clusterroles":        "/apis/rbac.authorization.k8s.io/v1",
		"clusterrolebindings": "/apis/rbac.authorization.k8s.io/v1",
		"namespaces":          "/api/v1",
		"nodes":               "/api/v1",
		"persistentvolumes":   "/api/v1",
		"resourcequotas":      "/api/v1",
		"roles":               "/apis/rbac.authorization.k8s.io/v1",
		"rolebindings":        "/apis/rbac.authorization.k8s.io/v1",
		"serviceaccounts":     "/api/v1",

		// ETC APIS
		"kubefedclusters":          "/apis/core.kubefed.io/v1beta1",
		"openmcpdeployments":       "/apis/openmcp.k8s.io/v1alpha1",
		"openmcpservices":          "/apis/openmcp.k8s.io/v1alpha1",
		"openmcpingresss":          "/apis/openmcp.k8s.io/v1alpha1",
		"openmcphybridautoscalers": "/apis/openmcp.k8s.io/v1alpha1",
		"openmcppolicys":           "/apis/openmcp.k8s.io/v1alpha1",
		"openmcpconfigmaps":        "/apis/openmcp.k8s.io/v1alpha1",
		"openmcpsecrets":           "/apis/openmcp.k8s.io/v1alpha1",
		"openmcpdnsendpoints":      "/apis/openmcp.k8s.io/v1alpha1",
	}

	KindMap = map[string]string{
		// WORKLOAD APIS
		"CronJob":                "cronjobs",
		"DaemonSet":              "daemonsets",
		"Deployment":             "deployments",
		"Job":                    "jobs",
		"Pod":                    "pods",
		"ReplicaSets":            "replicasets",
		"ReplicationControllers": "replicationcontrollers",
		"StatefulSets":           "statefulsets",

		// SERVICE APIS
		"Endpoints": "endpoints",
		"Ingress":   "ingresses",
		"Service":   "services",

		// CONFIG AND STORAGE APIS
		"ConfigMap":             "configmaps",
		"Secret":                "secrets",
		"PersistentVolumeClaim": "persistentvolumeclaims",
		"StorageClass":          "storageclasses",

		// METADATA APIS
		"CustomResourceDefinition": "customresourcedefinitions",
		"Event":                    "events",
		"HorizontalPodAutoscaler":  "horizontalpodautoscalers",
		"VerticalPodAutoscaler":    "verticalpodautoscalers",
		"PodDisruptionBudget":      "poddisruptionbudgets",

		// CLUSTER APIS
		"APIService":         "apiservices",
		"ClusterRole":        "clusterroles",
		"ClusterRoleBinding": "clusterrolebindings",
		"Namespace":          "namespaces",
		"Node":               "nodes",
		"PersistentVolume":   "persistentvolumes",
		"ResourceQuota":      "resourcequotas",
		"Role":               "roles",
		"RoleBinding":        "rolebindings",
		"ServiceAccount":     "serviceaccounts",

		// ETC APIS
		"KubeFedCluster":          "kubefedclusters",
		"OpenMCPDeployment":       "openmcpdeployments",
		"OpenMCPService":          "openmcpservices",
		"OpenMCPIngress":          "openmcpingresss",
		"OpenMCPHybridAutoScaler": "openmcphybridautoscalers",
		"OpenMCPPolicy":           "openmcppolicys",
		"OpenMCPConfigMap":        "openmcpconfigmaps",
		"OpenMCPSecret":           "openmcpsecrets",
		"OpenMCPDNSEndpoint":      "openmcpdnsendpoints",
	}

	noNamespaceResources = []string{"storageclasses", "customresourcedefinitions", "apiservices", "clusterroles", "clusterrolebindings", "namespaces", "nodes", "persistentvolumes"}
)

type MetaInfo struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `json:"metadata"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Details struct {
		Name  string `json:"name"`
		Group string `json:"group"`
		Kind  string `json:"kind"`
	} `json:"details"`
	Code int `json:"code"`
}

func CreateLinkParser(metainfo *MetaInfo) string {
	LINK := "http://" + OpenMCPAPIServer + apiGroup[KindMap[metainfo.Kind]] + "/namespaces/"

	if metainfo.Metadata.Namespace == "" {
		LINK += "default"
	} else {
		LINK += metainfo.Metadata.Namespace
	}
	LINK += "/" + KindMap[metainfo.Kind]

	if Option_context == "" {
		LINK += "?clustername=openmcp"
	} else {
		LINK += "?clustername=" + Option_context
	}

	return LINK
}
func ApplyLinkParser(metainfo *MetaInfo) string {
	LINK := "http://" + OpenMCPAPIServer + apiGroup[KindMap[metainfo.Kind]] + "/namespaces/"

	if metainfo.Metadata.Namespace == "" {
		LINK += "default"
	} else {
		LINK += metainfo.Metadata.Namespace
	}
	LINK += "/" + KindMap[metainfo.Kind]
	LINK += "/" + metainfo.Metadata.Name

	if Option_context == "" {
		LINK += "?clustername=openmcp"
	} else {
		LINK += "?clustername=" + Option_context
	}

	return LINK
}
func GetLinkParser(resourceKind, resourceName, resourceNamespace, clusterContext string) string {
	LINK := "http://" + OpenMCPAPIServer
	LINK += apiGroup[ResourceMap[resourceKind]]

	if contains(noNamespaceResources, ResourceMap[resourceKind]) {
		LINK += "/" + ResourceMap[resourceKind]
		if resourceName != "" {
			LINK += "/" + resourceName
		}
	} else {
		if Option_allnamespace {
			LINK += "/" + ResourceMap[resourceKind]
		} else {
			LINK += "/namespaces"

			if resourceNamespace != "" {
				LINK += "/" + resourceNamespace
			} else {
				LINK += "/default"
			}
			LINK += "/" + ResourceMap[resourceKind]

			if resourceName != "" {
				LINK += "/" + resourceName
			}
		}

	}
	LINK += "?clustername=" + clusterContext

	return LINK
}
func LogLinkParser(resourceKind, resourceName, clusterContext string) string {
	LINK := "http://" + OpenMCPAPIServer
	LINK += apiGroup[ResourceMap[resourceKind]]

	LINK += "/namespaces"

	if Option_namespace != "" {
		LINK += "/" + Option_namespace
	} else {
		LINK += "/default"
	}

	LINK += "/" + ResourceMap[resourceKind]

	LINK += "/" + resourceName
	LINK += "/log"

	LINK += "?clustername=" + clusterContext

	return LINK
}
func ExecLinkParser(resourceName string, clusterContext string) string {
	LINK := "http://" + OpenMCPAPIServer
	//LINK += "/omcpexec"
	LINK += apiGroup[ResourceMap["pods"]]

	LINK += "/namespaces"

	var namespace string

	if Option_namespace != "" {
		LINK += "/" + Option_namespace
		namespace = Option_namespace
	} else {
		LINK += "/default"
		namespace = "default"
	}

	LINK += "/" + ResourceMap["pods"]

	LINK += "/" + resourceName
	LINK += "/exec"

	if Option_context != "" {
		LINK += "?clustername=" + Option_context
	} else {
		LINK += "?clustername=openmcp"
	}

	LINK += "&podname=" + resourceName
	LINK += "&podnamespace=" + namespace

	if Option_containerName != "" {
		LINK += "&containername=" + Option_containerName
	}

	if strconv.FormatBool(Option_stdin) != "" {
		LINK += "&stdin=" + strconv.FormatBool(Option_stdin)
	} else {
		LINK += "&stdin=false"
	}

	if strconv.FormatBool(Option_tty) != "" {
		LINK += "&tty=" + strconv.FormatBool(Option_tty)
	} else {
		LINK += "&tty=false"
	}

	LINK += "&stderr=true"

	LINK += "&stdout=true"

	return LINK
}

func DeleteLinkParser(metainfo *MetaInfo, metainfoKindType string) string {
	var kindtype string

	if metainfoKindType == "kind" {
		kindtype = KindMap[metainfo.Kind]
	} else if metainfoKindType == "resource" {
		kindtype = ResourceMap[metainfo.Kind]
	}

	LINK := "http://" + OpenMCPAPIServer + apiGroup[kindtype] + "/namespaces/"

	if metainfo.Metadata.Namespace == "" {
		LINK += "default"
	} else {
		LINK += metainfo.Metadata.Namespace
	}

	if metainfo.Kind != "" {
		LINK += "/" + kindtype

		if metainfo.Metadata.Name != "" {
			LINK += "/" + metainfo.Metadata.Name
		}
	}

	if Option_context == "" {
		LINK += "?clustername=openmcp"
	} else {
		LINK += "?clustername=" + Option_context
	}

	return LINK
}
