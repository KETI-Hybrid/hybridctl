package cmd

// GKE terraform structure

type Cli struct {
	PlatformName string
	ClusterName  string
	NodeName     string
	NodeCount    string
	Version      string
}

type Cluster_info struct {
	Project_id    string `json:"project_id"`
	Cluster_name  string `json:"cluster_name"`
	Region        string `json:"region"`
	Gke_num_nodes uint64 `json:"gke_num_nodes"`
}

type TF struct {
	Resource *Resource `json:"resource"`
}

type TF_AKS struct {
	ResourceAksCluster *ResourceAksCluster `json:"resource"`
}

type ResourceAksCluster struct {
	AzurernKubernetesCluster *map[string]AksCluster `json:"azurerm_kubernetes_cluster"`
}

type Resource struct {
	Google_container_cluster *map[string]Cluster_type `json:"google_container_cluster"`
}

type SSHKey struct {
	KeyData string `json:"key_data"`
}
type LinuxProfile struct {
	AdminUsername string `json:"admin_username"`
	SSHKey        SSHKey `json:"ssh_key"`
}
type DefaultNodePool struct {
	Name      string `json:"name"`
	NodeCount int    `json:"node_count"`
	VMSize    string `json:"vm_size"`
}
type ServicePrincipal struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type OmsAgent struct {
	Enabled                 string `json:"enabled"`
	LogAnalyticsWorkspaceID string `json:"log_analytics_workspace_id"`
}
type AddonProfile struct {
	OmsAgent OmsAgent `json:"oms_agent"`
}
type NetworkProfile struct {
	LoadBalancerSku string `json:"load_balancer_sku"`
	NetworkPlugin   string `json:"network_plugin"`
}
type Tags struct {
	Environment string `json:"Environment"`
}
type AksCluster struct {
	Name               string           `json:"name"`
	Kubernetes_Version string           `json:"kubernetes_version"`
	Location           string           `json:"location"`
	ResourceGroupName  string           `json:"resource_group_name"`
	DNSPrefix          string           `json:"dns_prefix"`
	LinuxProfile       LinuxProfile     `json:"linux_profile"`
	DefaultNodePool    DefaultNodePool  `json:"default_node_pool"`
	ServicePrincipal   ServicePrincipal `json:"service_principal"`
	AddonProfile       AddonProfile     `json:"addon_profile"`
	NetworkProfile     NetworkProfile   `json:"network_profile"`
	Tags               Tags             `json:"tags"`
}

type Cluster_type struct {
	Name                     string     `json:"name"`
	Location                 string     `json:"location"`
	Remove_default_node_pool string     `json:"remove_default_node_pool"`
	Initial_node_count       int        `json:"initial_node_count"`
	Life_Cycle               Life_Cycle `json:"lifecycle"`
}

type Life_Cycle struct {
	Ignore_Changes []string `json:"ignore_changes"`
}

//--------------------------eks structure---------------------

type TF_EKS struct {
	ResourceEksCluster ResourceEksCluster `json:"resource"`
}
type VpcConfig struct {
	SecurityGroupIds      []string `json:"security_group_ids"`
	SubnetIds             string   `json:"subnet_ids"`
	EndpointPrivateAccess string   `json:"endpoint_private_access"`
	EndpointPublicAccess  string   `json:"endpoint_public_access"`
}

type EksCluster struct {
	Name                   string    `json:"name"`
	RoleArn                string    `json:"role_arn"`
	Version                string    `json:"version"`
	EnabledClusterLogTypes []string  `json:"enabled_cluster_log_types"`
	VpcConfig              VpcConfig `json:"vpc_config"`
	DependsOn              []string  `json:"depends_on"`
}

type ResourceEksCluster struct {
	AwsEksCluster *map[string]EksCluster `json:"aws_eks_cluster"`
}

//--------------------------eks structure end----------------

type TF_NodePool struct {
	NodePool_Resource *NodePool_Resource `json:"resource"`
}

type NodePool_Resource struct {
	Google_container_node_pool *map[string]Node_pool_type `json:"google_container_node_pool"`
}

type Node_pool_type struct {
	Name        string       `json:"name"`
	Location    string       `json:"location"`
	Cluster     string       `json:"cluster"`
	Node_count  string       `json:"node_count"`
	Node_config *Node_config `json:"node_config"`
}

type Labels struct {
	Env string `json:"env"`
}
type Node_config struct {
	Oauth_scopes []string  `json:"oauth_scopes"`
	Labels       *Labels   `json:"labels"`
	Machine_type string    `json:"machine_type"`
	Tags         []string  `json:"tags"`
	Metadata     *Metadata `json:"metadata"`
}

type Metadata struct {
	Disable_legacy_endpoints string `json:"disable-legacy-endpoints"`
}

//-------------------------eks structure--------------

type TF_EKS_Nopepool struct {
	EksResourceNode EksResourceNode `json:"resource"`
}
type EksLabels struct {
	Role string `json:"role"`
}
type ScalingConfig struct {
	DesiredSize string `json:"desired_size"`
	MinSize     string `json:"min_size"`
	MaxSize     string `json:"max_size"`
}
type EksTags struct {
	Name string `json:"Name"`
}
type EksClusterNode struct {
	ClusterName   string        `json:"cluster_name"`
	NodeGroupName string        `json:"node_group_name"`
	NodeRoleArn   string        `json:"node_role_arn"`
	SubnetIds     string        `json:"subnet_ids"`
	InstanceTypes []string      `json:"instance_types"`
	DiskSize      string        `json:"disk_size"`
	EksLabels     EksLabels     `json:"labels"`
	ScalingConfig ScalingConfig `json:"scaling_config"`
	DependsOn     []string      `json:"depends_on"`
	EksTags       EksTags       `json:"tags"`
}

