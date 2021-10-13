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

import (
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
			updateAddonInput.ClusterName = &args[0]
			updateAddonInput.AddonName = &args[1]
			updateAddon(updateAddonInput)
		}
	},
}

func updateAddon(updateAddonInput eks.UpdateAddonInput) {
	httpPostUrl := "http://10.0.5.83:8000/updateAddon"
	var output eks.UpdateAddonOutput
	getJson(httpPostUrl, updateAddonInput, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	RootCmd.AddCommand(updateAddonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//updateAddonCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
}
