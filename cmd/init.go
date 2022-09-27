package cmd

import "github.com/spf13/cobra"

// AKSCmd represents the AKS command
var AKSCmd = &cobra.Command{
	Use:   "aks",
	Short: "A brief description of your command",
	Long: ` 

	`,
}

// EKSCmd represents the EKS command
var EKSCmd = &cobra.Command{
	Use:   "eks",
	Short: "A brief description of your command",
	Long:  ``,
}

// GKECmd represents the GKE command
var GKECmd = &cobra.Command{
	Use:   "gke",
	Short: "A brief description of your command",
	Long: ` 

	`,
}

// AddCommand adds one or more commands to this parent command.
func init() {
	RootCmd.AddCommand(AKSCmd)
	AKSCmd.AddCommand(AKSAddonCmd)
	AKSAddonCmd.AddCommand(AKSDisableAddonsCmd)
	AKSAddonCmd.AddCommand(AKSEnableAddonsCmd)
	AKSAddonCmd.AddCommand(AKSListAddonsCmd)
	AKSAddonCmd.AddCommand(AKSListAddonsAvailableCmd)
	AKSAddonCmd.AddCommand(AKSShowAddonsCmd)
	AKSAddonCmd.AddCommand(AKSUpdateAddonsCmd)

	AKSCmd.AddCommand(AKSPodIdentityCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIAddCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIDeleteCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIListCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIExceptionCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionAddCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionDeleteCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionListCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionUpdateCmd)

	AKSCmd.AddCommand(StartCmd)
	AKSCmd.AddCommand(StopCmd)
	AKSCmd.AddCommand(RotateCertsCmd)
	AKSCmd.AddCommand(GetOSoptionsCmd)
	AKSCmd.AddCommand(MaintenanceconfigurationCmd)
	MaintenanceconfigurationCmd.AddCommand(MCAddCmd)
	MaintenanceconfigurationCmd.AddCommand(MCDeleteCmd)
	MaintenanceconfigurationCmd.AddCommand(MCUpdateCmd)
	MaintenanceconfigurationCmd.AddCommand(MCListCmd)
	MaintenanceconfigurationCmd.AddCommand(MCShowCmd)

	AKSCmd.AddCommand(AKSAppUpCmd)
	AKSCmd.AddCommand(AKSBrowseCmd)
	AKSCmd.AddCommand(AKSCheckAcrCmd)
	AKSCmd.AddCommand(AKSGetUpgradesCmd)
	AKSCmd.AddCommand(AKSGetVersionsCmd)
	AKSCmd.AddCommand(AKSKanalyzeCmd)
	AKSCmd.AddCommand(AKSKollectCmd)
	AKSCmd.AddCommand(Nodepool)
	Nodepool.AddCommand(AKSNodepoolGetUpgradesCmd)
	AKSCmd.AddCommand(AKSInstallCLICmd)

	AKSCmd.AddCommand(AKSConnectedK8sCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sConnectCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sDeleteCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sDisableFeaturesCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sEnableFeaturesCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sListCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sProxyCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sShowCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sUpdateCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sUpgradeCmd)

	AKSCmd.AddCommand(AKSk8sConfiguration)
	AKSk8sConfiguration.AddCommand(AKSConfigurationCreate)
	AKSk8sConfiguration.AddCommand(AKSConfigurationDelete)
	AKSk8sConfiguration.AddCommand(AKSConfigurationList)
	AKSk8sConfiguration.AddCommand(AKSConfigurationShow)
	aksFlags()

	// EKS
	RootCmd.AddCommand(EKSCmd)

	EKSCmd.AddCommand(EKSClusterCmd)
	EKSClusterCmd.AddCommand(EKSCreateClusterCmd)
	EKSClusterCmd.AddCommand(EKSDeleteClusterCmd)
	EKSClusterCmd.AddCommand(EKSDescribeClusterCmd)
	EKSClusterCmd.AddCommand(EKSListClusterCmd)
	EKSClusterCmd.AddCommand(EKSUpgradeClusterCmd)

	EKSCmd.AddCommand(EKSNodegroupCmd)
	EKSNodegroupCmd.AddCommand(EKSCreateNodegroupCmd)
	EKSNodegroupCmd.AddCommand(EKSDeleteNodegroupCmd)
	EKSNodegroupCmd.AddCommand(EKSDescribeNodegroupCmd)
	EKSNodegroupCmd.AddCommand(EKSListNodegroupCmd)

	EKSCmd.AddCommand(EKSAddonCmd)
	EKSAddonCmd.AddCommand(EKSCreateAddonCmd)
	EKSAddonCmd.AddCommand(EKSDeleteAddonCmd)
	EKSAddonCmd.AddCommand(EKSDescribeAddonCmd)
	EKSAddonCmd.AddCommand(EKSDescribeAddonVersionsCmd)
	EKSAddonCmd.AddCommand(EKSListAddonCmd)
	EKSAddonCmd.AddCommand(EKSUpdateAddonCmd)

	EKSCmd.AddCommand(EKSEncryptionConfigCmd)
	EKSEncryptionConfigCmd.AddCommand(EKSAssociateEncryptionConfigCmd)

	EKSCmd.AddCommand(EKSIdentityProviderConfigCmd)
	EKSIdentityProviderConfigCmd.AddCommand(EKSAssociateIdentityProviderConfigCmd)
	EKSIdentityProviderConfigCmd.AddCommand(EKSDisassociateIdentityProviderConfigCmd)
	EKSIdentityProviderConfigCmd.AddCommand(EKSDescribeIdentityProviderConfigCmd)
	EKSIdentityProviderConfigCmd.AddCommand(EKSListIdentityProviderConfigsCmd)

	EKSCmd.AddCommand(EKSResourceCmd)
	EKSResourceCmd.AddCommand(EKSTagResourceCmd)
	EKSResourceCmd.AddCommand(EKSUntagResourceCmd)
	EKSResourceCmd.AddCommand(EKSListTagsForResourceCmd)

	EKSCmd.AddCommand(EKSClusterConfigCmd)
	EKSClusterConfigCmd.AddCommand(EKSUpdateClusterConfigCmd)
	//EKSUpdateClusterConfigCmd.AddCommand(EKSUpdateClusterConfigCmd)

	EKSCmd.AddCommand(EKSNodegroupConfigCmd)
	EKSNodegroupConfigCmd.AddCommand(EKSUpdateNodegroupConfigCmd)

	EKSCmd.AddCommand(EKSDescribeUpdateCmd)
	EKSCmd.AddCommand(EKSListUpdateCmd)
	eksFlags()

	// GKE
	RootCmd.AddCommand(GKECmd)
	GKECmd.AddCommand(GKEInitCmd)
	GKECmd.AddCommand(GKEContainerCmd)
	GKEContainerCmd.AddCommand(GKEContainerImagesCmd)
	GKEContainerCmd.AddCommand(GKEContainerGetServerConfigCmd)
	GKEContainerCmd.AddCommand(GKEContainerNodePoolsCmd)
	GKEContainerNodePoolsCmd.AddCommand(GKENodePoolsRollbackCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesAddTagCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesDeleteCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesDescribeCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesListCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesListTagsCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesUnTagCmd)
	GKEContainerCmd.AddCommand(GKEContainerOperationsCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationDescribeCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationsListCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationsWaitCmd)
	GKECmd.AddCommand(GKEAuthCmd)
	GKEAuthCmd.AddCommand(GKEAuthConfigureDockerCmd)
	GKEAuthCmd.AddCommand(GKEAuthListCmd)
	GKEAuthCmd.AddCommand(GKEAuthLoginCmd)
	GKEAuthCmd.AddCommand(GKEAuthRevokeCmd)
	GKECmd.AddCommand(GKEDockerCmd)
	GKECmd.AddCommand(GKESourceCmd)
	GKESourceCmd.AddCommand(GKESourceProjectConfigsCmd)
	GKESourceProjectConfigsCmd.AddCommand(GKEProjectConfigsUpdateCmd)
	GKESourceProjectConfigsCmd.AddCommand(GKEProjectConfigsDescribeCmd)
	GKECmd.AddCommand(GKEConfigCmd)
	GKEConfigCmd.AddCommand(GKEConfigSetCmd)
	gkeFlags()
}
