package cmd

/*
import (
	"Hybrid_Cloud/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var updateNodegroupConfigInput eks.UpdateNodegroupConfigInput

// updateAddonCmd represents the updateAddon command
var updateNodegroupConfigCmd = &cobra.Command{
	Use:   "update-Nodegroup-config",
	Short: "A brief description of your command",
	Long: `
	- update-Nodegroup-config
		hybridctl update-Nodegroup-config <NodegroupName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl update-Nodegroup-config --help' to view all commands")
		} else {

			updateNodegroupConfigInput.ClusterName = &args[0]
			updateNodegroupConfigInput.NodegroupName = &args[1]

			jsonFileName, _ := cmd.Flags().GetString("labels")
			if jsonFileName != "" {
				var labels eks.UpdateLabelsPayload
				util.OpenAndReadJsonFile(jsonFileName, labels)
			}

			jsonFileName, _ = cmd.Flags().GetString("taints")
			if jsonFileName != "" {
				var taints eks.UpdateLabelsPayload
				util.OpenAndReadJsonFile(jsonFileName, taints)
			}

			jsonFileName, _ = cmd.Flags().GetString("scaling-config")
			if jsonFileName != "" {
				var scalingConfig eks.NodegroupScalingConfig
				util.OpenAndReadJsonFile(jsonFileName, scalingConfig)
			}

			jsonFileName, _ = cmd.Flags().GetString("update-config")
			if jsonFileName != "" {
				var updateConfig eks.NodegroupUpdateConfig
				util.OpenAndReadJsonFile(jsonFileName, updateConfig)
			}

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")

			if clientRequestToken != "" {
				updateNodegroupConfigInput.ClientRequestToken = &clientRequestToken
			}

			updateNodegroupConfig(updateNodegroupConfigInput)
		}
	},
}

func updateNodegroupConfig(input eks.UpdateNodegroupConfigInput) {
	httpPostUrl := "http://localhost:8080/updateNodegroupConfig"
	var output eks.UpdateNodegroupConfigOutput
	util.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(updateNodegroupConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//updateNodegroupConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	updateNodegroupConfigCmd.Flags().StringP("labels", "", "", "enter labels jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("taints", "", "", "enter taints jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("scaling-config", "", "", "enter resource-vpc-config jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("update-config", "", "", "enter logging jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
}
*/
