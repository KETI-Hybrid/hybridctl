package nks

type ClustersRes struct {
	// 클러스터 목록
	Clusters []Cluster `json:"clusters"`
}

type NodePoolsRes struct {
	// 노드풀 목록
	NodePool []NodePoolRes `json:"nodePool"`
}

type WorkerNodeRes struct {

	// 워커노드 목록
	Nodes []WorkerNode `json:"nodes"`
}

type Cluster struct {
	// 클러스터 uuid
	Uuid string `json:"uuid"`
	// 클러스터 acg 이름
	AcgName string `json:"acgName"`
	// 클러스터 이름
	Name string `json:"name"`
	// 클러스터 용량
	Capacity string `json:"capacity"`
	// 클러스터 타입
	ClusterType string `json:"clusterType"`
	// 등록된 노드 총 개수
	NodeCount int32 `json:"nodeCount"`
	// 사용할 수 있는 노드의 최대 개수
	NodeMaxCount int32 `json:"nodeMaxCount"`
	// cpu 개수
	CpuCount int32 `json:"cpuCount"`
	// 메모리 용량
	MemorySize int32 `json:"memorySize"`
	// 생성 일자
	CreatedAt string `json:"createdAt"`
	// Control Plane API 주소
	Endpoint string `json:"endpoint"`
	// 쿠버네티스 버전
	K8sVersion string `json:"k8sVersion"`
	// region의 코드
	RegionCode string `json:"regionCode"`
	// 클러스터의 상태
	Status string `json:"status"`
	// CNI Plugin Code
	KubeNetworkPlugin string `json:"kubeNetworkPlugin"`
	// 로드밸런서 전용 서브넷 이름
	SubnetLbName string `json:"subnetLbName"`
	// 로드밸런서 전용 서브넷 No
	SubnetLbNo int32 `json:"subnetLbNo"`
	// 로드밸런서 전용 Public Subnet No
	LbPublicSubnetNo int32 `json:"lbPublicSubnetNo,omitempty"`
	// 서브넷 이름
	SubnetName string `json:"subnetName"`
	// 서브넷 No 목록
	SubnetNoList []int32 `json:"subnetNoList"`
	// 최근 업데이트 일자
	UpdatedAt string `json:"updatedAt"`
	// vpc 이름
	VpcName string `json:"vpcName"`
	// vpc 번호
	VpcNo int32 `json:"vpcNo"`
	// zone 코드
	ZoneCode string `json:"zoneCode,omitempty"`
	// zone 번호
	ZoneNo int32 `json:"zoneNo,omitempty"`
	// 로그인 키 이름
	LoginKeyName string `json:"loginKeyName"`
	// log
	Log ClusterLogInput `json:"log"`
	// Public Network
	PublicNetwork bool `json:"publicNetwork"`
	// 노드풀
	NodePool []NodePoolRes `json:"nodePool"`
}

type ClusterLogInput struct {
	// audit log 설정
	Audit *bool `json:"audit,omitempty"`
}

type NodePoolRes struct {
	// 인스턴스 No
	InstanceNo int32 `json:"instanceNo"`
	// default pool 여부
	IsDefault bool `json:"isDefault"`
	// 노드풀 이름
	Name string `json:"name"`
	// 노드 개수
	NodeCount int32 `json:"nodeCount"`
	// Subnet no list
	SubnetNoList []int32 `json:"subnetNoList,omitempty"`
	// Subnet name list
	SubnetNameList []string `json:"subnetNameList,omitempty"`
	// 상품 코드
	ProductCode string `json:"productCode"`
	// 노드풀 상태
	Status string `json:"status"`
	// k8s version
	K8sVersion string `json:"k8sVersion,omitempty"`
	// 오토스케일
	Autoscale AutoscaleOption `json:"autoscale"`
}

type AutoscaleOption struct {
	// 오토스케일 가능여부
	Enabled bool `json:"enabled"`
	// 오토스케일 가능 최대 노드 수
	Max int32 `json:"max"`
	// 오토스케일 가능 최소 노드 수
	Min int32 `json:"min"`
}

type ClusterDefaultSet struct {
	Name         string  `json:"name"` //>> 매개변수로 받기
	ClusterType  string  `json:"clusterType"`
	LoginKeyName string  `json:"loginKeyName"`
	RegionCode   string  `json:"regionCode"`
	ZoneCode     string  `json:"zoneCode"`
	ZoneNo       int32   `json:"zoneNo"`
	VpcNo        int32   `json:"vpcNo"`
	SubnetNoList []int32 `json:"subnetNoList"`
	SubnetLbNo   int32   `json:"subnetLbno"`
}

type NodeDefaultSet struct {
	Name        string `json:"name"`
	NodeCount   int32  `json:"nodeCount"`
	ProductCode string `json:"productCode"`
}

type KubeconfigRes struct {

	// Kubeconfig
	Kubeconfig string `json:"kubeconfig"`
}

type WorkerNode struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	ServerName         string `json:"serverName"`
	ServerSpec         string `json:"serverSpec"`
	PrivateIp          string `json:"privateIp"`
	PublicIp           string `json:"publicIp"`
	ReturnProtectionYn string `json:"returnProtectionYn"`
	Status             string `json:"status"`
	StatusCode         string `json:"statusCode"`
	StatusIcon         string `json:"statusIcon"`
	StatusColor        string `json:"statusColor"`
	StatusName         string `json:"statusName"`
	ServerImageName    string `json:"serverImageName"`
	CpuCount           int32  `json:"cpuCount"`
	MemorySize         int32  `json:"memorySize"`
	SpecCode           string `json:"specCode"`
	LoginKeyName       string `json:"loginKeyName"`
	K8sStatus          string `json:"k8sStatus"`
	DockerVersion      string `json:"dockerVersion"`
	KernelVersion      string `json:"kernelVersion"`
	NodePoolName       string `json:"nodePoolName"`
}

type bucket struct {
	Bucket string `json:"bucket"`
}

type Ncp_registry struct {
	Registry string `json:"registry"`
}

type Ncp_image struct {
	Registry  string `json:"registry"`
	ImageName string `json:"imageName"`
}

type Ncp_image_description struct {
	Description      string `json:"description"`
	Full_description string `json:"full_description"`
}
