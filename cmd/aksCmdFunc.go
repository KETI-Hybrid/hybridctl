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
	}
}

func aksStart(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStart"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	if string(bytes) == "" {
		fmt.Println("Succeeded to start", p.ResourceName, "in", p.ResourceGroupName)
	} else {
		fmt.Println(string(bytes))
	}
}

func aksStop(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStop"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	if string(bytes) == "" {
		fmt.Println("Succeeded to stop", p.ResourceName, "in", p.ResourceGroupName)
	} else {
		fmt.Println(string(bytes))
	}
}

func aksRotateCerts(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksRotateCerts"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func aksGetOSoptions(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksGetOSoptions"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	CloudErr := util.CloudError{}
	errJson := json.Unmarshal(bytes, &CloudErr)
	fmt.Println(errJson)
	if errJson == nil {
		fmt.Println("Success")
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}

}

func maintenanceconfigurationCreateOrUpdate(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationCreateOrUpdate"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	CloudErr := util.CloudError{}
	errJson := json.Unmarshal(bytes, &CloudErr)
	fmt.Println(errJson)
	if errJson == nil {
		fmt.Println("Success")
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
}

func maintenanceconfigurationList(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationList"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	CloudErr := util.CloudError{}
	errJson := json.Unmarshal(bytes, &CloudErr)
	fmt.Println(errJson)
	if errJson == nil {
		fmt.Println("Success")
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
}

func maintenanceconfigurationDelete(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationDelete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	CloudErr := util.CloudError{}
	errJson := json.Unmarshal(bytes, &CloudErr)
	fmt.Println(errJson)
	if errJson == nil {
		fmt.Println("Success")
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
}

func maintenanceconfigurationShow(p util.EksAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationShow"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	CloudErr := util.CloudError{}
	errJson := json.Unmarshal(bytes, &CloudErr)
	fmt.Println(errJson)
	if errJson == nil {
		fmt.Println("Success")
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
}
