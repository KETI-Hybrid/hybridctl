package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"

	cmdpb "Hybrid_Cluster/protos/v1/cmd"

	"github.com/spf13/cobra"
)

//addon
var AddonCmd = &cobra.Command{
	Use:   "addon",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
}

var AKSDisableAddonsCmd = &cobra.Command{
	Use:   "disable",
	Short: "A brief description of your command",
	Long:  `hybridctl aks disable-addons`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		addon, _ := cmd.Flags().GetString("addon")

		AKSAddon := util.AKSAddon{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Addon:             addon,
		}
		addonDisable(AKSAddon)
	},
}

var AKSEnableAddonsCmd = &cobra.Command{
	Use:   "enable",
	Short: "A brief description of your command",
	Long:  `hybridctl aks addon enable`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		addon, _ := cmd.Flags().GetString("addon")

		AKSAddon := util.AKSAddon{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Addon:             addon,
		}
		addonEnable(AKSAddon)
	},
}

var AKSListAddonsCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hybridctl aks addon list`,
	Run: func(cmd *cobra.Command, args []string) {

		// resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAddon := cmdpb.AKSAddon{
			// ResourceGroupName: resourceGroupName,
			ClusterName: clusterName,
		}
		addonList(AKSAddon)
	},
}

var AKSListAddonsAvailableCmd = &cobra.Command{
	Use:   "list-available",
	Short: "A brief description of your command",
	Long:  `hybridctl aks addon list`,
	Run: func(cmd *cobra.Command, args []string) {
		addonListAvailable()
	},
}

var AKSShowAddonsCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long:  `hybridctl aks addon enable`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		addon, _ := cmd.Flags().GetString("addon")

		AKSAddon := util.AKSAddon{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Addon:             addon,
		}
		addonShow(AKSAddon)
	},
}

var AKSUpdateAddonsCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `hybridctl aks addon enable`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		addon, _ := cmd.Flags().GetString("addon")

		AKSAddon := util.AKSAddon{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Addon:             addon,
		}
		addonUpdate(AKSAddon)
	},
}

//pod-identity
var AKSPodIdentityCmd = &cobra.Command{
	Use:   "pod-identity",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity`,
}

var AKSPIAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity add`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		identityResourceID, _ := cmd.Flags().GetString("identity-resource-id")
		namespace, _ := cmd.Flags().GetString("namespace")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		podIdentityName, _ := cmd.Flags().GetString("name")
		bindingSelector, _ := cmd.Flags().GetString("addon")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName:  resourceGroupName,
			ClusterName:        clusterName,
			Namespace:          namespace,
			IdentityResourceID: identityResourceID,
		}
		if podIdentityName != "" {
			AKSPodIdentity.Name = podIdentityName
		}
		if bindingSelector != "" {
			AKSPodIdentity.BindingSelector = bindingSelector
		}
		podIdentityAdd(AKSPodIdentity)
	},
}

var AKSPIDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity delete`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		namespace, _ := cmd.Flags().GetString("namespace")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		podIdentityName, _ := cmd.Flags().GetString("name")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Namespace:         namespace,
			Name:              podIdentityName,
		}
		podIdentityDelete(AKSPodIdentity)
	},
}

var AKSPIListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity list`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
		}
		podIdentityList(AKSPodIdentity)
	},
}

var AKSPIExceptionCmd = &cobra.Command{
	Use:   "exception",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity exception`,
}

var AKSPIExceptionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity exception add`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		namespace, _ := cmd.Flags().GetString("namespace")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		podIdentityName, _ := cmd.Flags().GetString("name")
		podLabels, _ := cmd.Flags().GetString("podLabels")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Namespace:         namespace,
			PodLabels:         podLabels,
		}
		if podIdentityName != "" {
			AKSPodIdentity.Name = podIdentityName
		}
		podIdentityExceptionAdd(AKSPodIdentity)
	},
}

var AKSPIExceptionDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity exception delete`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		namespace, _ := cmd.Flags().GetString("namespace")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		podIdentityName, _ := cmd.Flags().GetString("name")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Namespace:         namespace,
			Name:              podIdentityName,
		}
		podIdentityExceptionDelete(AKSPodIdentity)
	},
}

var AKSPIExceptionListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity exception list`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
		}
		podIdentityExceptionList(AKSPodIdentity)
	},
}

var AKSPIExceptionUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `hybridctl aks pod-identity exception update`,
	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		namespace, _ := cmd.Flags().GetString("namespace")
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		podIdentityName, _ := cmd.Flags().GetString("name")
		podLabels, _ := cmd.Flags().GetString("podLabels")

		AKSPodIdentity := util.AKSPodIdentity{
			ResourceGroupName: resourceGroupName,
			ClusterName:       clusterName,
			Namespace:         namespace,
			PodLabels:         podLabels,
			Name:              podIdentityName,
		}
		podIdentityExceptionUpdate(AKSPodIdentity)
	},
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `hybridctl aks start --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksStart(EKSAPIParameter)

	},
}

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long:  `hybridctl aks stop --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksStop(EKSAPIParameter)

	},
}

var RotateCertsCmd = &cobra.Command{
	Use:   "rotate-certs",
	Short: "A brief description of your command",
	Long:  `hybridctl aks rotate-certs --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksRotateCerts(EKSAPIParameter)
	},
}

var GetOSoptionsCmd = &cobra.Command{
	Use:   "get-os-options",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")
		EKSAPIParameter := util.EKSAPIParameter{
			Location: location,
		}
		aksGetOSoptions(EKSAPIParameter)
	},
}

var MaintenanceconfigurationCmd = &cobra.Command{
	Use:   "maintenanceconfiguration",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
}

var MCAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		var config util.Config
		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		configName, _ := cmd.Flags().GetString("config-name")
		configFile, _ := cmd.Flags().GetString("config-file")
		// fmt.Println(configFile)
		// data, _ := ioutil.ReadFile(configFile)
		// fmt.Println(string(data))

		byteValue := util.OpenAndReadJsonFile(configFile)
		json.Unmarshal(byteValue, &config)

		fmt.Println(config)
		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
			ConfigFile:        config,
		}
		maintenanceconfigurationCreateOrUpdate(EKSAPIParameter)
	},
}

var MCDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		configName, _ := cmd.Flags().GetString("configname")
		if configName == "" {
			configName = "default"
		}

		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationDelete(EKSAPIParameter)
	},
}

var MCUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		configName, _ := cmd.Flags().GetString("configname")
		if configName == "" {
			configName = "default"
		}

		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationCreateOrUpdate(EKSAPIParameter)
	},
}

var MCListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		configName, _ := cmd.Flags().GetString("configname")
		if configName == "" {
			configName = "default"
		}

		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationList(EKSAPIParameter)
	},
}

var MCShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		configName, _ := cmd.Flags().GetString("configname")

		EKSAPIParameter := util.EKSAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationShow(EKSAPIParameter)
	},
}

// TODO: Github Path 입력 필수
//       사전에 Deploy-to-azure 다운받아야함.
var AKSAppUpCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your command",
	Long:  `hybridctl aks app up`,
	Run: func(cmd *cobra.Command, args []string) {

		AKSAPIParameter := util.AKSAPIParameter{}
		p, _ := cmd.Flags().GetString("acr")
		if p != "" {
			AKSAPIParameter.Acr = p
		}

		p, _ = cmd.Flags().GetString("aks-cluster")
		if p != "" {
			AKSAPIParameter.AksCluster = p
		}

		p, _ = cmd.Flags().GetString("branch-name")
		if p != "" {
			AKSAPIParameter.BranchName = p
		}

		p, _ = cmd.Flags().GetString("do-not-wait")
		if p != "" {
			AKSAPIParameter.DoNotWait = p
		}

		p, _ = cmd.Flags().GetString("port")
		if p != "" {
			AKSAPIParameter.Port = p
		}

		p, _ = cmd.Flags().GetString("repository")
		if p != "" {
			AKSAPIParameter.Repository = p
		}
		fmt.Println(AKSAPIParameter)
		appUp(AKSAPIParameter)
	},
}

