package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
		data, _ := ioutil.ReadFile(configFile)
		fmt.Println(string(data))

		json.Unmarshal([]byte(configFile), &config)

		fmt.Println(config)
		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
			ConfigFile:        config,
		}
		maintenanceconfigurationCreateOrUpdate(EksAPIParameter)
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

		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationDelete(EksAPIParameter)
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

		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationCreateOrUpdate(EksAPIParameter)
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

		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationList(EksAPIParameter)
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

		EksAPIParameter := util.EksAPIParameter{
			ResourceGroupName: resourceGroupName,
			ResourceName:      clusterName,
			ConfigName:        configName,
		}
		maintenanceconfigurationShow(EksAPIParameter)
	},
}
