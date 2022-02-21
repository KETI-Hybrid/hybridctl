package cmd

import (
	util "Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// // Pod-Identity
// func podIdentityAdd(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityAdd"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityDelete(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityDelete"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityList(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityList"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityExceptionAdd(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityExceptionAdd"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityExceptionDelete(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityExceptionDelete"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityExceptionList(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityExceptionList"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

// func podIdentityExceptionUpdate(p util.AKSPodIdentity) {
// 	httpPostUrl := "http://localhost:8080/podIdentityExceptionUpdate"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

func aksStart(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStart"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if string(bytes) == "" {
		fmt.Println("Succeeded to start", p.ClusterName, "in", p.ResourceGroupName)
	} else if errmsg.Error.Message == "" {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(errmsg.Error.Message)
	}
}

func aksStop(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStop"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if string(bytes) == "" {
		fmt.Println("Succeeded to stop", p.ClusterName, "in", p.ResourceGroupName)
	} else if errmsg.Error.Message == "" {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(errmsg.Error.Message)
	}
}

func HTTPPostRequest(p util.AKSAPIParameter, cmd string) {
	httpPostUrl := "http://localhost:8080/" + cmd
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	util.PrintOutput(bytes)
}

func HTTPPostRequestAPI(p util.AKSAPIParameter, cmd string) {
	httpPostUrl := "http://localhost:8080/" + cmd
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if errmsg.Error.Message == "" {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(errmsg.Error.Message)
	}
}
