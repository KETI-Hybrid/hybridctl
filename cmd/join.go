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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	resourcev1alpha1 "Hybrid_Cluster/apis/clusterRegister/v1alpha1"
	clusterRegisterv1alpha1 "Hybrid_Cluster/clientset/clusterRegister/v1alpha1"

	cobrautil "Hybrid_Cluster/hybridctl/util"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var checkAKS, checkEKS, checkGKE = false, false, false
var master_config, mcerr = cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
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
	- cluster join
		hybridctl join <platformName> <clusterName>

		- platformName
			- aks (azure kubernetes service)
			- eks (elastic kubernetes service)
			- gke (google kuberntes engine)

	- cluster register
		hybridctl join register <platformName>

		- platformName
			- aks (azure kubernetes service)
			- eks (elastic kubernetes service)
			- gke (google kuberntes engine)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 || len(args) == 1 {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else if args[1] == "" {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else {
			switch args[0] {
			case "aks":
				fallthrough
			case "eks":
				fallthrough
			case "gke":
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])
				cli := Cli{args[0], args[1]}
				join(cli)
			case "register":
				master_config, _ := cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
				clusterRegisterClientSet, err := clusterRegisterv1alpha1.NewForConfig(master_config)
				var region string
				var clustername string
				if mcerr != nil {
					log.Println(err)
				}
				createPlatformNamespace()
				switch args[1] {
				case "aks":
					var resourcegroup string
					fmt.Printf("Enter cluster region : ")
					fmt.Scanln(&region)
					fmt.Printf("Enter resourcegroup : ")
					fmt.Scanln(&resourcegroup)
					fmt.Printf("Enter clustername : ")
					fmt.Scanln(&clustername)
					newclusterRegister := &resourcev1alpha1.ClusterRegister{
						TypeMeta: metav1.TypeMeta{
							Kind:       "ClusterRegister",
							APIVersion: "hcp.k8s.io/v1alpha1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      clustername,
							Namespace: "aks",
						},
						Spec: resourcev1alpha1.ClusterRegisterSpec{
							Clustername:   clustername,
							Region:        region,
							Platform:      "aks",
							Resourcegroup: resourcegroup,
						},
					}

					_, err = clusterRegisterClientSet.ClusterRegister("aks").Create(newclusterRegister)

					if err != nil {
						log.Println(err)
					}
				case "eks":
					fmt.Printf("Enter cluster region : ")
					fmt.Scanln(&region)
					fmt.Printf("Enter clustername : ")
					fmt.Scanln(&clustername)
					newclusterRegister := &resourcev1alpha1.ClusterRegister{
						TypeMeta: metav1.TypeMeta{
							Kind:       "ClusterRegister",
							APIVersion: "hcp.k8s.io/v1alpha1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      clustername,
							Namespace: "eks",
						},
						Spec: resourcev1alpha1.ClusterRegisterSpec{
							Clustername: clustername,
							Region:      region,
							Platform:    "eks",
						},
					}
					_, err = clusterRegisterClientSet.ClusterRegister("eks").Create(newclusterRegister)

					if err != nil {
						log.Println(err)
					}
				case "gke":
					var projectid string
					fmt.Printf("Enter projectid : ")
					fmt.Scanln(&projectid)
					fmt.Printf("Enter cluster region : ")
					fmt.Scanln(&region)
					fmt.Printf("Enter clustername : ")
					fmt.Scanln(&clustername)
					newclusterRegister := &resourcev1alpha1.ClusterRegister{
						TypeMeta: metav1.TypeMeta{
							Kind:       "ClusterRegister",
							APIVersion: "hcp.k8s.io/v1alpha1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      clustername,
							Namespace: "gke",
						},
						Spec: resourcev1alpha1.ClusterRegisterSpec{
							Clustername: clustername,
							Region:      region,
							Platform:    "gke",
							Projectid:   projectid,
						},
					}
					_, err = clusterRegisterClientSet.ClusterRegister("gke").Create(newclusterRegister)

					if err != nil {
						log.Println(err)
					}
				default:
					fmt.Println("Run 'hybridctl join --help' to view all commands")
				}
			default:
				fmt.Println("Run 'hybridctl join --help' to view all commands")
			}
		}
	},
}

func join(info Cli) {
	httpPostUrl := "http://10.0.5.83:8000/join"
	jsonData, _ := json.Marshal(&info)

	buff := bytes.NewBuffer(jsonData)
	request, _ := http.NewRequest("POST", httpPostUrl, buff)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
