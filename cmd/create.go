package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/KETI-Hybrid/hcp-pkg/apis/resource/v1alpha1"
	resourcev1alpha1scheme "github.com/KETI-Hybrid/hcp-pkg/client/resource/v1alpha1/clientset/versioned/scheme"

	"github.com/KETI-Hybrid/hybridctl-v1/pkg/nks"
	cobrautil "github.com/KETI-Hybrid/hybridctl-v1/util"

	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	klog "k8s.io/klog/v2"
)

type HCPResource struct {
	TargetCluster string
	RealResource  interface{}
}

// createCmd represents the create command
var Createcmd = &cobra.Command{
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

var CreateResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "A brief description of your command",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		CreateResource()
	},
}

func CreateResource() {
	yaml, err := ReadFile()
	if err != nil {
		println(err)
		return
	}

	fmt.Println("here")
	obj, gvk, err := GetObject(yaml)
	if err != nil {
		println(err)
		return
	}

	RequestCreateResource(obj, gvk)
}

func ReadFile() ([]byte, error) {
	file_name := cobrautil.Option_file
	yaml, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return yaml, err
}

func GetObject(yaml []byte) (runtime.Object, *schema.GroupVersionKind, error) {

	utilruntime.Must(resourcev1alpha1scheme.AddToScheme(scheme.Scheme))
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, gvk, err := decode([]byte(yaml), nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return obj, gvk, err
}

func RequestCreateResource(obj runtime.Object, gvk *schema.GroupVersionKind) ([]byte, error) {
	LINK := "/resources"

	// check context flag
	flag_context := cobrautil.Option_context
	var target_cluster string
	var resource HCPResource

	if flag_context == "" {
		target_cluster = ""
	} else {
		target_cluster = flag_context
	}

	// match obj kind
	switch gvk.Kind {
	case "Deployment":
		real_resource := obj.(*appsv1.Deployment)
		namespace := real_resource.Namespace
		if namespace == "" {
			namespace = "default"
		}
		resource.TargetCluster = target_cluster
		resource.RealResource = real_resource
		LINK += "/namespaces/" + namespace + "/deployments"
	case "Pod":
		LINK += "/pod"
		real_resource := obj.(*v1.Pod)
		resource.TargetCluster = target_cluster
		resource.RealResource = real_resource
	case "HCPHybridAutoScaler":
		real_resource := obj.(*v1alpha1.HCPHybridAutoScaler)
		namespace := "hcp" // hcp로 고정
		real_resource.Namespace = namespace
		resource.TargetCluster = "master"
		resource.RealResource = real_resource
		LINK += "/namespaces/" + namespace + "/hcphybridautoscalers"
	}

	fmt.Println(LINK)
	bytes, err := cobrautil.GetResponseBody("POST", LINK, &resource)
	if err != nil {
		fmt.Println(err)
	}

	return bytes, err
}

// create common cli

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
		} else if platform_name == "nks" {
			klog.Infoln("call create_nks func")
			nks.NksCreateCluster(cli.ClusterName)
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_LCW/Hybrid_Cloud/terraform/eks/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../../../terraform/eks/"

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

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cloud/terraform/gke/"+cluster+"/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cloud/terraform/aks/"+info.ClusterName+"-"+info.NodeName+".tf.json", []byte(string(send)), os.FileMode(0644))
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_LCW/Hybrid_Cloud/terraform/eks/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../../../terraform/eks"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		klog.Infoln(err)
		klog.Infoln(string(output))
	} else {
		fmt.Println(string(output))
		klog.Infoln(string(output))
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

	err = ioutil.WriteFile("/root/go/src/Hybrid_LCW/Hybrid_Cloud/terraform/aks/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../../../terraform/aks/"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		klog.Infoln(err)
		klog.Infoln(string(output))
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

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cloud/terraform/gke/"+cluster+"/"+info.ClusterName+".tfvars.json", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

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

	err = ioutil.WriteFile("/root/go/src/Hybrid_LCW/Hybrid_Cloud/terraform/gke/"+cluster+"/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
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
	RootCmd.AddCommand(Createcmd)
	Createcmd.AddCommand(CreateResourceCmd)
	Createcmd.AddCommand(createClusterCmd)
	Createcmd.AddCommand(createNodeCmd)
	Createcmd.AddCommand(uniCmd)
	CreateResourceCmd.Flags().StringVarP(&cobrautil.Option_file, "file", "f", "", "FILENAME")
	CreateResourceCmd.MarkFlagRequired("file")
	CreateResourceCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "", "", "CLUSTERNAME")

	uniCmd.Flags().String("pod-name", "", "input your uni-pod name")
	createClusterCmd.Flags().String("platform", "", "input your platform name")
	createClusterCmd.Flags().String("cluster-name", "", "input your cluster name")

	createNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	createNodeCmd.Flags().String("platform", "", "input your platform name")
	createNodeCmd.Flags().String("cluster-name", "", "input your cluster name")
}
