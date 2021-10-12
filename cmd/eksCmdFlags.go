package cmd

func eksFlags() {
	associateEncryptionConfigCmd.Flags().StringP("encryption-config", "", "", "enter your encryption-config Jsonfile name")
	associateEncryptionConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	associateIdentityProviderConfigCmd.Flags().StringP("oidc", "", "", "enter your oidc Jsonfile name")
	associateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
	associateIdentityProviderConfigCmd.Flags().StringP("tags", "", "", "enter your tags Jsonfile name")

	createAddonCmd.Flags().StringP("addon-version", "", "", "enter addon version")
	createAddonCmd.Flags().StringP("service-account-role-arn", "", "", "enter service account rolearn")
	createAddonCmd.Flags().StringP("resolve-conflicts", "", "", "enter addon version")
	createAddonCmd.Flags().StringP("client-request-token", "", "", "enter client request token")
	createAddonCmd.Flags().StringP("tags", "", "", "enter your tags Jsonfile name")

	describeAddonVersionsCmd.Flags().StringP("addon-name", "", "", "enter kubernetes version")
	describeAddonVersionsCmd.Flags().StringP("kubernetes-version", "", "", "enter kubernetes version")
	describeAddonVersionsCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	describeAddonVersionsCmd.Flags().StringP("next-token", "", "", "enter next token")

	describeIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")

	describeUpdateCmd.Flags().StringP("nodegroup-name", "", "", "enter nodegroupName")
	describeUpdateCmd.Flags().StringP("addon-name", "", "", "enter addonName")

	disassociateIdentityProviderConfigCmd.Flags().StringP("identity-provider-config", "", "", "enter your Jsonfile name")
	disassociateIdentityProviderConfigCmd.Flags().StringP("client-request-token", "", "", "enter client request token")

	listAddonCmd.Flags().Int64P("max-result", "", 0, "enter maxresult")
	listAddonCmd.Flags().StringP("next-token", "", "", "enter next token")

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
