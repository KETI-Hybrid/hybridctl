package main

import (
	"Hybrid_Cloud/hybridctl/util"
	"log"
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

// operations
type Operations struct {
	OPERATION_ID string
}

func (o *Operations) Describe() {
	o = &Operations{
		OPERATION_ID: "1189332694316803667",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/operations/describe"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, o)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (o *Operations) List() {
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/operations/list"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, o)
	checkErr(err)
	util.PrintOutput(bytes)
}

func (o *Operations) Wait() {
	o = &Operations{
		OPERATION_ID: "",
	}
	httpPostUrl := "http://localhost:3001" + GKE_CONTAINER_PATH + "/operations/wait"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, o)
	checkErr(err)
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
	httpPostUrl := "http://localhost:3001/gke/docker"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, d)
	checkErr(err)
	util.PrintOutput(bytes)
}

func main() {
	//var images Images
	//var operations Operations
	var docker Docker
	docker.Docker()
}
