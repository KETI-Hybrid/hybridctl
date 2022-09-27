package cmd

func eksFlags() {

	EKSCreateClusterCmd.Flags().StringP("cluster-name", "c", "", "The unique name to give to your cluster.")
	EKSCreateClusterCmd.MarkFlagRequired("cluster-name")
	EKSCreateClusterCmd.Flags().StringP("role-arn", "", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to Amazon Web Services API operations on your behalf.")
	EKSCreateClusterCmd.MarkFlagRequired("role-arn")
	EKSCreateClusterCmd.Flags().StringP("resources-vpc-config", "", "", "The VPC configuration that’s used by the cluster control plane.")
	EKSCreateClusterCmd.MarkFlagRequired("resources-vpc-config")
	EKSCreateClusterCmd.Flags().StringP("kubernetes-version", "", "", "The desired Kubernetes version for your cluster.")
	EKSCreateClusterCmd.Flags().StringP("region", "", "", "The region to place your cluster")

	EKSDeleteClusterCmd.Flags().StringP("cluster-name", "c", "", "The unique name to give to your cluster.")
	EKSDeleteClusterCmd.MarkFlagRequired("cluster-name")
	EKSDeleteClusterCmd.Flags().StringP("region", "", "", "The region to place your cluster")
	EKSDeleteClusterCmd.MarkFlagRequired("region")

	EKSDescribeClusterCmd.Flags().StringP("cluster-name", "c", "", "The unique name to give to your cluster.")
	EKSDescribeClusterCmd.MarkFlagRequired("cluster-name")
	EKSDescribeClusterCmd.Flags().StringP("region", "", "", "The region to place your cluster")
	EKSDescribeClusterCmd.MarkFlagRequired("region")

	EKSListClusterCmd.Flags().StringP("region", "", "", "The region to place your cluster")
	EKSListClusterCmd.MarkFlagRequired("region")

	EKSUpgradeClusterCmd.Flags().StringP("cluster-name", "c", "", "The unique name to give to your cluster.")
	EKSUpgradeClusterCmd.MarkFlagRequired("cluster-name")
	EKSUpgradeClusterCmd.Flags().StringP("region", "", "", "The region to place your cluster")
	EKSUpgradeClusterCmd.MarkFlagRequired("region")
	EKSUpgradeClusterCmd.Flags().StringP("version", "", "", "The k8s version upgrade your cluster ")
	EKSUpgradeClusterCmd.MarkFlagRequired("version")

	EKSCreateNodegroupCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to create the node group in.")
	EKSCreateNodegroupCmd.MarkFlagRequired("cluster-name")
	EKSCreateNodegroupCmd.Flags().StringP("nodegroup-name", "", "", "The unique name to give your node group.")
	EKSCreateNodegroupCmd.MarkFlagRequired("nodegroup-name")
	EKSCreateNodegroupCmd.Flags().StringP("node-role", "", "", "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.")
	EKSCreateNodegroupCmd.MarkFlagRequired("node-role")
	EKSCreateNodegroupCmd.Flags().StringArray("subnets", []string{}, "The subnets to use for the Auto Scaling group that is created for your node group.")
	EKSCreateNodegroupCmd.MarkFlagRequired("subnets")
	EKSCreateNodegroupCmd.Flags().StringP("region", "", "", "The region to place your cluster")
	EKSCreateNodegroupCmd.MarkFlagRequired("region")

	EKSDeleteNodegroupCmd.Flags().StringP("cluster-name", "c", "", "The name of the Amazon EKS cluster that is associated with your node group.")
	EKSDeleteNodegroupCmd.MarkFlagRequired("cluster-name")
	EKSDeleteNodegroupCmd.Flags().StringP("nodegroup-name", "", "", "The name of the node group to delete.")
	EKSDeleteNodegroupCmd.MarkFlagRequired("nodegroup-name")
	EKSDeleteNodegroupCmd.Flags().StringP("region", "", "", "The region to use.")
	EKSDeleteNodegroupCmd.MarkFlagRequired("region")

	EKSDescribeNodegroupCmd.Flags().StringP("cluster-name", "c", "", "The name of the Amazon EKS cluster that is associated with your node group.")
	EKSDescribeNodegroupCmd.MarkFlagRequired("cluster-name")
	EKSDescribeNodegroupCmd.Flags().StringP("nodegroup-name", "", "", "The name of the node group to describe.")
	EKSDescribeNodegroupCmd.MarkFlagRequired("nodegroup-name")
	EKSDescribeNodegroupCmd.Flags().StringP("region", "", "", "The region to use.")
	EKSDescribeNodegroupCmd.MarkFlagRequired("region")

	EKSListNodegroupCmd.Flags().StringP("cluster-name", "c", "", "The name of the Amazon EKS cluster that is associated with your node group.")
	EKSListNodegroupCmd.MarkFlagRequired("cluster-name")
	EKSListNodegroupCmd.Flags().StringP("region", "", "", "The region to use.")
	EKSListNodegroupCmd.MarkFlagRequired("region")

	EKSAssociateEncryptionConfigCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster that you are associating with encryption configuration.")
	EKSAssociateEncryptionConfigCmd.MarkFlagRequired("cluster-name")
	EKSAssociateEncryptionConfigCmd.Flags().StringP("encryption-config", "", "", "The configuration you are using for encryption.")
	EKSAssociateEncryptionConfigCmd.MarkFlagRequired("encryption-config")
	EKSAssociateEncryptionConfigCmd.Flags().StringP("client-request-token", "", "", "The client request token you are using with the encryption configuration.")

	EKSAssociateIdentityProviderConfigCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to associate the configuration to.")
	EKSAssociateIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	EKSAssociateIdentityProviderConfigCmd.Flags().StringP("oidc", "", "", `An object that represents an OpenID Connect (OIDC) identity provider configuration. 
- JSON Syntax (Enter the path of the json file.) : 
{
	"identityProviderConfigName": "string",
	"issuerUrl": "string",
	"clientId": "string",
	"usernameClaim": "string",
	"usernamePrefix": "string",
	"groupsClaim": "string",
	"groupsPrefix": "string",
	"requiredClaims": {"string": "string"
	...}
}`)
	EKSAssociateIdentityProviderConfigCmd.MarkFlagRequired("oidc")
	EKSAssociateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	EKSAssociateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", `The metadata to apply to the configuration to assist with categorization and organization. 
Each tag consists of a key and an optional value. 
You define both. 
> key -> (string) 
> value -> (string) 
- Shorthand Syntax: 
	KeyName1=string,KeyName2=string
- JSON Syntax (Enter the path of the json file.) : 
	{"string": "string"...} 
`)

	EKSCreateAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to create the add-on for.")
	EKSCreateAddonCmd.MarkFlagRequired("cluster-name")
	EKSCreateAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by \"hybridctl eks addon describe-versions\" CLI")
	EKSCreateAddonCmd.MarkFlagRequired("addon-name")
	EKSCreateAddonCmd.Flags().StringP("addon-version", "", "", "The version of the add-on. The version must match one of the versions returned by \"hybridctl eks addon describe-versions\" CLI")
	EKSCreateAddonCmd.Flags().StringP("service-account-role-arn", "", "", "The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.")
	EKSCreateAddonCmd.Flags().StringP("resolve-conflicts", "", "", "How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on. Possible values: OVERWRITE, NONE")
	EKSCreateAddonCmd.Flags().StringP("client-request-token", "", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	EKSCreateAddonCmd.Flags().StringP("tags", "", "", "The metadata to apply to the cluster to assist with categorization and organization. Shorthand Syntax: KeyName1=string,KeyName2=string")

	EKSDeleteAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to delete the add-on from.")
	EKSDeleteAddonCmd.MarkFlagRequired("cluster-name")
	EKSDeleteAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ListAddons.")
	EKSDeleteAddonCmd.MarkFlagRequired("addon-name")

	EKSDescribeAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	EKSDescribeAddonCmd.MarkFlagRequired("cluster-name")
	EKSDescribeAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ListAddons.")
	EKSDescribeAddonCmd.MarkFlagRequired("addon-name")

	EKSDescribeAddonVersionsCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ListAddons.")
	EKSDescribeAddonVersionsCmd.Flags().StringP("kubernetes-version", "", "", "The Kubernetes versions that the add-on can be used with.")
	EKSDescribeAddonVersionsCmd.Flags().Int64P("max-results", "", 0, "The maximum number of results to return.")
	EKSDescribeAddonVersionsCmd.Flags().StringP("next-token", "", "", "The nextToken value returned from a previous paginated DescribeAddonVersionsRequest where maxResults was used and the results exceeded the value of that parameter.")

	EKSDescribeUpdateCmd.Flags().StringP("name", "c", "", "The name of the Amazon EKS cluster associated with the update.")
	EKSDescribeUpdateCmd.MarkFlagRequired("name")
	EKSDescribeUpdateCmd.Flags().StringP("update-id", "", "", "")
	EKSDescribeUpdateCmd.MarkFlagRequired("update-id")
	EKSDescribeUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	EKSDescribeUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")

	EKSDisassociateIdentityProviderConfigCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to disassociate an identity provider from.")
	EKSDisassociateIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	EKSDisassociateIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", `An object that represents an identity provider configuration.
> type -> (string)
  The type of the identity provider configuration. The only type available is oidc .
> name -> (string)
  The name of the identity provider configuration.
- Shorthand Syntax
  type=string,name=string
- JSON Syntax (Enter the path of the json file.) : 
   {
	  "type": "string",
	  "name": "string"
   }
`)
	EKSDisassociateIdentityProviderConfigCmd.MarkFlagRequired("identity-provider-config")
	EKSDisassociateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")

	EKSListAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	EKSListAddonCmd.MarkFlagRequired("cluster-name")
	EKSListAddonCmd.Flags().Int64P("max-results", "", 0, "The maximum number of add-on results returned by ListAddonsRequest in paginated.")
	EKSListAddonCmd.Flags().StringP("next-token", "", "", "The nextToken value returned from a previous paginated ListAddonsRequest.")

	EKSListIdentityProviderConfigsCmd.Flags().StringP("cluster-name", "c", "", "The cluster name that you want to list identity provider configurations for.")
	EKSListIdentityProviderConfigsCmd.MarkFlagRequired("cluster-name")
	EKSListIdentityProviderConfigsCmd.Flags().Int64P("max-results", "", 0, "The maximum number of add-on results returned by ListAddonsRequest in paginated.")
	EKSListIdentityProviderConfigsCmd.Flags().StringP("next-token", "", "", "The nextToken value returned from a previous paginated IdentityProviderConfigsRequest")

	EKSDescribeIdentityProviderConfigCmd.Flags().StringP("cluster-name", "c", "", "The cluster name that the identity provider configuration is associated to.")
	EKSDescribeIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	EKSDescribeIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", `An object that represents an identity provider configuration.
> type -> (string)
  The type of the identity provider configuration. The only type available is oidc .
> name -> (string)
  The name of the identity provider configuration.
- Shorthand Syntax
  type=string,name=string
- JSON Syntax (Enter the path of the json file.) : 
   {
	  "type": "string",
	  "name": "string"
   }
   `)
	EKSDescribeIdentityProviderConfigCmd.MarkFlagRequired("identity-provider-config")

	EKSListTagsForResourceCmd.Flags().StringP("resource-arn", "", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
	EKSListTagsForResourceCmd.MarkFlagRequired("resource-arn")

	EKSListUpdateCmd.Flags().StringP("name", "c", "", "The name of the Amazon EKS cluster associated with the update.")
	EKSListUpdateCmd.MarkFlagRequired("name")
	EKSListUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	EKSListUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")
	EKSListUpdateCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	EKSListUpdateCmd.Flags().StringP("next-token", "", "", "enter next token")

	EKSTagResourceCmd.Flags().StringP("tags", "t", "", "The tags to add to the resource. A tag is an array of key-value pairs. ")
	EKSTagResourceCmd.MarkFlagRequired("tags")
	EKSTagResourceCmd.Flags().StringP("resource-arn", "", "", "The Amazon Resource Name (ARN) of the resource to which to add tags. Shorthand Syntax: KeyName1=string,KeyName2=string")
	EKSTagResourceCmd.MarkFlagRequired("resource-arn")

	EKSUntagResourceCmd.Flags().StringP("resource-arn", "", "", "The Amazon Resource Name (ARN) of the resource from which to delete tags.")
	EKSUntagResourceCmd.Flags().StringP("tag-keys", "t", "", "The keys of the tags to be removed. Syntax: KeyName1,KeyName2")
	EKSUntagResourceCmd.MarkFlagRequired("tag-keys")
	EKSUntagResourceCmd.MarkFlagRequired("resource-arn")

	EKSUpdateAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	EKSUpdateAddonCmd.MarkFlagRequired("cluster-name")
	EKSUpdateAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ListAddons.")
	EKSUpdateAddonCmd.MarkFlagRequired("addon-name")
	EKSUpdateAddonCmd.Flags().StringP("addon-version", "", "", "The version of the add-on. The version must match one of the versions returned by DescribeAddonVersions")
	EKSUpdateAddonCmd.Flags().StringP("service-account-role-arn", "", "", "The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.")
	EKSUpdateAddonCmd.Flags().StringP("resolve-conflicts", "", "", "How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on. Possible values: OVERWRITE, NONE")
	EKSUpdateAddonCmd.Flags().StringP("client-request-token", "", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")

	EKSUpdateClusterConfigCmd.Flags().StringP("cluster-name", "c", "", "The name of the Amazon EKS cluster associated with the update.")
	EKSUpdateClusterConfigCmd.MarkFlagRequired("cluster-name")
	EKSUpdateClusterConfigCmd.Flags().StringP("resource-vpc-config", "", "", "An object representing the VPC configuration to use for an Amazon EKS cluster.")
	EKSUpdateClusterConfigCmd.Flags().StringP("logging", "", "", "Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs.")
	EKSUpdateClusterConfigCmd.Flags().StringP("client-request-token", "", "", `Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
	By default, cluster control plane logs aren't exported to CloudWatch Logs.`)

	EKSUpdateNodegroupConfigCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	EKSUpdateNodegroupConfigCmd.MarkFlagRequired("cluster-name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	EKSUpdateNodegroupConfigCmd.MarkFlagRequired("nodegroup-name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("labels", "", "", "enter labels jsonfile name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("taints", "", "", "enter taints jsonfile name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("scaling-config", "", "", "enter resource-vpc-config jsonfile name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("update-config", "", "", "enter logging jsonfile name")
	EKSUpdateNodegroupConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

}
