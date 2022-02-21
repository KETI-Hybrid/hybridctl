package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	cobrautil "Hybrid_Cluster/hybridctl/util"
	resourcev1alpha1 "Hybrid_Cluster/pkg/apis/resource/v1alpha1"
	resourcev1alpha1clientset "Hybrid_Cluster/pkg/client/resource/v1alpha1/clientset/versioned"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

type Resource struct {
	TargetCluster string
	RealResource  interface{}
}

// CreateCmd represents the Create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		CreateResource()
	},
}

func CreateResource() {

	fmt.Println("1")
	yaml, err := ReadFile()
	if err != nil {
		println(err)
		return
	}
	obj, gvk, err := GetObject(yaml)
	if err != nil {
		println(err)
		return
	}

	RequestCreateResource(obj, gvk)
}

func ReadFile() ([]byte, error) {
	fmt.Println("2")
	file_name := cobrautil.Option_file

	if file_name == "" {
		fmt.Println("Run 'hybridctl create --help' to view all commands")
		return nil, nil
	}

	yaml, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return yaml, err
}

func GetObject(yaml []byte) (runtime.Object, *schema.GroupVersionKind, error) {
	fmt.Println("3")
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, gvk, err := decode([]byte(yaml), nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return obj, gvk, err
}

func RequestCreateResource(obj runtime.Object, gvk *schema.GroupVersionKind) ([]byte, error) {
	fmt.Println("4")
	LINK := "http://10.0.5.86:8080/resources"

	// check context flag
	flag_context := cobrautil.Option_context
	var target_cluster string
	var resource Resource

	if flag_context == "" {
		target_cluster = "undefined"
	} else {
		target_cluster = flag_context
	}

	// match obj kind
	switch gvk.Kind {
	case "Deployment":
		LINK += "/deployment"
		real_resource := obj.(*appsv1.Deployment)
		fmt.Println(real_resource)
		resource.TargetCluster = target_cluster
		resource.RealResource = real_resource
	}

	fmt.Println(resource)
	fmt.Println(LINK)
	bytes, err := util.GetResponseBody("POST", LINK, &resource)
	if err != nil {
		fmt.Println(err)
	}

	return bytes, err
}

func CreateHCPDeployment(resource *appsv1.Deployment) {
	hcp_resource := new(resourcev1alpha1.HCPDeployment)
	clientset, err := resourcev1alpha1clientset.NewForConfig(master_config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check context flag
	flag_context := cobrautil.Option_context
	var target_cluster string
	if flag_context == "" {
		target_cluster = "Undefined"
	} else {
		target_cluster = flag_context
	}

	// typemeta
	hcp_resource.TypeMeta.APIVersion = "hcp.crd.com/v1alpha1"
	hcp_resource.TypeMeta.Kind = "HCPDeployment"

	// objectmeta
	hcp_resource.ObjectMeta.Name = resource.Name

	// spec
	hcp_resource.Spec.TargetCluster = target_cluster
	hcp_resource.Spec.Replicas = resource.Spec.Replicas
	hcp_resource.Spec.Selector = resource.Spec.Selector
	hcp_resource.Spec.Template = resource.Spec.Template
	hcp_resource.Spec.RealDeploymentSpec = resource.Spec

	r, err := clientset.HcpV1alpha1().HCPDeployments("hcp").Create(context.TODO(), hcp_resource, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("success to create hcp_resource %s \n", r.Name)
	}
}

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.Flags().StringVarP(&cobrautil.Option_file, "file", "f", "", "FILENAME")
	CreateCmd.MarkFlagRequired("file")
	CreateCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "", "", "CLUSTERNAME")
}
