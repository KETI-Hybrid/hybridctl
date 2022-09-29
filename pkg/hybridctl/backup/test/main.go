package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	apiserverutil "hcp-apiserver/pkg/util"

	"hybridctl/util"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/option"
	"google.golang.org/api/sourcerepo/v1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

var GKE_CONTAINER_PATH = "/gke/container"
var GKE_AUTH_PATH = "/gke/auth"
var GKE_CONFIG_PATH = "/gke/config"

// images
type Images struct {
	SRC_IMAGE  string
	DEST_IMAGE string
	IMAGE_NAME string
}

func (i *Images) AddTag() {
	i = &Images{
		SRC_IMAGE:  "gcr.io/keti-container/busybox",
		DEST_IMAGE: "gcr.io/keti-container/busybox:mytag3",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/addTag"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (i *Images) Delete() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/delete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (i *Images) Describe() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/describe"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (i *Images) List() {
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (i *Images) ListTags() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/listTags"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (i *Images) UnTags() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox:mytag3",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/images/unTags"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

// gcloud auth configure-docker
func ConfigureDocker() {
	httpPostUrl := "http://localhost:3080" + GKE_AUTH_PATH + "/configureDocker"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

type Auth struct {
	CRED_FILE string
}

func (a *Auth) List() {
	httpPostUrl := "http://localhost:3080" + GKE_AUTH_PATH + "/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (a *Auth) Revoke() {
	httpPostUrl := "http://localhost:3080" + GKE_AUTH_PATH + "/revoke"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func (a *Auth) Login() {
	a = &Auth{
		CRED_FILE: "/root/hcp-key.json",
	}
	httpPostUrl := "http://localhost:3080" + GKE_AUTH_PATH + "/login"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, a)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
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

func GetServerConfig() {
	input := &containerpb.GetServerConfigRequest{
		ProjectId: "keti-container",
		Zone:      "us-central1-a",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/getServerConfig"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	if err != nil {
		log.Println(err)
		return
	}

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

func RollbackNodePoolUpgrade() {
	input := &containerpb.RollbackNodePoolUpgradeRequest{
		ProjectId: "keti-container",
		Zone:      "us-central1-a",
		ClusterId: "hcp-cluster",
		Name:      "pool-1",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/rollbackNodePoolUpgrade"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, input)
	if err != nil {
		log.Println(err)
		return
	}

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

func OperationsDescribe() {
	op := &containerpb.GetOperationRequest{
		ProjectId:   "keti-container",
		Zone:        "us-central1-a",
		OperationId: "operation-1648309236003-34160983",
		Name:        "operation-1648309236003-34160983",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/operations/describe"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	if err != nil {
		log.Println(err)
		return
	}

	var output apiserverutil.Output
	json.Unmarshal(bytes, &output)
	if output.Stderr != nil {
		fmt.Println(string(output.Stderr))
	}

	if output.Stdout != nil {
		stdout := output.Stdout
		var resp *containerpb.Operation
		json.Unmarshal(stdout, &resp)
		header := []string{"EndTime", "Name", "OperationType", "SelfLink", "StartTime", "Status", "TargetLink", "Zone"}
		for _, i := range header {
			fmt.Printf("%s: %s\n", i, reflect.ValueOf(resp).Elem().FieldByName(i))
		}
	}
}

func OperationsList() {
	op := &containerpb.ListOperationsRequest{
		ProjectId: "keti-container",
		Zone:      "-",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/operations/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	if err != nil {
		log.Println(err)
		return
	}

	var output apiserverutil.Output
	json.Unmarshal(bytes, &output)
	if output.Stderr != nil {
		fmt.Println(string(output.Stderr))
	}

	if output.Stdout != nil {
		stdout := output.Stdout
		var resp *containerpb.ListOperationsResponse
		json.Unmarshal(stdout, &resp)
		table := tablewriter.NewWriter(os.Stdout)
		header := []string{"NAME", "TYPE", "LOCATION", "TARGET", "STATUS_MESSAGE", "STATUS", "START_TIME", "END_TIME"}
		table.SetHeader(header)
		for _, v := range resp.Operations {
			targetLink := v.GetTargetLink()
			target := targetLink[strings.LastIndex(targetLink, "/")+1:]
			fmt.Println(target)
			temp := []string{v.Name, v.OperationType.String(), v.Location, target, v.StatusMessage, v.Status.String(), v.StartTime, v.EndTime}
			table.Append(temp)
		}
		table.Render()
	}
}

type Operations struct {
	ProjectId   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Zone        string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	OperationId string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Name        string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func OperationWait() {
	op := &Operations{
		OperationId: "operation-1648309236003-34160983",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/operations/wait"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
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
	if err != nil {
		log.Println(err)
		return
	}
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
	if err != nil {
		log.Println(err)
		return
	}
	util.PrintOutput(bytes)
}

func main() {
	//var images Images
	//var operations Operations

	// var auth Auth
	// auth.Login()
	ConfigSet()
	//GetServerConfig()
}
