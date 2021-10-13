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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

// listClusterCmd represents the listCluster command
var listClusterCmd = &cobra.Command{
	Use:   "listCluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("listCluster called")

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl create --help' to view all commands")
		} else if args[0] == "gke" {

			fmt.Println("kubernetes engine Name : ", args[0])

		} else if args[0] == "eks" {

			fmt.Println("kubernetes engine Name : ", args[0])

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-2")},
			)
			if err != nil {
				fmt.Print(err)
			}

			svc := eks.New(sess)
			input := &eks.ListClustersInput{}

			result, err := svc.ListClusters(input)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case eks.ErrCodeInvalidParameterException:
						fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
					case eks.ErrCodeClientException:
						fmt.Println(eks.ErrCodeClientException, aerr.Error())
					case eks.ErrCodeServerException:
						fmt.Println(eks.ErrCodeServerException, aerr.Error())
					case eks.ErrCodeServiceUnavailableException:
						fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
					default:
						fmt.Println(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
				}
				return
			}

			fmt.Println(result)

		} else if args[0] == "aks" {

			fmt.Println("kubernetes engine Name : ", args[0])

		}
	},
}

func init() {
	RootCmd.AddCommand(listClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listClusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
