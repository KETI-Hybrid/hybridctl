package nks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ncpapi "github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"

	// monitoring "github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/monitoring"
	klog "k8s.io/klog/v2"
)

var (
	ctx               context.Context
	Error_Clustername = errors.New("Cluster Name not Exist")
	Error_Nodename    = errors.New("Node Name not Exist")
) // Def context, error object
func NksgetClusterlist() {
	var mainlist = ClustersRes{}
	var sublist = Cluster{}
	ncpclient := SetNksClient()
	nksClusterlist, err := ncpclient.V2Api.ClustersGet(ctx)
	if err != nil {
		klog.Error(err)
	}
	lenClusterlist := len(nksClusterlist.Clusters)
	for i := 0; i < lenClusterlist; i++ {
		var l = &NodePoolRes{}
		sublist.SubnetNoList = nil
		sublist.NodePool = nil
		l.SubnetNameList = nil
		l.SubnetNoList = nil
		sublist.Uuid = *nksClusterlist.Clusters[i].Uuid
		sublist.AcgName = *nksClusterlist.Clusters[i].AcgName
		sublist.Name = *nksClusterlist.Clusters[i].Name
		sublist.Capacity = *nksClusterlist.Clusters[i].Capacity
		sublist.ClusterType = *nksClusterlist.Clusters[i].ClusterType
		klog.Infoln("clustertype : ", sublist.ClusterType)
		sublist.NodeCount = *nksClusterlist.Clusters[i].NodeCount
		sublist.NodeMaxCount = *nksClusterlist.Clusters[i].NodeMaxCount
		sublist.CpuCount = *nksClusterlist.Clusters[i].CpuCount
		sublist.MemorySize = *nksClusterlist.Clusters[i].MemorySize
		sublist.CreatedAt = *nksClusterlist.Clusters[i].CreatedAt
		sublist.Endpoint = *nksClusterlist.Clusters[i].Endpoint
		sublist.K8sVersion = *nksClusterlist.Clusters[i].K8sVersion
		sublist.RegionCode = *nksClusterlist.Clusters[i].RegionCode
		klog.Infoln("regioncode : ", sublist.RegionCode)
		sublist.Status = *nksClusterlist.Clusters[i].Status
		sublist.KubeNetworkPlugin = *nksClusterlist.Clusters[i].KubeNetworkPlugin
		sublist.SubnetLbName = *nksClusterlist.Clusters[i].SubnetLbName
		sublist.SubnetLbNo = *nksClusterlist.Clusters[i].SubnetLbNo
		if nksClusterlist.Clusters[i].LbPublicSubnetNo != nil {
			sublist.LbPublicSubnetNo = *nksClusterlist.Clusters[i].LbPublicSubnetNo
		}
		sublist.SubnetName = *nksClusterlist.Clusters[i].SubnetName
		lenSubnoList := len(nksClusterlist.Clusters[i].SubnetNoList)
		if lenSubnoList > 0 {
			for j := 0; j < lenSubnoList; j++ {
				sublist.SubnetNoList = append(sublist.SubnetNoList, *nksClusterlist.Clusters[i].SubnetNoList[j])
				klog.Infoln(*nksClusterlist.Clusters[i].SubnetNoList[j])
			}
		}
		klog.Infoln("subnoliust : ", sublist.SubnetNoList)
		sublist.UpdatedAt = *nksClusterlist.Clusters[i].UpdatedAt
		sublist.VpcName = *nksClusterlist.Clusters[i].VpcName
		sublist.VpcNo = *nksClusterlist.Clusters[i].VpcNo
		klog.Infoln("vpcno : ", sublist.VpcNo)

		sublist.ZoneCode = *nksClusterlist.Clusters[i].ZoneCode
		sublist.ZoneNo = *nksClusterlist.Clusters[i].ZoneNo
		sublist.LoginKeyName = *nksClusterlist.Clusters[i].LoginKeyName
		klog.Infoln("loginkeyName : ", sublist.LoginKeyName)
		sublist.Log.Audit = nksClusterlist.Clusters[i].Log.Audit
		sublist.PublicNetwork = *nksClusterlist.Clusters[i].PublicNetwork
		// NodePool 추가 부분
		lenNodepool := len(nksClusterlist.Clusters[i].NodePool)

		for a := 0; a < lenNodepool; a++ {
			l.InstanceNo = *nksClusterlist.Clusters[i].NodePool[a].InstanceNo
			//IsDefault는 nil로 나와서 뺌
			l.Name = *nksClusterlist.Clusters[i].NodePool[a].Name
			l.NodeCount = *nksClusterlist.Clusters[i].NodePool[a].NodeCount
			lenNodeSubnolist := len(nksClusterlist.Clusters[i].NodePool[a].SubnetNoList)
			if lenNodeSubnolist > 0 {
				for b := 0; b < lenNodeSubnolist; b++ {
					l.SubnetNoList = append(l.SubnetNoList, *nksClusterlist.Clusters[i].NodePool[a].SubnetNoList[b])
				}
			}
			lenNodeSubNamelist := len(nksClusterlist.Clusters[i].NodePool[a].SubnetNameList)
			if lenNodeSubNamelist > 0 {
				for c := 0; c < lenNodeSubNamelist; c++ {
					l.SubnetNameList = append(l.SubnetNameList, *nksClusterlist.Clusters[i].NodePool[a].SubnetNameList[c])
				}
			}
			l.ProductCode = *nksClusterlist.Clusters[i].NodePool[a].ProductCode
			l.Status = *nksClusterlist.Clusters[i].NodePool[a].Status
			l.K8sVersion = *nksClusterlist.Clusters[i].NodePool[a].K8sVersion
			l.Autoscale.Enabled = *nksClusterlist.Clusters[i].NodePool[a].Autoscale.Enabled
			l.Autoscale.Max = *nksClusterlist.Clusters[i].NodePool[a].Autoscale.Max
			l.Autoscale.Min = *nksClusterlist.Clusters[i].NodePool[a].Autoscale.Min
			sublist.NodePool = append(sublist.NodePool, *l)
		}
		mainlist.Clusters = append(mainlist.Clusters, sublist)
	}
	lenMainClusterlist := len(mainlist.Clusters)
	for g := 1; g < lenMainClusterlist+1; g++ {
		fmt.Println(g, "번째 클러스터 : ", mainlist.Clusters[g-1])
	}
}

