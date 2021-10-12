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
	"net/http"

	mappingTable "Hybrid_Cluster/hcp-apiserver/converter"

	"github.com/spf13/cobra"
)

// unjoinCmd represents the unjoin command
var unjoinCmd = &cobra.Command{
	Use:   "unjoin",
	Short: "A brief description of your command",
	Long: ` 
NAME 
	hybridctl unjoin PLATFORM CLUSTER

DESCRIPTION
	
	>> cluster unjoin PLATFORM CLUSTER <<


	PLATFORM means the Kubernetes platform of the cluster to unjoin.
	The types of platforms offered are as follows.

	- aks   azure kubernetes service
	- eks   elastic kubernetes service
	- gke   google kuberntes engine

	* PLATFORM mut be written in LOWERCASE letters

	CLUSTER means the name of the cluster on the specified platform.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) < 2 {
			fmt.Println("Run 'hybridctl unjoin --help' to view all commands")
		} else {
			switch args[0] {
			case "aks":
				fallthrough
			case "eks":
				fallthrough
			case "gke":
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])
				cli := mappingTable.ClusterInfo{
					PlatformName: args[0],
					ClusterName:  args[1]}
				unjoin(cli)
			default:
				fmt.Println("Run 'hybridctl join --help' to view all commands")
			}
		}
	},
}

func unjoin(info mappingTable.ClusterInfo) {
	httpPostUrl := "http://localhost:8080/unjoin"
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
}

func init() {
	RootCmd.AddCommand(unjoinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
