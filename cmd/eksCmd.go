package cmd

import (
	"Hybrid_Cloud/hybridctl/util"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

func EKSCommonPrintOption(output interface{}, bytes []byte) {
	niloutput := output
	json.Unmarshal(bytes, &output)
	if output == niloutput {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(output)
	}
}

var EKSAddonCmd = &cobra.Command{
	Use:   "addon",
	Short: "Amazon EKS add-ons help to automate the provisioning and lifecycle management of common operational software for Amazon EKS clusters.",
	Long: `	
	Amazon EKS add-ons help to automate the provisioning and lifecycle management of common operational software for Amazon EKS clusters. 
	Amazon EKS add-ons require clusters running version 1.18 or later because Amazon EKS add-ons rely on the Server-side Apply Kubernetes feature, 
	which is only available in Kubernetes 1.18 and later.
	For more information, see Amazon EKS add-ons in the Amazon EKS User Guide .`,
}

var EKSCreateAddonCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an Amazon EKS add-on.",
	Long: `	
	- create
		hybridctl addon create

	- flags
		[--addon-version <value>]
		[--service-account-role-arn <value>]
		[--resolve-conflicts <value>]
		[--client-request-token <value>]
		[--tags <value>]

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		addonVersion, err := cmd.Flags().GetString("addon-version")
		util.CheckERR(err)
		serviceAccountRoleArn, err := cmd.Flags().GetString("service-account-role-arn")
		util.CheckERR(err)
		resolveConflicts, err := cmd.Flags().GetString("resolve-conflicts")
		util.CheckERR(err)
		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		createAddonInput.ClusterName = &clusterName
		createAddonInput.AddonName = &addonName
		if addonVersion != "" {
			createAddonInput.AddonVersion = &addonVersion
		}
		if serviceAccountRoleArn != "" {
			createAddonInput.ServiceAccountRoleArn = &serviceAccountRoleArn
		}
		if resolveConflicts != "" {
			createAddonInput.ResolveConflicts = &resolveConflicts
		}
		if clientRequestToken != "" {
			createAddonInput.ClientRequestToken = &clientRequestToken
		}
		tags, err := cmd.Flags().GetString("tags")
		util.CheckERR(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			createAddonInput.Tags = tagsMap
		}

		httpPostUrl := "/eks/addon/create"
		bytes := util.HTTPPostRequest(createAddonInput, httpPostUrl)
		var output eks.CreateAddonOutput
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDeleteAddonCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an Amazon EKS add-on.",
	Long: `	
	- delete
		hybridctl delete addon 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		deleteAddonInput.ClusterName = &clusterName
		deleteAddonInput.AddonName = &addonName

		httpPostUrl := "/eks/addon/delete"
		bytes := util.HTTPPostRequest(deleteAddonInput, httpPostUrl)
		var output eks.DeleteAddonOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var EKSDescribeAddonCmd = &cobra.Command{
	Use:   "describe",
	Short: "A brief description of your command",
	Long: `	
	- describe
		hybridctl addon describe <clusterName> <addonName> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		describeAddonInput.ClusterName = &clusterName
		describeAddonInput.AddonName = &addonName

		httpPostUrl := "/eks/addon/describe"
		bytes := util.HTTPPostRequest(describeAddonInput, httpPostUrl)
		var output eks.DescribeAddonOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var EKSDescribeAddonVersionsCmd = &cobra.Command{
	Use:   "describe-addon-versions",
	Short: "A brief description of your command",
	Long: `	
	- describe-addon-versions
		hybridctl describe-addon-versions 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		kubernetesVersion, err := cmd.Flags().GetString("kubernetes-version")
		util.CheckERR(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		util.CheckERR(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		util.CheckERR(err)
		if addonName != "" {
			describeAddonVersionsInput.AddonName = &addonName
			fmt.Println(addonName)
		}
		if kubernetesVersion != "" {
			describeAddonVersionsInput.KubernetesVersion = &kubernetesVersion
		}
		if maxResults != 0 {
			describeAddonVersionsInput.MaxResults = &maxResults
		}
		if nextToken != "" {
			describeAddonVersionsInput.NextToken = &nextToken
		}

		httpPostUrl := "/eks/addon/describe-versions"
		bytes := util.HTTPPostRequest(describeAddonVersionsInput, httpPostUrl)
		var output eks.DescribeAddonVersionsOutput
		json.Unmarshal(bytes, &output)
		if output.Addons == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}
	},
}

var EKSListAddonCmd = &cobra.Command{
	Use:   "list-addon",
	Short: "A brief description of your command",
	Long: `	
	- list-addon
		hybridctl eks list-addon 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		util.CheckERR(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		util.CheckERR(err)
		listAddonInput.ClusterName = &clusterName
		if maxResults != 0 {
			if maxResults < 1 || maxResults > 100 {
				fmt.Println("MaxResults can be between 1 and 100.")
				return
			} else {
				listAddonInput.MaxResults = &maxResults
			}
		}
		if nextToken != "" {
			listAddonInput.NextToken = &nextToken
		}

		httpPostUrl := "/eks/addon/list"
		bytes := util.HTTPPostRequest(listAddonInput, httpPostUrl)
		var output eks.ListAddonsOutput
		json.Unmarshal(bytes, &output)
		if output.Addons == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}
	},
}

var associateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "associate-identity-provider-config",
	Short: "A brief description of your command",
	Long: `	
	- associate-identity-provider-config
		hybridctl associate-identity-provider-config 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		associateIdentityProviderConfigInput.ClusterName = &clusterName
		// json parsing
		oidc, err := cmd.Flags().GetString("oidc")
		util.CheckERR(err)
		byteValue := util.OpenAndReadJsonFile(oidc)
		json.Unmarshal(byteValue, &oidcRequest)
		associateIdentityProviderConfigInput.Oidc = &oidcRequest

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		if clientRequestToken != "" {
			associateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		tags, err := cmd.Flags().GetString("tags")
		util.CheckERR(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			associateIdentityProviderConfigInput.Tags = tagsMap
		}

		httpPostUrl := "/eks/identity-provider-config/associate"
		bytes := util.HTTPPostRequest(associateIdentityProviderConfigInput, httpPostUrl)
		var output eks.AssociateIdentityProviderConfigOutput
		json.Unmarshal(bytes, &output)
		if output.Tags == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}

	},
}

var disassociateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "disassociate-identity-provider-config",
	Short: "A brief description of your command",
	Long: `	
	- disassociate-identity-provider-config
		hybridctl disassociate-identity-provider-config <clusterName> <oidc> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		disassociateIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		jsonFileName, err := cmd.Flags().GetString("identity-provider-config")
		util.CheckERR(err)
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &IdentityProviderConfig)
		if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
			fmt.Println("identityProviderConfig format is wrong.")
			return
		}
		disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		if clientRequestToken != "" {
			disassociateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/identity-provider-config/disassociate"
		bytes := util.HTTPPostRequest(disassociateIdentityProviderConfigInput, httpPostUrl)
		var output eks.DisassociateIdentityProviderConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var listIdentityProviderConfigsCmd = &cobra.Command{
	Use:   "list-identity-provider-configs",
	Short: "A brief description of your command",
	Long: `	
	- list-identity-provider-configs
		hybridctl list-identity-provider-configs --cluster-name

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		listIdentityProviderConfigsInput.ClusterName = &clusterName
		maxResults, err := cmd.Flags().GetInt64("max-result")
		util.CheckERR(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		util.CheckERR(err)
		if maxResults != 0 {
			listIdentityProviderConfigsInput.MaxResults = &maxResults
		}
		if nextToken != "" {
			listIdentityProviderConfigsInput.NextToken = &nextToken
		}

		httpPostUrl := "/eks/identity-provider-config/list"
		bytes := util.HTTPPostRequest(listIdentityProviderConfigsInput, httpPostUrl)
		var output eks.ListIdentityProviderConfigsOutput
		json.Unmarshal(bytes, &output)
		if output.IdentityProviderConfigs == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}

	},
}

var describeIdentityProviderConfigCmd = &cobra.Command{
	Use:   "describe-identity-provider-config",
	Short: "A brief description of your command",
	Long: `	
	- describe-identity-provider-config
		hybridctl describe-identity-provider-config <clusterName> <oidc> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		describeIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		jsonFileName, err := cmd.Flags().GetString("identity-provider-config")
		util.CheckERR(err)
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &IdentityProviderConfig)
		if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
			fmt.Println("identityProviderConfig format is wrong.")
			return
		}
		describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

		httpPostUrl := "/eks/identity-provider-config/describe"
		bytes := util.HTTPPostRequest(describeIdentityProviderConfigInput, httpPostUrl)
		var output eks.DescribeIdentityProviderConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var associateEncryptionConfigCmd = &cobra.Command{
	Use:   "associate-encryption-config",
	Short: "A brief description of your command",
	Long: `	
	- associate-encryption-config
		hybridctl associate-encryption-config <clusterName> --encryption-config <jsonfile>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		associateEncryptionConfigInput.ClusterName = &clusterName

		// json parsing
		jsonFileName, err := cmd.Flags().GetString("encryption-config")
		util.CheckERR(err)
		var encryptionConfig []*eks.EncryptionConfig
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &encryptionConfig)
		associateEncryptionConfigInput.EncryptionConfig = encryptionConfig

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		if clientRequestToken != "" {
			associateEncryptionConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/encryption-config/associate"
		bytes := util.HTTPPostRequest(associateEncryptionConfigInput, httpPostUrl)
		var output eks.AssociateEncryptionConfigOutput
		EKSCommonPrintOption(output, bytes)
	},
}

var listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A brief description of your command",
	Long: `	
	- list-tags-for-resource
		hybridctl list-tags-for-resource --resource-arn

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		resourceArn, err := cmd.Flags().GetString("resource-arn")
		util.CheckERR(err)
		listTagsForResourceInput.ResourceArn = &resourceArn

		httpPostUrl := "/eks/resource/list"
		bytes := util.HTTPPostRequest(listTagsForResourceInput, httpPostUrl)
		var output eks.ListTagsForResourceOutput
		json.Unmarshal(bytes, &output)
		if output.Tags == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}
	},
}

var tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "A brief description of your command",
	Long: `	
	- tage-resource
		hybridctl tag-resource --resource-arn <value> --tags <jsonfile>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		var tagResourceInput eks.TagResourceInput
		resourceArn, err := cmd.Flags().GetString("resource-arn")
		util.CheckERR(err)
		if resourceArn == "" || resourceArn == "--tags" {
			fmt.Println("resourceArn must not be nil")
		}
		tagResourceInput.ResourceArn = &resourceArn

		tags, err := cmd.Flags().GetString("tags")
		util.CheckERR(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			tagResourceInput.Tags = tagsMap
		}

		httpPostUrl := "/eks/resource/tag"
		bytes := util.HTTPPostRequest(tagResourceInput, httpPostUrl)
		util.PrintErrMsg(bytes)
	},
}

var untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "A brief description of your command",
	Long: `	
	- untage-resource
		hybridctl untag-resource --resource-arn <value> --tag-keys <key,key>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		var untagResourceInput eks.UntagResourceInput
		resourceArn, err := cmd.Flags().GetString("resource-arn")
		util.CheckERR(err)
		if resourceArn == "" || resourceArn == "--tag-keys" {
			fmt.Println("resourceArn must not be nil")
		}
		untagResourceInput.ResourceArn = &resourceArn

		keys, err := cmd.Flags().GetString("tag-keys")
		util.CheckERR(err)
		slice := strings.Split(keys, ",")
		keyList := []*string{}
		for i := 0; i < len(slice); i++ {
			s := append(keyList, &slice[i])
			keyList = s
		}

		untagResourceInput.TagKeys = keyList

		httpPostUrl := "/eks/resource/untag"
		bytes := util.HTTPPostRequest(untagResourceInput, httpPostUrl)
		util.PrintErrMsg(bytes)
	},
}
var describeUpdateCmd = &cobra.Command{
	Use:   "describe-update",
	Short: "A brief description of your command",
	Long: `	
	- describe-update
		hybridctl describe-update <clusterName> <updateID>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		name, err := cmd.Flags().GetString("name")
		util.CheckERR(err)
		describeUpdateInput.Name = &name
		updateId, err := cmd.Flags().GetString("update-id")
		util.CheckERR(err)
		describeUpdateInput.UpdateId = &updateId
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		util.CheckERR(err)
		if nodegroupName != "" {
			describeUpdateInput.NodegroupName = &nodegroupName
		}
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		if addonName != "" {
			describeUpdateInput.AddonName = &addonName
		}

		httpPostUrl := "/eks/describe/update"
		bytes := util.HTTPPostRequest(describeUpdateInput, httpPostUrl)
		var output eks.DescribeUpdateOutput
		EKSCommonPrintOption(output, bytes)
	},
}

