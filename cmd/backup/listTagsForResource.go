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
import (
	"Hybrid_Cloud/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var listTagsForResourceInput eks.ListTagsForResourceInput

// listTagsForResourceCmd represents the listTagsForResource command
var listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A brief description of your command",
	Long: `
	- list-tags-for-resource
		hybridctl list-tags-for-resource --resource-arn

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		listTagsForResourceInput.ResourceArn = &resourceArn
		listTagsForResource(listTagsForResourceInput)
	},
}

func listTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) {
	httpPostUrl := "http://localhost:8080/listTagsForResource"
	var output eks.ListTagsForResourceOutput
	util.GetJson(httpPostUrl, listTagsForResourceInput, &output)
	fmt.Printf("%+v\n", output)
}

func init() {
	EksCmd.AddCommand(listTagsForResourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//listTagsForResourceCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
	listTagsForResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")
}
*/
