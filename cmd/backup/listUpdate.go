package cmd

/*
import (
	"Hybrid_Cloud/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var listUpdateInput eks.ListUpdatesInput

// listAddonCmd represents the listAddon command
var listUpdateCmd = &cobra.Command{
	Use:   "list-update",
	Short: "A brief description of your command",
	Long: `
	- list-update
		hybridctl list-update <clusterName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl list-update --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl list-update --help' to view all commands")
		} else {
			listUpdateInput.Name = &args[0]
			nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")
			addonName, _ := cmd.Flags().GetString("addon-name")
			maxResults, _ := cmd.Flags().GetInt64("max-result")
			nextToken, _ := cmd.Flags().GetString("next-token")
			if nodegroupName != "" {
				listUpdateInput.NodegroupName = &nodegroupName
			}
			if addonName != "" {
				listUpdateInput.AddonName = &addonName
			}
			if maxResults != 0 {
				listAddonInput.MaxResults = &maxResults
			}
			if nextToken != "" {
				listAddonInput.NextToken = &nextToken
			}
			listUpdate(listUpdateInput)
		}
	},
}

func listUpdate(listUpdateInput eks.ListUpdatesInput) {
	httpPostUrl := "http://localhost:8080/listUpdate"
	var output eks.ListUpdatesOutput
	util.GetJson(httpPostUrl, listUpdateInput, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(listUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//listAddonCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	listUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	listUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")
	listUpdateCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	listUpdateCmd.Flags().StringP("next-token", "", "", "enter next token")
}
*/
