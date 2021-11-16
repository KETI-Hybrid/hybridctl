package cmd

func aksFlags() {
	aksCmd.PersistentFlags().StringP("resource-group", "g", "", "Name of resource group.")
	aksCmd.PersistentFlags().StringP("name", "n", "", "Name of the managed cluster.")

	//addon
	AddonCmd.PersistentFlags().StringP("addon", "a", "", "Specify the Kubernetes addon")

	AKSDisableAddonsCmd.MarkFlagRequired("resource-group")
	AKSDisableAddonsCmd.MarkFlagRequired("name")
	AKSDisableAddonsCmd.MarkFlagRequired("addon")

	AKSEnableAddonsCmd.MarkFlagRequired("resource-group")
	AKSEnableAddonsCmd.MarkFlagRequired("name")
	AKSEnableAddonsCmd.MarkFlagRequired("addon")

	// AKSListAddonsCmd.MarkFlagRequired("resource-group")
	AKSListAddonsCmd.MarkFlagRequired("name")

	AKSShowAddonsCmd.MarkFlagRequired("resource-group")
	AKSShowAddonsCmd.MarkFlagRequired("name")
	AKSShowAddonsCmd.MarkFlagRequired("addon")

	AKSUpdateAddonsCmd.MarkFlagRequired("resource-group")
	AKSUpdateAddonsCmd.MarkFlagRequired("name")
	AKSUpdateAddonsCmd.MarkFlagRequired("addon")

	//Pod-Identity

	AKSPodIdentityCmd.PersistentFlags().String("cluster-name", "", "The cluster name.")
	AKSPodIdentityCmd.PersistentFlags().String("namespace", "", "The pod identity namespace.")
	AKSPodIdentityCmd.PersistentFlags().StringP("name", "", "n", "The pod identity name. Generate if not specified.")

	AKSPIAddCmd.MarkPersistentFlagRequired("resource-group")
	AKSPIAddCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIAddCmd.MarkPersistentFlagRequired("namespace")
	AKSPIAddCmd.MarkPersistentFlagRequired("name")
	AKSPIAddCmd.Flags().String("identity-resource-id", "", "Resource id of the identity to use.")
	AKSPIAddCmd.Flags().String("binding-selector", "", "Optional binding selector to use.")

	AKSPIDeleteCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("namespace")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("name")
	AKSPIDeleteCmd.MarkPersistentFlagRequired("resource-group")

	AKSPIListCmd.MarkPersistentFlagRequired("cluster-name")
	AKSPIListCmd.MarkPersistentFlagRequired("resource-group")

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

	StopCmd.MarkPersistentFlagRequired("resource-group")
	StopCmd.MarkPersistentFlagRequired("name")

	GetOSoptionsCmd.PersistentFlags().StringP("location", "l", "", "location")
	GetOSoptionsCmd.MarkPersistentFlagRequired("location")

	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("resource-group")
	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("name")

	MCAddCmd.Flags().StringP("config-name", "c", "", "configname")
	MCAddCmd.MarkFlagRequired("config-name")
	MCAddCmd.Flags().StringP("config-file", "", "", "configfile")
	MCAddCmd.MarkFlagRequired("config-file")

	MCDeleteCmd.Flags().StringP("config-name", "c", "", "configname")
	MCDeleteCmd.MarkFlagRequired("config-name")

	MCUpdateCmd.Flags().StringP("config-name", "c", "", "configname")
	MCUpdateCmd.MarkFlagRequired("config-name")
	MCUpdateCmd.Flags().StringP("config-file", "", "", "configfile")
	MCUpdateCmd.MarkFlagRequired("config-file")

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

	AKSCheckAcrCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSCheckAcrCmd.Flags().String("acr", "", "The Azure Container Registry name used to push the image.")

	AKSGetUpgradesCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")

	AKSGetVersionsCmd.Flags().String("location", "l", "Location")
	AKSGetVersionsCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")

	AKSNodepoolGetUpgradesCmd.Flags().String("nodepool-name", "", "The name of the node pool.")
	AKSNodepoolGetUpgradesCmd.Flags().String("subscription", "", "The name or ID of the subscription. ")
	AKSNodepoolGetUpgradesCmd.Flags().String("cluster-name", "", "The cluster name.")

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
