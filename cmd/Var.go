package cmd

import (
	"Hybrid_Cloud/hybridctl/util"

	"github.com/aws/aws-sdk-go/service/eks"
	"k8s.io/client-go/kubernetes"
)

var master_config, _ = util.BuildConfigFromFlags("kube-master", "/root/.kube/config")
var master_client = kubernetes.NewForConfigOrDie(master_config)

// eks
var associateEncryptionConfigInput eks.AssociateEncryptionConfigInput
var associateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput
var oidcRequest eks.OidcIdentityProviderConfigRequest
var createAddonInput eks.CreateAddonInput
var deleteAddonInput eks.DeleteAddonInput
var describeAddonInput eks.DescribeAddonInput
var describeAddonVersionsInput eks.DescribeAddonVersionsInput
var describeIdentityProviderConfigInput eks.DescribeIdentityProviderConfigInput
var describeUpdateInput eks.DescribeUpdateInput
var disassociateIdentityProviderConfigInput eks.DisassociateIdentityProviderConfigInput
var listAddonInput eks.ListAddonsInput
var listIdentityProviderConfigsInput eks.ListIdentityProviderConfigsInput
var listTagsForResourceInput eks.ListTagsForResourceInput
var listUpdateInput eks.ListUpdatesInput
var updateAddonInput eks.UpdateAddonInput
var updateClusterConfigInput eks.UpdateClusterConfigInput
var updateNodegroupConfigInput eks.UpdateNodegroupConfigInput
