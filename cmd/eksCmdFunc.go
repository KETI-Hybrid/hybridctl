package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	cobrautil "Hybrid_Cluster/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
)

func AssociateEncryptionConfig(AssociateEncryptionConfigInput eks.AssociateEncryptionConfigInput) {
	httpPostUrl := "http://localhost:8080/associateEncryptionConfig"
	var output eks.AssociateEncryptionConfigOutput
	cobrautil.GetJson(httpPostUrl, AssociateEncryptionConfigInput, &output)
	fmt.Printf("%+v\n", output)
}

func AssociateIdentityProviderConfig(AssociateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/associateIdentityProviderConfig"
	var output eks.AssociateIdentityProviderConfigOutput
	cobrautil.GetJson(httpPostUrl, AssociateIdentityProviderConfigInput, &output)
	fmt.Printf("%+v\n", output)
}

func createAddon(createAddonInput eks.CreateAddonInput) {
	httpPostUrl := "http://localhost:8080/createAddon"
	var output eks.CreateAddonOutput
	util.GetJson(httpPostUrl, createAddonInput, output)

	// var message util.Addon
	// response, _ := util.GetJson(httpPostUrl, createAddonInput, output)
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Println(bytes)
	// }
	// defer response.Body.Close()
	// json.Unmarshal(bytes, &message)
	// json.Unmarshal(bytes, &output)
	// if output.Addon == nil {
	// 	json.Unmarshal(bytes, &message)
	// 	fmt.Println(message.Message_)
	// } else {
	// 	fmt.Println(output)
	// }
}

func deleteAddon(deleteAddonInput eks.DeleteAddonInput) {
	httpPostUrl := "http://localhost:8080/deleteAddon"
	var output eks.DeleteAddonOutput
	// var message util.Addon
	util.GetJson(httpPostUrl, deleteAddonInput, &output)
	fmt.Println(output)
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Println(bytes)
	// }
	// defer response.Body.Close()
	// json.Unmarshal(bytes, &output)
	// if output.Addon == nil {
	// 	json.Unmarshal(bytes, &message)
	// 	fmt.Println(message.Message_)
	// } else {
	// 	fmt.Println(output)
	// }
}

func describeAddon(describeAddonInput eks.DescribeAddonInput) {
	httpPostUrl := "http://localhost:8080/describeAddon"
	var output eks.DescribeAddonOutput
	// var message util.Addon
	util.GetJson(httpPostUrl, describeAddonInput, &output)
	fmt.Println(output)
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Println(bytes)
	// }
	// defer response.Body.Close()
	// json.Unmarshal(bytes, &output)
	// fmt.Println(string(bytes))
	// if output.Addon == nil {
	// 	fmt.Println("A")
	// 	json.Unmarshal(bytes, &message)
	// 	fmt.Println(message.Message_)
	// } else {
	// 	fmt.Println(output)
	// }
}

func describeAddonVersions(describeAddonVersionsInput eks.DescribeAddonVersionsInput) {
	httpPostUrl := "http://localhost:8080/describeAddonVersions"
	var output eks.DescribeAddonVersionsOutput
	util.GetJson(httpPostUrl, describeAddonVersionsInput, &output)
	fmt.Printf("%+v\n", output)
}

func describeIdentityProvicerConfig(input eks.DescribeIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/describeIdentityProviderConfig"
	var output eks.DescribeIdentityProviderConfigOutput
	cobrautil.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func describeUpdate(describeUpdateInput eks.DescribeUpdateInput) {
	httpPostUrl := "http://localhost:8080/describeUpdate"
	var output eks.DescribeUpdateOutput
	util.GetJson(httpPostUrl, describeUpdateInput, &output)
	fmt.Printf("%+v\n", output)
}

func disassociateIdentityProvicerConfig(input eks.DisassociateIdentityProviderConfigInput) {
	httpPostUrl := "http://localhost:8080/disassociateIdentityProviderConfig"
	var output eks.DisassociateIdentityProviderConfigOutput
	cobrautil.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func listAddon(listAddonInput eks.ListAddonsInput) {
	httpPostUrl := "http://localhost:8080/listAddon"
	var output eks.ListAddonsOutput
	util.GetJson(httpPostUrl, listAddonInput, &output)
	fmt.Printf("%+v\n", output)
}

func listIdentityProviderConfigs(input eks.ListIdentityProviderConfigsInput) {
	httpPostUrl := "http://localhost:8080/listIdentityProviderConfigs"
	var output eks.ListIdentityProviderConfigsOutput
	util.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func listTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) {
	httpPostUrl := "http://localhost:8080/listTagsForResource"
	var output eks.ListTagsForResourceOutput
	util.GetJson(httpPostUrl, listTagsForResourceInput, &output)
	fmt.Printf("%+v\n", output)
}

func listUpdate(listUpdateInput eks.ListUpdatesInput) {
	httpPostUrl := "http://localhost:8080/listUpdate"
	var output eks.ListUpdatesOutput
	util.GetJson(httpPostUrl, listUpdateInput, &output)
	fmt.Printf("%+v\n", output)
}

func TagResource(input eks.TagResourceInput) {
	httpPostUrl := "http://localhost:8080/tagResource"
	var output eks.TagResourceOutput
	util.GetJson(httpPostUrl, input, &output)
}

func unTagResource(input eks.UntagResourceInput) {
	httpPostUrl := "http://localhost:8080/untagResource"
	var output eks.UntagResourceOutput
	util.GetJson(httpPostUrl, input, &output)
	// fmt.Printf("%+v\n", output)
}

func updateAddon(updateAddonInput eks.UpdateAddonInput) {
	httpPostUrl := "http://localhost:8080/updateAddon"
	var output eks.UpdateAddonOutput
	util.GetJson(httpPostUrl, updateAddonInput, &output)
	fmt.Printf("%+v\n", output)
}

func updateClusterConfig(input eks.UpdateClusterConfigInput) {
	httpPostUrl := "http://localhost:8080/updateClusterConfig"
	var output eks.UpdateClusterConfigOutput
	util.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}

func updateNodegroupConfig(input eks.UpdateNodegroupConfigInput) {
	httpPostUrl := "http://localhost:8080/updateNodegroupConfig"
	var output eks.UpdateNodegroupConfigOutput
	util.GetJson(httpPostUrl, input, &output)
	fmt.Printf("%+v\n", output)
}
