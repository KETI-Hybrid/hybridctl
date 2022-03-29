package main

import (
	apiserverutil "Hybrid_Cloud/hcp-apiserver/pkg/util"
	"Hybrid_Cloud/hybridctl/util"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

var GKE_CONTAINER_PATH = "/gke/container"

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// images
type Images struct {
	SRC_IMAGE  string
	DEST_IMAGE string
	IMAGE_NAME string
}

func (i *Images) AddTag() {
	i = &Images{
		SRC_IMAGE:  "gcr.io/keti-container/busybox",
		DEST_IMAGE: "gcr.io/keti-container/busybox:mytag2",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/addTag"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (i *Images) Delete() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/delete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (i *Images) Describe() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/describe"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (i *Images) List() {
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (i *Images) ListTags() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/listTags"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (i *Images) UnTags() {
	i = &Images{
		IMAGE_NAME: "gcr.io/keti-container/busybox:mytag2",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/images/unTags"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, i)
	checkErr(err)
	util.PrintOutput(bytes)
}

func OperationsDescribe() {
	op := &containerpb.GetOperationRequest{
		ProjectId:   "keti-container",
		Zone:        "us-central1-a",
		OperationId: "operation-1648309236003-341609",
		Name:        "operation-1648309236003-34160983",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/operations/describe"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	checkErr(err)

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

func GetServerConfig() {}

func OperationsList() {
	op := &containerpb.ListOperationsRequest{
		ProjectId: "keti-container",
		Zone:      "-",
	}
	httpPostUrl := "http://localhost:3080" + GKE_CONTAINER_PATH + "/operations/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	checkErr(err)

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

/*
func (op *Operations) Wait() {
	op = &Operations{
		ProjectId:   "keti-container",
		Zone:        "us-central1-a",
		OperationId: "operation-1648309236003-34160983",
		Name:        "operation-1648309236003-34160983",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/operations/wait"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, op)
	checkErr(err)
	util.PrintOutput(bytes)
}
*/

type Docker struct {
	AUTHORIZE_ONLY bool
	DOCKER_HOST    string
	SERVER         string
}

func (d *Docker) Docker() {
	d = &Docker{
		AUTHORIZE_ONLY: false,
	}
	httpPostUrl := "http://localhost:3001/gke/docker"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, d)
	checkErr(err)
	util.PrintOutput(bytes)
}

func main() {
	//var images Images
	//var operations Operations
	OperationsList()
}
