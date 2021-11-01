package cmd

import (
	"Hybrid_Cluster/hybridctl/util"
	cobrautil "Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

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

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl associate-encryption-config --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl associate-encryption-config --help' to view all commands")
		} else {
			associateEncryptionConfigInput.ClusterName = &args[0]

			// json parsing
			jsonFileName, _ := cmd.Flags().GetString("encryption-config")
			var encryptionConfig []*eks.EncryptionConfig
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &encryptionConfig)
			associateEncryptionConfigInput.EncryptionConfig = encryptionConfig

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
			if clientRequestToken != "" {
				associateEncryptionConfigInput.ClientRequestToken = &clientRequestToken
			}
			// AssociateEncryptionConfig(associateEncryptionConfigInput)
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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		associateIdentityProviderConfigInput.ClusterName = &clusterName
		// json parsing
		oidc, _ := cmd.Flags().GetString("oidc")
		byteValue := cobrautil.OpenAndReadJsonFile(oidc)
		json.Unmarshal(byteValue, &oidcRequest)
		associateIdentityProviderConfigInput.Oidc = &oidcRequest

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			associateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := cobrautil.OpenAndReadJsonFile(tags)
			json.Unmarshal(byteValue, &tagsMap)
			associateIdentityProviderConfigInput.Tags = tagsMap
		}
		AssociateIdentityProviderConfig(associateIdentityProviderConfigInput)

	},
}

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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		addonName, _ := cmd.Flags().GetString("addon-name")
		addonVersion, _ := cmd.Flags().GetString("addon-version")
		serviceAccountRoleArn, _ := cmd.Flags().GetString("service-account-role-arn")
		resolveConflicts, _ := cmd.Flags().GetString("resolve-conflicts")
		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
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
		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := cobrautil.OpenAndReadJsonFile(tags)
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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		addonName, _ := cmd.Flags().GetString("addon-name")
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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		addonName, _ := cmd.Flags().GetString("addon-name")
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

		addonName, _ := cmd.Flags().GetString("addon-name")
		kubernetesVersion, _ := cmd.Flags().GetString("kubernetes-version")
		maxResults, _ := cmd.Flags().GetInt64("max-result")
		nextToken, _ := cmd.Flags().GetString("next-token")
		if addonName != "" {
			describeAddonVersionsInput.AddonName = &addonName
			fmt.Printf(addonName)
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

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl describe-identity-provider-config --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl describe-identity-provider-config --help' to view all commands")
		} else {
			describeIdentityProviderConfigInput.ClusterName = &args[0]

			// json parsing
			var IdentityProviderConfig eks.IdentityProviderConfig
			jsonFileName, _ := cmd.Flags().GetString("identity-provider-config")
			byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &IdentityProviderConfig)
			describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

			// describeIdentityProvicerConfig(describeIdentityProviderConfigInput)
		}
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

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl describe-update --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl describe-update --help' to view all commands")
		} else {
			describeUpdateInput.Name = &args[0]
			describeUpdateInput.UpdateId = &args[1]
			nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")
			addonName, _ := cmd.Flags().GetString("addon-name")
			if nodegroupName != "" {
				describeUpdateInput.NodegroupName = &nodegroupName
			}
			if addonName != "" {
				describeUpdateInput.AddonName = &addonName
			}
			// describeUpdate(describeUpdateInput)
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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		disassociateIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		jsonFileName, _ := cmd.Flags().GetString("identity-provider-config")

		byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &IdentityProviderConfig)
		disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			disassociateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		disassociateIdentityProvicerConfig(disassociateIdentityProviderConfigInput)

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

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		maxResults, _ := cmd.Flags().GetInt64("max-result")
		nextToken, _ := cmd.Flags().GetString("next-token")
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

