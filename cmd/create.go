// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Kubernetes engine resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")
	},
}

var uniCmd = &cobra.Command{
	Use:   "uni",
	Short: "Command to deploy Uni Kernel sample Pods on uni-master cluster",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")
	},
}
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Command to create a cluster corresponding to each kubernetes engine platform ",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")

		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}

		var cli = Cli{}
		cli.ClusterName = cluster_name

		cli.PlatformName = platform_name

		if platform_name == "aks" {
			fmt.Println("call create_aks func")
			create_aks(cli)
		} else if platform_name == "eks" {
			fmt.Println("call create_eks func")
			create_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call create_gke func")
			create_gke(cli)
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}
	},
}

var createNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Command to deploy nodepool corresponding to each cluster",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")

		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}

		nodepool_name, err := cmd.Flags().GetString("nodepool-name")
		if err != nil {
			panic(err)
		}

		var cli = Cli{}
		cli.ClusterName = cluster_name

		cli.PlatformName = platform_name

		cli.NodeName = nodepool_name

		if platform_name == "aks" {
			fmt.Println("call create_aks func")
			// create_aks(cli)
			createNodepool_aks(cli)
		} else if platform_name == "eks" {
			fmt.Println("call create_eks func")
			createNodepool_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call create_gke func")
			createNodepool_gke(cli)
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}
	},
}

