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

	"hcp-pkg/util/clientset"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// unjoinCmd represents the unjoin command
var unjoinCmd = &cobra.Command{
	Use:   "unjoin",
	Short: "A brief description of your command",
	Long: ` 
	NAME 
		hybridctl unjoin CLUSTER_NAME

	DESCRIPTION
		
		>> cluster unjoin CLUSTER_NAME <<

	CLUSTER_NAME means the name of the cluster on the specified platform.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println(cmd.Help())
		} else {

			fmt.Printf("Cluster Name : %s\n", args[0])
			clustername := args[0]
			cluster, err := clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Get(context.TODO(), clustername, metav1.GetOptions{})
			if err != nil {
				log.Println(err)
			}

			joinstatus := cluster.Spec.JoinStatus
			if joinstatus == "UNJOIN" {
				fmt.Println("ERROR: This is an already unjoined cluster.")
				return
			} else if joinstatus == "UNJOINING" {
				fmt.Println("ERROR: Cluster is already waiting to unjoin")
				return
			} else if joinstatus == "JOINING" {
				fmt.Println("ERROR: Cluster is waiting to join")
				return
			} else {
				cluster.Spec.JoinStatus = "UNJOINING"
				_, err = clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Update(context.TODO(), cluster, metav1.UpdateOptions{})
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(unjoinCmd)
}
