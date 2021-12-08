package cmd

import (
	"github.com/spf13/cobra"
)

// createAddonCmd represents the createAddon command
var EksCmd = &cobra.Command{
	Use:   "eks",
	Short: "A brief description of your command",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(EksCmd)
	EksCmd.AddCommand(associateEncryptionConfigCmd)
	EksCmd.AddCommand(associateIdentityProviderConfigCmd)
	EksCmd.AddCommand(createAddonCmd)
	EksCmd.AddCommand(deleteAddonCmd)
	EksCmd.AddCommand(describeAddonCmd)
	EksCmd.AddCommand(describeAddonVersionsCmd)
	EksCmd.AddCommand(describeIdentityProviderConfigCmd)
	EksCmd.AddCommand(describeUpdateCmd)
	EksCmd.AddCommand(disassociateIdentityProviderConfigCmd)
	EksCmd.AddCommand(listAddonCmd)
	EksCmd.AddCommand(listIdentityProviderConfigsCmd)
	EksCmd.AddCommand(listTagsForResourceCmd)
	EksCmd.AddCommand(listUpdateCmd)
	EksCmd.AddCommand(tagResourceCmd)
	EksCmd.AddCommand(untagResourceCmd)
	EksCmd.AddCommand(updateAddonCmd)
	EksCmd.AddCommand(updateClusterConfigCmd)
	EksCmd.AddCommand(updateNodegroupConfigCmd)
	eksFlags()
}
