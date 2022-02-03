package cmd

func aksFlags() {

	//addon

	AKSDisableAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSDisableAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSDisableAddonsCmd.Flags().StringP("addon", "a", "", "Specify the Kubernetes addon to disable.")
	AKSDisableAddonsCmd.MarkFlagRequired("resource-group")
	AKSDisableAddonsCmd.MarkFlagRequired("name")
	AKSDisableAddonsCmd.MarkFlagRequired("addon")

	AKSEnableAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSEnableAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSEnableAddonsCmd.Flags().StringP("addon", "a", "", "Specify the Kubernetes addon to enable.")
	AKSEnableAddonsCmd.MarkFlagRequired("resource-group")
	AKSEnableAddonsCmd.MarkFlagRequired("name")
	AKSEnableAddonsCmd.MarkFlagRequired("addon")

	AKSListAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSListAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSListAddonsCmd.MarkFlagRequired("resource-group")
	AKSListAddonsCmd.MarkFlagRequired("name")

	AKSShowAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSShowAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSShowAddonsCmd.Flags().StringP("addon", "a", "", "Specify the Kubernetes addon.")
	AKSShowAddonsCmd.MarkFlagRequired("resource-group")
	AKSShowAddonsCmd.MarkFlagRequired("name")
	AKSShowAddonsCmd.MarkFlagRequired("addon")

	AKSUpdateAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSUpdateAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSUpdateAddonsCmd.Flags().StringP("addon", "a", "", "")
	AKSUpdateAddonsCmd.MarkFlagRequired("resource-group")
	AKSUpdateAddonsCmd.MarkFlagRequired("name")
	AKSUpdateAddonsCmd.MarkFlagRequired("addon")

	//Pod-Identity

	AKSPIAddCmd.Flags().StringP("resource-group", "g", "", "Name of resource group. ")
	AKSPIAddCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIAddCmd.Flags().StringP("namespace", "", "", "The pod identity namespace.")
	AKSPIAddCmd.Flags().StringP("identity-resource-id", "", "", "Resource id of the identity to use.")
	AKSPIAddCmd.Flags().StringP("binding-selector", "", "", "Optional binding selector to use.")
	AKSPIAddCmd.Flags().StringP("name", "n", "", "The pod identity name. Generate if not specified.")
	AKSPIAddCmd.MarkFlagRequired("resource-group")
	AKSPIAddCmd.MarkFlagRequired("cluster-name")
	AKSPIAddCmd.MarkFlagRequired("namespace")
	AKSPIAddCmd.MarkFlagRequired("identity-resource-id")

	AKSPIDeleteCmd.Flags().StringP("resource-group", "g", "", "Name of resource group. ")
	AKSPIDeleteCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIDeleteCmd.Flags().StringP("namespace", "", "", "The pod identity namespace.")
	AKSPIDeleteCmd.Flags().String("name", "", "The pod identity name. Generate if not specified.")
	AKSPIDeleteCmd.MarkFlagRequired("cluster-name")
	AKSPIDeleteCmd.MarkFlagRequired("namespace")
	AKSPIDeleteCmd.MarkFlagRequired("name")
	AKSPIDeleteCmd.MarkFlagRequired("resource-group")

	AKSPIListCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSPIListCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIListCmd.MarkFlagRequired("resource-group")
	AKSPIListCmd.MarkFlagRequired("cluster-name")

	AKSPIExceptionCmd.Flags().String("pod-labels", "", "Space-separated labels: key=value [key=value ...].")

	AKSPIExceptionAddCmd.Flags().String("pod-labels", "", "Space-separated labels: key=value [key=value ...].")
	AKSPIExceptionAddCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSPIExceptionAddCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIExceptionAddCmd.Flags().StringP("namespace", "", "", "The pod identity namespace.")
	AKSPIExceptionAddCmd.Flags().StringP("name", "n", "", "The pod identity exception name. Generate if not specified.")
	AKSPIExceptionAddCmd.MarkFlagRequired("cluster-name")
	AKSPIExceptionAddCmd.MarkFlagRequired("namespace")
	AKSPIExceptionAddCmd.MarkFlagRequired("pod-labels")
	AKSPIExceptionAddCmd.MarkFlagRequired("resource-group")

	AKSPIExceptionDeleteCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSPIExceptionDeleteCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIExceptionDeleteCmd.Flags().StringP("namespace", "", "", "The pod identity exception namespace to remove.")
	AKSPIExceptionDeleteCmd.Flags().StringP("name", "n", "", "The pod identity exception name to remove.")
	AKSPIExceptionDeleteCmd.MarkFlagRequired("cluster-name")
	AKSPIExceptionDeleteCmd.MarkFlagRequired("name")
	AKSPIExceptionDeleteCmd.MarkFlagRequired("namespace")
	AKSPIExceptionDeleteCmd.MarkFlagRequired("resource-group")

	AKSPIExceptionListCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSPIExceptionListCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIExceptionListCmd.MarkFlagRequired("cluster-name")
	AKSPIExceptionListCmd.MarkFlagRequired("resource-group")

	AKSPIExceptionUpdateCmd.Flags().String("pod-labels", "", "Space-separated labels: key=value [key=value ...].")
	AKSPIExceptionUpdateCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSPIExceptionUpdateCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	AKSPIExceptionUpdateCmd.Flags().StringP("namespace", "", "", "The pod identity namespace.")
	AKSPIExceptionUpdateCmd.Flags().StringP("name", "n", "", "The pod identity exception name to remove.")
	AKSPIExceptionUpdateCmd.MarkFlagRequired("cluster-name")
	AKSPIExceptionUpdateCmd.MarkFlagRequired("name")
	AKSPIExceptionUpdateCmd.MarkFlagRequired("namespace")
	AKSPIExceptionUpdateCmd.MarkFlagRequired("resource-group")
	AKSPIExceptionUpdateCmd.MarkFlagRequired("pod-labels")

	// maintenanceconfiguration

	MCAddCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	MCAddCmd.Flags().StringP("name", "n", "", "The config name.")
	MCAddCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	MCAddCmd.Flags().StringP("config-file", "", "", "The maintenance configuration json file.")
	MCAddCmd.MarkFlagRequired("cluster-name")
	MCAddCmd.MarkFlagRequired("name")
	MCAddCmd.MarkFlagRequired("resource-group")

	MCDeleteCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	MCDeleteCmd.Flags().StringP("name", "n", "", "The config name.")
	MCDeleteCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	MCDeleteCmd.MarkFlagRequired("cluster-name")
	MCDeleteCmd.MarkFlagRequired("name")
	MCDeleteCmd.MarkFlagRequired("resource-group")

	MCUpdateCmd.Flags().StringP("cluster-name", "", "", "")
	MCUpdateCmd.Flags().StringP("name", "n", "", "")
	MCUpdateCmd.Flags().StringP("resource-group", "g", "", "")
	MCUpdateCmd.Flags().StringP("config-file", "", "", "configfile")
	MCUpdateCmd.MarkFlagRequired("cluster-name")
	MCUpdateCmd.MarkFlagRequired("name")
	MCUpdateCmd.MarkFlagRequired("resource-group")

	MCListCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	MCListCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	MCListCmd.MarkFlagRequired("resource-group")
	MCListCmd.MarkFlagRequired("cluster-name")

	MCShowCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	MCShowCmd.Flags().StringP("cluster-name", "", "", "The cluster name.")
	MCShowCmd.Flags().StringP("name", "", "", "The config name.")
	MCShowCmd.MarkFlagRequired("cluster-name")
	MCShowCmd.MarkFlagRequired("resource-group")
	MCShowCmd.MarkFlagRequired("name")

	// k8sconfiguration

	AKSConfigurationCreate.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConfigurationCreate.Flags().StringP("name", "n", "", "The name of the Kubernetes configuration.")
	AKSConfigurationCreate.Flags().StringP("cluster-name", "c", "", "Name of the managed cluster.")
	AKSConfigurationCreate.Flags().StringP("cluster-type", "", "", "Specifies an Arc cluster or AKS-managed cluster.")
	AKSConfigurationCreate.Flags().StringP("repository-url", "u", "", "The URL of the source control repository.")
	AKSConfigurationCreate.Flags().String("scope", "", "Scope the operator to either 'namespace' or 'cluster'. Allowed values: cluster, namespace")

	AKSConfigurationDelete.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConfigurationDelete.Flags().StringP("name", "n", "", "The name of the Kubernetes configuration.")
	AKSConfigurationDelete.Flags().StringP("cluster-name", "c", "", "Name of the managed cluster.")
	AKSConfigurationDelete.Flags().StringP("cluster-type", "", "", "Specifies an Arc cluster or AKS-managed cluster.")
	AKSConfigurationDelete.MarkFlagRequired("resource-group")
	AKSConfigurationDelete.MarkFlagRequired("name")
	AKSConfigurationDelete.MarkFlagRequired("cluster-name")
	AKSConfigurationDelete.MarkFlagRequired("cluster-type")

	AKSConfigurationList.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConfigurationList.Flags().StringP("cluster-name", "c", "", "Name of the managed cluster.")
	AKSConfigurationList.Flags().StringP("cluster-type", "", "", "Specifies an Arc cluster or AKS-managed cluster.")
	AKSConfigurationList.MarkFlagRequired("resource-group")
	AKSConfigurationList.MarkFlagRequired("cluster-name")
	AKSConfigurationList.MarkFlagRequired("cluster-type")

	AKSConfigurationShow.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConfigurationShow.Flags().StringP("name", "n", "", "The name of the Kubernetes configuration.")
	AKSConfigurationShow.Flags().StringP("cluster-name", "c", "", "Name of the managed cluster.")
	AKSConfigurationShow.Flags().StringP("cluster-type", "", "", "Specifies an Arc cluster or AKS-managed cluster.")
	AKSConfigurationShow.MarkFlagRequired("resource-group")
	AKSConfigurationShow.MarkFlagRequired("name")
	AKSConfigurationShow.MarkFlagRequired("cluster-name")
	AKSConfigurationShow.MarkFlagRequired("cluster-type")

	// connectedk8s

	AKSConnectedk8sConnectCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sConnectCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sConnectCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sConnectCmd.MarkFlagRequired("name")

	AKSConnectedk8sDeleteCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sDeleteCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sDeleteCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sDeleteCmd.MarkFlagRequired("name")

	AKSConnectedk8sDisableFeaturesCmd.Flags().StringSliceVarP(&slice, "features", "s", []string{}, "")
	AKSConnectedk8sDisableFeaturesCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sDisableFeaturesCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sDisableFeaturesCmd.MarkFlagRequired("features")
	AKSConnectedk8sDisableFeaturesCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sDisableFeaturesCmd.MarkFlagRequired("name")

	AKSConnectedk8sEnableFeaturesCmd.Flags().StringSliceVarP(&slice, "features", "s", []string{}, "")
	AKSConnectedk8sEnableFeaturesCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sEnableFeaturesCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sEnableFeaturesCmd.MarkFlagRequired("features")
	AKSConnectedk8sEnableFeaturesCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sEnableFeaturesCmd.MarkFlagRequired("name")

	AKSConnectedk8sListCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")

	AKSConnectedk8sProxyCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sProxyCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sProxyCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sProxyCmd.MarkFlagRequired("name")

	AKSConnectedk8sShowCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sShowCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sShowCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sShowCmd.MarkFlagRequired("name")

	AKSConnectedk8sUpdateCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sUpdateCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sUpdateCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sUpdateCmd.MarkFlagRequired("name")

	AKSConnectedk8sUpgradeCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedk8sUpgradeCmd.Flags().StringP("name", "n", "", "The name of the connected cluster.")
	AKSConnectedk8sUpgradeCmd.MarkFlagRequired("resource-group")
	AKSConnectedk8sUpgradeCmd.MarkFlagRequired("name")

	// etc

	StartCmd.Flags().StringP("resource-group", "g", "", "Name of resource group. ")
	StartCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	StartCmd.MarkFlagRequired("resource-group")
	StartCmd.MarkFlagRequired("name")

	StopCmd.Flags().StringP("resource-group", "g", "", "Name of resource group. ")
	StopCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	StopCmd.MarkFlagRequired("resource-group")
	StopCmd.MarkFlagRequired("name")

	RotateCertsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	RotateCertsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	RotateCertsCmd.MarkFlagRequired("resource-group")
	RotateCertsCmd.MarkFlagRequired("name")

	GetOSoptionsCmd.Flags().StringP("location", "l", "", "Location.")
	GetOSoptionsCmd.MarkFlagRequired("location")

	AKSAppUpCmd.Flags().String("acr", "", "The Azure Container Registry name used to push the image.")
	AKSAppUpCmd.Flags().String("aks-cluster", "", "The name of the cluster to select for deployment.")
	AKSAppUpCmd.Flags().String("branch-name", "", "The new branch name to create to check in the file and raise the PR.")
	AKSAppUpCmd.Flags().String("do-not-wait", "", "It does not wait for the workflow to complete.")
	AKSAppUpCmd.Flags().String("port", "", "The port on which the application runs. The default is 8080.")
	AKSAppUpCmd.Flags().StringP("repository", "r", "", "The URL of your GitHub repository")

	AKSBrowseCmd.Flags().Bool("disable-browser", false, "Don't launch a web browser after establishing port-forwarding.")
	AKSBrowseCmd.Flags().String("listen-address", "", "The listening address of the dashboard. Default: 127.0.0.1")
	AKSBrowseCmd.Flags().String("listen-port", "", "The listening port for the dashboard. Default: 8001")
	AKSBrowseCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSBrowseCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSBrowseCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSBrowseCmd.MarkFlagRequired("name")
	AKSBrowseCmd.MarkFlagRequired("resource-group")

	AKSCheckAcrCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSCheckAcrCmd.Flags().String("acr", "", "The FQDN of the ACR.")
	AKSCheckAcrCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSCheckAcrCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSCheckAcrCmd.MarkFlagRequired("acr")
	AKSCheckAcrCmd.MarkFlagRequired("name")
	AKSCheckAcrCmd.MarkFlagRequired("resource-group")

	AKSGetUpgradesCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSGetUpgradesCmd.Flags().StringP("name", "n", "", "The cluster name.")
	AKSGetUpgradesCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSGetUpgradesCmd.MarkFlagRequired("resource-group")
	AKSGetUpgradesCmd.MarkFlagRequired("name")

	AKSGetVersionsCmd.Flags().StringP("location", "l", "", "Location")
	AKSGetVersionsCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSGetVersionsCmd.MarkFlagRequired("location")

	AKSNodepoolGetUpgradesCmd.Flags().String("nodepool-name", "", "The name of the node pool.")
	AKSNodepoolGetUpgradesCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSNodepoolGetUpgradesCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSNodepoolGetUpgradesCmd.Flags().String("cluster-name", "", "The cluster name.")
	AKSNodepoolGetUpgradesCmd.MarkFlagRequired("nodepool-name")
	AKSNodepoolGetUpgradesCmd.MarkFlagRequired("resource-group")
	AKSNodepoolGetUpgradesCmd.MarkFlagRequired("cluster-name")

	AKSInstallCLICmd.Flags().String("base-src-url", "", "Default download source URL for Kubectl releases.")
	AKSInstallCLICmd.Flags().String("client-version", "", "The version of kubectl to install. Default: latest")
	AKSInstallCLICmd.Flags().String("install-location", "", "The path where you want to install Kubectl. Default: ~/.azure-kubectl/kubectl.exe")
	AKSInstallCLICmd.Flags().StringP("kubelogin-base-src-url", "l", "", "Default download source URL for Kubelgin releases.")
	AKSInstallCLICmd.Flags().String("kubelogin-install-location", "", "The path to install Kubelogin. Default: ~/.azure-kubelogin/kubelogin.exe")
	AKSInstallCLICmd.Flags().String("kubelogin-version", "", "The version of kubelogin to install. Default: latest")
	AKSInstallCLICmd.Flags().String("subscription", "", "The name or ID of the subscription.")

	AKSKanalyzeCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSKanalyzeCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSKanalyzeCmd.MarkFlagRequired("resource-group")
	AKSKanalyzeCmd.MarkFlagRequired("name")

	AKSKollectCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSKollectCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSKollectCmd.Flags().String("storage-account", "", "Name or ID of the storage account to save the diagnostic information.")
	AKSKollectCmd.MarkFlagRequired("resource-group")
	AKSKollectCmd.MarkFlagRequired("name")
}
