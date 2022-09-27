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
	"os"
	"os/exec"

	"github.com/KETI-Hybrid/hcp-pkg/util/clientset"

	hcpclusterapis "github.com/KETI-Hybrid/hcp-pkg/apis/hcpcluster/v1alpha1"
	resource "github.com/KETI-Hybrid/hcp-pkg/kube-resource/namespace"

	"github.com/KETI-Hybrid/hybridctl-v1/pkg/nks"
	cobrautil "github.com/KETI-Hybrid/hybridctl-v1/util"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

// joinCmd represents the join command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: ` 
	NAME 
		hybridctl register PLATFORM CLUSTER_NAME

	DESCRIPTION
		
	>> hybridctl register PLATFORM CLUSTERNAME <<

	* This command registers the cluster you want to manage, 
	For each platform, you must fill in the information below.
	Please refer to the INFO section

	PLATFORM means the Kubernetes platform of the cluster to register.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	  hybridctl register aks CLUSTER_NAME --resource-group RESOURCEGROUP

	- eks   elastic kubernetes service
	  hybridctl register eks CLUSTER_NAME --region REGION

	- gke   google kuberntes engine
	  hybridctl register gke CLUSTER_NAME
	
	- nks naver kubernetes service
	hybridctl register nks CLUSTER_NAME --region REGION

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
			arguments = append(arguments, "bin/bash", platform, clustername)
			var region string
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
				region, _ = cmd.Flags().GetString("region")
				if region != "" {
					arguments = append(arguments, region)
				} else {
					fmt.Println("ERROR: Enter region")
					return
				}
				fallthrough
			case "gke":
				command := &exec.Cmd{
					Path:   "/root/go/src/Hybrid_LCW/github.com/KETI-Hybrid/hybridctl-v1/cmd/register/register",
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

				resource.CreateNamespace(clientset.MasterClienset, HCP_NAMESPACE)
				if CreateHCPCluster(platform, clustername, region) {
					err := cobrautil.ChangeConfigClusterName(HCP_NAMESPACE, clustername)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Printf("fail to create HCPCluster %s\n", clustername)
					return
				}
			case "nks":
				region, _ = cmd.Flags().GetString("region")
				if region != "" {
					arguments = append(arguments, region)
				} else {
					fmt.Println("ERROR: Enter region")
					return
				}
				// klog.Infoln(arguments)
				// arguments = ["bin/bash", "platform", "cluster-name", "region"]
				nks_clustername := arguments[2]

				// klog.Infoln(ncp_clustername)
				real_clustername, err := nks.NksGetClusterName(nks_clustername)
				// klog.Infoln(real_clustername)
				if err != nil {
					klog.Error(err)
				}
				arguments[2] = real_clustername

				command := &exec.Cmd{
					Path:   "/root/go/src/Hybrid_LCW/github.com/KETI-Hybrid/hybridctl-v1/cmd/register/register",
					Args:   arguments,
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				}

				err = command.Start()
				if err != nil {
					fmt.Println(err)
					return
				}

				err = command.Wait()
				if err != nil {
					// fmt.Println(err)
					return
				}

				resource.CreateNamespace(clientset.MasterClienset, HCP_NAMESPACE)
				if CreateHCPCluster(platform, clustername, region) {
					err := cobrautil.ChangeConfigClusterName(HCP_NAMESPACE, clustername)
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

func CreateHCPCluster(platform string, clustername string, region string) bool {

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
			Region:          region,
			KubeconfigInfo:  data,
			JoinStatus:      "UNJOIN",
		},
	}

	fmt.Println(cluster.Spec.Region)
	newhcpcluster, err := clientset.HCPClusterClientset.HcpV1alpha1().HCPClusters(HCP_NAMESPACE).Create(context.TODO(), &cluster, metav1.CreateOptions{})

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
	registerCmd.Flags().StringP("region", "", "", "")
}
