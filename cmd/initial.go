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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/KETI-Hybrid/hybridctl-v1/util"

	"github.com/spf13/cobra"
)

const CONFIGFILE_DIR string = "/root/.hcp/config"

// joinCmd represents the join command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long: ` 
NAME 
	hybridctl configure PLATFORM

DESCRIPTION
	
	>> hybridctl configure PLATFORM <<


	PLATFORM means the Kubernetes platform of the cluster to config.
	The types of platforms offered are as follows.

	- default 
	- aks   azure kubernetes service
	- eks   elastic kubernetes service
	- gke   google kuberntes engine

	* PLATFORM must be written in LOWERCASE letters
	`,
	Run: func(cmd *cobra.Command, args []string) {

		var hcp_config util.HCPConfig
		var max_cluster_cpu int
		var max_cluster_mem int
		var default_node_option string
		var extra int

		//.hcp 디렉터리 확인 및 생성
		if _, err := os.Stat("/root/.hcp/"); errors.Is(err, os.ErrNotExist) {
			if err = os.MkdirAll("/root/.hcp/", 0666); err != nil {
				fmt.Println(err)
				return
			}
		}

		// .hcp/config 파일 확인 및 생성
		if _, err := os.Stat(CONFIGFILE_DIR); errors.Is(err, os.ErrNotExist) {
			if _, err = os.Create(CONFIGFILE_DIR); err != nil {
				fmt.Println(err)
				return
			}
		}

		hcp_config.Section = args[0]

		fmt.Printf("Enter the maximum number of CPUs to allocate in the cluster. [NanoCores] : ")
		fmt.Scanln(&max_cluster_cpu)
		hcp_config.MaxClusterCpu = max_cluster_cpu

		fmt.Printf("Enter the maximum amount of memory to allocate in the cluster.. : ")
		fmt.Scanln(&max_cluster_mem)
		hcp_config.MaxClusterMem = max_cluster_mem

	CON3:
		fmt.Printf("Enter the node option to use as default [ Low / Middle / High ]: ")
		fmt.Scanln(&default_node_option)
		if !(default_node_option == "Low" || default_node_option == "Middle" || default_node_option == "High") {
			goto CON3
		}
		hcp_config.DefaultNodeOption = default_node_option

		fmt.Printf("Enter the percentage of free resources to use when automatically creating a node. : ")
		fmt.Scanln(&extra)
		hcp_config.Extra = extra

		err := UpdateHCPConfigFile(hcp_config)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func UpdateHCPConfigFile(new_hcp_config util.HCPConfig) error {

	var exist bool = false

	command := &exec.Cmd{
		Path:   "/root/go/src/Hybrid_LCW/github.com/KETI-Hybrid/hybridctl-v1/pkg/hybridctl/sh/parse-config.sh",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err := command.Start()
	if err != nil {
		return err
	}

	err = command.Wait()
	if err != nil {
		return err
	}

	old_hcp_config, err := LoadHCPConfig("/root/go/src/Hybrid_LCW/github.com/KETI-Hybrid/hybridctl-v1/pkg/hybridctl/tmp.json")
	if err != nil {
		return err
	}

	for _, config := range old_hcp_config {
		if config.Section == new_hcp_config.Section {
			exist = true
			break
		}
	}

	if exist {
		err = DeleteHCPConfig(new_hcp_config.Section)
		if err != nil {
			return err
		}

		err = CreateHCPConfig(new_hcp_config)
		if err != nil {
			return err
		}

	} else {
		err := CreateHCPConfig(new_hcp_config)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteHCPConfig(section string) error {

	var arguments []string
	arguments = append(arguments, section)

	command := &exec.Cmd{
		Path:   "/root/go/src/Hybrid_LCW/github.com/KETI-Hybrid/hybridctl-v1/pkg/hybridctl/sh/delete-config.sh",
		Args:   arguments,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err := command.Start()
	if err != nil {
		return err
	}

	err = command.Wait()
	if err != nil {
		return err
	}

	return nil
}

func CreateHCPConfig(hcp_config util.HCPConfig) error {
	f, err := os.OpenFile(CONFIGFILE_DIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	var str string
	str += "[" + hcp_config.Section + "]" + "\n"
	str += "max_cluster_cpu" + "=" + strconv.Itoa(hcp_config.MaxClusterCpu) + "\n"
	str += "max_cluster_mem" + "=" + strconv.Itoa(hcp_config.MaxClusterMem) + "\n"
	str += "default_node_option" + "=" + hcp_config.DefaultNodeOption + "\n"
	str += "extra" + "=" + strconv.Itoa(hcp_config.Extra) + "\n"
	_, err = f.WriteString(str)

	return err
}

func LoadHCPConfig(configFile string) ([]util.HCPConfig, error) {

	var hcp_config []util.HCPConfig
	// we have a config so parse it.
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &hcp_config)
	if err != nil {
		return nil, err
	}
	fmt.Println(hcp_config)

	return hcp_config, nil
}

func init() {
	RootCmd.AddCommand(configureCmd)
}