// TODO: disable-browser Boolean 처리
var AKSBrowseCmd = &cobra.Command{
	Use:   "browse",
	Short: "A brief description of your command",
	Long:  `hybridctl aks browse`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		p, _ := cmd.Flags().GetBool("disable-browser")
		fmt.Printf("%t", p)
		if !p {
			AKSAPIParameter.DisableBrowser = p
		}

		t, _ := cmd.Flags().GetString("listen-address")
		if t != "" {
			AKSAPIParameter.ListenAddress = t
		}

		t, _ = cmd.Flags().GetString("listen-port")
		if t != "" {
			AKSAPIParameter.ListenPort = t
		}

		t, _ = cmd.Flags().GetString("subscription")
		if t != "" {
			AKSAPIParameter.Subscription = t
		}
		fmt.Println(AKSAPIParameter)
		browse(AKSAPIParameter)
	},
}

var AKSCheckAcrCmd = &cobra.Command{
	Use:   "check-acr",
	Short: "A brief description of your command",
	Long:  `hybridctl aks check-acr`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		acr, _ := cmd.Flags().GetString("acr")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
			Acr:           acr,
		}
		p, _ := cmd.Flags().GetString("subscription")
		if p != "" {
			AKSAPIParameter.Subscription = p
		}
		fmt.Println(AKSAPIParameter)
		checkAcr(AKSAPIParameter)
	},
}

// TODO: get-upgrades 명령어 처리 이전에 az login되어 있어야함
var AKSGetUpgradesCmd = &cobra.Command{
	Use:   "get-upgrades",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-upgrades`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		p, _ := cmd.Flags().GetString("subscription")
		if p != "" {
			AKSAPIParameter.Subscription = p
		}
		fmt.Println(AKSAPIParameter)
		getUpgrades(AKSAPIParameter)
	},
}

var AKSGetVersionsCmd = &cobra.Command{
	Use:   "get-versions",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-versions`,
	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")
		AKSAPIParameter := util.AKSAPIParameter{
			Location: location,
		}
		p, _ := cmd.Flags().GetString("subscription")
		if p != "" {
			AKSAPIParameter.Subscription = p
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "getVersions")
	},
}

var AKSKanalyzeCmd = &cobra.Command{
	Use:   "kanalyze",
	Short: "A brief description of your command",
	Long:  `hybridctl aks kanalyze`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "kanalyze")
	},
}

var Nodepool = &cobra.Command{
	Use:   "nodepool",
	Short: "A brief description of your command",
	Long:  `hybridctl aks nodepool`,
}

var AKSNodepoolGetUpgradesCmd = &cobra.Command{
	Use:   "get-upgrades",
	Short: "A brief description of your command",
	Long:  `hybridctl aks nodepool get-upgrades`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		nodepoolName, _ := cmd.Flags().GetString("nodepool-name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
			NodepoolName:  nodepoolName,
		}
		p, _ := cmd.Flags().GetString("subscription")
		if p != "" {
			AKSAPIParameter.Subscription = p
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "nodepoolGetUpgrades")
	},
}

