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
	Acr             string
	AksCluster      string
	BranchName      string
	DoNotWait       string
	Port            string
	BindingSelector string
	Repository      string
	Name            string
	ResourceGroup   string
	DisableBrowser  bool
	ListenAddress   string
	ListenPort      string
	Subscription    string
	Location        string
}

type AKSAddon struct {
	ResourceGroupName          string
	ClusterName                string
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
	ResourceGroupName  string
	ClusterName        string
	Namespace          string
	IdentityResourceID string
	Name               string
	BindingSelector    string
	PodLabels          string
}

// type AKSAppUp struct {
// 	Acr             string
// 	AksCluster      string
// 	BranchName      string
// 	DoNotWait       string
// 	Port            string
// 	BindingSelector string
// 	Repository      string
// }
// type AKSBrowse struct {
// 	Name           string
// 	ResourceGroup  string
// 	DisableBrowser bool
// 	ListenAddress  string
// 	ListenPort     string
// 	Subscription   string
// }

// type AKSAcr struct {
// 	Name          string
// 	ResourceGroup string
// 	Acr           string
// 	Subscription  string
// }

// type AKSCheckAcr struct {
// 	Acr           string
// 	Name          string
// 	ResourceGroup string
// 	Subscription  string
// }

type CloudError struct {
	// Error - Details about the error.
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
