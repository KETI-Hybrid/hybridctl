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
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("delete called")
		cli := Cli{args[0], args[1]}
		fmt.Println(cli)
		if len(args) == 0 {
			fmt.Println("Run 'hybridctl delete --help' to view all commands")
		} else if args[0] == "gke" {
			if args[1] == "" {

				delete_gke(cli)
				fmt.Println("Run 'hybridctl delete --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				delete_gke(cli)
			}
		} else if args[0] == "aks" {
			if args[1] == "" {

				fmt.Println("Run 'hybridctl delete --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				delete_aks(cli)
			}
		} else if args[0] == "eks" {
			if args[1] == "" {

				fmt.Println("Run 'hybridctl delete --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				delete_eks(cli)
			}
		}
	},
}

func delete_eks(info Cli) {
	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/eks"

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool.tf.json")
	cmd_rm_nodepool.Dir = "../terraform/eks"

	output, err = cmd_rm_nodepool.Output()
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

func delete_aks(info Cli) {
	// num := 1
	// data := make([]Cluster_info, 1)
	// cluster := "cluster"

	// fmt.Println("!", info.ClusterName, "!")

	// data[0].Project_id = "keti-container"
	// data[0].Cluster_name = info.ClusterName
	// data[0].Region = "us-central1-a"
	// data[0].Gke_num_nodes = uint64(num)

	// doc, _ := json.Marshal(data)

	// fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/create/", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/aks"

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	// cmd := exec.Command("terraform", "destroy", "-auto-approve")

	// cmd1 := exec.Command("terraform", "plan", "-lock=false")
	// cmd1.Dir = "../terraform/aks"
	// output, err = cmd1.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(output))
	// }
	cmd2 := exec.Command("terraform", "apply", "-auto-approve")

	cmd2.Dir = "../terraform/aks"

	output, err = cmd2.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func delete_gke(info Cli) {
	num := 1
	data := make([]Cluster_info, 1)
	cluster := "cluster"

	fmt.Println("!", info.ClusterName, "!")

	data[0].Project_id = "keti-container"
	data[0].Cluster_name = info.ClusterName
	data[0].Region = "us-central1-a"
	data[0].Gke_num_nodes = uint64(num)

	doc, _ := json.Marshal(data)

	fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/create/", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/gke/" + cluster

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodePool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/gke/" + cluster

	output, err = cmd_rm_nodepool.Output()
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

func init() {
	RootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
