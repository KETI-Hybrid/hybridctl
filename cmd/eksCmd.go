package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var createAddonCmd = &cobra.Command{
	Use:   "create-addon",
	Short: "A brief description of your command",
	Long: `	
	- create-addon
		hybridctl create-addon 

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
		checkErr(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		addonVersion, err := cmd.Flags().GetString("addon-version")
		checkErr(err)
		serviceAccountRoleArn, err := cmd.Flags().GetString("service-account-role-arn")
		checkErr(err)
		resolveConflicts, err := cmd.Flags().GetString("resolve-conflicts")
		checkErr(err)
		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)
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
		checkErr(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			createAddonInput.Tags = tagsMap
		}

		createAddon(createAddonInput)
	},
}

var deleteAddonCmd = &cobra.Command{
	Use:   "delete-addon",
	Short: "A brief description of your command",
	Long: `	
	- delete-addon
		hybridctl delete-addon 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		checkErr(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		deleteAddonInput.ClusterName = &clusterName
		deleteAddonInput.AddonName = &addonName
		deleteAddon(deleteAddonInput)

	},
}

var describeAddonCmd = &cobra.Command{
	Use:   "describe-addon",
	Short: "A brief description of your command",
	Long: `	
	- describe-addon
		hybridctl describe-addon <clusterName> <addonName> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		checkErr(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		describeAddonInput.ClusterName = &clusterName
		describeAddonInput.AddonName = &addonName
		describeAddon(describeAddonInput)

	},
}

var describeAddonVersionsCmd = &cobra.Command{
	Use:   "describe-addon-versions",
	Short: "A brief description of your command",
	Long: `	
	- describe-addon-versions
		hybridctl describe-addon-versions 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		kubernetesVersion, err := cmd.Flags().GetString("kubernetes-version")
		checkErr(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		checkErr(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		checkErr(err)
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
		describeAddonVersions(describeAddonVersionsInput)
	},
}

var listAddonCmd = &cobra.Command{
	Use:   "list-addon",
	Short: "A brief description of your command",
	Long: `	
	- list-addon
		hybridctl eks list-addon 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, err := cmd.Flags().GetString("cluster-name")
		checkErr(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		checkErr(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		checkErr(err)
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
		listAddon(listAddonInput)

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
		checkErr(err)
		associateIdentityProviderConfigInput.ClusterName = &clusterName
		// json parsing
		oidc, err := cmd.Flags().GetString("oidc")
		checkErr(err)
		byteValue := util.OpenAndReadJsonFile(oidc)
		json.Unmarshal(byteValue, &oidcRequest)
		associateIdentityProviderConfigInput.Oidc = &oidcRequest

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)
		if clientRequestToken != "" {
			associateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		tags, err := cmd.Flags().GetString("tags")
		checkErr(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			associateIdentityProviderConfigInput.Tags = tagsMap
		}
		AssociateIdentityProviderConfig(associateIdentityProviderConfigInput)

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
		checkErr(err)
		disassociateIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		jsonFileName, err := cmd.Flags().GetString("identity-provider-config")
		checkErr(err)
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &IdentityProviderConfig)
		if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
			fmt.Println("identityProviderConfig format is wrong.")
			return
		}
		disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)
		if clientRequestToken != "" {
			disassociateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		disassociateIdentityProviderConfig(disassociateIdentityProviderConfigInput)

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
		checkErr(err)
		listIdentityProviderConfigsInput.ClusterName = &clusterName
		maxResults, err := cmd.Flags().GetInt64("max-result")
		checkErr(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		checkErr(err)
		if maxResults != 0 {
			listIdentityProviderConfigsInput.MaxResults = &maxResults
		}
		if nextToken != "" {
			listIdentityProviderConfigsInput.NextToken = &nextToken
		}
		listIdentityProviderConfigs(listIdentityProviderConfigsInput)

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
		checkErr(err)
		describeIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		jsonFileName, err := cmd.Flags().GetString("identity-provider-config")
		checkErr(err)
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &IdentityProviderConfig)
		if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
			fmt.Println("identityProviderConfig format is wrong.")
			return
		}
		describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

		describeIdentityProviderConfig(describeIdentityProviderConfigInput)

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
		checkErr(err)
		associateEncryptionConfigInput.ClusterName = &clusterName

		// json parsing
		jsonFileName, err := cmd.Flags().GetString("encryption-config")
		checkErr(err)
		var encryptionConfig []*eks.EncryptionConfig
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &encryptionConfig)
		associateEncryptionConfigInput.EncryptionConfig = encryptionConfig

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)
		if clientRequestToken != "" {
			associateEncryptionConfigInput.ClientRequestToken = &clientRequestToken
		}
		AssociateEncryptionConfig(associateEncryptionConfigInput)

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
		checkErr(err)
		listTagsForResourceInput.ResourceArn = &resourceArn
		listTagsForResource(listTagsForResourceInput)
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
		checkErr(err)
		if resourceArn == "" || resourceArn == "--tags" {
			fmt.Println("resourceArn must not be nil")
		}
		tagResourceInput.ResourceArn = &resourceArn

		tags, err := cmd.Flags().GetString("tags")
		checkErr(err)
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			tagResourceInput.Tags = tagsMap
		}
		TagResource(tagResourceInput)
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
		checkErr(err)
		if resourceArn == "" || resourceArn == "--tag-keys" {
			fmt.Println("resourceArn must not be nil")
		}
		untagResourceInput.ResourceArn = &resourceArn

		keys, err := cmd.Flags().GetString("tag-keys")
		checkErr(err)
		slice := strings.Split(keys, ",")
		keyList := []*string{}
		for i := 0; i < len(slice); i++ {
			s := append(keyList, &slice[i])
			keyList = s
		}

		untagResourceInput.TagKeys = keyList

		unTagResource(untagResourceInput)
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
		checkErr(err)
		describeUpdateInput.Name = &name
		updateId, err := cmd.Flags().GetString("update-id")
		checkErr(err)
		describeUpdateInput.UpdateId = &updateId
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		checkErr(err)
		if nodegroupName != "" {
			describeUpdateInput.NodegroupName = &nodegroupName
		}
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		if addonName != "" {
			describeUpdateInput.AddonName = &addonName
		}
		describeUpdate(describeUpdateInput)

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
		checkErr(err)
		listUpdateInput.Name = &name
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		checkErr(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		maxResults, err := cmd.Flags().GetInt64("max-result")
		checkErr(err)
		nextToken, err := cmd.Flags().GetString("next-token")
		checkErr(err)
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
		listUpdate(listUpdateInput)

	},
}

var updateAddonCmd = &cobra.Command{
	Use:   "update-addon",
	Short: "A brief description of your command",
	Long: `	
	- update-addon
		hybridctl update-addon  

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmd.Flags().GetString("name")
		checkErr(err)
		addonName, err := cmd.Flags().GetString("addon-name")
		checkErr(err)
		addonVersion, err := cmd.Flags().GetString("addon-version")
		checkErr(err)
		serviceAccountRoleArn, err := cmd.Flags().GetString("service-account-role-arn")
		checkErr(err)
		resolveConflicts, err := cmd.Flags().GetString("resolve-conflicts")
		checkErr(err)
		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)
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
		updateAddon(updateAddonInput)

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
		checkErr(err)
		updateClusterConfigInput.Name = &name

		jsonFileName, err := cmd.Flags().GetString("resource-vpc-config")
		checkErr(err)
		if jsonFileName != "" {
			var resourcesVpcConfig eks.VpcConfigRequest
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &resourcesVpcConfig)
			updateClusterConfigInput.ResourcesVpcConfig = &resourcesVpcConfig
		}

		jsonFileName, err = cmd.Flags().GetString("logging")
		checkErr(err)
		if jsonFileName != "" {
			var logging eks.Logging
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &logging)
			updateClusterConfigInput.Logging = &logging
		}

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		checkErr(err)

		if clientRequestToken != "" {
			updateClusterConfigInput.ClientRequestToken = &clientRequestToken
		}

		updateClusterConfig(updateClusterConfigInput)

	},
}