var listUpdateCmd = &cobra.Command{
	Use:   "list-update",
	Short: "A brief description of your command",
	Long: `	
	- list-update
		hybridctl list-update <clusterName> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		name, err := cmd.Flags().GetString("name")
		util.CheckERR(err)
		listUpdateInput.Name = &name
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		util.CheckERR(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		util.CheckERR(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		util.CheckERR(err)
		if nodegroupName != "" {
			listUpdateInput.NodegroupName = &nodegroupName
		}
		if addonName != "" {
			listUpdateInput.AddonName = &addonName
		}
		if maxResults != 0 {
			listAddonInput.MaxResults = &maxResults
		}
		if nextToken != "" {
			listAddonInput.NextToken = &nextToken
		}

		httpPostUrl := "/eks/list/update"
		bytes := util.HTTPPostRequest(listUpdateInput, httpPostUrl)
		var output eks.ListUpdatesOutput
		json.Unmarshal(bytes, &output)
		if output.UpdateIds == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}

	},
}

var EKSUpdateAddonCmd = &cobra.Command{
	Use:   "update-addon",
	Short: "A brief description of your command",
	Long: `	
	- update-addon
		hybridctl update-addon  

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmd.Flags().GetString("name")
		util.CheckERR(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		util.CheckERR(err)
		addonVersion, err := cmd.Flags().GetString("addon-version")
		util.CheckERR(err)
		serviceAccountRoleArn, err := cmd.Flags().GetString("service-account-role-arn")
		util.CheckERR(err)
		resolveConflicts, err := cmd.Flags().GetString("resolve-conflicts")
		util.CheckERR(err)
		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		updateAddonInput.ClusterName = &name
		updateAddonInput.AddonName = &addonName
		if addonVersion != "" {
			updateAddonInput.AddonVersion = &addonVersion
		}
		if serviceAccountRoleArn != "" {
			updateAddonInput.ServiceAccountRoleArn = &serviceAccountRoleArn
		}
		if resolveConflicts != "" {
			updateAddonInput.ResolveConflicts = &resolveConflicts
		}
		if clientRequestToken != "" {
			updateAddonInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/addon/update"
		bytes := util.HTTPPostRequest(updateAddonInput, httpPostUrl)
		var output eks.UpdateAddonOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var updateClusterConfigCmd = &cobra.Command{
	Use:   "update-cluster-config",
	Short: "A brief description of your command",
	Long: `	
	- update-cluster-config
		hybridctl update-cluster-config <clusterName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		name, err := cmd.Flags().GetString("name")
		util.CheckERR(err)
		updateClusterConfigInput.Name = &name

		jsonFileName, err := cmd.Flags().GetString("resource-vpc-config")
		util.CheckERR(err)
		if jsonFileName != "" {
			var resourcesVpcConfig eks.VpcConfigRequest
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &resourcesVpcConfig)
			updateClusterConfigInput.ResourcesVpcConfig = &resourcesVpcConfig
		}

		jsonFileName, err = cmd.Flags().GetString("logging")
		util.CheckERR(err)
		if jsonFileName != "" {
			var logging eks.Logging
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &logging)
			updateClusterConfigInput.Logging = &logging
		}

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)

		if clientRequestToken != "" {
			updateClusterConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/cluster-config/update"
		bytes := util.HTTPPostRequest(updateClusterConfigInput, httpPostUrl)
		var output eks.UpdateClusterConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var updateNodegroupConfigCmd = &cobra.Command{
	Use:   "update-nodegroup-config",
	Short: "A brief description of your command",
	Long: `	
	- update-Nodegroup-config
		hybridctl update-nodegroup-config --cluster-name <value> --nodegroup-name <value>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		util.CheckERR(err)
		updateNodegroupConfigInput.ClusterName = &clusterName

		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		util.CheckERR(err)
		updateNodegroupConfigInput.NodegroupName = &nodegroupName
		jsonFileName, err := cmd.Flags().GetString("labels")
		util.CheckERR(err)
		if jsonFileName != "" {
			var labels eks.UpdateLabelsPayload
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &labels)
			updateNodegroupConfigInput.Labels = &labels
		}

		jsonFileName, err = cmd.Flags().GetString("taints")
		util.CheckERR(err)
		if jsonFileName != "" {
			var taints eks.UpdateLabelsPayload
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &taints)
			// updateNodegroupConfigInput.Taints = taints
		}
		jsonFileName, err = cmd.Flags().GetString("scaling-config")
		util.CheckERR(err)
		if jsonFileName != "" {
			var scalingConfig eks.NodegroupScalingConfig
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &scalingConfig)
		}

		jsonFileName, err = cmd.Flags().GetString("update-config")
		util.CheckERR(err)
		if jsonFileName != "" {
			var updateConfig eks.NodegroupUpdateConfig
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &updateConfig)
		}

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		util.CheckERR(err)
		if clientRequestToken != "" {
			updateNodegroupConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/nodegroup-config/update"
		bytes := util.HTTPPostRequest(updateNodegroupConfigInput, httpPostUrl)
		var output eks.UpdateNodegroupConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}
