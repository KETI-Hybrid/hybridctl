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
	// "hybridctl/pkg/handler"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"hcp-analytic-engine/pkg/handler"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get Kubernetes engine reousrce information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

	},
}

var metricCmd = &cobra.Command{
	Use:   "metric",
	Short: "get metric information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

	},
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "get node metric information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("get node metric called")
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

		// var jsonarray PodMetric
		var cluster_list = []string{cli.ClusterName}
		// cluster_list 생성 우선 gke-cluster, aks-cluster, eks-cluster 가 저장되어있다고 가정
		var podNum = []int{21}

		if platform_name == "aks" {
			fmt.Println("call create_aks func")
			create_aks(cli)
		} else if platform_name == "eks" {
			fmt.Println("call create_eks func")
			create_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call create_gke func")
			Print_NodeMetric(podNum[0], cluster_list[0])
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}
	},
}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "get pod metric information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("get pod metric called")

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
		// var jsonarray PodMetric
		var cluster_list = []string{cli.ClusterName}
		// cluster_list 생성 우선 gke-cluster, aks-cluster, eks-cluster 가 저장되어있다고 가정
		var podNum = []int{21}

		if platform_name == "aks" {
			fmt.Println("call create_aks func")
			create_aks(cli)
		} else if platform_name == "eks" {
			fmt.Println("call create_eks func")
			create_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call create_gke func")
			Print_PodMetric(podNum[0], cluster_list[0])
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}
	},
}

type NodeMetric struct {
	Nodemetrics []struct {
		Time    time.Time `json:"time"`
		Cluster string    `json:"cluster"`
		Node    string    `json:"node"`
		CPU     struct {
			CPUUsageNanoCores string `json:"CPUUsageNanoCores"`
		} `json:"cpu"`
		Memory struct {
			MemoryAvailableBytes  string `json:"MemoryAvailableBytes"`
			MemoryUsageBytes      string `json:"MemoryUsageBytes"`
			MemoryWorkingSetBytes string `json:"MemoryWorkingSetBytes"`
		} `json:"memory"`
		Fs struct {
			FsAvailableBytes string `json:"FsAvailableBytes"`
			FsCapacityBytes  string `json:"FsCapacityBytes"`
			FsUsedBytes      string `json:"FsUsedBytes"`
		} `json:"fs"`
		Network struct {
			NetworkRxBytes string `json:"NetworkRxBytes"`
			NetworkTxBytes string `json:"NetworkTxBytes"`
		} `json:"network"`
	} `json:"nodemetrics"`
}

type PodMetric struct {
	Podmetrics []struct {
		Time      time.Time `json:"time"`
		Cluster   string    `json:"cluster"`
		Namespace string    `json:"namespace"`
		Node      string    `json:"node"`
		Pod       string    `json:"pod"`
		CPU       struct {
			CPUUsageNanoCores string `json:"CPUUsageNanoCores"`
		} `json:"cpu"`
		Memory struct {
			MemoryAvailableBytes  string `json:"MemoryAvailableBytes"`
			MemoryUsageBytes      string `json:"MemoryUsageBytes"`
			MemoryWorkingSetBytes string `json:"MemoryWorkingSetBytes"`
		} `json:"memory"`
		Fs struct {
			FsAvailableBytes string `json:"FsAvailableBytes"`
			FsCapacityBytes  string `json:"FsCapacityBytes"`
			FsUsedBytes      string `json:"FsUsedBytes"`
		} `json:"fs"`
		Network struct {
			NetworkRxBytes string `json:"NetworkRxBytes"`
			NetworkTxBytes string `json:"NetworkTxBytes"`
		} `json:"network"`
	} `json:"podmetrics"`
}