var updateNodegroupConfigCmd = &cobra.Command{
	Use:   "update-Nodegroup-config",
	Short: "A brief description of your command",
	Long: `	
	- update-Nodegroup-config
		hybridctl update-Nodegroup-config --cluster-name <value> --nodegroup-name <value>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl update-Nodegroup-config --help' to view all commands")
		} else {

			updateNodegroupConfigInput.ClusterName = &args[0]
			updateNodegroupConfigInput.NodegroupName = &args[1]

			jsonFileName, err := cmd.Flags().GetString("labels")
			checkErr(err)
			if jsonFileName != "" {
				var labels eks.UpdateLabelsPayload
				byteValue := util.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &labels)
			}

			jsonFileName, err = cmd.Flags().GetString("taints")
			checkErr(err)
			if jsonFileName != "" {
				var taints eks.UpdateLabelsPayload
				byteValue := util.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &taints)
			}
			jsonFileName, err = cmd.Flags().GetString("scaling-config")
			checkErr(err)
			if jsonFileName != "" {
				var scalingConfig eks.NodegroupScalingConfig
				byteValue := util.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &scalingConfig)
			}

			jsonFileName, err = cmd.Flags().GetString("update-config")
			checkErr(err)
			if jsonFileName != "" {
				var updateConfig eks.NodegroupUpdateConfig
				byteValue := util.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &updateConfig)
			}

			clientRequestToken, err := cmd.Flags().GetString("client-request-token")
			checkErr(err)
			if clientRequestToken != "" {
				updateNodegroupConfigInput.ClientRequestToken = &clientRequestToken
			}

			updateNodegroupConfig(updateNodegroupConfigInput)
		}
	},
}
