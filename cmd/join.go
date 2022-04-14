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
	"log"

	hcpclusterv1alpha1 "Hybrid_Cloud/pkg/client/hcpcluster/v1alpha1/clientset/versioned"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var HCP_NAMESPACE string = "hcp"

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "A brief description of your command",
	Long: ` 
	NAME 
		hybridctl join CLUSTER

	DESCRIPTION
	
	>> cluster join CLUSTER <<

	CLUSTER means the name of the cluster on the specified platform.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println(cmd.Help())
		} else {
			fmt.Printf("Cluster Name : %s\n", args[0])
			clustername := args[0]
			if !CheckHCPClusterListToJoin(clustername) {
				return
			}
		}
	},
}

func CheckHCPClusterListToJoin(clustername string) bool {
	hcp_cluster, err := hcpclusterv1alpha1.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
	}
	cluster_list, err := hcp_cluster.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return false
	}

	for _, cluster := range cluster_list.Items {
		joinstatus := cluster.Spec.JoinStatus
		if cluster.Name == clustername {
			if joinstatus == "UNJOIN" {
				cluster.Spec.JoinStatus = "JOINING"
				_, err = hcp_cluster.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Update(context.TODO(), &cluster, metav1.UpdateOptions{})
				if err != nil {
					fmt.Println(err)
					return false
				}
				return true
			} else if joinstatus == "JOINING" {
				fmt.Println("ERROR: Cluster is already waiting to join")
				return false
			} else if joinstatus == "JOIN" {
				fmt.Println("ERROR: This is an already joined cluster.")
				return false
			} else {
				fmt.Println("ERROR: JOINSTATUS is wrong")
				return false
			}
		}
	}
	fmt.Println("ERROR: no such Cluster")
	fmt.Println("you must register yout cluster to join")
	fmt.Println("ex) kubectl register <platform> <clustername>")
	return false
}

func init() {
	RootCmd.AddCommand(joinCmd)
}