type EksResourceNode struct {
	EksResourceNode *map[string]EksClusterNode `json:"aws_eks_node_group"`
}

//-------------------------eks structure end---------

//-------------------------aks structure--------------

type TF_AKS_Nodepool struct {
	AksNodepoolResource AksNodepoolResource `json:"resource"`
}
type AksNodeTags struct {
	Environment string `json:"environment"`
}
type AksNodepool struct {
	Name                 string      `json:"name"`
	KubernetesClusterID  string      `json:"kubernetes_cluster_id"`
	Orchestrator_Version string      `json:"orchestrator_version"`
	VMSize               string      `json:"vm_size"`
	NodeCount            string      `json:"node_count"`
	AksNodeTags          AksNodeTags `json:"tags"`
}

type AksNodepoolResource struct {
	AzurermKubernetesClusterNodePool *map[string]AksNodepool `json:"azurerm_kubernetes_cluster_node_pool"`
}

//AKS-API RESPONSE STRUCT
type AKS_Cluster_API struct {
	Value []struct {
		ID       string `json:"id"`
		Location string `json:"location"`
		Name     string `json:"name"`
		Tags     struct {
			Environment string `json:"Environment"`
		} `json:"tags"`
		Type       string `json:"type"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
			PowerState        struct {
				Code string `json:"code"`
			} `json:"powerState"`
			KubernetesVersion string `json:"kubernetesVersion"`
			DNSPrefix         string `json:"dnsPrefix"`
			Fqdn              string `json:"fqdn"`
			AzurePortalFQDN   string `json:"azurePortalFQDN"`
			AgentPoolProfiles []struct {
				Name              string `json:"name"`
				Count             int    `json:"count"`
				VMSize            string `json:"vmSize"`
				OsDiskSizeGB      int    `json:"osDiskSizeGB"`
				OsDiskType        string `json:"osDiskType"`
				KubeletDiskType   string `json:"kubeletDiskType"`
				MaxPods           int    `json:"maxPods"`
				Type              string `json:"type"`
				EnableAutoScaling bool   `json:"enableAutoScaling"`
				ProvisioningState string `json:"provisioningState"`
				PowerState        struct {
					Code string `json:"code"`
				} `json:"powerState"`
				OrchestratorVersion    string `json:"orchestratorVersion"`
				EnableNodePublicIP     bool   `json:"enableNodePublicIP"`
				Mode                   string `json:"mode"`
				EnableEncryptionAtHost bool   `json:"enableEncryptionAtHost"`
				OsType                 string `json:"osType"`
				OsSKU                  string `json:"osSKU"`
				NodeImageVersion       string `json:"nodeImageVersion"`
				EnableFIPS             bool   `json:"enableFIPS"`
			} `json:"agentPoolProfiles"`
			LinuxProfile struct {
				AdminUsername string `json:"adminUsername"`
				SSH           struct {
					PublicKeys []struct {
						KeyData string `json:"keyData"`
					} `json:"publicKeys"`
				} `json:"ssh"`
			} `json:"linuxProfile"`
			ServicePrincipalProfile struct {
				ClientID string `json:"clientId"`
			} `json:"servicePrincipalProfile"`
			AddonProfiles struct {
				Omsagent struct {
					Enabled bool `json:"enabled"`
					Config  struct {
						LogAnalyticsWorkspaceResourceID string `json:"logAnalyticsWorkspaceResourceID"`
					} `json:"config"`
				} `json:"omsagent"`
			} `json:"addonProfiles"`
			NodeResourceGroup string `json:"nodeResourceGroup"`
			EnableRBAC        bool   `json:"enableRBAC"`
			NetworkProfile    struct {
				NetworkPlugin       string `json:"networkPlugin"`
				LoadBalancerSku     string `json:"loadBalancerSku"`
				LoadBalancerProfile struct {
					ManagedOutboundIPs struct {
						Count int `json:"count"`
					} `json:"managedOutboundIPs"`
					EffectiveOutboundIPs []struct {
						ID string `json:"id"`
					} `json:"effectiveOutboundIPs"`
				} `json:"loadBalancerProfile"`
				PodCidr          string `json:"podCidr"`
				ServiceCidr      string `json:"serviceCidr"`
				DNSServiceIP     string `json:"dnsServiceIP"`
				DockerBridgeCidr string `json:"dockerBridgeCidr"`
				OutboundType     string `json:"outboundType"`
			} `json:"networkProfile"`
			MaxAgentPools          int `json:"maxAgentPools"`
			APIServerAccessProfile struct {
				EnablePrivateCluster bool `json:"enablePrivateCluster"`
			} `json:"apiServerAccessProfile"`
		} `json:"properties"`
		Sku struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku"`
	} `json:"value"`
}

type bearerToken struct {
	Token_type     string `json:"token_type" protobuf:"bytes,1,opt,name=token_type"`
	Expires_in     string `json:"expires_in" protobuf:"bytes,2,opt,name=expires_in"`
	Ext_expires_in string `json:"ext_expires_in" protobuf:"bytes,3,opt,name=ext_expires_in"`
	Expires_on     string `json:"expires_on" protobuf:"bytes,4,opt,name=expires_on"`
	Not_before     string `json:"not_before" protobuf:"bytes,5,opt,name=not_before"`
	Resource       string `json:"resource" protobuf:"bytes,6,opt,name=resource"`
	Access_token   string `json:"access_token" protobuf:"bytes,7,opt,name=access_token"`
}

type EKS_Custer_List struct {
	Clusters []string `json:"clusters"`
}
