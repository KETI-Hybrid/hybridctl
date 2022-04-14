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

// var checkAKS, checkEKS, checkGKE = false, false, false
// var master_config, _ = util.BuildConfigFromFlags("kube-master", "/root/.kube/config")
// var master_client = kubernetes.NewForConfigOrDie(master_config)

// type Cli struct {
// 	PlatformName string
// 	ClusterName  string
// }

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
			hcp_cluster, err := hcpclusterv1alpha1.NewForConfig(master_config)
			if err != nil {
				log.Println(err)
			}

			cluster, err := hcp_cluster.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Get(context.TODO(), clustername, metav1.GetOptions{})
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
				_, err = hcp_cluster.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Update(context.TODO(), cluster, metav1.UpdateOptions{})
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
