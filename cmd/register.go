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
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	resource "Hybrid_Cloud/kube-resource/namespace"
	hcpclusterapis "Hybrid_Cloud/pkg/apis/hcpcluster/v1alpha1"
	hcpclusterv1alpha1 "Hybrid_Cloud/pkg/client/hcpcluster/v1alpha1/clientset/versioned"
	u "Hybrid_Cloud/util"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// joinCmd represents the join command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: ` 
	NAME 
		hybridctl register PLATFORM CLUSTER_NAME

	DESCRIPTION
		
	>> cluster register PLATFORM CLUSTER <<

	* This command registers the cluster you want to manage, 
	For each platform, you must fill in the information below.
	Please refer to the INFO section

	PLATFORM means the Kubernetes platform of the cluster to register.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	  hybridctl register aks CLUSTER_NAME --resource-group RESOURCEGROUP

	- eks   elastic kubernetes service
	  hybridctl register eks CLUSTER_NAME

	- gke   google kuberntes engine
	  hybridctl egister gke CLUSTER_NAME 

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if len(args) < 2 {
			fmt.Println(cmd.Help())
		} else {
			platform := args[0]
			if platform == "" {
				fmt.Println("ERROR: Enter Platform")
				return
			}

			clustername := args[1]
			if clustername == "" {
				fmt.Println("ERROR: Enter Clustername")
				return
			}

			var arguments []string
			arguments = append(arguments, "bash/sh", platform, clustername)

			switch platform {
			case "aks":
				resource_group, _ := cmd.Flags().GetString("resource-group")
				if resource_group != "" {
					arguments = append(arguments, resource_group)
				} else {
					fmt.Println("ERROR: Enter resource group")
					return
				}
				fallthrough
			case "eks":
				fallthrough
			case "gke":
				command := &exec.Cmd{
					Path:   "/root/go/src/Hybrid_Cloud/hybridctl/cmd/register",
					Args:   arguments,
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				}

				err := command.Start()
				if err != nil {
					fmt.Println(err)
					return
				}

				err = command.Wait()
				if err != nil {
					// fmt.Println(err)
					return
				}

				resource.CreateNamespace("kube-master", HCP_NAMESPACE)
				if CreateHCPCluster(platform, clustername) {
					err := u.ChangeConfigClusterName(HCP_NAMESPACE, clustername)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Printf("fail to create HCPCluster %s\n", clustername)
					return
				}
			default:

			}
		}
	},
}

func CreateHCPCluster(platform string, clustername string) bool {
	hcp_cluster, err := hcpclusterv1alpha1.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
		return false
	}
	data, err := ioutil.ReadFile("/root/.kube/kubeconfig")
	if err != nil {
		fmt.Println("File reading error", err)
		return false
	}
	if data == nil {
		fmt.Printf("fail to get config about %s\n", clustername)
		return false
	}
	cluster := hcpclusterapis.HCPCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HCPCluster",
			APIVersion: "hcp.crd.com",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clustername,
			Namespace: HCP_NAMESPACE,
		},
		Spec: hcpclusterapis.HCPClusterSpec{
			ClusterPlatform: platform,
			KubeconfigInfo:  data,
			JoinStatus:      "UNJOIN",
		},
	}
	newhcpcluster, err := hcp_cluster.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Create(context.TODO(), &cluster, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Printf("success to register %s in %s\n", newhcpcluster.Name, newhcpcluster.Namespace)
		return true
	}
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("resource-group", "g", "", "")
}
