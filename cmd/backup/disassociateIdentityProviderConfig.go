package cmd

/*
import (
	"fmt"

	"Hybrid_Cloud/hybridctl/util"
	cobrautil "Hybrid_Cloud/hybridctl/util"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var disassociateIdentityProviderConfigInput eks.DisassociateIdentityProviderConfigInput

// disassociateIdentityProvicerConfigCmd represents the disassociateIdentityProvicerConfig command
var disassociateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "disassociate-identity-provider-config",
	Short: "A brief description of your command",
	Long: `
	- disassociate-identity-provider-config
		hybridctl disassociate-identity-provider-config <clusterName> <oidc>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl disassociate-identity-provider-config --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl disassociate-identity-provider-config --help' to view all commands")
		} else {
			disassociateIdentityProviderConfigInput.ClusterName = &args[0]

			// json parsing
			var IdentityProviderConfig eks.IdentityProviderConfig
			jsonFileName, _ := cmd.Flags().GetString("identity-provider-config")
			util.OpenAndReadJsonFile(jsonFileName, IdentityProviderConfig)
			disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
			if clientRequestToken != "" {
				disassociateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
			}

			disassociateIdentityProvicerConfig(disassociateIdentityProviderConfigInput)
		}
	},
}

func disassociateIdentityProvicerConfig(input eks.DisassociateIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/disassociateIdentityProviderConfig"
	var output eks.DisassociateIdentityProviderConfigOutput
	cobrautil.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(disassociateIdentityProviderConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// disassociateIdentityProvicerConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	// disassociateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter tags")
	disassociateIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")
	disassociateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
}
*/
