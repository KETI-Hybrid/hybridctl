package nks

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	ncpapi "github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	klog "k8s.io/klog/v2"
)

func SetNksClient() *ncpapi.APIClient {
	keys := ncloud.Keys()
	configure := ncpapi.NewConfiguration("KR", keys)
	apiclient := ncpapi.NewAPIClient(configure)
	if apiclient == nil {
		klog.Infoln("NCP API is Empty")
	}
	return apiclient
}

// func SetVpcClient() *vpcapi.APIClient {
// 	keys := ncloud.Keys()
// 	configure := vpcapi.NewConfiguration(keys)
// 	apiclient := vpcapi.NewAPIClient(configure)
// 	if apiclient == nil {
// 		klog.Infoln("vpc api is empty")
// 	}
// 	return apiclient
// }
