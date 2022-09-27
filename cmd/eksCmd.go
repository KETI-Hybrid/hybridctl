package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/KETI-Hybrid/hybridctl-v1/util"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"
)

func EKSCommonPrintOption(generic interface{}, bytes []byte) {
	ptr := reflect.New(reflect.TypeOf(generic))
	output := ptr.Elem().Interface()
	json.Unmarshal(bytes, &output)
	str, _ := json.MarshalIndent(output, "", "    ")
	if output == "" {
		util.PrintErrMsg(bytes)
	} else {
		fmt.Println(string(str))
	}
}

var EKSClusterCmd = &cobra.Command{
	Use: "cluster",
}

var EKSCreateClusterCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an Amazon EKS control plane.",
	Long: `
		
	hybridctl eks cluster create

	- flags
		--cluster-name <value>
		--role-arn <value>
		--resources-vpc-config <value>
		--region <value>
		[--kubernetes-version <value>]
		[--kubernetes-network-config <value>]
		[--logging <value>]
		[--client-request-token <value>]
		[--tags <value>]
		[--encryption-config]`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		roleArn, _ := cmd.Flags().GetString("role-arn")
		resourcesVPCConfig, _ := cmd.Flags().GetString("resources-vpc-config")
		kubernetesVersion, _ := cmd.Flags().GetString("kubernetes-version")
		region, _ := cmd.Flags().GetString("region")

		createClusterInput.EKSCreateClusterInput.Name = &clusterName          // eks-keti-cluster1
		createClusterInput.EKSCreateClusterInput.RoleArn = &roleArn           // eksClusterRole
		createClusterInput.EKSCreateClusterInput.Version = &kubernetesVersion // 1.22
		createClusterInput.Region = region

		byteValue := util.OpenAndReadJsonFile(resourcesVPCConfig)
		json.Unmarshal(byteValue, &createClusterInput.EKSCreateClusterInput.ResourcesVpcConfig)

		var output eks.CreateClusterOutput
		httpPostUrl := "/eks/cluster/create"
		bytes := util.HTTPPostRequest(createClusterInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDeleteClusterCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the Amazon EKS cluster control plane.",
	Long: `
		
	hybridctl eks cluster delete

	- flags
		--cluster-name <value>
		--region <value>
		`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")

		deleteClusterInput.EKSDeleteClusterInput.Name = &clusterName
		deleteClusterInput.Region = region

		fmt.Println(deleteClusterInput)

		var output eks.DeleteClusterOutput
		httpPostUrl := "/eks/cluster/delete"
		bytes := util.HTTPPostRequest(deleteClusterInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDescribeClusterCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe the Amazon EKS Cluster Control Plane.",
	Long: `
	hybridctl eks cluster describe 

	- flags
	  --cluster-name <value>
	  --region <value>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infoln("describe eks")
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")

		describeClusterInput.EKSDescribeClusterInput.Name = &clusterName
		describeClusterInput.Region = region

		fmt.Println(describeClusterInput)

		var output eks.DescribeClusterOutput
		httpPostUrl := "/eks/cluster/describe"
		bytes := util.HTTPPostRequest(describeClusterInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)

	},
}

var EKSListClusterCmd = &cobra.Command{
	Use:   "list",
	Short: "list the Amazon EKS Cluster Control Plane.",
	Long: `
	hybridctl eks cluster list

	- flags
	  --region <value>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infoln("list eks")
		region, _ := cmd.Flags().GetString("region")

		listClusterInput.Region = region

		fmt.Println(listClusterInput)

		var output eks.DescribeClusterOutput
		httpPostUrl := "/eks/cluster/list"
		bytes := util.HTTPPostRequest(listClusterInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSUpgradeClusterCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the Amazon EKS Cluster Control Plane.",
	Long: `

	hybridctl eks cluster upgrade

	- flags
		--cluster-name <value>
		--region <value>
		--version <value>
		`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infoln("Upgrade Eks")
		region, _ := cmd.Flags().GetString("region")
		cluster, _ := cmd.Flags().GetString("cluster-name")
		version, _ := cmd.Flags().GetString("version")

		upgradeClusterInput.Region = region
		upgradeClusterInput.EKSUpdateClusterVersionInput.Name = &cluster
		upgradeClusterInput.EKSUpdateClusterVersionInput.Version = &version

		var output eks.UpdateClusterVersionInput
		httpPostUrl := "/eks/cluster/upgrade"
		bytes := util.HTTPPostRequest(upgradeClusterInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

// nodegroup

var EKSNodegroupCmd = &cobra.Command{
	Use: "nodegroup",
}

var EKSCreateNodegroupCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a managed node group for an Amazon EKS cluster.",
	Long: `

	hybridctl eks nodegroup create

	- flags
		--cluster-name <value>
		--region <value>
		--nodegroup-name <value>
		--node-role <value>
		--subnets <value>
		`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")
		nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")
		nodeRole, _ := cmd.Flags().GetString("node-role")
		subnets, _ := cmd.Flags().GetStringArray("subnets")

		createNodegroupInput.EKSCreateNodegroupInput.NodegroupName = &nodegroupName
		createNodegroupInput.EKSCreateNodegroupInput.ClusterName = &clusterName
		createNodegroupInput.EKSCreateNodegroupInput.NodeRole = &nodeRole
		createNodegroupInput.Region = region
		for _, subnet := range subnets {
			list := strings.Split(subnet, ",")

			for _, m := range list {
				var tmp string = m
				createNodegroupInput.EKSCreateNodegroupInput.Subnets = append(createNodegroupInput.EKSCreateNodegroupInput.Subnets, &tmp)
			}
		}

		var output eks.CreateNodegroupOutput
		httpPostUrl := "/eks/nodegroup/create"
		bytes := util.HTTPPostRequest(createNodegroupInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDeleteNodegroupCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an Amazon EKS node group for a cluster.",
	Long: `

	hybridctl eks nodegroup delete

	- flags
		--cluster-name <value>
		--nodegroup-name <value>
		--region <value>
		`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")
		nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")

		deleteNodegroupInput.EKSDeleteNodegroupInput.NodegroupName = &nodegroupName
		deleteNodegroupInput.EKSDeleteNodegroupInput.ClusterName = &clusterName
		deleteNodegroupInput.Region = region

		var output eks.DeleteNodegroupInput
		httpPostUrl := "/eks/nodegroup/delete"
		bytes := util.HTTPPostRequest(deleteNodegroupInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDescribeNodegroupCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe an Amazon EKS node group for a cluster.",
	Long: `

	hybridctl eks nodegroup describe

	- flags
		--cluster-name <value>
		--nodegroup-name <value>
		--region <value>
		`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")
		nodegroupName, _ := cmd.Flags().GetString("nodegroup-name")

		describeNodegroupInput.EKSDescribeNodegroupInput.NodegroupName = &nodegroupName
		describeNodegroupInput.EKSDescribeNodegroupInput.ClusterName = &clusterName
		describeNodegroupInput.Region = region

		var output eks.DescribeNodegroupInput
		httpPostUrl := "/eks/nodegroup/describe"
		bytes := util.HTTPPostRequest(describeNodegroupInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSListNodegroupCmd = &cobra.Command{
	Use:   "list",
	Short: "list an Amazon EKS node group for a cluster.",
	Long: `

	hybridctl eks nodegroup list

	- flags
		--cluster-name <value>
		--region <value>
		`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		region, _ := cmd.Flags().GetString("region")

		listNodegroupInput.EKSListNodegroupInput.ClusterName = &clusterName
		listNodegroupInput.Region = region

		var output eks.DescribeNodegroupInput
		httpPostUrl := "/eks/nodegroup/list"
		bytes := util.HTTPPostRequest(listNodegroupInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
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

	hybridctl eks addon create

	- flags
		--cluster-name <value>
		--addon-name <value>
		[--addon-version <value>]
		[--service-account-role-arn <value>]
		[--resolve-conflicts <value>]
		[--client-request-token <value>]
		[--tags <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		createAddonInput.ClusterName = &clusterName

		addonName, _ := cmd.Flags().GetString("addon-name")
		createAddonInput.AddonName = &addonName

		addonVersion, _ := cmd.Flags().GetString("addon-version")
		if addonVersion != "" {
			createAddonInput.AddonVersion = &addonVersion
		}

		serviceAccountRoleArn, _ := cmd.Flags().GetString("service-account-role-arn")
		if serviceAccountRoleArn != "" {
			createAddonInput.ServiceAccountRoleArn = &serviceAccountRoleArn
		}

		resolveConflicts, _ := cmd.Flags().GetString("resolve-conflicts")
		if resolveConflicts != "" {
			createAddonInput.ResolveConflicts = &resolveConflicts
		}

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			createAddonInput.ClientRequestToken = &clientRequestToken
		}

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap = map[string]*string{}
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			fmt.Println(byteValue)
			if len(byteValue) == 0 {
				tagList := strings.Split(tags, ",")
				for _, tag := range tagList {
					pair := strings.Split(tag, "=")
					if len(pair) != 2 {
						fmt.Println("Flag Tags format is wrong")
						return
					}
					str := pair[1]
					tagsMap[pair[0]] = &str
				}
				createAddonInput.Tags = tagsMap
			} else {
				json.Unmarshal(byteValue, &tagsMap)
				createAddonInput.Tags = tagsMap
			}
		}

		var output eks.CreateAddonOutput
		httpPostUrl := "/eks/addon/create"
		bytes := util.HTTPPostRequest(createAddonInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDeleteAddonCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an Amazon EKS add-on.",
	Long: `	

	hybridctl eks addon delete 

	- flags
		--cluster-name <value>
		--addon-name <value>`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		deleteAddonInput.ClusterName = &clusterName

		addonName, _ := cmd.Flags().GetString("addon-name")
		deleteAddonInput.AddonName = &addonName

		var output eks.DeleteAddonOutput
		httpPostUrl := "/eks/addon/delete"
		bytes := util.HTTPPostRequest(deleteAddonInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSDescribeAddonCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describes an Amazon EKS add-on.",
	Long: `	

	hybridctl eks addon describe

	-flags
		--cluster-name <value>
		--addon-name <value>`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		describeAddonInput.ClusterName = &clusterName

		addonName, _ := cmd.Flags().GetString("addon-name")
		describeAddonInput.AddonName = &addonName

		var output eks.DescribeAddonOutput
		httpPostUrl := "/eks/addon/describe"
		bytes := util.HTTPPostRequest(describeAddonInput, httpPostUrl)
		EKSCommonPrintOption(reflect.TypeOf(output), bytes)
	},
}

var EKSDescribeAddonVersionsCmd = &cobra.Command{
	Use:   "describe-versions",
	Short: "Describes the Kubernetes versions that the add-on can be used with.",
	Long: `	

	hybridctl eks addon describe-versions

	- flags
		[--kubernetes-version <value>]
		[--addon-name <value>]
		[--next-token <value>]
		[--max-results <value>]`,

	Run: func(cmd *cobra.Command, args []string) {

		addonName, _ := cmd.Flags().GetString("addon-name")
		if addonName != "" {
			describeAddonVersionsInput.AddonName = &addonName
		}

		kubernetesVersion, _ := cmd.Flags().GetString("kubernetes-version")
		if kubernetesVersion != "" {
			describeAddonVersionsInput.KubernetesVersion = &kubernetesVersion
		}

		maxResults, _ := cmd.Flags().GetInt64("max-results")
		if maxResults != 0 {
			if maxResults < 1 || maxResults > 100 {
				fmt.Println("MaxResults can be between 1 and 100.")
				return
			} else {
				describeAddonVersionsInput.MaxResults = &maxResults
			}
		}
		nextToken, _ := cmd.Flags().GetString("next-token")
		if nextToken != "" {
			describeAddonVersionsInput.NextToken = &nextToken
		}

		var output eks.DescribeAddonVersionsOutput
		httpPostUrl := "/eks/addon/describe-versions"
		bytes := util.HTTPPostRequest(describeAddonVersionsInput, httpPostUrl)
		json.Unmarshal(bytes, &output)
		if output.Addons == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}
	},
}

var EKSListAddonCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available add-ons.",
	Long: `	

	hybridctl eks addon list

	- flags
		--cluster-name <value>
		[--max-results <value>]
		[--next-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		listAddonInput.ClusterName = &clusterName

		maxResults, _ := cmd.Flags().GetInt64("max-results")
		if maxResults != 0 {
			if maxResults < 1 || maxResults > 100 {
				fmt.Println("MaxResults can be between 1 and 100.")
				return
			} else {
				listAddonInput.MaxResults = &maxResults
			}
		}

		nextToken, _ := cmd.Flags().GetString("next-token")
		if nextToken != "" {
			listAddonInput.NextToken = &nextToken
		}

		var output eks.ListAddonsOutput
		httpPostUrl := "/eks/addon/list"
		bytes := util.HTTPPostRequest(listAddonInput, httpPostUrl)
		json.Unmarshal(bytes, &output)
		if output.Addons == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output.String())
		}
	},
}

var EKSUpdateAddonCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an Amazon EKS add-on.",
	Long: `	
	
	hybridctl eks addon update 

	- flags
		--cluster-name <value>
		--addon-name <value>
		[--addon-version <value>]
		[--service-account-role-arn <value>]
		[--resolve-conflicts <value>]
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("cluster-name")
		updateAddonInput.ClusterName = &name

		addonName, _ := cmd.Flags().GetString("addon-name")
		updateAddonInput.AddonName = &addonName

		addonVersion, _ := cmd.Flags().GetString("addon-version")
		if addonVersion != "" {
			updateAddonInput.AddonVersion = &addonVersion
		}

		serviceAccountRoleArn, _ := cmd.Flags().GetString("service-account-role-arn")
		if serviceAccountRoleArn != "" {
			updateAddonInput.ServiceAccountRoleArn = &serviceAccountRoleArn
		}

		resolveConflicts, _ := cmd.Flags().GetString("resolve-conflicts")
		if resolveConflicts != "" {
			updateAddonInput.ResolveConflicts = &resolveConflicts
		}

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			updateAddonInput.ClientRequestToken = &clientRequestToken
		}

		var output eks.UpdateAddonOutput
		httpPostUrl := "/eks/addon/update"
		bytes := util.HTTPPostRequest(updateAddonInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSIdentityProviderConfigCmd = &cobra.Command{
	Use:   "identity-provider-config",
	Short: "",
	Long:  "",
}

var EKSAssociateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate an identity provider configuration to a cluster.",
	Long: `	

	hybridctl eks identity-provider-config associate

	- flags
		--cluster-name <value>
		--oidc <value>
		[--tags <value>]
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		associateIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		oidc, _ := cmd.Flags().GetString("oidc")
		byteValue := util.OpenAndReadJsonFile(oidc)
		json.Unmarshal(byteValue, &oidcRequest)
		associateIdentityProviderConfigInput.Oidc = &oidcRequest

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			associateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap map[string]*string
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			if byteValue == nil {
				tagList := strings.Split(tags, ",")
				for _, tag := range tagList {
					pair := strings.Split(tag, "=")
					if len(pair) != 2 {
						fmt.Println("Flag Tags format is wrong")
						return
					}
					tagsMap[pair[0]] = tagsMap[pair[1]]
				}
				associateIdentityProviderConfigInput.Tags = tagsMap
			} else {
				json.Unmarshal(byteValue, &tagsMap)
				associateIdentityProviderConfigInput.Tags = tagsMap
			}
		}

		var output eks.AssociateIdentityProviderConfigOutput
		httpPostUrl := "/eks/identity-provider-config/associate"
		bytes := util.HTTPPostRequest(associateIdentityProviderConfigInput, httpPostUrl)
		json.Unmarshal(bytes, &output)
		if output.Tags == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}

	},
}

var EKSDisassociateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "disassociate",
	Short: "Disassociates an identity provider configuration from a cluster.",
	Long: `	
	
	hybridctl eks identity-provider-config disassociate

	- flags
		--cluster-name <value>
		--identity-provider-config <value>
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		disassociateIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		identityProviderConfig, _ := cmd.Flags().GetString("identity-provider-config")
		fmt.Println(identityProviderConfig)
		byteValue := util.OpenAndReadJsonFile(identityProviderConfig)
		if len(byteValue) == 0 {
			list := strings.Split(identityProviderConfig, ",")

			tmp := map[string]string{}
			for _, e := range list {
				value := strings.Split(e, "=")
				tmp[value[0]] = value[1]
			}

			typ, type_ok := tmp["type"]
			name, name_ok := tmp["name"]
			fmt.Println(typ, name)
			if len(list) == 2 || type_ok || name_ok {
				IdentityProviderConfig.Name = &name
				IdentityProviderConfig.Type = &typ
				disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig
			} else {
				fmt.Println("Flag identityProviderConfig format is wrong")
				return
			}
		} else {
			json.Unmarshal(byteValue, &IdentityProviderConfig)
			if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
				fmt.Println("identityProviderConfig format is wrong.")
				return
			}
			disassociateIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig
		}

		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			disassociateIdentityProviderConfigInput.ClientRequestToken = &clientRequestToken
		}

		var output eks.DisassociateIdentityProviderConfigOutput
		httpPostUrl := "/eks/identity-provider-config/disassociate"
		bytes := util.HTTPPostRequest(disassociateIdentityProviderConfigInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSListIdentityProviderConfigsCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of identity provider configurations.",
	Long: `	

	hybridctl eks identity-provider-config list

	- flags
		--cluster-name <value>
		[--max-results <value>]
		[--next-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		clusterName, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			log.Println(err)
			return
		}
		listIdentityProviderConfigsInput.ClusterName = &clusterName
		maxResults, err := cmd.Flags().GetInt64("max-results")
		if err != nil {
			log.Println(err)
			return
		}

		if maxResults != 0 {
			if maxResults < 1 || maxResults > 100 {
				fmt.Println("MaxResults can be between 1 and 100.")
				return
			} else {
				listIdentityProviderConfigsInput.MaxResults = &maxResults
			}
		}

		nextToken, err := cmd.Flags().GetString("next-token")
		if err != nil {
			log.Println(err)
			return
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

var EKSDescribeIdentityProviderConfigCmd = &cobra.Command{
	Use:   "describe",
	Short: "Returns descriptive information about an identity provider configuration.",
	Long: `	

	hybridctl eks identity-provider-config describe

	- flags
		--cluster-name <value>
		--identity-provider-config <value>`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			log.Println(err)
			return
		}
		describeIdentityProviderConfigInput.ClusterName = &clusterName

		// json parsing
		var IdentityProviderConfig eks.IdentityProviderConfig
		identityProviderConfig, _ := cmd.Flags().GetString("identity-provider-config")

		byteValue := util.OpenAndReadJsonFile(identityProviderConfig)
		if len(byteValue) == 0 {
			list := strings.Split(identityProviderConfig, ",")

			tmp := map[string]string{}
			for _, e := range list {
				value := strings.Split(e, "=")
				tmp[value[0]] = value[1]
			}

			typ, type_ok := tmp["type"]
			name, name_ok := tmp["name"]

			if len(list) == 2 || type_ok || name_ok {
				IdentityProviderConfig.Name = &name
				IdentityProviderConfig.Type = &typ
				describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig
			} else {
				fmt.Println("Flag identityProviderConfig format is wrong")
				return
			}
		} else {
			json.Unmarshal(byteValue, &IdentityProviderConfig)
			if (IdentityProviderConfig == eks.IdentityProviderConfig{}) {
				fmt.Println("identityProviderConfig format is wrong.")
				return
			}
			describeIdentityProviderConfigInput.IdentityProviderConfig = &IdentityProviderConfig
		}

		httpPostUrl := "/eks/identity-provider-config/describe"
		bytes := util.HTTPPostRequest(describeIdentityProviderConfigInput, httpPostUrl)
		var output eks.DescribeIdentityProviderConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var EKSEncryptionConfigCmd = &cobra.Command{
	Use:   "encryption-config",
	Short: "Associate encryption configuration to an existing cluster.",
	Long: `	
	You can use this API to enable encryption on existing clusters which do not have encryption already enabled. 
	This allows you to implement a defense-in-depth security strategy without migrating applications to new Amazon EKS clusters.`,
}

var EKSAssociateEncryptionConfigCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate encryption configuration to an existing cluster.",
	Long: `	

	hybridctl eks encryption-config associate <clusterName> --encryption-config <jsonfile>

	- flags
		--cluster-name <value>
		--encryption-config <jsonfile>
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		clientRequestToken, _ := cmd.Flags().GetString("client-request-token")
		if clientRequestToken != "" {
			associateEncryptionConfigInput.ClientRequestToken = &clientRequestToken
		}

		clusterName, _ := cmd.Flags().GetString("cluster-name")
		associateEncryptionConfigInput.ClusterName = &clusterName

		// json parsing
		jsonFileName, _ := cmd.Flags().GetString("encryption-config")
		var encryptionConfig []*eks.EncryptionConfig
		byteValue := util.OpenAndReadJsonFile(jsonFileName)
		json.Unmarshal(byteValue, &encryptionConfig)
		associateEncryptionConfigInput.EncryptionConfig = encryptionConfig

		var output eks.AssociateEncryptionConfigOutput
		httpPostUrl := "/eks/encryption-config/associate"
		bytes := util.HTTPPostRequest(associateEncryptionConfigInput, httpPostUrl)
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSListTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "List the tags for an Amazon EKS resource.",
	Long: `	
	
	hybridctl eks resource list-tags

	- flags
		--resource-arn <value>`,

	Run: func(cmd *cobra.Command, args []string) {
		resourceArn, _ := cmd.Flags().GetString("resource-arn")
		listTagsForResourceInput.ResourceArn = &resourceArn

		var output eks.ListTagsForResourceOutput
		httpPostUrl := "/eks/resource/list"
		bytes := util.HTTPPostRequest(listTagsForResourceInput, httpPostUrl)
		json.Unmarshal(bytes, &output)
		if output.Tags == nil {
			util.PrintErrMsg(bytes)
		} else {
			fmt.Println(output)
		}
	},
}

var EKSResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "",
	Long:  "",
}

var EKSTagResourceCmd = &cobra.Command{
	Use:   "tag",
	Short: "Associates the specified tags to a resource with the specified resourceArn.",
	Long: `	
	
	hybridctl eks resource tag

	- flags
		--resource-arn <value>
		--tags <value>`,

	Run: func(cmd *cobra.Command, args []string) {
		var tagResourceInput eks.TagResourceInput
		resourceArn, err := cmd.Flags().GetString("resource-arn")
		if err != nil {
			fmt.Println(err)
			return
		}
		tagResourceInput.ResourceArn = &resourceArn

		tags, _ := cmd.Flags().GetString("tags")
		var tagsMap = map[string]*string{}
		if tags != "" {
			byteValue := util.OpenAndReadJsonFile(tags)
			if len(byteValue) == 0 {
				tagList := strings.Split(tags, ",")
				for _, tag := range tagList {
					pair := strings.Split(tag, "=")
					if len(pair) != 2 {
						fmt.Println("Flag Tags format is wrong")
						return
					}
					str := pair[1]
					tagsMap[pair[0]] = &str
				}
				tagResourceInput.Tags = tagsMap
			} else {
				json.Unmarshal(byteValue, &tagsMap)
				tagResourceInput.Tags = tagsMap
			}
		}

		httpPostUrl := "/eks/resource/tag"
		bytes := util.HTTPPostRequest(tagResourceInput, httpPostUrl)
		util.PrintErrMsg(bytes)
	},
}

var EKSUntagResourceCmd = &cobra.Command{
	Use:   "untag",
	Short: "Deletes specified tags from a resource.",
	Long: `	
	
	hybridctl eks resource untag

	- flags
		--resource-arn <value>
		--tag-keys <value>`,

	Run: func(cmd *cobra.Command, args []string) {

		var untagResourceInput eks.UntagResourceInput
		resourceArn, err := cmd.Flags().GetString("resource-arn")
		if err != nil {
			fmt.Println(err)
			return
		}
		untagResourceInput.ResourceArn = &resourceArn

		keys, err := cmd.Flags().GetString("tag-keys")
		if err != nil {
			fmt.Println(err)
			return
		}
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
var EKSDescribeUpdateCmd = &cobra.Command{
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
		if err != nil {
			log.Println(err)
			return
		}
		describeUpdateInput.Name = &name
		updateId, err := cmd.Flags().GetString("update-id")
		if err != nil {
			log.Println(err)
			return
		}
		describeUpdateInput.UpdateId = &updateId
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		if err != nil {
			log.Println(err)
			return
		}
		if nodegroupName != "" {
			describeUpdateInput.NodegroupName = &nodegroupName
		}
		addonName, err := cmd.Flags().GetString("addon-name")
		if err != nil {
			log.Println(err)
			return
		}
		if addonName != "" {
			describeUpdateInput.AddonName = &addonName
		}

		httpPostUrl := "/eks/describe/update"
		bytes := util.HTTPPostRequest(describeUpdateInput, httpPostUrl)
		var output eks.DescribeUpdateOutput
		EKSCommonPrintOption(output, bytes)
	},
}

var EKSListUpdateCmd = &cobra.Command{
	Use:   "list-updates",
	Short: "A brief description of your command",
	Long: `	

	hybridctl list-update <clusterName> 

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Println(err)
			return
		}
		listUpdateInput.Name = &name
		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		if err != nil {
			log.Println(err)
			return
		}
		addonName, err := cmd.Flags().GetString("addon-name")
		if err != nil {
			log.Println(err)
			return
		}
		maxResults, err := cmd.Flags().GetInt64("max-result")
		if err != nil {
			log.Println(err)
			return
		}
		nextToken, err := cmd.Flags().GetString("next-token")
		if err != nil {
			log.Println(err)
			return
		}
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

var EKSClusterConfigCmd = &cobra.Command{
	Use:   "cluster-config",
	Short: "",
	Long:  "",
}

var EKSUpdateClusterConfigCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an Amazon EKS cluster configuration.",
	Long: `	

	hybridctl eks cluster-config update

	- flags
		--cluster-name <value>
		[--resources-vpc-config <value>]
		[--logging <value>]
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			log.Println(err)
			return
		}
		updateClusterConfigInput.Name = &clusterName

		jsonFile, err := cmd.Flags().GetString("resource-vpc-config")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFile != "" {
			var resourcesVpcConfig eks.VpcConfigRequest
			byteValue := util.OpenAndReadJsonFile(jsonFile)
			json.Unmarshal(byteValue, &resourcesVpcConfig)
			updateClusterConfigInput.ResourcesVpcConfig = &resourcesVpcConfig
		}

		jsonFile, err = cmd.Flags().GetString("logging")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFile != "" {
			var logging eks.Logging
			byteValue := util.OpenAndReadJsonFile(jsonFile)
			json.Unmarshal(byteValue, &logging)
			updateClusterConfigInput.Logging = &logging
		}

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		if err != nil {
			log.Println(err)
			return
		}

		if clientRequestToken != "" {
			updateClusterConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/cluster-config/update"
		bytes := util.HTTPPostRequest(updateClusterConfigInput, httpPostUrl)
		var output eks.UpdateClusterConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}

var EKSNodegroupConfigCmd = &cobra.Command{
	Use:   "nodegroup-config",
	Short: "",
	Long:  "",
}

var EKSUpdateNodegroupConfigCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an Amazon EKS managed node group configuration.",
	Long: `	

	hybridctl eks nodegroup-config update 

	- flags
		--cluster-name <value>
		--nodegroup-name <value>
		[--labels <value>]
		[--taints <value>]
		[--scaling-config <value>]
		[--update-config <value>]
		[--client-request-token <value>]`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		clusterName, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			log.Println(err)
			return
		}
		updateNodegroupConfigInput.ClusterName = &clusterName

		nodegroupName, err := cmd.Flags().GetString("nodegroup-name")
		if err != nil {
			log.Println(err)
			return
		}
		updateNodegroupConfigInput.NodegroupName = &nodegroupName
		jsonFileName, err := cmd.Flags().GetString("labels")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFileName != "" {
			var labels eks.UpdateLabelsPayload
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &labels)
			updateNodegroupConfigInput.Labels = &labels
		}

		jsonFileName, err = cmd.Flags().GetString("taints")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFileName != "" {
			var taints eks.UpdateLabelsPayload
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &taints)
			// updateNodegroupConfigInput.Taints = taints
		}
		jsonFileName, err = cmd.Flags().GetString("scaling-config")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFileName != "" {
			var scalingConfig eks.NodegroupScalingConfig
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &scalingConfig)
		}

		jsonFileName, err = cmd.Flags().GetString("update-config")
		if err != nil {
			log.Println(err)
			return
		}
		if jsonFileName != "" {
			var updateConfig eks.NodegroupUpdateConfig
			byteValue := util.OpenAndReadJsonFile(jsonFileName)
			json.Unmarshal(byteValue, &updateConfig)
		}

		clientRequestToken, err := cmd.Flags().GetString("client-request-token")
		if err != nil {
			log.Println(err)
			return
		}
		if clientRequestToken != "" {
			updateNodegroupConfigInput.ClientRequestToken = &clientRequestToken
		}

		httpPostUrl := "/eks/nodegroup-config/update"
		bytes := util.HTTPPostRequest(updateNodegroupConfigInput, httpPostUrl)
		var output eks.UpdateNodegroupConfigOutput
		EKSCommonPrintOption(output, bytes)

	},
}