// func DescribeCluster_ncp(clustername string)
func NksDescribeCluster(clustername string) {

	// uuid, err := NcpgetClustername(clustername)
	// if uuid == "" && err != nil {
	// 	klog.Infoln("Cluster Not Exist!")
	// 	klog.Errorln(err)
	// }
	// cluster_name := clustername
	var pointer_uuid *string
	pointer_uuid = &clustername

	var sublist = Cluster{}
	ncpclient := SetNksClient()
	nksClusterlist, err := ncpclient.V2Api.ClustersUuidGet(ctx, pointer_uuid)
	if err != nil {
		klog.Error(err)
	}
	var l = &NodePoolRes{}
	sublist.SubnetNoList = nil
	sublist.NodePool = nil
	l.SubnetNameList = nil
	l.SubnetNoList = nil
	sublist.Uuid = *nksClusterlist.Cluster.Uuid
	sublist.AcgName = *nksClusterlist.Cluster.AcgName
	sublist.Name = *nksClusterlist.Cluster.Name
	sublist.Capacity = *nksClusterlist.Cluster.Capacity
	sublist.ClusterType = *nksClusterlist.Cluster.ClusterType
	sublist.NodeCount = *nksClusterlist.Cluster.NodeCount
	sublist.NodeMaxCount = *nksClusterlist.Cluster.NodeMaxCount
	sublist.CpuCount = *nksClusterlist.Cluster.CpuCount
	sublist.MemorySize = *nksClusterlist.Cluster.MemorySize
	sublist.CreatedAt = *nksClusterlist.Cluster.CreatedAt
	sublist.Endpoint = *nksClusterlist.Cluster.Endpoint
	sublist.K8sVersion = *nksClusterlist.Cluster.K8sVersion
	sublist.RegionCode = *nksClusterlist.Cluster.RegionCode
	sublist.Status = *nksClusterlist.Cluster.Status
	sublist.KubeNetworkPlugin = *nksClusterlist.Cluster.KubeNetworkPlugin
	sublist.SubnetLbName = *nksClusterlist.Cluster.SubnetLbName
	sublist.SubnetLbNo = *nksClusterlist.Cluster.SubnetLbNo
	if nksClusterlist.Cluster.LbPublicSubnetNo != nil {
		sublist.LbPublicSubnetNo = *nksClusterlist.Cluster.LbPublicSubnetNo
	}
	sublist.SubnetName = *nksClusterlist.Cluster.SubnetName
	lenSubnoList := len(nksClusterlist.Cluster.SubnetNoList)
	if lenSubnoList > 0 {
		for j := 0; j < lenSubnoList; j++ {
			sublist.SubnetNoList = append(sublist.SubnetNoList, *nksClusterlist.Cluster.SubnetNoList[j])
		}
	}
	sublist.UpdatedAt = *nksClusterlist.Cluster.UpdatedAt
	sublist.VpcName = *nksClusterlist.Cluster.VpcName
	sublist.VpcNo = *nksClusterlist.Cluster.VpcNo
	sublist.ZoneCode = *nksClusterlist.Cluster.ZoneCode
	sublist.ZoneNo = *nksClusterlist.Cluster.ZoneNo
	sublist.LoginKeyName = *nksClusterlist.Cluster.LoginKeyName
	sublist.Log.Audit = nksClusterlist.Cluster.Log.Audit
	sublist.PublicNetwork = *nksClusterlist.Cluster.PublicNetwork
	// NodePool 추가 부분
	lenNodepool := len(nksClusterlist.Cluster.NodePool)

	for a := 0; a < lenNodepool; a++ {
		l.InstanceNo = *nksClusterlist.Cluster.NodePool[a].InstanceNo
		//IsDefault는 nil로 나와서 뺌
		l.Name = *nksClusterlist.Cluster.NodePool[a].Name
		l.NodeCount = *nksClusterlist.Cluster.NodePool[a].NodeCount
		lenNodeSubnolist := len(nksClusterlist.Cluster.NodePool[a].SubnetNoList)
		if lenNodeSubnolist > 0 {
			for b := 0; b < lenNodeSubnolist; b++ {
				l.SubnetNoList = append(l.SubnetNoList, *nksClusterlist.Cluster.NodePool[a].SubnetNoList[b])
			}
		}
		lenNodeSubNamelist := len(nksClusterlist.Cluster.NodePool[a].SubnetNameList)
		if lenNodeSubNamelist > 0 {
			for c := 0; c < lenNodeSubNamelist; c++ {
				l.SubnetNameList = append(l.SubnetNameList, *nksClusterlist.Cluster.NodePool[a].SubnetNameList[c])
			}
		}
		l.ProductCode = *nksClusterlist.Cluster.NodePool[a].ProductCode
		l.Status = *nksClusterlist.Cluster.NodePool[a].Status
		l.K8sVersion = *nksClusterlist.Cluster.NodePool[a].K8sVersion
		l.Autoscale.Enabled = *nksClusterlist.Cluster.NodePool[a].Autoscale.Enabled
		l.Autoscale.Max = *nksClusterlist.Cluster.NodePool[a].Autoscale.Max
		l.Autoscale.Min = *nksClusterlist.Cluster.NodePool[a].Autoscale.Min
		sublist.NodePool = append(sublist.NodePool, *l)
		// uuid_String := *uuid
		klog.Infoln("Cluster", sublist.Name, "Infomation : ", sublist)
		klog.Infoln(sublist.SubnetNoList)
	}
}

