package cmd

// import (
// 	"fmt"

// 	cobrautil "Hybrid_Cluster/hybridctl/util"

// 	"github.com/aws/aws-sdk-go/service/eks"
// 	"github.com/spf13/cobra"
// )

// var associateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput
// var oidcRequest eks.OidcIdentityProviderConfigRequest

// // AssociateIdentityProvicerConfigCmd represents the AssociateIdentityProvicerConfig command
// var associateIdentityProviderConfigCmd = &cobra.Command{
// 	Use:   "associate-identity-provider-config",
// 	Short: "A brief description of your command",
// 	Long: `
// 	- associate-identity-provider-config
// 		hybridctl associate-identity-provider-config <clusterName> <oidc>

// 	- platform
// 		- eks (elastic kubernetes service)`,

// 	Run: func(cmd *cobra.Command, args []string) {
// 		// TODO: Work your own magic here

// 		if len(args) == 0 {
// 			fmt.Println("Run 'hybridctl associate-identity-provider-config --help' to view all commands")
// 		} else if args[0] == "" {
// 			fmt.Println("Run 'hybridctl associate-identity-provider-config --help' to view all commands")
// 		} else {
// 			associateIdentityProviderConfigInput.ClusterName = &args[0]

// 			// json parsing
// 			oidc, _ := cmd.Flags().GetString("oidc")
// 			cobrautil.OpenAndReadJsonFile(oidc, oidcRequest)
// 			associateIdentityProviderConfigInput.Oidc = &oidcRequest

// 			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
// 			if clientRequestToken != "" {
// 				associateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
// 			}

// 			tags, _ := cmd.Flags().GetString("tags")
// 			var tagsMap map[string]*string
// 			if tags != "" {
// 				cobrautil.OpenAndReadJsonFile(tags, &tagsMap)
// 				associateIdentityProviderConfigInput.Tags = tagsMap
// 			}
// 			AssociateIdentityProvicerConfig(associateIdentityProviderConfigInput)
// 		}
// 	},
// }

// func AssociateIdentityProvicerConfig(AssociateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput) {
// 	httpPostUrl := "http://localhost:8080/associateIdentityProviderConfig"
// 	var output eks.AssociateIdentityProviderConfigOutput
// 	cobrautil.GetJson(httpPostUrl, AssociateIdentityProviderConfigInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func init() {
// 	EksCmd.AddCommand(associateIdentityProviderConfigCmd)

// 	associateIdentityProviderConfigCmd.Flags().StringP("oidc", "", "", "enter your oidc Jsonfile name")
// 	associateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
// 	associateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter your tags Jsonfile name")
// }
