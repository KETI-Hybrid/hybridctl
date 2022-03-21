package cmd

import (
	"Hybrid_Cloud/hybridctl/util"
	cobrautil "Hybrid_Cloud/hybridctl/util"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
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
		resource.TargetCluster = target_cluster
		resource.RealResource = real_resource
	case "Pod":
		LINK += "/pod"
		real_resource := obj.(*v1.Pod)
		resource.TargetCluster = target_cluster
		resource.RealResource = real_resource
	}

	fmt.Println(LINK)
	bytes, err := util.GetResponseBody("POST", LINK, &resource)
	if err != nil {
		fmt.Println(err)
	}

	return bytes, err
}

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.Flags().StringVarP(&cobrautil.Option_file, "file", "f", "", "FILENAME")
	CreateCmd.MarkFlagRequired("file")
	CreateCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "", "", "CLUSTERNAME")
}
