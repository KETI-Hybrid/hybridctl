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
	"strconv"

	hcppolicyapis "Hybrid_Cluster/pkg/apis/hcppolicy/v1alpha1"

	hcppolicyv1alpha1 "Hybrid_Cluster/pkg/client/policy/v1alpha1/clientset/versioned"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// var master_config, _ = util.BuildConfigFromFlags("kube-master", "/root/.kube/config")

// var master_client = kubernetes.NewForConfigOrDie(master_config)

// joinCmd represents the join command
var initialCmd = &cobra.Command{
	Use:   "initial-setting",
	Short: "A brief description of your command",
	Long: ` 
NAME 
	hybridctl join PLATFORM CLUSTER
	hybridctl join register PLATFORM

DESCRIPTION
	
	>> cluster join PLATFORM CLUSTER <<


	PLATFORM means the Kubernetes platform of the cluster to join.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	- eks   elastic kubernetes service
	- gke   google kuberntes engine

	* PLATFORM mut be written in LOWERCASE letters

	CLUSTER means the name of the cluster on the specified platform.

	>> hybridctl join register PLATFORM <<

	* This command registers the cluster you want to manage, 
	For each platform, you must fill in the information below.
	Please refer to the INFO section

	PLATFORM means the Kubernetes platform of the cluster to join.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	- eks   elastic kubernetes service
	- gke   google kuberntes engine

	[INFO]

		GKE 
		- projectid    the ID of GKE cloud project to use. 
		- clustername  the name of the cluster on the specified platform.
		- region       choose Google Compute Zone from 1 to 85.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		polices := make([]hcppolicyapis.HCPPolicies, 5)
		var cost int
		var max_cpu int
		var max_memory int
		var default_node_option string
		var extra int
		var exist bool = false
		hcp_policy, err := hcppolicyv1alpha1.NewForConfig(master_config)
		if err != nil {
			log.Println(err)
		}

		list, err := hcp_policy.HcpV1alpha1().HCPPolicies("hcp").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Println(err)
		}
		for _, policy := range list.Items {
			if "initial-setting" == policy.Name {
				exist = true
			}
		}

		if !exist {
			fmt.Printf("cost : ")
			fmt.Scanln(&cost)
			polices[0].Type = "cost"
			_ = append(polices[0].Value, strconv.Itoa(cost))

			fmt.Printf("Enter the maximum number of CPUs to allocate in the cluster. : ")
			fmt.Scanln(&max_cpu)
			polices[1].Type = "max_cpu"
			_ = append(polices[1].Value, strconv.Itoa(max_cpu))

			fmt.Printf("Enter the maximum amount of memory to allocate in the cluster.. : ")
			fmt.Scanln(&max_memory)
			polices[2].Type = "max_memory"
			_ = append(polices[2].Value, strconv.Itoa(max_memory))

			fmt.Printf("Enter the node option to use as default : ")
			fmt.Scanln(&default_node_option)
			polices[3].Type = "default_node_option"
			_ = append(polices[3].Value, default_node_option)

			fmt.Printf("Enter the percentage of free resources to use when automatically creating a node. : ")
			fmt.Scanln(&extra)
			polices[4].Type = "extra"
			_ = append(polices[4].Value, strconv.Itoa(extra))

			policy := hcppolicyapis.HCPPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "HCPPolicy",
					APIVersion: "hcp.k8s.io/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initial-setting",
					Namespace: "hcp",
				},
				Spec: hcppolicyapis.HCPPolicySpec{
					Template: hcppolicyapis.HCPPolicyTemplate{
						Spec: hcppolicyapis.HCPPolicyTemplateSpec{
							Policies: polices,
						},
					},
				},
			}
			hcp_policy.HcpV1alpha1().HCPPolicies("hcp").Create(context.TODO(), &policy, metav1.CreateOptions{})
		} else {
			var answer string
			fmt.Printf("Do you want to update initial-setting? [y/n] : ")
			fmt.Scanln(&answer)
			if answer == "y" {
				fmt.Printf("cost : ")
				fmt.Scanln(&cost)
				polices[0].Type = "cost"
				_ = append(polices[0].Value, strconv.Itoa(cost))

				fmt.Printf("Enter the maximum number of CPUs to allocate in the cluster. : ")
				fmt.Scanln(&max_cpu)
				polices[1].Type = "max_cpu"
				_ = append(polices[1].Value, strconv.Itoa(max_cpu))

				fmt.Printf("Enter the maximum amount of memory to allocate in the cluster.. : ")
				fmt.Scanln(&max_memory)
				polices[2].Type = "max_memory"
				_ = append(polices[2].Value, strconv.Itoa(max_memory))

				fmt.Printf("Enter the node option to use as default : ")
				fmt.Scanln(&default_node_option)
				polices[3].Type = "default_node_option"
				_ = append(polices[3].Value, default_node_option)

				fmt.Printf("Enter the percentage of free resources to use when automatically creating a node. : ")
				fmt.Scanln(&extra)
				polices[4].Type = "extra"
				_ = append(polices[4].Value, strconv.Itoa(extra))

				policy := hcppolicyapis.HCPPolicy{
					TypeMeta: metav1.TypeMeta{
						Kind:       "HCPPolicy",
						APIVersion: "hcp.k8s.io/v1alpha1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "initial-setting",
						Namespace: "hcp",
					},
					Spec: hcppolicyapis.HCPPolicySpec{
						Template: hcppolicyapis.HCPPolicyTemplate{
							Spec: hcppolicyapis.HCPPolicyTemplateSpec{
								Policies: polices,
							},
						},
					},
				}
				hcp_policy.HcpV1alpha1().HCPPolicies("hcp").Update(context.TODO(), &policy, metav1.UpdateOptions{})
			} else if answer == "n" {
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(initialCmd)
}
