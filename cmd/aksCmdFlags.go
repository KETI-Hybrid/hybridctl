package cmd

func aksFlags() {
	// aksCmd.PersistentFlags().StringP("resource-group", "g", "", "Name of resource group.")
	// aksCmd.PersistentFlags().StringP("name", "n", "", "Name of the managed cluster.")

	//addon
	// AddonCmd.PersistentFlags().StringP("addon", "a", "", "Specify the Kubernetes addon")

	AKSDisableAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSDisableAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSDisableAddonsCmd.Flags().StringP("addon", "a", "", "")
	AKSDisableAddonsCmd.MarkFlagRequired("resource-group")
	AKSDisableAddonsCmd.MarkFlagRequired("name")
	AKSDisableAddonsCmd.MarkFlagRequired("addon")

	AKSEnableAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSEnableAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSEnableAddonsCmd.Flags().StringP("addon", "a", "", "")
	AKSEnableAddonsCmd.MarkFlagRequired("resource-group")
	AKSEnableAddonsCmd.MarkFlagRequired("name")
	AKSEnableAddonsCmd.MarkFlagRequired("addon")

	AKSListAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSListAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSListAddonsCmd.MarkFlagRequired("resource-group")
	AKSListAddonsCmd.MarkFlagRequired("name")

	AKSShowAddonsCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSShowAddonsCmd.Flags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSShowAddonsCmd.Flags().StringP("addon", "a", "", "")
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

	AKSPodIdentityCmd.PersistentFlags().String("cluster-name", "", "The cluster name.")
	AKSPodIdentityCmd.PersistentFlags().String("namespace", "", "The pod identity namespace.")
	AKSPodIdentityCmd.PersistentFlags().StringP("name", "", "n", "The pod identity name. Generate if not specified.")

	AKSPIAddCmd.Flags().StringP("resource-group", "g", "", "The pod identity name. Generate if not specified.")
	AKSPIAddCmd.Flags().StringP("cluster-name", "c", "", "The cluster name.")
	AKSPIAddCmd.Flags().StringP("namespace", "", "", "The pod identity namespace.")
	AKSPIAddCmd.Flags().StringP("identity-resource-id", "n", "", "Resource id of the identity to use.")
	AKSPIAddCmd.MarkFlagRequired("resource-group")
	AKSPIAddCmd.MarkFlagRequired("cluster-name")
	AKSPIAddCmd.MarkFlagRequired("namespace")
	AKSPIAddCmd.MarkFlagRequired("name")
	AKSPIAddCmd.Flags().String("name", "", "Pod id 이름입니다. 지정 하지 않으면 생성 합니다.")

	AKSPIDeleteCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("namespace")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("name")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("resource-group")

	AKSPIListCmd.Flags().StringP("resource-group", "g", "", "The pod identity name. Generate if not specified.")
	AKSPIListCmd.Flags().StringP("cluster-name", "c", "", "The cluster name.")
	AKSPIListCmd.MarkPersistentFlagRequired("resource-group")
	AKSPIListCmd.MarkPersistentFlagRequired("cluster-name")

	AKSPIExceptionCmd.Flags().String("pod-labels", "", "Space-separated labels: key=value [key=value ...].")
	AKSPIExceptionAddCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIExceptionAddCmd.MarkPersistentFlagRequired("namespace")
	AKSPIExceptionAddCmd.MarkPersistentFlagRequired("pod-labels")
	AKSPIExceptionAddCmd.MarkPersistentFlagRequired("resource-group")
	AKSPIExceptionAddCmd.MarkPersistentFlagRequired("name")

	AKSPIExceptionDeleteCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIExceptionDeleteCmd.MarkPersistentFlagRequired("name")
	AKSPIExceptionDeleteCmd.MarkPersistentFlagRequired("namespace")
	AKSPIExceptionDeleteCmd.MarkPersistentFlagRequired("resource-group")

	AKSPIExceptionListCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIExceptionListCmd.MarkPersistentFlagRequired("resource-group")

	AKSPIExceptionUpdateCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIExceptionUpdateCmd.MarkPersistentFlagRequired("name")
	AKSPIExceptionUpdateCmd.MarkPersistentFlagRequired("namespace")
	AKSPIExceptionUpdateCmd.MarkPersistentFlagRequired("resource-group")
	AKSPIExceptionUpdateCmd.MarkPersistentFlagRequired("pod-labels")

	StartCmd.Flags().StringP("resource-group", "g", "", "")
	StartCmd.Flags().StringP("name", "n", "", "")
	StartCmd.MarkFlagRequired("resource-group")
	StartCmd.MarkFlagRequired("name")

	StopCmd.Flags().StringP("resource-group", "g", "", "")
	StopCmd.Flags().StringP("name", "n", "", "")
	StopCmd.MarkFlagRequired("resource-group")
	StopCmd.MarkFlagRequired("name")

	RotateCertsCmd.Flags().StringP("resource-group", "g", "", "")
	RotateCertsCmd.Flags().StringP("name", "n", "", "")
	RotateCertsCmd.MarkFlagRequired("resource-group")
	RotateCertsCmd.MarkFlagRequired("name")

	GetOSoptionsCmd.PersistentFlags().StringP("location", "l", "", "location")
	GetOSoptionsCmd.MarkPersistentFlagRequired("location")

	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("resource-group")
	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("name")

	MCAddCmd.Flags().StringP("cluster-name", "", "", "")
	MCAddCmd.Flags().StringP("name", "", "", "")
	MCAddCmd.Flags().StringP("resource-group", "g", "", "")
	MCAddCmd.Flags().StringP("config-file", "", "", "configfile")
	MCAddCmd.MarkFlagRequired("cluster-name")
	// MCAddCmd.MarkFlagRequired("name")
	MCAddCmd.MarkFlagRequired("resource-group")

	MCDeleteCmd.Flags().StringP("cluster-name", "", "", "")
	MCDeleteCmd.Flags().StringP("name", "", "", "")
	MCDeleteCmd.Flags().StringP("resource-group", "g", "", "")
	MCDeleteCmd.MarkFlagRequired("cluster-name")
	MCDeleteCmd.MarkFlagRequired("name")
	MCDeleteCmd.MarkFlagRequired("resource-group")

	MCUpdateCmd.Flags().StringP("cluster-name", "", "", "")
	MCUpdateCmd.Flags().StringP("name", "", "", "")
	MCUpdateCmd.Flags().StringP("resource-group", "g", "", "")
	MCUpdateCmd.Flags().StringP("config-file", "", "", "configfile")
	MCUpdateCmd.MarkFlagRequired("cluster-name")
	MCUpdateCmd.MarkFlagRequired("name")
	MCUpdateCmd.MarkFlagRequired("resource-group")

	MCListCmd.Flags().StringP("resource-group", "g", "", "")
	MCListCmd.Flags().StringP("name", "n", "", "")
	MCListCmd.Flags().StringP("config-name", "c", "", "configname")
	MCListCmd.MarkFlagRequired("resource-group")
	MCListCmd.MarkFlagRequired("name")

	MCShowCmd.Flags().StringP("config-name", "c", "", "configname")
	MCShowCmd.MarkFlagRequired("config-name")

	AKSAppUpCmd.Flags().String("acr", "", "The Azure Container Registry name used to push the image.")
	AKSAppUpCmd.Flags().String("aks-cluster", "", "The name of the cluster to select for deployment.")
	AKSAppUpCmd.Flags().String("branch-name", "", "The new branch name to create to check in the file and raise the PR.")
	AKSAppUpCmd.Flags().String("do-not-wait", "", "It does not wait for the workflow to complete.")
	AKSAppUpCmd.Flags().String("port", "", "The port on which the application runs. The default is 8080.")
	AKSAppUpCmd.Flags().StringP("repository", "r", "", "The URL of your GitHub repository")

	AKSBrowseCmd.Flags().Bool("disable-browser", false, "Do not start the web browser after setting up port forwarding.")
	AKSBrowseCmd.Flags().String("listen-address", "", "The listening address of the dashboard. Default: 127.0.0.1")
	AKSBrowseCmd.Flags().String("listen-port", "", "The listening port for the dashboard. Default: 8001")
	AKSBrowseCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSBrowseCmd.Flags().String("name", "", "The Azure Container Registry name used to push the image.")
	AKSBrowseCmd.Flags().String("resource-group", "", "The Azure Container Registry name used to push the image.")
	AKSBrowseCmd.MarkFlagRequired("name")
	AKSBrowseCmd.MarkFlagRequired("resource-group")

	AKSCheckAcrCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSCheckAcrCmd.Flags().String("acr", "", "The Azure Container Registry name used to push the image.")
	AKSCheckAcrCmd.Flags().String("name", "", "The Azure Container Registry name used to push the image.")
	AKSCheckAcrCmd.Flags().String("resource-group", "", "The Azure Container Registry name used to push the image.")
	AKSCheckAcrCmd.MarkFlagRequired("acr")
	AKSCheckAcrCmd.MarkFlagRequired("name")
	AKSCheckAcrCmd.MarkFlagRequired("resource-group")

	AKSGetUpgradesCmd.Flags().StringP("resource-group", "", "g", "")
	AKSGetUpgradesCmd.Flags().String("cluster-name", "", "The cluster name.")
	AKSGetUpgradesCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSGetUpgradesCmd.MarkFlagRequired("resource-group")
	AKSGetUpgradesCmd.MarkFlagRequired("cluster-name")

	AKSGetVersionsCmd.Flags().String("location", "l", "Location")
	AKSGetVersionsCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSGetVersionsCmd.MarkFlagRequired("location")
	AKSGetVersionsCmd.MarkFlagRequired("subscription")

	AKSNodepoolGetUpgradesCmd.Flags().String("nodepool-name", "", "The name of the node pool.")
	AKSNodepoolGetUpgradesCmd.Flags().StringP("resource-group", "", "g", "")
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

	AKSConnectedK8sCmd.PersistentFlags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSConnectedK8sCmd.PersistentFlags().StringP("name", "n", "", "Name of the managed cluster.")
	AKSConnectedDisableFeaturesCmd.Flags().StringSliceVarP(&slice, "features", "s", []string{}, "")
	AKSConnectedDisableFeaturesCmd.Flags().StringP("resource-group", "g", "", "Name of resource group.")

	AKSk8sConfiguration.PersistentFlags().StringP("resource-group", "g", "", "Name of resource group.")
	AKSk8sConfiguration.PersistentFlags().StringP("name", "n", "", "The name of the Kubernetes configuration.")
	AKSk8sConfiguration.PersistentFlags().StringP("cluster-name", "c", "", "Name of the managed cluster.")
	AKSk8sConfiguration.PersistentFlags().StringP("cluster-type", "", "", "Specifies an Arc cluster or AKS-managed cluster.")
	AKSConfigurationCreate.Flags().StringP("repository-url", "u", "", "The URL of the source control repository.")
	AKSConfigurationCreate.Flags().String("scope", "", "Scope the operator to either 'namespace' or 'cluster'. Allowed values: cluster, namespace")

}
