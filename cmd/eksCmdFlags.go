package cmd

func eksFlags() {
	associateEncryptionConfigCmd.Flags().StringP("encryption-config", "", "", "enter your encryption-config Jsonfile name")
	associateEncryptionConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	associateIdentityProviderConfigCmd.Flags().StringP("oidc", "", "", "enter your oidc Jsonfile name")
	associateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
	associateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter your tags Jsonfile name")

	createAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster to create the add-on for.")
	createAddonCmd.MarkFlagRequired("cluster-name")
	createAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ` ListAddons")
	createAddonCmd.MarkFlagRequired("addon-name")
	createAddonCmd.Flags().StringP("addon-version", "", "", "The version of the add-on. The version must match one of the versions returned")
	createAddonCmd.Flags().StringP("service-account-role-arn", "", "", "The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.")
	createAddonCmd.Flags().StringP("resolve-conflicts", "", "", "How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.")
	createAddonCmd.Flags().StringP("client-request-token", "", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	createAddonCmd.Flags().StringP("tags", "", "", "The metadata to apply to the cluster to assist with categorization and organization.")

	deleteAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	deleteAddonCmd.MarkFlagRequired("cluster-name")
	deleteAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ` ListAddons")
	deleteAddonCmd.MarkFlagRequired("addon-name")

	describeAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	describeAddonCmd.MarkFlagRequired("cluster-name")
	describeAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ` ListAddons")
	describeAddonCmd.MarkFlagRequired("addon-name")

	describeAddonVersionsCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ` ListAddons")
	describeAddonVersionsCmd.MarkFlagRequired("addon-name")
	describeAddonVersionsCmd.Flags().StringP("kubernetes-version", "", "", "The Kubernetes versions that the add-on can be used with.")
	describeAddonVersionsCmd.Flags().Int64P("max-result", "", 0, "The maximum number of results to return.")
	describeAddonVersionsCmd.Flags().StringP("next-token", "", "", "This token should be treated as an opaque identifier that is used only to retrieve the next items in a list and not for other programmatic purposes.")

	describeIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")

	describeUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	describeUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")

	disassociateIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")
	disassociateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	listAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	listAddonCmd.MarkFlagRequired("cluster-name")
	listAddonCmd.Flags().Int64P("max-result", "", 0, "The maximum number of add-on results returned by ListAddonsRequest in paginated")
	listAddonCmd.Flags().StringP("next-token", "", "", "The nextToken value returned from a previous paginated ListAddonsRequest")

	listIdentityProviderConfigsCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	listIdentityProviderConfigsCmd.Flags().StringP("next-token", "", "", "enter next token")

	listTagsForResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")

	listUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	listUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")
	listUpdateCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	listUpdateCmd.Flags().StringP("next-token", "", "", "enter next token")

	tagResourceCmd.Flags().StringP("tags", "t", "", "enter your tags Jsonfile name")
	tagResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")
	tagResourceCmd.MarkPersistentFlagRequired("tags")
	tagResourceCmd.MarkPersistentFlagRequired("resource-arn")

	untagResourceCmd.Flags().StringP("resource-arn", "", "", "Enter resource-arn")
	untagResourceCmd.Flags().StringP("tag-keys", "t", "", "enter your tag-keys list")
	untagResourceCmd.MarkPersistentFlagRequired("tag-keys")
	untagResourceCmd.MarkPersistentFlagRequired("resource-arn")

	updateAddonCmd.Flags().StringP("cluster-name", "c", "", "The name of the cluster.")
	updateAddonCmd.MarkFlagRequired("cluster-name")
	updateAddonCmd.Flags().StringP("addon-name", "a", "", "The name of the add-on. The name must match one of the names returned by ` ListAddons")
	updateAddonCmd.MarkFlagRequired("addon-name")
	updateAddonCmd.Flags().StringP("addon-version", "", "", "enter addon version")
	updateAddonCmd.Flags().StringP("service-account-role-arn", "", "", "enter service account rolearn")
	updateAddonCmd.Flags().StringP("resolve-conflicts", "", "", "enter addon version")
	updateAddonCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	updateClusterConfigCmd.Flags().StringP("resource-vpc-config", "", "", "enter resource-vpc-config jsonfile name")
	updateClusterConfigCmd.Flags().StringP("logging", "", "", "enter logging jsonfile name")
	updateClusterConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	updateNodegroupConfigCmd.Flags().StringP("labels", "", "", "enter labels jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("taints", "", "", "enter taints jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("scaling-config", "", "", "enter resource-vpc-config jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("update-config", "", "", "enter logging jsonfile name")
	updateNodegroupConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

}