var AKSInstallCLICmd = &cobra.Command{
	Use:   "install-cli",
	Short: "A brief description of your command",
	Long:  `hybridctl aks install-cli`,
	Run: func(cmd *cobra.Command, args []string) {

		AKSAPIParameter := util.AKSInstallCLI{}
		p, _ := cmd.Flags().GetString("base-src-url")
		if p != "" {
			AKSAPIParameter.BaseSrcURL = p
		}
		p, _ = cmd.Flags().GetString("client-version")
		if p != "" {
			AKSAPIParameter.ClientVersion = p
		}
		p, _ = cmd.Flags().GetString("install-location")
		if p != "" {
			AKSAPIParameter.InstallLocation = p
		}
		p, _ = cmd.Flags().GetString("kubelogin-base-src-url")
		if p != "" {
			AKSAPIParameter.KubeloginBaseSrcURL = p
		}
		p, _ = cmd.Flags().GetString("kubelogin-install-location")
		if p != "" {
			AKSAPIParameter.KubeloginInstallLocation = p
		}
		p, _ = cmd.Flags().GetString("kubelogin-version")
		if p != "" {
			AKSAPIParameter.KubeloginVersion = p
		}
		p, _ = cmd.Flags().GetString("subscription")
		if p != "" {
			AKSAPIParameter.Subscription = p
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequestCLI(AKSAPIParameter, "installCLI")
	},
}

var AKSConnectedK8sCmd = &cobra.Command{
	Use:   "connectedk8s",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s`,
}

var AKSConnectedConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s connect`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedConnect")
	},
}

var AKSConnectedDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s delete`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedk8sDelete")
	},
}

var slice []string
var AKSConnectedDisableFeaturesCmd = &cobra.Command{
	Use:   "disable-features",
	Short: "A brief description of your command",
	Long:  `hybridctl connetedk8s disable-features`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		features, _ := cmd.Flags().GetStringSlice("features")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
			Features:      features,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedDisableFeatures")
	},
}

var AKSConnectedEnableFeaturesCmd = &cobra.Command{
	Use:   "enable-features",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s enable-features`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedEnableFeatures")
	},
}

var AKSConnectedListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s list`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		AKSAPIParameter := util.AKSAPIParameter{
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedList")
	},
}

var AKSConnectedProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s proxy`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedProxy")
	},
}

var AKSConnectedShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s show`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedShow")
	},
}

var AKSConnectedUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s update`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedUpdate")
	},
}

var AKSConnectedUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "A brief description of your command",
	Long:  `hybridctl aks connectedk8s upgrade`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSAPIParameter{
			Name:          clusterName,
			ResourceGroup: resourceGroupName,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequest(AKSAPIParameter, "connectedUpgrade")
	},
}

var AKSk8sConfiguration = &cobra.Command{
	Use:   "k8sconfiguration",
	Short: "A brief description of your command",
	Long:  "hybridctl aks k8sconfiguration",
}

var AKSConfigurationCreate = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `hybridctl aks k8sconfiguration create`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		clusterType, _ := cmd.Flags().GetString("cluster-type")
		if clusterType != "connectedClusters" && clusterType != "managedClusters" {
			fmt.Println("Allowed values: connectedClusters, managedClusters")
			return
		}
		name, _ := cmd.Flags().GetString("name")
		repositoryURL, _ := cmd.Flags().GetString("repository-url")
		scope, _ := cmd.Flags().GetString("scope")
		if scope != "cluster" && scope != "namespace" {
			fmt.Println("Scope the operator to either 'namespace' or 'cluster'.")
			return
		}
		AKSAPIParameter := util.AKSk8sConfiguration{
			Name:          name,
			ResourceGroup: resourceGroupName,
			ClusterName:   clusterName,
			ClusterType:   clusterType,
			RepositoryURL: repositoryURL,
			Scope:         scope,
		}

		fmt.Println(AKSAPIParameter)
		HTTPPostRequestConfig(AKSAPIParameter, "configurationCreate")
	},
}

var AKSConfigurationDelete = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `hybridctl aks k8sconfiguration delete`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		clusterType, _ := cmd.Flags().GetString("cluster-type")
		if clusterType != "connectedClusters" && clusterType != "managedClusters" {
			fmt.Println("Allowed values: connectedClusters, managedClusters")
			return
		}
		name, _ := cmd.Flags().GetString("name")
		AKSAPIParameter := util.AKSk8sConfiguration{
			Name:          name,
			ResourceGroup: resourceGroupName,
			ClusterName:   clusterName,
			ClusterType:   clusterType,
		}
		fmt.Println(AKSAPIParameter)
		HTTPPostRequestConfig(AKSAPIParameter, "configurationDelete")
	},
}
