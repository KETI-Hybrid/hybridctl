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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"Hybrid_Cluster/hybridctl/util"

	hcpclusterapis "Hybrid_Cluster/pkg/apis/hcpcluster/v1alpha1"
	hcpclusterv1alpha1 "Hybrid_Cluster/pkg/client/hcpcluster/v1alpha1/clientset/versioned"
	u "Hybrid_Cluster/util"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var checkAKS, checkEKS, checkGKE = false, false, false
var master_config, _ = util.BuildConfigFromFlags("kube-master", "/root/.kube/config")
var master_client = kubernetes.NewForConfigOrDie(master_config)

type Cli struct {
	PlatformName string
	ClusterName  string
}

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
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

		if len(args) == 0 {
		} else {
			switch args[0] {
			case "aks":
				fallthrough
			case "eks":
				fallthrough
			case "gke":
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])
				platform := args[0]
				clustername := args[1]
				if !CheckHCPClusterListToJoin(platform, clustername) {
					return
				}
			case "register":
				platform := args[1]
				if platform == "" {
					fmt.Println("ERROR: Input Platform")
				}
				clustername := args[2]
				if clustername == "" {
					fmt.Println("ERROR: Input Clustername")
				}
				createPlatformNamespace()
				switch platform {
				case "aks":
					fallthrough
				case "eks":
					fallthrough
				case "gke":
					if CreateHCPCluster(platform, clustername) {
						err := u.ChangeConfigClusterName(platform, clustername)
						if err != nil {
							fmt.Println(err)
						}
					}
					return
				default:
					return
				}
			default:
				fmt.Println("Run 'hybridctl join --help' to view all commands")
			}
			//
		}
	},
}

func CheckHCPClusterListToJoin(platform string, clustername string) bool {
	hcp_cluster, err := hcpclusterv1alpha1.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
	}
	cluster_list, err := hcp_cluster.HcpV1alpha1().HCPClusters(platform).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, cluster := range cluster_list.Items {
		joinstatus := cluster.Spec.JoinStatus
		if cluster.Spec.ClusterPlatform == platform && cluster.Name == clustername {
			if joinstatus == "UNJOIN" {
				cluster.Spec.JoinStatus = "JOINING"
				_, err = hcp_cluster.HcpV1alpha1().HCPClusters(platform).Update(context.TODO(), &cluster, metav1.UpdateOptions{})
				fmt.Println(cluster.Spec.JoinStatus)
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
			APIVersion: "hcp.k8s.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clustername,
			Namespace: platform,
		},
		Spec: hcpclusterapis.HCPClusterSpec{
			ClusterPlatform: platform,
			KubeconfigInfo:  data,
			JoinStatus:      "UNJOIN",
		},
	}
	newhcpcluster, err := hcp_cluster.HcpV1alpha1().HCPClusters(platform).Create(context.TODO(), &cluster, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Printf("success to register %s in %s\n", newhcpcluster.Name, newhcpcluster.Namespace)
		return true
	}
}

func createPlatformNamespace() {

	namespaceList, _ := master_client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for i := range namespaceList.Items {
		if checkAKS && checkEKS && checkGKE {
			break
		}
		switch namespaceList.Items[i].Name {
		case "aks":
			checkAKS = true
			continue
		case "eks":
			checkEKS = true
			continue
		case "gke":
			checkGKE = true
			continue
		default:
			continue
		}
	}
	checkAndCreateNamespace(checkAKS, "aks")
	checkAndCreateNamespace(checkEKS, "eks")
	checkAndCreateNamespace(checkGKE, "gke")
}

func checkAndCreateNamespace(PlatformCheck bool, platformName string) {
	if !PlatformCheck {
		Namespace := corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: platformName,
			},
		}
		master_client.CoreV1().Namespaces().Create(context.TODO(), &Namespace, metav1.CreateOptions{})
	}
}

func init() {
	RootCmd.AddCommand(joinCmd)
}
