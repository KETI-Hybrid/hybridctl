package cmd

import (
	"Hybrid_Cluster/hybridctl/util"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `hybridctl aks start --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksStart(EksAPIParameter)

	},
}

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long:  `hybridctl aks stop --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksStop(EksAPIParameter)

	},
}

var RotateCertsCmd = &cobra.Command{
	Use:   "rotate-certs",
	Short: "A brief description of your command",
	Long:  `hybridctl aks rotate-certs --name <clusterName> --resource-group <ResourceGroupName>`,
	Run: func(cmd *cobra.Command, args []string) {

		resourceGroupName, _ := cmd.Flags().GetString("resource-group")
		clusterName, _ := cmd.Flags().GetString("name")
		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
		}
		aksRotateCerts(EksAPIParameter)
	},
}

var GetOSoptionsCmd = &cobra.Command{
	Use:   "get-os-options",
	Short: "A brief description of your command",
	Long:  `hybridctl aks get-os-options --location`,
	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")
		EksAPIParameter := util.EksAPIParameter{
			Location: location,
		}
		aksGetOSoptions(EksAPIParameter)
	},
}
