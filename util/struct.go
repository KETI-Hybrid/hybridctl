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

// images
type GKEImages struct {
	SRC_IMAGE  string
	DEST_IMAGE []string
	IMAGE_NAME []string

	// Images List
	REPOSITORY string
	FILTER     string
	LIMIT      string
	PAGE_SIZE  string
	SORT_BY    string
	URI        bool

	// Images Delete
	FORCE_DELETE_TAGS bool
}

type GKEAuth struct {
	CRED_FILE  string
	REGISTRIES string
	ACCOUNTS   string
	ALL        bool

	// List
	FILTER_ACCOUNT string
	FILTER         string
	LIMIT          string
	PAGE_SIZE      string
	SORT_BY        string
}

type GKEOperations struct {
	PROJECT_ID   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ZONE         string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	OPERATION_ID string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	NAME         string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

type pushblock int

const (
	Enable  pushblock = 1
	Disable pushblock = 0
	None    pushblock = -1
)

func (p pushblock) String() string {
	if p == Enable {
		return "enable"
	} else if p == Disable {
		return "disable"
	}
	return "none"
}

type GKESource struct {
	PUSHBLOCK       pushblock
	MESSAGE_FORMAT  string
	SERVICE_ACCOUNT string
	TOPIC_PROJECT   string
	ADD_TOPIC       string
	REMOVE_TOPIC    string
	UPDATE_TOPIC    string
}

type GKESetProperty struct {
	SECTION      string
	PROPERTY     string
	VALUE        string
	INSTALLATION bool
}
