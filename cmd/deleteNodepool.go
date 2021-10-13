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
	"os/exec"

	"github.com/spf13/cobra"
)

// deleteNodepoolCmd represents the deleteNodepool command
var deleteNodepoolCmd = &cobra.Command{
	Use:   "deleteNodepool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("deleteNodepool called")
		cli := Cli{args[0], args[1]}
		fmt.Println(cli)
		if len(args) == 0 {
			fmt.Println("Run 'hybridctl deleteNodepool --help' to view all commands")
		} else if args[0] == "gke" {
			if args[1] == "" {

				delete_gke(cli)
				fmt.Println("Run 'hybridctl deleteNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				deleteNodepool_gke(cli)
			}
		} else if args[0] == "aks" {
			if args[1] == "" {

				fmt.Println("Run 'hybridctl deleteNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				deleteNodepool_aks(cli)
			}
		} else if args[0] == "eks" {
			if args[1] == "" {

				fmt.Println("Run 'hybridctl deleteNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				deleteNodepool_eks(cli)
			}
		}
	},
}

func deleteNodepool_gke(info Cli) {

	cluster := "cluster"

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/gke/" + cluster

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/gke/" + cluster

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func deleteNodepool_eks(info Cli) {

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/eks"

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/eks"

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func deleteNodepool_aks(info Cli) {

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/aks"

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/aks"

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func init() {
	RootCmd.AddCommand(deleteNodepoolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteNodepoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteNodepoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
