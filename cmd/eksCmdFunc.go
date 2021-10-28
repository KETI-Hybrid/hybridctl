package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
	"github.com/aws/aws-sdk-go/service/eks"
)

// func AssociateEncryptionConfig(AssociateEncryptionConfigInput eks.AssociateEncryptionConfigInput) {
// 	httpPostUrl := "http://localhost:8080/associateEncryptionConfig"
// 	var output eks.AssociateEncryptionConfigOutput
// 	cobrautil.GetJson(httpPostUrl, AssociateEncryptionConfigInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func AssociateIdentityProviderConfig(AssociateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput) {
// 	httpPostUrl := "http://localhost:8080/associateIdentityProviderConfig"
// 	var output eks.AssociateIdentityProviderConfigOutput
// 	cobrautil.GetJson(httpPostUrl, AssociateIdentityProviderConfigInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

func createAddon(input eks.CreateAddonInput) {
	httpPostUrl := "http://localhost:8080/createAddon"
	var output eks.CreateAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Addon == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

func deleteAddon(input eks.DeleteAddonInput) {
	httpPostUrl := "http://localhost:8080/deleteAddon"
	var output eks.DeleteAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Addon == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

func describeAddon(input eks.DescribeAddonInput) {
	httpPostUrl := "http://localhost:8080/describeAddon"
	var output eks.DescribeAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Addon == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

func describeAddonVersions(input eks.DescribeAddonVersionsInput) {
	httpPostUrl := "http://localhost:8080/describeAddonVersions"
	var output eks.DescribeAddonVersionsOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Addons == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

// func describeIdentityProvicerConfig(input eks.DescribeIdentityProviderConfigInput) {
// 	httpPostUrl := "http://localhost:8080/describeIdentityProviderConfig"
// 	var output eks.DescribeIdentityProviderConfigOutput
// 	cobrautil.GetJson(httpPostUrl, input, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func describeUpdate(describeUpdateInput eks.DescribeUpdateInput) {
// 	httpPostUrl := "http://localhost:8080/describeUpdate"
// 	var output eks.DescribeUpdateOutput
// 	util.GetJson(httpPostUrl, describeUpdateInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func disassociateIdentityProvicerConfig(input eks.DisassociateIdentityProviderConfigInput) {
// 	httpPostUrl := "http://localhost:8080/disassociateIdentityProviderConfig"
// 	var output eks.DisassociateIdentityProviderConfigOutput
// 	cobrautil.GetJson(httpPostUrl, input, &output)
// 	fmt.Printf("%+v\n", output)
// }

func listAddon(input eks.ListAddonsInput) {
	httpPostUrl := "http://localhost:8080/listAddon"
	var output eks.ListAddonsOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Addons == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

// func listIdentityProviderConfigs(input eks.ListIdentityProviderConfigsInput) {
// 	httpPostUrl := "http://localhost:8080/listIdentityProviderConfigs"
// 	var output eks.ListIdentityProviderConfigsOutput
// 	util.GetJson(httpPostUrl, input, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func listTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) {
// 	httpPostUrl := "http://localhost:8080/listTagsForResource"
// 	var output eks.ListTagsForResourceOutput
// 	util.GetJson(httpPostUrl, listTagsForResourceInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func listUpdate(listUpdateInput eks.ListUpdatesInput) {
// 	httpPostUrl := "http://localhost:8080/listUpdate"
// 	var output eks.ListUpdatesOutput
// 	util.GetJson(httpPostUrl, listUpdateInput, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func TagResource(input eks.TagResourceInput) {
// 	httpPostUrl := "http://localhost:8080/tagResource"
// 	var output eks.TagResourceOutput
// 	util.GetJson(httpPostUrl, input, &output)
// }

// func unTagResource(input eks.UntagResourceInput) {
// 	httpPostUrl := "http://localhost:8080/untagResource"
// 	var output eks.UntagResourceOutput
// 	util.GetJson(httpPostUrl, input, &output)
// 	// fmt.Printf("%+v\n", output)
// }

func updateAddon(input eks.UpdateAddonInput) {
	httpPostUrl := "http://localhost:8080/updateAddon"
	var output eks.UpdateAddonOutput
	bytes, _ := util.GetResponseBody("POST", httpPostUrl, input)
	json.Unmarshal(bytes, &output)
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		panic(err)
	}
	if output.Update == nil {
		fmt.Println(jsonParsed.Path("Message_").Data())
	} else {
		fmt.Println(output)
	}
}

// func updateClusterConfig(input eks.UpdateClusterConfigInput) {
// 	httpPostUrl := "http://localhost:8080/updateClusterConfig"
// 	var output eks.UpdateClusterConfigOutput
// 	util.GetJson(httpPostUrl, input, &output)
// 	fmt.Printf("%+v\n", output)
// }

// func updateNodegroupConfig(input eks.UpdateNodegroupConfigInput) {
// 	httpPostUrl := "http://localhost:8080/updateNodegroupConfig"
// 	var output eks.UpdateNodegroupConfigOutput
// 	util.GetJson(httpPostUrl, input, &output)
// 	fmt.Printf("%+v\n", output)
// }