func NksCreateCluster(clustername string) {
	// klog.Infoln("call ncp create cluster func")
	var clusterset = ClusterDefaultSet{}
	var inputdata = ncpapi.ClusterInputBody{}

	defset, err := ioutil.ReadFile("/root/.ncloud/default_cluster.json")
	if err != nil {
		klog.Infoln(err)
	}
	json.Unmarshal(defset, &clusterset)
	// klog.Infoln(clusterset)
	klog.Infoln("Create cluster Name : ", clustername)
	clusterset.Name = clustername

	inputdata.Name = &clusterset.Name
	inputdata.ClusterType = &clusterset.ClusterType
	inputdata.LoginKeyName = &clusterset.LoginKeyName
	inputdata.RegionCode = &clusterset.RegionCode
	inputdata.ZoneCode = &clusterset.ZoneCode
	inputdata.ZoneNo = &clusterset.ZoneNo
	inputdata.VpcNo = &clusterset.VpcNo
	lensub := len(clusterset.SubnetNoList)
	if lensub > 0 {
		for i := 0; i < lensub; i++ {
			inputdata.SubnetNoList = append(inputdata.SubnetNoList, &clusterset.SubnetNoList[i])
			klog.Infoln(clusterset.SubnetNoList[i])
		}
	}

	klog.Infoln(inputdata.SubnetNoList)
	inputdata.SubnetLbNo = &clusterset.SubnetLbNo

	klog.Infoln(clusterset)

	// klog.Infoln(inputdata)
	// var sublist = Cluster{}
	ncpclient := SetNksClient()
	nksClusterlist, err := ncpclient.V2Api.ClustersPost(ctx, &inputdata)
	if err != nil {
		klog.Error(err)
		klog.Infoln(&nksClusterlist)
		klog.Infoln(nksClusterlist)
	}
	// klog.Infoln("uuid : ", &nksClusterlist)
}

