package cmd

/*
import (
	"Hybrid_Cluster/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var updateClusterConfigInput eks.UpdateClusterConfigInput

// updateAddonCmd represents the updateAddon command
var updateClusterConfigCmd = &cobra.Command{
	Use:   "update-cluster-config",
	Short: "A brief description of your command",
	Long: `
	- update-cluster-config
		hybridctl update-cluster-config <clusterName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl update-cluster-config --help' to view all commands")
		} else {

			updateClusterConfigInput.Name = &args[0]

			jsonFileName, _ := cmd.Flags().GetString("resource-vpc-config")
			if jsonFileName != "" {
				var resourcesVpcConfig eks.VpcConfigRequest
				util.OpenAndReadJsonFile(jsonFileName, resourcesVpcConfig)
			}

			jsonFileName, _ = cmd.Flags().GetString("logging")
			if jsonFileName != "" {
				var logging eks.Logging
				util.OpenAndReadJsonFile(jsonFileName, logging)
			}

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")

			if clientRequestToken != "" {
				updateClusterConfigInput.ClientRequestToken = &clientRequestToken
			}

			updateClusterConfig(updateClusterConfigInput)
		}
	},
}

func updateClusterConfig(input eks.UpdateClusterConfigInput) {
	httpPostUrl := "http://localhost:8080/updateClusterConfig"
	var output eks.UpdateClusterConfigOutput
	util.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(updateClusterConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//updateClusterConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	updateClusterConfigCmd.Flags().StringP("resource-vpc-config", "", "", "enter resource-vpc-config jsonfile name")
	updateClusterConfigCmd.Flags().StringP("logging", "", "", "enter logging jsonfile name")
	updateClusterConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
}
*/
