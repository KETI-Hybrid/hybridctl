package cmd

import (
	apiserverutil "Hybrid_Cloud/hcp-apiserver/pkg/util"
	"Hybrid_Cloud/hybridctl/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sourcerepo/v1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func HTTPPostRequest(input interface{}, httpPostUrl string) []byte {
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}

func PrintServerConfig(resp containerpb.ServerConfig) {
	//	var field string
	fmt.Println("channels:")
	for _, c := range resp.Channels {
		fmt.Println("- channel:", c.GetChannel())
		fmt.Printf("  defaultVersion: %s\n", c.GetDefaultVersion())
		fmt.Println("  validVersions:")
		for _, j := range c.GetValidVersions() {
			fmt.Println("  - ", j)
		}
	}

	fmt.Println("defaultClusterVersion: ", resp.DefaultClusterVersion)
	fmt.Println("defaultImageType: ", resp.DefaultImageType)

	fmt.Println("validImageTypes:")
	for _, c := range resp.ValidImageTypes {
		fmt.Println("- ", c)
	}

	fmt.Println("validMasterVersions:")
	for _, c := range resp.ValidMasterVersions {
		fmt.Println("- ", c)
	}

	fmt.Println("validNodeVersions:")
	for _, c := range resp.ValidNodeVersions {
		fmt.Println("- ", c)
	}
}

func GetServerConfig(input *containerpb.GetServerConfigRequest) {
	// input := &containerpb.GetServerConfigRequest{
	// 	ProjectId: "keti-container",
	// 	Zone:      "us-central1-a",
	// }
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/getServerConfig"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	checkErr(err)

	var output apiserverutil.Output
	json.Unmarshal(bytes, &output)
	if output.Stderr != nil {
		fmt.Println(string(output.Stderr))
	}

	if output.Stdout != nil {
		stdout := output.Stdout
		var resp containerpb.ServerConfig
		json.Unmarshal(stdout, &resp)
		fmt.Printf("Fetching server config for %s\n", input.Zone)
		PrintServerConfig(resp)
	}
}

func RollbackNodePoolUpgrade(input *containerpb.RollbackNodePoolUpgradeRequest) {
	// input := &containerpb.RollbackNodePoolUpgradeRequest{
	// 	ProjectId: "keti-container",
	// 	Zone:      "us-central1-a",
	// 	ClusterId: "hcp-cluster",
	// 	Name:      "pool-1",
	// }
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/rollbackNodePoolUpgrade"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	checkErr(err)

	var output apiserverutil.Output
	json.Unmarshal(bytes, &output)
	if output.Stderr != nil {
		fmt.Println(string(output.Stderr))
	}

	if output.Stdout != nil {
		stdout := output.Stdout
		var resp containerpb.Operation
		json.Unmarshal(stdout, &resp)
		fmt.Printf("Updated [%s]\n", resp.TargetLink)
		fmt.Printf("operationId: %s\nprojectId: %s\nzone: %s\n", resp.GetName(), resp.GetZone(), input.GetProjectId())
	}
}

type Docker struct {
	AUTHORIZE_ONLY bool
	DOCKER_HOST    string
	SERVER         string
}

func (d *Docker) Docker() {
	d = &Docker{
		AUTHORIZE_ONLY: false,
	}
	httpPostUrl := "http://localhost:3080/gke/docker"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, d)
	checkErr(err)
	util.PrintOutput(bytes)
}

func GetProjectConfig() {

	ctx := context.TODO()
	sourcerepoService, err := sourcerepo.NewService(ctx, option.WithCredentialsFile("/root/hcp-key.json"))
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("gcloud", "config", "get-value", "project")
	data, err := apiserverutil.GetOutput(cmd)
	if err != nil {
		fmt.Println(err)
	}

	var output apiserverutil.Output
	json.Unmarshal(data, &output)
	projectsService := sourcerepo.NewProjectsService(sourcerepoService)
	fmt.Printf("projectsService: %v\n", projectsService)

	project_id := string(output.Stdout)
	project_id = strings.TrimSuffix(project_id, "\n")
	project_name := "projects/" + project_id

	call := projectsService.GetConfig(project_name)
	resp, err := call.Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Name)

}

/*
// https://pkg.go.dev/cloud.google.com/go/pubsub#section-readme
// https://github.com/googleapis/google-cloud-go/blob/5a2ed6b2cd1c304e0f59daa29959863bff9b5c29/pubsub/example_test.go
func UpdateProjectConfig() {

	ctx := context.TODO()
	sourcerepoService, err := sourcerepo.NewService(ctx, option.WithCredentialsFile("/root/hcp-key.json"))
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("gcloud", "config", "get-value", "project")
	data, err := apiserverutil.GetOutput(cmd)
	if err != nil {
		fmt.Println(err)
	}

	var output apiserverutil.Output
	json.Unmarshal(data, &output)
	projectsService := sourcerepo.NewProjectsService(sourcerepoService)

	project_id := string(output.Stdout)
	project_id = strings.TrimSuffix(project_id, "\n")
	project_name := "projects/" + project_id

	fmt.Println(project_name)
	var req *sourcerepo.UpdateProjectConfigRequest
	current := req.ProjectConfig.PubsubConfigs
	current["new"] = "new_topic"
	// gke.SetGKERequest(r, &req)

	call := projectsService.UpdateConfig(project_name, req)
	resp, err := call.Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Name)

}
*/

type SetProperty struct {
	SECTION  string
	PROPERTY string
	VALUE    string
}

func ConfigSet() {
	input := SetProperty{
		SECTION:  "compute",
		PROPERTY: "zone",
		VALUE:    "us-central1-a",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONFIG_PATH + "/set"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	checkErr(err)
	util.PrintOutput(bytes)
}
