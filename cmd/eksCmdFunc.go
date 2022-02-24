package cmd

import (
	"Hybrid_Cloud/hybridctl/util"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
)

func AssociateEncryptionConfig(input eks.AssociateEncryptionConfigInput) {
	httpPostUrl := "http://localhost:8080/associateEncryptionConfig"
	var output eks.AssociateEncryptionConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func createAddon(input eks.CreateAddonInput) {
	httpPostUrl := "http://localhost:8080/createAddon"
	var output eks.CreateAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func deleteAddon(input eks.DeleteAddonInput) {
	httpPostUrl := "http://localhost:8080/deleteAddon"
	var output eks.DeleteAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func describeAddon(input eks.DescribeAddonInput) {
	httpPostUrl := "http://localhost:8080/describeAddon"
	var output eks.DescribeAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func describeAddonVersions(input eks.DescribeAddonVersionsInput) {
	httpPostUrl := "http://localhost:8080/describeAddonVersions"
	var output eks.DescribeAddonVersionsOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	// niloutput := output
	json.Unmarshal(bytes, &output)
	if output.Addons == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func listAddon(input eks.ListAddonsInput) {
	httpPostUrl := "http://localhost:8080/listAddon"
	var output eks.ListAddonsOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	// niloutput := output
	json.Unmarshal(bytes, &output)
	if output.Addons == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func updateAddon(input eks.UpdateAddonInput) {
	httpPostUrl := "http://localhost:8080/updateAddon"
	var output eks.UpdateAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func AssociateIdentityProviderConfig(input eks.AssociateIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/associateIdentityProviderConfig"
	var output eks.AssociateIdentityProviderConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	// niloutput := output
	json.Unmarshal(bytes, &output)
	if output.Tags == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func disassociateIdentityProviderConfig(input eks.DisassociateIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/disassociateIdentityProviderConfig"
	var output eks.DisassociateIdentityProviderConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func describeIdentityProviderConfig(input eks.DescribeIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/describeIdentityProviderConfig"
	var output eks.DescribeIdentityProviderConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func listIdentityProviderConfigs(input eks.ListIdentityProviderConfigsInput) {
	httpPostUrl := "http://localhost:8080/listIdentityProviderConfigs"
	var output eks.ListIdentityProviderConfigsOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	// niloutput := output
	json.Unmarshal(bytes, &output)
	if output.IdentityProviderConfigs == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func describeUpdate(input eks.DescribeUpdateInput) {
	httpPostUrl := "http://localhost:8080/describeUpdate"
	var output eks.DescribeUpdateOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func listTagsForResource(input eks.ListTagsForResourceInput) {
	httpPostUrl := "http://localhost:8080/listTagsForResource"
	var output eks.ListTagsForResourceOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	// niloutput := output
	json.Unmarshal(bytes, &output)
	if output.Tags == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func listUpdate(input eks.ListUpdatesInput) {
	httpPostUrl := "http://localhost:8080/listUpdate"
	var output eks.ListUpdatesOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	if output.UpdateIds == nil {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func TagResource(input eks.TagResourceInput) {
	httpPostUrl := "http://localhost:8080/tagResource"
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	util.PrintErrMsg(bytes)
}

func unTagResource(input eks.UntagResourceInput) {
	httpPostUrl := "http://localhost:8080/untagResource"
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	util.PrintErrMsg(bytes)
}

func updateClusterConfig(input eks.UpdateClusterConfigInput) {
	httpPostUrl := "http://localhost:8080/updateClusterConfig"
	var output eks.UpdateClusterConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

func updateNodegroupConfig(input eks.UpdateNodegroupConfigInput) {
	httpPostUrl := "http://localhost:8080/updateNodegroupConfig"
	var output eks.UpdateNodegroupConfigOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}
