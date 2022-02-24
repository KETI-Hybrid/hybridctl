package cmd

/*
import (
	"Hybrid_Cloud/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var updateAddonInput eks.UpdateAddonInput

// updateAddonCmd represents the updateAddon command
var updateAddonCmd = &cobra.Command{
	Use:   "update-addon",
	Short: "A brief description of your command",
	Long: `
	- update-addon
		hybridctl update-addon <clusterName> <addonName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 || len(args) == 1 {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else if args[1] == "" {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else {
			addonVersion, _ := cmd.Flags().GetString("addon-version")
			serviceAccountRoleArn, _ := cmd.Flags().GetString("service-account-role-arn")
			resolveConflicts, _ := cmd.Flags().GetString("resolve-conflicts")
			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
			updateAddonInput.ClusterName = &args[0]
			updateAddonInput.AddonName = &args[1]
			if addonVersion != "" {
				updateAddonInput.AddonVersion = &addonVersion
				fmt.Printf(addonVersion)
			}
			if serviceAccountRoleArn != "" {
				updateAddonInput.ServiceAccountRoleArn = &serviceAccountRoleArn
			}
			if resolveConflicts != "" {
				updateAddonInput.ResolveConflicts = &resolveConflicts
			}
			if clientRequestToken != "" {
				updateAddonInput.ClientRequestToken = &clientRequestToken
			}
			updateAddon(updateAddonInput)
		}
	},
}

func updateAddon(updateAddonInput eks.UpdateAddonInput) {
	httpPostUrl := "http://localhost:8080/updateAddon"
	var output eks.UpdateAddonOutput
	util.GetJson(httpPostUrl, updateAddonInput, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(updateAddonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//updateAddonCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	updateAddonCmd.Flags().StringP("addon-version", "", "", "enter addon version")
	updateAddonCmd.Flags().StringP("service-account-role-arn", "", "", "enter service account rolearn")
	updateAddonCmd.Flags().StringP("resolve-conflicts", "", "", "enter addon version")
	updateAddonCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
}
*/
