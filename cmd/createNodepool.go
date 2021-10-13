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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createNodepoolCmd represents the createNodepool command
var createNodepoolCmd = &cobra.Command{
	Use:   "createNodepool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		cli := Cli{args[0], args[1]}
		fmt.Println("createNodepool called")
		if len(args) == 0 {
			fmt.Println("Run 'hybridctl createNodepool --help' to view all commands")
		} else if args[0] == "gke" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl createNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				createNodepool_gke(args[1])
			}
		} else if args[0] == "aks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl createNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				createNodepool_aks(cli)
			}
		} else if args[0] == "eks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl createNodepool --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name: ", args[0])
				fmt.Printf("Cluster Name: %s\n", args[1])
				createNodepool_eks(cli)
			}
		}

	},
}

type TF_NodePool struct {
	NodePool_Resource *NodePool_Resource `json:"resource"`
}

type NodePool_Resource struct {
	Google_container_node_pool *map[string]Node_pool_type `json:"google_container_node_pool"`
}

type Node_pool_type struct {
	Name        string       `json:"name"`
	Location    string       `json:"location"`
	Cluster     string       `json:"cluster"`
	Node_count  int          `json:"node_count"`
	Node_config *Node_config `json:"node_config"`
}

type Labels struct {
	Env string `json:"env"`
}
type Node_config struct {
	Oauth_scopes []string  `json:"oauth_scopes"`
	Labels       *Labels   `json:"labels"`
	Machine_type string    `json:"machine_type"`
	Tags         []string  `json:"tags"`
	Metadata     *Metadata `json:"metadata"`
}

type Metadata struct {
	Disable_legacy_endpoints string `json:"disable-legacy-endpoints"`
}

//-------------------------eks structure--------------

type TF_EKS_Nopepool struct {
	EksResourceNode EksResourceNode `json:"resource"`
}
type EksLabels struct {
	Role string `json:"role"`
}
type ScalingConfig struct {
	DesiredSize string `json:"desired_size"`
	MinSize     string `json:"min_size"`
	MaxSize     string `json:"max_size"`
}
type EksTags struct {
	Name string `json:"Name"`
}
type EksClusterNode struct {
	ClusterName   string        `json:"cluster_name"`
	NodeGroupName string        `json:"node_group_name"`
	NodeRoleArn   string        `json:"node_role_arn"`
	SubnetIds     string        `json:"subnet_ids"`
	InstanceTypes []string      `json:"instance_types"`
	DiskSize      string        `json:"disk_size"`
	EksLabels     EksLabels     `json:"labels"`
	ScalingConfig ScalingConfig `json:"scaling_config"`
	DependsOn     []string      `json:"depends_on"`
	EksTags       EksTags       `json:"tags"`
}

type EksResourceNode struct {
	EksResourceNode *map[string]EksClusterNode `json:"aws_eks_node_group"`
}

//-------------------------eks structure end---------

//-------------------------aks structure--------------

type TF_AKS_Nodepool struct {
	AksNodepoolResource AksNodepoolResource `json:"resource"`
}
type AksNodeTags struct {
	Environment string `json:"environment"`
}
type AksNodepool struct {
	Name                string      `json:"name"`
	KubernetesClusterID string      `json:"kubernetes_cluster_id"`
	VMSize              string      `json:"vm_size"`
	NodeCount           string      `json:"node_count"`
	AksNodeTags         AksNodeTags `json:"tags"`
}

type AksNodepoolResource struct {
	AzurermKubernetesClusterNodePool *map[string]AksNodepool `json:"azurerm_kubernetes_cluster_node_pool"`
}

//-------------------------aks structure end---------

func createNodepool_eks(info Cli) {
	send_js_cluster := TF_EKS_Nopepool{
		EksResourceNode: EksResourceNode{
			EksResourceNode: &map[string]EksClusterNode{
				info.ClusterName: {
					ClusterName:   "${aws_eks_cluster." + info.ClusterName + ".name}",
					NodeGroupName: "terraform-eks-m5-large",
					NodeRoleArn:   "${aws_iam_role.terraform-eks-node.arn}",
					SubnetIds:     "${aws_subnet.terraform-eks-private-subnet[*].id}",
					InstanceTypes: []string{"m5.large"},
					DiskSize:      "50",
					EksLabels: EksLabels{
						Role: "terraform-eks-m5-large",
					},
					ScalingConfig: ScalingConfig{
						DesiredSize: "3",
						MinSize:     "1",
						MaxSize:     "3",
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/eks/"+info.ClusterName+"nodepool.tf.json", []byte(string(send)), os.FileMode(0644))
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

func createNodepool_aks(info Cli) {

	send_js_cluster := TF_AKS_Nodepool{
		AksNodepoolResource: AksNodepoolResource{
			AzurermKubernetesClusterNodePool: &map[string]AksNodepool{
				info.ClusterName: {
					Name:                "default",
					KubernetesClusterID: "${azurerm_kubernetes_cluster." + info.ClusterName + ".id}",
					VMSize:              "Standard_DS2_v2",
					NodeCount:           "1",
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/aks/"+info.ClusterName+"nodepool.tf.json", []byte(string(send)), os.FileMode(0644))
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

func createNodepool_gke(clusterName string) {
	cluster := "cluster"
	num := 1

	send_js_nodePool := TF_NodePool{
		NodePool_Resource: &NodePool_Resource{
			Google_container_node_pool: &map[string]Node_pool_type{
				clusterName + "_nodes": {
					Name:       "${google_container_cluster." + clusterName + ".name}-node-pool",
					Location:   "us-central1-a",
					Cluster:    "${google_container_cluster." + clusterName + ".name}",
					Node_count: num,
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+clusterName+"nodePool"+".tf.json", []byte(string(send)), os.FileMode(0644))
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

func init() {
	RootCmd.AddCommand(createNodepoolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createNodepoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createNodepoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
