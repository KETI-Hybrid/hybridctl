package cmd

/*
var describeIdentityProviderConfigInput eks.DescribeIdentityProviderConfigInput

// describeIdentityProvicerConfigCmd represents the describeIdentityProvicerConfig command
var describeIdentityProviderConfigCmd = &cobra.Command{
	Use:   "describe-identity-provider-config",
	Short: "A brief description of your command",
	Long: `
	- describe-identity-provider-config
		hybridctl describe-identity-provider-config <clusterName> <oidc>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl describe-identity-provider-config --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl describe-identity-provider-config --help' to view all commands")
		} else {
			describeIdentityProviderConfigInput.ClusterName = &args[0]

			// json parsing
			var IdentityProviderConfig eks.IdentityProviderConfig
			jsonFileName, _ := cmd.Flags().GetString("identity-provider-config")
			util.OpenAndReadJsonFile(jsonFileName, IdentityProviderConfig)
			describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

			describeIdentityProvicerConfig(describeIdentityProviderConfigInput)
		}
	},
}

func describeIdentityProvicerConfig(input eks.DescribeIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/describeIdentityProviderConfig"
	var output eks.DescribeIdentityProviderConfigOutput
	cobrautil.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(describeIdentityProviderConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// describeIdentityProvicerConfigCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	// describeIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter tags")
	describeIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")
}
*/