func createNodepool_eks(info Cli) {
	if info.NodeCount == "" {
		info.NodeCount = "1"
	}
	if info.Version == "" {
		info.Version = "1.21"
	}
	send_js_cluster := TF_EKS_Nopepool{
		EksResourceNode: EksResourceNode{
			EksResourceNode: &map[string]EksClusterNode{
				info.NodeName: {
					ClusterName:   "${aws_eks_cluster." + info.ClusterName + ".name}",
					NodeGroupName: info.NodeName,
					NodeRoleArn:   "${aws_iam_role.terraform-eks-node.arn}",
					SubnetIds:     "${aws_subnet.terraform-eks-private-subnet[*].id}",
					InstanceTypes: []string{"m5.large"},
					DiskSize:      "50",
					EksLabels: EksLabels{
						Role: "terraform-eks-m5-large",
					},
					ScalingConfig: ScalingConfig{
						DesiredSize: info.NodeCount,
						MinSize:     info.NodeCount,
						MaxSize:     info.NodeCount,
					},
					DependsOn: []string{"aws_iam_role_policy_attachment.terraform-eks-node-AmazonEKSWorkerNodePolicy",
						"aws_iam_role_policy_attachment.terraform-eks-node-AmazonEKS_CNI_Policy",
						"aws_iam_role_policy_attachment.terraform-eks-node-AmazonEC2ContainerRegistryReadOnly"},
					EksTags: EksTags{
						Name: "${aws_eks_cluster." + info.ClusterName + ".name}-terraform-eks-m5-large-Node",
					},
				},
			},
		},
	}
	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/eks/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/eks/"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func createNodepool_gke(info Cli) {
	cluster := "cluster"
	if info.NodeCount == "" {
		info.NodeCount = "1"
	}

	send_js_nodePool := TF_NodePool{
		NodePool_Resource: &NodePool_Resource{
			Google_container_node_pool: &map[string]Node_pool_type{
				info.NodeName: {
					Name:       info.NodeName,
					Location:   "us-central1-a",
					Cluster:    "${google_container_cluster." + info.ClusterName + ".name}",
					Node_count: info.NodeCount,
					Node_config: &Node_config{
						Labels: &Labels{
							Env: "keti-container",
						},
						Metadata: &Metadata{
							Disable_legacy_endpoints: "true",
						},
						Tags:         []string{"gke-node", "keti-container-gke"},
						Machine_type: "n1-standard-1",
						Oauth_scopes: []string{"https://www.googleapis.com/auth/logging.write", "https://www.googleapis.com/auth/monitoring"},
					},
				},
			},
		},
	}

	send, err := json.MarshalIndent(send_js_nodePool, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/gke/" + cluster

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func createNodepool_aks(info Cli) {
	if info.Version == "" {
		info.Version = "1.21.9"
	}
	if info.NodeCount == "" {
		info.NodeCount = "1"
	}
	send_js_cluster := TF_AKS_Nodepool{
		AksNodepoolResource: AksNodepoolResource{
			AzurermKubernetesClusterNodePool: &map[string]AksNodepool{
				info.NodeName: {
					Name:                 info.NodeName,
					KubernetesClusterID:  "${azurerm_kubernetes_cluster." + info.ClusterName + ".id}",
					Orchestrator_Version: info.Version,
					VMSize:               "Standard_DS2_v2",
					NodeCount:            info.NodeCount,
					AksNodeTags: AksNodeTags{
						Environment: "Production",
					},
				},
			},
		},
	}
	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/aks/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	// cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd := exec.Command("./build.sh")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/aks/"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func create_eks(info Cli) {
	fmt.Println("!")
	if info.Version == "" {
		info.Version = "1.21"
	}
	send_js_cluster := TF_EKS{
		ResourceEksCluster: ResourceEksCluster{
			AwsEksCluster: &map[string]EksCluster{
				info.ClusterName: {
					Name:                   info.ClusterName,
					RoleArn:                "${aws_iam_role.terraform-eks-cluster.arn}",
					Version:                info.Version,
					EnabledClusterLogTypes: []string{"api", "audit", "authenticator", "controllerManager", "scheduler"},
					VpcConfig: VpcConfig{
						SubnetIds:             "${concat(aws_subnet.terraform-eks-public-subnet[*].id, aws_subnet.terraform-eks-private-subnet[*].id)}",
						EndpointPrivateAccess: "true",
						EndpointPublicAccess:  "true",
					},
					DependsOn: []string{"aws_iam_role_policy_attachment.terraform-eks-cluster-AmazonEKSClusterPolicy", "aws_iam_role_policy_attachment.terraform-eks-cluster-AmazonEKSVPCResourceController"},
				},
			},
		},
	}

	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/eks/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/eks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
func create_aks(info Cli) {

	if info.Version == "" {
		info.Version = "1.21.9"
	}
	send_js_cluster := TF_AKS{
		ResourceAksCluster: &ResourceAksCluster{
			AzurernKubernetesCluster: &map[string]AksCluster{
				info.ClusterName: {
					Name:               info.ClusterName,
					Kubernetes_Version: info.Version,
					Location:           "${azurerm_resource_group.k8s.location}",
					ResourceGroupName:  "${azurerm_resource_group.k8s.name}",
					DNSPrefix:          "${var.dns_prefix}",
					LinuxProfile: LinuxProfile{
						AdminUsername: "ubuntu",
						SSHKey: SSHKey{
							KeyData: "${file(var.ssh_public_key)}",
						},
					},
					DefaultNodePool: DefaultNodePool{
						Name:      "agentpool",
						NodeCount: 1,
						VMSize:    "Standard_D2_v2",
					},
					ServicePrincipal: ServicePrincipal{
						ClientID:     "${var.appId}",
						ClientSecret: "${var.password}",
					},
					AddonProfile: AddonProfile{
						OmsAgent: OmsAgent{
							Enabled:                 "true",
							LogAnalyticsWorkspaceID: "${azurerm_log_analytics_workspace.test.id}",
						},
					},
					NetworkProfile: NetworkProfile{
						LoadBalancerSku: "Standard",
						NetworkPlugin:   "kubenet",
					},
					Tags: Tags{
						Environment: "Development",
					},
				},
			},
		},
	}
	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/aks/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/aks/"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func create_gke(info Cli) {

	cluster := "cluster"
	num := 1
	// data := make([]Cluster_info, 1)

	// data[0].Project_id = "keti-container"
	// data[0].Cluster_name = info.ClusterName
	// data[0].Region = "us-central1-a"
	// data[0].Gke_num_nodes = uint64(num)

	// doc, _ := json.Marshal(data)

	// fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+".tfvars.json", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	send_js_cluster := TF{
		Resource: &Resource{
			Google_container_cluster: &map[string]Cluster_type{
				info.ClusterName: {
					Name:                     info.ClusterName,
					Location:                 "us-central1-a",
					Remove_default_node_pool: "true",
					Initial_node_count:       num,
					Life_Cycle: Life_Cycle{
						Ignore_Changes: []string{"ip_allocation_policy"},
					},
				},
			},
		},
	}

	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	// src, _ := json.Marshal([]byte(string(resource)))

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	// cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/gke/cluster"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createClusterCmd)
	createCmd.AddCommand(createNodeCmd)
	createCmd.AddCommand(uniCmd)

	uniCmd.Flags().String("pod-name", "", "input your uni-pod name")
	createClusterCmd.Flags().String("platform", "", "input your platform name")
	createClusterCmd.Flags().String("cluster-name", "", "input your cluster name")

	createNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	createNodeCmd.Flags().String("platform", "", "input your platform name")
	createNodeCmd.Flags().String("cluster-name", "", "input your cluster name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
