package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DeleteCmd represents the Delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `hybridctl delete deployment <name> -n <namespace> --context <cluster_name>
	hybridctl delete -f deployment `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		file_name := util.Option_file
		if file_name == "" {
			if len(args) < 2 {
				fmt.Println("hybridctl delete --help")
			} else {
				util.Option_Resource = args[0]
				util.Option_Name = args[1]

				DeleteResource()
			}
		} else {
			DeleteResource()
		}
	},
}

func DeleteResource() {

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

	RequestDeleteResource(obj, gvk)
}

func RequestDeleteResource(obj runtime.Object, gvk *schema.GroupVersionKind) ([]byte, error) {
	fmt.Println("4")
	LINK := "http://10.0.5.86:8080/resources"

	// check context flag
	flag_context := util.Option_context
	var target_cluster string
	var resource Resource

	if flag_context == "" {
		target_cluster = "kube-master"
	} else {
		target_cluster = flag_context
	}

	// match obj kind
	switch gvk.Kind {
	case "Deployment":
		LINK += "/deployments"

	}

	LINK += "/?cluster=" + target_cluster
	LINK += "&namespaces=" + util.Option_Namespace
	LINK += "&name=" + util.Option_Name
	fmt.Println(resource)
	fmt.Println(LINK)
	bytes, err := util.GetResponseBody("DELETE", LINK, &resource)
	if err != nil {
		fmt.Println(err)
	}

	return bytes, err
}

func init() {
	RootCmd.AddCommand(DeleteCmd)
	DeleteCmd.Flags().StringVarP(&util.Option_file, "file", "f", "", "FILENAME")
	DeleteCmd.MarkFlagRequired("file")
	DeleteCmd.Flags().StringVarP(&util.Option_context, "context", "", "", "CLUSTERNAME")
}
