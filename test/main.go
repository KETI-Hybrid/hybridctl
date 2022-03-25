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

func main() {
	var images Images
	images.AddTag()
	images.UnTags()
}
