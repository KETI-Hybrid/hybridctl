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

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade Kubernetes engine reosource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("upgrade called")
	},
}
var upgradeClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("upgrade called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}
		k8s_version, err := cmd.Flags().GetString("version")
		if err != nil {
			panic(err)
		}
		var cli = Cli{}
		cli.ClusterName = cluster_name
		cli.PlatformName = platform_name
		cli.Version = k8s_version

		if cli.PlatformName == "aks" {
			// cmd := exec.Command("az", "aks", "upgrade", "--resource-group", "hcp",
			// 	"--name", cli.ClusterName, "--kubernetes-version", cli.Version, "-y")
			// output, err := cmd.Output()
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(string(output))
			// }
			create_aks(cli)
		} else if cli.PlatformName == "gke" {
			cmd := exec.Command("gcloud", "container", "clusters", "upgrade", cli.ClusterName, "--cluster-version",
				cli.Version, "--region", "us-central1-a", "-q")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(output))
			}
		} else if cli.PlatformName == "eks" {
			// cmd := exec.Command("aws", "eks", "update-cluster-version", "--name", cli.ClusterName, "--kubernetes-version",
			// 	cli.Version)
			// output, err := cmd.Output()
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(string(output))
			// }
			create_eks(cli)
		}

	},
}
var upgradeNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("upgrade called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}
		k8s_version, err := cmd.Flags().GetString("version")
		if err != nil {
			panic(err)
		}

		node_name, err := cmd.Flags().GetString("nodepool-name")
		if err != nil {
			panic(err)
		}
		var cli = Cli{}
		cli.ClusterName = cluster_name
		cli.NodeName = node_name
		cli.PlatformName = platform_name
		cli.Version = k8s_version

		if cli.PlatformName == "aks" {
			// cmd := exec.Command("az", "aks", "nodepool", "upgrade", "--resource-group", "hcp",
			// 	"--cluster-name", cli.ClusterName, "--name", cli.NodeName, "--kubernetes-version", cli.Version)
			// output, err := cmd.Output()

			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(string(output))
			// }
			createNodepool_aks(cli)
		} else if cli.PlatformName == "gke" {
			cmd := exec.Command("gcloud", "container", "clusters", "upgrade", cli.ClusterName, "--node-pool", cli.NodeName, "--cluster-version",
				cli.Version, "--region", "us-central1-a", "-q")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(output))
			}
		} else if cli.PlatformName == "eks" {
			// cmd := exec.Command("aws", "eks", "update-nodegroup-version", "--name", cli.ClusterName, "--kubernetes-version",
			// 	cli.Version)
			// output, err := cmd.Output()
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(string(output))
			// }
			createNodepool_eks(cli)
		}

	},
}

func init() {
	RootCmd.AddCommand(upgradeCmd)
	upgradeCmd.AddCommand(upgradeClusterCmd)
	upgradeCmd.AddCommand(upgradeNodeCmd)

	upgradeClusterCmd.Flags().String("platform", "", "input your platform name")
	upgradeClusterCmd.Flags().String("cluster-name", "", "input your cluster name")
	upgradeClusterCmd.Flags().String("version", "", "input your k8s version")

	upgradeNodeCmd.Flags().String("platform", "", "input your platform name")
	upgradeNodeCmd.Flags().String("cluster-name", "", "input your cluster name")
	upgradeNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	upgradeNodeCmd.Flags().String("version", "", "input your k8s version")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
