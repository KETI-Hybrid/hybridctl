package util

type Config struct {
	Properties struct {
		TimeInWeek []struct {
			Day       string `json:"day"`
			HourSlots []int  `json:"hourSlots"`
		} `json:"timeInWeek"`
		NotAllowedTime []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"notAllowedTime"`
	} `json:"properties"`
}

type Error struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type EKSAPIParameter struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	ApiVersion        string
	Location          string
	ConfigName        string
	ConfigFile        Config
}

type AKSAPIParameter struct {
	ResourceGroupName string
	ClusterName       string
	ConfigName        string
	ConfigFile        Config
	Acr               string
	AksCluster        string
	BranchName        string
	DoNotWait         string
	Port              string
	BindingSelector   string
	Repository        string
	DisableBrowser    bool
	ListenAddress     string
	ListenPort        string
	Subscription      string
	Location          string
	NodepoolName      string
	StorageAccount    string
	Features          []string
	Addon             AKSAddon            `json:"addon"`
	PodIdentity       AKSPodIdentity      `json:"podIdentity"`
	Install           AKSInstallCLI       `json:"install"`
	K8sConfiguration  AKSk8sConfiguration `json:"k8sConfiguration"`
}

type AKSAddon struct {
	Addon                      string
	AppgwID                    string
	AppgwName                  string
	AppgwSubnetCidr            string
	AppgwSubnetID              string
	AppgwSubnetPrefix          string
	AppgwWatchNamespace        string
	EnableMsiAuthForMonitoring bool
	EnableSecretRotation       bool
	EnableSgxquotehelper       bool
	SubnetName                 string
	WorkspaceResourceID        string
}

type AKSPodIdentity struct {
	Namespace          string
	IdentityResourceID string
	Name               string
	BindingSelector    string
	PodLabels          string
}

type AKSInstallCLI struct {
	BaseSrcURL               string
	ClientVersion            string
	InstallLocation          string
	KubeloginBaseSrcURL      string
	KubeloginInstallLocation string
	KubeloginVersion         string
	Subscription             string
}

type AKSk8sConfiguration struct {
	ClusterType   string
	Name          string
	RepositoryURL string
	Scope         string
}