func Print_NodeMetric(podNum int, clusterName string) {
	var jsonarray NodeMetric
	// ns := "hcp"
	jsonByteArray := handler.GetResource(podNum, clusterName, "nodes")
	stringArray := string(jsonByteArray[:])

	if err := json.Unmarshal([]byte(stringArray), &jsonarray); err != nil {
		panic(err)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Print Cluster : ", clusterName)
	fmt.Println("")

	for i := 0; i < 3; i++ {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("[", i+1, "] Node:", jsonarray.Nodemetrics[i].Node, " Metric Information")
		fmt.Println("Time              :", jsonarray.Nodemetrics[i].Time)
		fmt.Println("Cluster           :", jsonarray.Nodemetrics[i].Cluster)
		fmt.Println("Node              :", jsonarray.Nodemetrics[i].Node)
		fmt.Println("CpuUsage          :", jsonarray.Nodemetrics[i].CPU.CPUUsageNanoCores)
		fmt.Println("MemoryAvailable   :", jsonarray.Nodemetrics[i].Memory.MemoryAvailableBytes)
		fmt.Println("MemoryUsage       :", jsonarray.Nodemetrics[i].Memory.MemoryUsageBytes)
		fmt.Println("MemWorkingSet     :", jsonarray.Nodemetrics[i].Memory.MemoryWorkingSetBytes)
		fmt.Println("FsAvailable       :", jsonarray.Nodemetrics[i].Fs.FsAvailableBytes)
		fmt.Println("FsCapacity        :", jsonarray.Nodemetrics[i].Fs.FsCapacityBytes)
		fmt.Println("FsUsed            :", jsonarray.Nodemetrics[i].Fs.FsUsedBytes)
		fmt.Println("NetworkRx         :", jsonarray.Nodemetrics[i].Network.NetworkRxBytes)
		fmt.Println("NetworkTx         :", jsonarray.Nodemetrics[i].Network.NetworkTxBytes)
	}
	fmt.Println("------------------------------------------------------------------")
}

func Print_PodMetric(podNum int, clusterName string) {
	var jsonarray PodMetric
	// ns := "hcp"
	jsonByteArray := handler.GetResource(podNum, clusterName, "pods")
	stringArray := string(jsonByteArray[:])
	if err := json.Unmarshal([]byte(stringArray), &jsonarray); err != nil {
		panic(err)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Print Cluster : ", clusterName)
	fmt.Println("")

	//파드 개수를 가져와서 for문의 변수로 넣어주어야 함

	for i := 0; i < podNum; i++ {

		fmt.Println("------------------------------------------------------------------")
		fmt.Println("[", i+1, "] Pod:", jsonarray.Podmetrics[i].Pod, " Metric Information")
		fmt.Println("Time              :", jsonarray.Podmetrics[i].Time)
		fmt.Println("Cluster           :", jsonarray.Podmetrics[i].Cluster)
		fmt.Println("Namespace         :", jsonarray.Podmetrics[i].Namespace)
		fmt.Println("Node              :", jsonarray.Podmetrics[i].Node)
		fmt.Println("PodMetric         :", jsonarray.Podmetrics[i].Pod)
		fmt.Println("CpuUsage          :", jsonarray.Podmetrics[i].CPU.CPUUsageNanoCores)
		fmt.Println("MemoryAvailable   :", jsonarray.Podmetrics[i].Memory.MemoryAvailableBytes)
		fmt.Println("MemoryUsage       :", jsonarray.Podmetrics[i].Memory.MemoryUsageBytes)
		fmt.Println("MemWorkingSet     :", jsonarray.Podmetrics[i].Memory.MemoryWorkingSetBytes)
		fmt.Println("FsAvailable       :", jsonarray.Podmetrics[i].Fs.FsAvailableBytes)
		fmt.Println("FsCapacity        :", jsonarray.Podmetrics[i].Fs.FsCapacityBytes)
		fmt.Println("FsUsed            :", jsonarray.Podmetrics[i].Fs.FsUsedBytes)
		fmt.Println("NetworkRx         :", jsonarray.Podmetrics[i].Network.NetworkRxBytes)
		fmt.Println("NetworkTx         :", jsonarray.Podmetrics[i].Network.NetworkTxBytes)
		// fmt.Print(jsonarray.Podmetrics[i])

	}
	fmt.Println("------------------------------------------")
}
func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.AddCommand(metricCmd)
	metricCmd.AddCommand(nodeCmd)
	metricCmd.AddCommand(podCmd)

	podCmd.Flags().String("platform", "", "input your platform name")
	podCmd.Flags().String("cluster-name", "", "input your cluster name")

	nodeCmd.Flags().String("platform", "", "input your platform name")
	nodeCmd.Flags().String("cluster-name", "", "input your cluster name")

	os.Setenv("INFLUX_IP", "10.0.5.83")
	os.Setenv("INFLUX_PORT", "31051")
	os.Setenv("INFLUX_USERNAME", "root")
	os.Setenv("INFLUX_PASSWORD", "root")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
