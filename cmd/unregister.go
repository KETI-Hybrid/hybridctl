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
	"time"

	"hcp-pkg/hcp-resource/hcpcluster"
	"hcp-pkg/util/clientset"

	cobrautil "hybridctl/util"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	kubefed "sigs.k8s.io/kubefed/pkg/client/generic"
)

// joinCmd represents the join command
var unregisterCmd = &cobra.Command{
	Use:   "unregister",
	Short: "A brief description of your command",
	Long: ` 
	NAME 
		hybridctl unregister PLATFORM CLUSTER_NAME

	DESCRIPTION
		
	>> hybridctl unregister PLATFORM CLUSTERNAME <<

	* This command registers the cluster you want to manage, 
	For each platform, you must fill in the information below.
	Please refer to the INFO section

	PLATFORM means the Kubernetes platform of the cluster to register.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	  hybridctl unregister aks CLUSTER_NAME 

	- eks   elastic kubernetes service
	  hybridctl unregister eks CLUSTER_NAME

	- gke   google kuberntes engine
	  hybridctl unregister gke CLUSTER_NAME 

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

			switch platform {
			case "aks":
				fallthrough
			case "eks":
				fallthrough
			case "gke":
				fallthrough
			case "nks":
				if hcpcluster.FindHCPClusterList(clientset.HCPClusterClientset, clustername) {
					HCP_NAMESPACE = "hcp"
					if Iskubefedcluster(clustername) {
						fmt.Println(">>> unjoin cluster")
						cluster, err := clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Get(context.TODO(), clustername, metav1.GetOptions{})
						if err != nil {
							log.Println(err)
						}
						cluster.Spec.JoinStatus = "UNJOINING"
						_, err = clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Update(context.TODO(), cluster, metav1.UpdateOptions{})
						if err != nil {
							fmt.Println(err)
						}
					}

					fmt.Println(">>> delete hcpcluster")
					for {
						cluster, err := clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Get(context.TODO(), clustername, metav1.GetOptions{})
						if err != nil {
							log.Println(err)
						}

						if cluster.Spec.JoinStatus == "UNJOINING" {
							timer := time.NewTimer(time.Second * 7)
							<-timer.C
							fmt.Printf("Failed to UNJOIN cluster %s", clustername)
						}

						fmt.Println(">>> delete config in kubeconfig")
						err = cobrautil.DeleteConfig(platform, clustername)
						if err != nil {
							fmt.Println(err)
						}

						err = clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Delete(context.TODO(), clustername, metav1.DeleteOptions{})
						if err != nil {
							log.Println(err)
						} else {
							break
						}
					}

				} else {
					fmt.Printf("%s does not exist\n", clustername)
				}
			default:
			}
		}
	},
}

func Iskubefedcluster(clustername string) bool {
	clientset := kubefed.NewForConfigOrDie(clientset.MasterConfig)
	tempClusterList := &fedv1b1.KubeFedClusterList{}
	_ = clientset.List(context.TODO(), tempClusterList, "kube-federation-system", &client.ListOptions{})
	for _, i := range tempClusterList.Items {
		if i.Name == clustername {
			return true
		}
	}
	return false
}

func init() {
	RootCmd.AddCommand(unregisterCmd)
}
