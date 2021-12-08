package cmd

/*
import (
	"Hybrid_Cluster/hybridctl/util"
	"strings"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

// AssociateIdentityProvicerConfigCmd represents the AssociateIdentityProvicerConfig command
var untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "A brief description of your command",
	Long: `
	- untage-resource
		hybridctl untag-resource --resource-arn <value> --tag-keys <key,key>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		var untagResourceInput eks.UntagResourceInput
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		untagResourceInput.ResourceArn = &resourceArn

		keys, _ := cmd.Flags().GetString("tag-keys")
		slice := strings.Split(keys, ",")
		keyList := []*string{}
		for i := 0; i < len(slice); i++ {
			s := append(keyList, &slice[i])
			keyList = s
		}

		untagResourceInput.TagKeys = keyList

		unTagResource(untagResourceInput)
	},
}

func unTagResource(input eks.UntagResourceInput) {
	httpPostUrl := "http://localhost:8080/untagResource"
	var output eks.UntagResourceOutput
	util.GetJson(httpPostUrl, input, &output)
	// fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(untagResourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// AssociateIdentityProvicerConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	// associateIdentityProviderConfigCmd.Flags().StringP("untags", "", "", "enter untags")
	untagResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")
	untagResourceCmd.Flags().StringP("tag-keys", "t", "", "enter your tag-keys list")
	untagResourceCmd.MarkPersistentFlagRequired("tag-keys")
	untagResourceCmd.MarkPersistentFlagRequired("resource-arn")
}
*/
