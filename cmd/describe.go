// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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

	"github.com/KETI-Hybrid/hybridctl-v1/pkg/nks"

	klog "k8s.io/klog/v2"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe Kubernetes engine resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:


Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("describe called")
	},
}

var describeClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("describe called")

		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			fmt.Println("여기")
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}

		var cli = Cli{}
		cli.ClusterName = cluster_name
		cli.PlatformName = platform_name

		if cli.PlatformName == "gke" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeCluster_gke(cli.ClusterName)
			}
		} else if cli.PlatformName == "eks" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeCluster_eks(cli.ClusterName)
			}
		} else if cli.PlatformName == "aks" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeCluster_aks(cli.ClusterName)
			}
		} else if cli.PlatformName == "nks" {
			if cli.ClusterName == "" {
				klog.Infoln("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				real_clustername, err := nks.NksGetClusterName(cli.ClusterName)
				if err != nil {
					klog.Error(err)
				}
				klog.Infoln(real_clustername)
				if real_clustername != "" {
					nks.NksDescribeCluster(real_clustername)
				}
			}
		}
	},
}

func describeCluster_eks(clusterName string) {
	cluster_name_dir := "aws_eks_cluster." + clusterName
	cmd := exec.Command("terraform", "state", "show", cluster_name_dir)
	cmd.Dir = "../terraform/eks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func describeCluster_aks(clusterName string) {
	cluster_name_dir := "azurerm_kubernetes_cluster." + clusterName
	cmd := exec.Command("terraform", "state", "show", cluster_name_dir)
	cmd.Dir = "../terraform/aks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func describeCluster_gke(clusterName string) {
	cluster := "cluster"

	cluster_name_dir := "google_container_cluster." + clusterName
	cmd := exec.Command("terraform", "state", "show", cluster_name_dir)
	cmd.Dir = "../terraform/gke/" + cluster

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

var describeNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("describe called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}
		nodepool_name, err := cmd.Flags().GetString("nodepool-name")
		if err != nil {
			panic(err)
		}

		var cli = Cli{}
		cli.ClusterName = cluster_name
		cli.PlatformName = platform_name
		cli.NodeName = nodepool_name

		if cli.PlatformName == "gke" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeNodepool_gke(cli.ClusterName)
			}
		} else if cli.PlatformName == "aks" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeNodepool_aks(cli.ClusterName)
			}
		} else if cli.PlatformName == "eks" {
			if cli.ClusterName == "" {
				fmt.Println("Run 'hybridctl describeCluster --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", cli.PlatformName)
				fmt.Printf("Cluster Name: %s\n", cli.ClusterName)
				describeNodepool_eks(cli.ClusterName)
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

func describeNodepool_eks(clusterName string) {

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

func describeNodepool_aks(clusterName string) {

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
	RootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describeClusterCmd)
	describeCmd.AddCommand(describeNodeCmd)

	describeClusterCmd.Flags().String("platform", "", "input your platform name")
	describeClusterCmd.Flags().String("cluster-name", "", "input your cluster name")

	describeNodeCmd.Flags().String("platform", "", "input your platform name")
	describeNodeCmd.Flags().String("cluster-name", "", "input your cluster name")
	describeNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