var listIdentityProviderConfigsCmd = &cobra.Command{
	Use:   "list-identity-provider-configs",
	Short: "A brief description of your command",
	Long: `	
	- list-identity-provider-configs
		hybridctl list-identity-provider-configs --cluster-name

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		listIdentityProviderConfigsInput.ClusterName = &clusterName
		maxResults, _ := cmd.Flags().GetInt64("max-result")
		nextToken, _ := cmd.Flags().GetString("next-token")
		if maxResults != 0 {
			listIdentityProviderConfigsInput.MaxResults = &maxResults
		}
		if nextToken != "" {
			listIdentityProviderConfigsInput.NextToken = &nextToken
		}
		listIdentityProviderConfigs(listIdentityProviderConfigsInput)

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
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		listTagsForResourceInput.ResourceArn = &resourceArn
		listTagsForResource(listTagsForResourceInput)
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

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl list-update --help' to view all commands")
		} else if args[0] == "" {
			fmt.Println("Run 'hybridctl list-update --help' to view all commands")
		} else {
			listUpdateInput.Name = &args[0]
			nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")
			addonName, _ := cmd.Flags().GetString("addon-name")
			maxResults, _ := cmd.Flags().GetInt64("max-result")
			nextToken, _ := cmd.Flags().GetString("next-token")
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
			// listUpdate(listUpdateInput)
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
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		tagResourceInput.ResourceArn = &resourceArn

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := cobrautil.OpenAndReadJsonFile(tags)
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
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		untagResourceInput.ResourceArn = &resourceArn

		keys, _ := cmd.Flags().GetString("tag-keys")
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

var updateAddonCmd = &cobra.Command{
	Use:   "update-addon",
	Short: "A brief description of your command",
	Long: `	
	- update-addon
		hybridctl update-addon  

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		addonName, _ := cmd.Flags().GetString("addon-name")
		addonVersion, _ := cmd.Flags().GetString("addon-version")
		serviceAccountRoleArn, _ := cmd.Flags().GetString("service-account-role-arn")
		resolveConflicts, _ := cmd.Flags().GetString("resolve-conflicts")
		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		updateAddonInput.ClusterName = &clusterName
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

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl update-cluster-config --help' to view all commands")
		} else {

			updateClusterConfigInput.Name = &args[0]

			jsonFileName, _ := cmd.Flags().GetString("resource-vpc-config")
			if jsonFileName != "" {
				var resourcesVpcConfig eks.VpcConfigRequest
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &resourcesVpcConfig)
			}

			jsonFileName, _ = cmd.Flags().GetString("logging")
			if jsonFileName != "" {
				var logging eks.Logging
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &logging)
			}

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")

			if clientRequestToken != "" {
				updateClusterConfigInput.ClientRequestToken = &clientRequestToken
			}

			// updateClusterConfig(updateClusterConfigInput)
		}
	},
}

var updateNodegroupConfigCmd = &cobra.Command{
	Use:   "update-Nodegroup-config",
	Short: "A brief description of your command",
	Long: `	
	- update-Nodegroup-config
		hybridctl update-Nodegroup-config <NodegroupName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 {
			fmt.Println("Run 'hybridctl update-Nodegroup-config --help' to view all commands")
		} else {

			updateNodegroupConfigInput.ClusterName = &args[0]
			updateNodegroupConfigInput.NodegroupName = &args[1]

			jsonFileName, _ := cmd.Flags().GetString("labels")
			if jsonFileName != "" {
				var labels eks.UpdateLabelsPayload
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &labels)
			}

			jsonFileName, _ = cmd.Flags().GetString("taints")
			if jsonFileName != "" {
				var taints eks.UpdateLabelsPayload
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &taints)
			}
			jsonFileName, _ = cmd.Flags().GetString("scaling-config")
			if jsonFileName != "" {
				var scalingConfig eks.NodegroupScalingConfig
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &scalingConfig)
			}

			jsonFileName, _ = cmd.Flags().GetString("update-config")
			if jsonFileName != "" {
				var updateConfig eks.NodegroupUpdateConfig
				byteValue := cobrautil.OpenAndReadJsonFile(jsonFileName)
				json.Unmarshal(byteValue, &updateConfig)
			}

			clientRequestToken, _ := cmd.Flags().GetString("client-request-token")

			if clientRequestToken != "" {
				updateNodegroupConfigInput.ClientRequestToken = &clientRequestToken
			}

			// updateNodegroupConfig(updateNodegroupConfigInput)
		}
	},
}