func NksDeleteCluster(uuid string) {
	ncpclient := SetNksClient()
	var pointer_uuid *string
	pointer_uuid = &uuid
	err := ncpclient.V2Api.ClustersUuidDelete(ctx, pointer_uuid)
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infoln(uuid, "Cluster is Delete !")
	}
}

func NksGetClusterName(clustername string) (uuid string, err error) {
	/*
		클러스터 이름을 받으면
		전체 클러스터 정보를 조회 후
		해당 클러스터 이름에 대응되는 uuid를 찾아서
		리턴값으로 uuid를 넘겨주기

		return "" 시 에러
	*/
	var mainlist = ClustersRes{}
	var sublist = Cluster{}
	ncpclient := SetNksClient()
	nksClusterlist, err := ncpclient.V2Api.ClustersGet(ctx)
	if err != nil {
		klog.Error(err)
	}
	lenClusterlist := len(nksClusterlist.Clusters)
	for i := 0; i < lenClusterlist; i++ {
		sublist.Uuid = *nksClusterlist.Clusters[i].Uuid
		sublist.Name = *nksClusterlist.Clusters[i].Name
		mainlist.Clusters = append(mainlist.Clusters, sublist)
	}
	lenMainClusterlist := len(mainlist.Clusters)
	// klog.Infoln(lenMainClusterlist)
	for g := 0; g < lenMainClusterlist; g++ {
		if mainlist.Clusters[g].Name == clustername {
			uuid := mainlist.Clusters[g].Uuid
			return uuid, err
		}
	}
	err = Error_Clustername
	return "", err
}

