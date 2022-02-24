package cmd

import (
	util "Hybrid_Cloud/hybridctl/util"
	cmdpb "Hybrid_Cloud/protos/v1/cmd"
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

//addon
func addonDisable(p util.AKSAddon) {
	httpPostUrl := "http://localhost:8080/addonDisable"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func addonEnable(p util.AKSAddon) {
	httpPostUrl := "http://localhost:8080/addonEnable"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func addonList(p cmdpb.AKSAddon) {
	httpPostUrl := "http://localhost:8080/addonList"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func addonListAvailable() {
	httpPostUrl := "http://localhost:8080/addonListAvailable"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
	checkErr(err)
	fmt.Println(string(bytes))
}

func addonShow(p util.AKSAddon) {
	httpPostUrl := "http://localhost:8080/addonShow"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func addonUpdate(p util.AKSAddon) {
	httpPostUrl := "http://localhost:8080/addonUpdate"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

// Pod-Identity
func podIdentityAdd(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityAdd"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityDelete(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityDelete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityList(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityList"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityExceptionAdd(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityExceptionAdd"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, nil)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityExceptionDelete(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityExceptionDelete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityExceptionList(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityExceptionList"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func podIdentityExceptionUpdate(p util.AKSPodIdentity) {
	httpPostUrl := "http://localhost:8080/podIdentityExceptionUpdate"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	fmt.Println(string(bytes))
}

func aksStart(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStart"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if string(bytes) == "" {
		fmt.Println("Succeeded to start", p.ResourceName, "in", p.ResourceGroupName)
	} else if errmsg.Error.Message == "" {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(errmsg.Error.Message)
	}
}

func aksStop(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksStop"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if string(bytes) == "" {
		fmt.Println("Succeeded to stop", p.ResourceName, "in", p.ResourceGroupName)
	} else if errmsg.Error.Message == "" {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(errmsg.Error.Message)
	}
}

func aksRotateCerts(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksRotateCerts"
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

func aksGetOSoptions(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/aksGetOSoptions"
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

func maintenanceconfigurationCreateOrUpdate(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationCreateOrUpdate"
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

func maintenanceconfigurationList(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationList"
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

func maintenanceconfigurationDelete(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationDelete"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	// var errmsg util.Error
	// json.Unmarshal(bytes, &errmsg)
	// if string(bytes) == "" {
	// 	fmt.Println("Succeeded to delete", p.ResourceName, "in", p.ResourceGroupName)
	// } else if errmsg.Error.Message == "" {
	fmt.Println(string(bytes))
	// } else {
	// 	fmt.Println(errmsg.Error.Message)
	// }
}

func maintenanceconfigurationShow(p util.EKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/maintenanceconfigurationShow"
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

func appUp(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/appUp"
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

func browse(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/browse"
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

func checkAcr(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/checkAcr"
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	var errmsg util.Error
	json.Unmarshal(bytes, &errmsg)
	if errmsg.Error.Message == "" {
		fmt.Printf(string(bytes))
	} else {
		fmt.Printf(errmsg.Error.Message)
	}
}

func getUpgrades(p util.AKSAPIParameter) {
	httpPostUrl := "http://localhost:8080/getUpgrades"
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

// func getVersions(p util.AKSAPIParameter) {
// 	httpPostUrl := "http://localhost:8080/getUpgrades"
// 	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
// 	checkErr(err)
// 	fmt.Println(string(bytes))
// }

func HTTPPostRequest(p util.AKSAPIParameter, cmd string) {
	httpPostUrl := "http://localhost:8080/" + cmd
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	util.PrintOutput(bytes)
}

func HTTPPostRequestCLI(p util.AKSInstallCLI, cmd string) {
	httpPostUrl := "http://localhost:8080/" + cmd
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	util.PrintOutput(bytes)
}

func HTTPPostRequestConfig(p util.AKSk8sConfiguration, cmd string) {
	httpPostUrl := "http://localhost:8080/" + cmd
	bytes, err := util.GetResponseBody("POST", httpPostUrl, p)
	checkErr(err)
	util.PrintOutput(bytes)
}
