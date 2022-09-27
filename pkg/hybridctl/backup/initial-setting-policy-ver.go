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

package backup

import (
	"fmt"
	"strconv"

	hcppolicyapis "github.com/KETI-Hybrid/hcp-pkg/apis/hcppolicy/v1alpha1"

	"github.com/KETI-Hybrid/hcp-pkg/util/clientset"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
		polices := make([]hcppolicyapis.HCPPolicies, 4)
		var max_cpu int
		var max_memory int
		var default_node_option string
		var extra int
		var exist bool = false

		list, _ := clientset.HCPPolicyClientset.HcpV1alpha1().HCPPolicies("hcp").List(context.TODO(), metav1.ListOptions{})
		for _, policy := range list.Items {
			if policy.Name == "initial-setting" {
				exist = true
			}
		}

		if !exist {
			fmt.Printf("Enter the maximum number of CPUs to allocate in the cluster.[NanoCores] : ")
			fmt.Scanln(&max_cpu)
			polices[0].Type = "max_cpu"
			polices[0].Value = append(polices[0].Value, strconv.Itoa(max_cpu))

			fmt.Printf("Enter the maximum amount of memory to allocate in the cluster.. : ")
			fmt.Scanln(&max_memory)
			polices[1].Type = "max_memory"
			polices[1].Value = append(polices[1].Value, strconv.Itoa(max_memory))

			fmt.Printf("Enter the node option to use as default [ Low / Middle / High ]: ")
			fmt.Scanln(&default_node_option)
			polices[2].Type = "default_node_option"
			polices[2].Value = append(polices[2].Value, default_node_option)

			fmt.Printf("Enter the percentage of free resources to use when automatically creating a node. : ")
			fmt.Scanln(&extra)
			polices[3].Type = "extra"
			polices[3].Value = append(polices[3].Value, strconv.Itoa(extra))

			fmt.Println(polices)

			policy := hcppolicyapis.HCPPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "HCPPolicy",
					APIVersion: "hcp.crd.com/v1alpha1",
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
					PolicyStatus: "Enabled",
				},
			}
			_, err := clientset.HCPPolicyClientset.HcpV1alpha1().HCPPolicies("hcp").Create(context.TODO(), &policy, metav1.CreateOptions{})
			if err != nil {
				fmt.Println(err)
			}
		} else {
			var answer string
			fmt.Printf("Do you want to update initial-setting? [y/n] : ")
			fmt.Scanln(&answer)
			if answer == "y" {

				fmt.Printf("Enter the maximum number of CPUs to allocate in the cluster.[NanoCores] : ")
				fmt.Scanln(&max_cpu)
				polices[0].Type = "max_cpu"
				polices[0].Value = append(polices[0].Value, strconv.Itoa(max_cpu))

				fmt.Printf("Enter the maximum amount of memory to allocate in the cluster. : ")
				fmt.Scanln(&max_memory)
				polices[1].Type = "max_memory"
				polices[1].Value = append(polices[1].Value, strconv.Itoa(max_memory))

				for (default_node_option != "Low") && (default_node_option != "Middle") && (default_node_option != "High") {
					fmt.Printf("Enter the node option to use as default [ Low / Middle / High ]: ")
					fmt.Scanln(&default_node_option)
					polices[2].Type = "default_node_option"
					polices[2].Value = append(polices[2].Value, default_node_option)
				}

				fmt.Printf("Enter the percentage of free resources to use when automatically creating a node. : ")
				fmt.Scanln(&extra)
				polices[3].Type = "extra"
				polices[3].Value = append(polices[3].Value, strconv.Itoa(extra))

				fmt.Println(polices)
				policy, _ := clientset.HCPPolicyClientset.HcpV1alpha1().HCPPolicies("hcp").Get(context.TODO(), "initial-setting", metav1.GetOptions{})
				policy.Spec.Template.Spec.Policies = polices
				_, err := clientset.HCPPolicyClientset.HcpV1alpha1().HCPPolicies("hcp").Update(context.TODO(), policy, metav1.UpdateOptions{})
				if err != nil {
					fmt.Println(err)
				}
			} else if answer == "n" {
				return
			}
		}
	},
}