func NksGetNodeName(clustername string, nodepoolname string) (instanceNo int32, err error) {
	/* 클러스터 이름과 노드풀 이름을 받으면
	해당 클러스터의 노드풀을 조회 후
	이에 대응되는 uuid를 찾아서
	리턴값으로 넘겨주기

	return 0 시 에러
	*/

	var mainlist = NodePoolsRes{}
	var sublist = NodePoolRes{}
	cluster_uuid, err := NksGetClusterName(clustername)
	if err != nil {
		klog.Errorln(err)
	}

	ncpclient := SetNksClient()
	ncpNodelist, err := ncpclient.V2Api.ClustersUuidNodePoolGet(ctx, &cluster_uuid)
	if err != nil {
		klog.Errorln(err)
	}
	lenNodelist := len(ncpNodelist.NodePool)
	for i := 0; i < lenNodelist; i++ {
		sublist.Name = *ncpNodelist.NodePool[i].Name
		sublist.InstanceNo = *ncpNodelist.NodePool[i].InstanceNo
		mainlist.NodePool = append(mainlist.NodePool, sublist)
	}
	lenMainNodelist := len(mainlist.NodePool)
	for g := 0; g < lenMainNodelist; g++ {
		if mainlist.NodePool[g].Name == nodepoolname {
			instanceNo := mainlist.NodePool[g].InstanceNo
			klog.Infoln(instanceNo)
			return instanceNo, err
		}
	}
	err = Error_Nodename
	return 0, err
}

func NksGetNodeList(uuid string) {
	ncpclient := SetNksClient()
	var pointer_uuid *string
	var mainlist = NodePoolsRes{}
	var sublist = NodePoolRes{}
	pointer_uuid = &uuid

	nksClusterlist, err := ncpclient.V2Api.ClustersUuidNodePoolGet(ctx, pointer_uuid)
	if err != nil {
		klog.Error(err)
	}

	lenClusterlist := len(nksClusterlist.NodePool)
	for i := 0; i < lenClusterlist; i++ {
		sublist.SubnetNoList = nil
		sublist.InstanceNo = *nksClusterlist.NodePool[i].InstanceNo
		sublist.Name = *nksClusterlist.NodePool[i].Name
		sublist.NodeCount = *nksClusterlist.NodePool[i].NodeCount

		lenSubnoList := len(nksClusterlist.NodePool[i].SubnetNoList)
		if lenSubnoList > 0 {
			for j := 0; j < lenSubnoList; j++ {
				sublist.SubnetNoList = append(sublist.SubnetNoList, *nksClusterlist.NodePool[i].SubnetNoList[j])
			}
		}
		sublist.ProductCode = *nksClusterlist.NodePool[i].ProductCode
		sublist.Status = *nksClusterlist.NodePool[i].Status
		sublist.Autoscale.Enabled = *nksClusterlist.NodePool[i].Autoscale.Enabled
		sublist.Autoscale.Max = *nksClusterlist.NodePool[i].Autoscale.Max
		sublist.Autoscale.Min = *nksClusterlist.NodePool[i].Autoscale.Min
	}
	mainlist.NodePool = append(mainlist.NodePool, sublist)
	lenMainClusterlist := len(mainlist.NodePool)
	for g := 1; g < lenMainClusterlist+1; g++ {
		fmt.Println(g, "번째 클러스터의 노드풀 : ", mainlist.NodePool[g-1])
	}
}

func NksCreateNode(nodepool_name string, uuid string) {
	ncpclient := SetNksClient()
	var pointer_uuid *string
	pointer_uuid = &uuid

	var nodeset = NodeDefaultSet{}
	var inputdata = ncpapi.NodePoolCreationBody{}

	defset, err := ioutil.ReadFile("/root/.ncloud/default_node.json")
	if err != nil {
		klog.Infoln(err)
	}
	json.Unmarshal(defset, &nodeset)

	inputdata.Name = &nodepool_name
	inputdata.NodeCount = &nodeset.NodeCount
	inputdata.ProductCode = &nodeset.ProductCode

	err = ncpclient.V2Api.ClustersUuidNodePoolPost(ctx, &inputdata, pointer_uuid)
	if err != nil {
		klog.Error(err)
	}
}

func NksDeleteNode(uuid string, instanceNo int32) {
	ncpclient := SetNksClient()
	instanceNo_String := fmt.Sprint(instanceNo)
	klog.Infoln(instanceNo_String)

	err := ncpclient.V2Api.ClustersUuidNodePoolInstanceNoDelete(ctx, &uuid, &instanceNo_String)
	if err != nil {
		klog.Errorln(err)
	} else {
		klog.Infoln("NodePool Delete !")
	}
}

