package cmd

/*
import (
	"Hybrid_Cloud/hybridctl/util"
	cobrautil "Hybrid_Cloud/hybridctl/util"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

// AssociateIdentityProvicerConfigCmd represents the AssociateIdentityProvicerConfig command
var tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "A brief description of your command",
	Long: `
	- tage-resource
		hybridctl tag-resource --resource-arn <value> --tags <jsonfile>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		var tagResourceInput eks.TagResourceInput
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		tagResourceInput.ResourceArn = &resourceArn

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap map[string]*string
		if tags != "" {
			cobrautil.OpenAndReadJsonFile(tags, &tagsMap)
			tagResourceInput.Tags = tagsMap
		}
		TagResource(tagResourceInput)
	},
}

func TagResource(input eks.TagResourceInput) {
	httpPostUrl := "http://localhost:8080/tagResource"
	var output eks.TagResourceOutput
	util.GetJson(httpPostUrl, input, &output)
}

func init() {
	EksCmd.AddCommand(tagResourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// AssociateIdentityProvicerConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	// associateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter tags")
	tagResourceCmd.Flags().StringP("tags", "t", "", "enter your tags Jsonfile name")
	tagResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")
	tagResourceCmd.MarkPersistentFlagRequired("tags")
	tagResourceCmd.MarkPersistentFlagRequired("resource-arn")
}
*/
