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

// describeNodepoolCmd represents the describeNodepool command
var describeNodepoolCmd = &cobra.Command{
	Use:   "describeNodepool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("describeNodepool called")
		if len(args) == 0 {
			fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
		} else if args[0] == "gke" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				describeNodepool_gke(args[1])
			}
		} else if args[0] == "aks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				describeNodepool_aks(args[1])
			}
		} else if args[0] == "eks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				describeNodepool_eks(args[1])
			}
		}
	},
}

func describeNodepool_gke(clusterName string) {
	cluster := "cluster"

	nodepool_name_dir := "google_container_node_pool." + clusterName + "_nodes"
	cmd := exec.Command("terraform", "state", "show", nodepool_name_dir)
	cmd.Dir = "../terraform/gke/" + cluster

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func describeNodepool_aks(clusterName string) {

	nodepool_name_dir := "aws_eks_node_group." + clusterName
	cmd := exec.Command("terraform", "state", "show", nodepool_name_dir)
	cmd.Dir = "../terraform/eks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func describeNodepool_eks(clusterName string) {

	nodepool_name_dir := "azurerm_kubernetes_cluster_node_pool." + clusterName
	cmd := exec.Command("terraform", "state", "show", nodepool_name_dir)
	cmd.Dir = "../terraform/aks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func init() {
	RootCmd.AddCommand(describeNodepoolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeNodepoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeNodepoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