func NksGetKubeconfig(uuid string) (kconfig KubeconfigRes) {
	ncpclient := SetNksClient()
	var kubeconfig = KubeconfigRes{}

	config, err := ncpclient.V2Api.ClustersUuidKubeconfigGet(ctx, &uuid)
	if err != nil {
		klog.Errorln(err)
	}

	klog.Infoln(config)
	kubeconfig.Kubeconfig = *config.Kubeconfig

	return kubeconfig
}

func NksRegisterConfig(uuid string) {
	config := NksGetKubeconfig(uuid)
	stringconfig := fmt.Sprint(config)
	stringconfig = strings.TrimLeft(stringconfig, "{")
	stringconfig = strings.TrimRight(stringconfig, "}")
	byteconfig := []byte(stringconfig)

	file, err := os.OpenFile(
		"configtest",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644))
	if err != nil {
		klog.Infoln(err)
	}
	num, err := file.Write(byteconfig)
	if err != nil {
		klog.Errorln(err)
	}
	klog.Infoln(num)
}

// func NcpRegister

// func NksCreateClustertest(clustername string) {
// 	var client = &http.Client{}
// 	getKey := ncloud.Keys()
// 	ak := getKey.AccessKey
// 	sk := getKey.SecretKey

// 	clusterType := "SVR.VNKS.STAND.C002.M008.NET.SSD.B050.G002"
// 	lk := "keti-nks"
// 	regionCode := "KR"
// 	// vpcnum := 25793
// 	var vpcnum int32 = 25793
// 	var subno int32 = 54819
// 	var lbno int32 = 54820
// 	zoneCode := "KR-1"

// 	var inputdata = ncpapi.ClusterInputBody{}

// 	inputdata.Name = &clustername
// 	inputdata.ClusterType = &clusterType
// 	inputdata.LoginKeyName = &lk
// 	inputdata.RegionCode = &regionCode
// 	inputdata.VpcNo = &vpcnum
// 	inputdata.SubnetNoList = append(inputdata.SubnetNoList, &subno)
// 	inputdata.SubnetLbNo = &lbno
// 	inputdata.ZoneCode = &zoneCode

// 	js, err := json.Marshal(inputdata)
// 	if err != nil {
// 		klog.Errorln(err)
// 	}
// 	buff := bytes.NewBuffer(js)

// 	method := "POST"
// 	url := "https://nks.apigw.ntruss.com/vnks/v2/clusters"
// 	signature, timestamp := NcrCreateSigKey(method, url, ak, sk)

// 	req, err := http.NewRequest(method, url, buff)
// 	if err != nil {
// 		klog.Errorln(err)
// 	}
// 	req.Header.Add("x-ncp-apigw-timestamp", timestamp)
// 	req.Header.Add("x-ncp-iam-access-key", ak)
// 	req.Header.Add("x-ncp-apigw-signature-v1", signature)
// 	req.Header.Add("Content-Type", "application/json")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		klog.Errorln(err)
// 	}
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	str := string(bytes)
// 	klog.Infoln(str)
// }

// func NcrCreateSigKey(Method string, Url string, ak string, sk string) (signature string, timestamp string) {
// 	// getKey := ncloud.Keys()
// 	timestamp = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
// 	signer := hmac.NewSigner(sk, crypto.SHA256)
// 	signature, _ = signer.Sign(Method, Url, ak, timestamp)

// 	return signature, timestamp
// }

// func Ncptestfunc() {
// 	ncpclient := SetVpcClient()
// 	subreq := vpc.GetSubnetListRequest{}
// 	rgcode := "KR"
// 	subreq.RegionCode = &rgcode

// 	nksClusterlist, err := ncpclient.V2Api.GetSubnetList(&subreq)
// 	if err != nil {
// 		klog.Error(err)
// 	}
// 	klog.Infoln(nksClusterlist)
// }
