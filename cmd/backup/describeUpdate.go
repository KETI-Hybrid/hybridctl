// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

/*
var describeUpdateInput eks.DescribeUpdateInput

// describeAddonCmd represents the describeAddon command
var describeUpdateCmd = &cobra.Command{
	Use:   "describe-update",
	Short: "A brief description of your command",
	Long: `
	- describe-update
		hybridctl describe-update <clusterName> <updateID>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl describe-update --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl describe-update --help' to view all commands")
		} else {
			describeUpdateInput.Name = &args[0]
			describeUpdateInput.UpdateId = &args[1]
			nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")
			addonName, _ := cmd.Flags().GetString("addon-name")
			if nodegroupName != "" {
				describeUpdateInput.NodegroupName = &nodegroupName
			}
			if addonName != "" {
				describeUpdateInput.AddonName = &addonName
			}
			describeUpdate(describeUpdateInput)
		}
	},
}

func describeUpdate(describeUpdateInput eks.DescribeUpdateInput) {
	httpPostUrl := "http://localhost:8080/describeUpdate"
	var output eks.DescribeUpdateOutput
	util.GetJson(httpPostUrl, describeUpdateInput, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(describeUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//describeAddonCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	describeUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	describeUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")
}
*/
