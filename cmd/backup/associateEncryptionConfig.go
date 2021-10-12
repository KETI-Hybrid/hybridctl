package cmd

// import (
// 	"Hybrid_Cluster/hybridctl/util"
// 	cobrautil "Hybrid_Cluster/hybridctl/util"
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/service/eks"
// 	"github.com/spf13/cobra"
// )

// var associateEncryptionConfigInput eks.AssociateEncryptionConfigInput

// // AssociateIdentityProvicerConfigCmd represents the AssociateIdentityProvicerConfig command
// var associateEncryptionConfigCmd = &cobra.Command{
// 	Use:   "associate-encryption-config",
// 	Short: "A brief description of your command",
// 	Long: `
// 	- associate-encryption-config
// 		hybridctl associate-encryption-config <clusterName> --encryption-config <jsonfile>

// 	- platform
// 		- eks (elastic kubernetes service)`,

// 	Run: func(cmd *cobra.Command, args []string) {
// 		// TODO: Work your own magic here

// 		if len(args) == 0 {
// 			fmt.Println("Run 'hybridctl associate-encryption-config --help' to view all commands")
// 		} else if args[0] == "" {
// 			fmt.Println("Run 'hybridctl associate-encryption-config --help' to view all commands")
// 		} else {
// 			associateEncryptionConfigInput.ClusterName = &args[0]

// 			// json parsing
// 			jsonFileName, _ := cmd.Flags().GetString("encryption-config")
// 			var encryptionConfig []*eks.EncryptionConfig
// 			util.UnmarshalJsonFile(jsonFileName, encryptionConfig)
// 			associateEncryptionConfigInput.EncryptionConfig = encryptionConfig

// 			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
// 			if clientRequestToken != "" {
// 				associateEncryptionConfigInput.ClientRequestToken = &clientRequestToken
// 			}
// 			AssociateEncryptionConfig(associateEncryptionConfigInput)
// 		}
// 	},
// }

// func AssociateEncryptionConfig(AssociateEncryptionConfigInput eks.AssociateEncryptionConfigInput) {
// 	httpPostUrl := "http://localhost:8080/associateEncryptionConfig"
// 	var output eks.AssociateEncryptionConfigOutput
// 	cobrautil.GetJson(httpPostUrl, AssociateEncryptionConfigInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func init() {
// 	EksCmd.AddCommand(associateEncryptionConfigCmd)
// 	associateEncryptionConfigCmd.Flags().StringP("encryption-config", "", "", "enter your encryption-config Jsonfile name")
// 	associateEncryptionConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
// }
