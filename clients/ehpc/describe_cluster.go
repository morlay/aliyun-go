package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterRequest) Invoke(client *sdk.Client) (resp *DescribeClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "DescribeCluster", "ehs", "")
	resp = &DescribeClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterResponse struct {
	responses.BaseResponse
	RequestId   string
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
	Id               string
	RegionId         string
	Name             string
	Description      string
	Status           string
	OsTag            string
	AccountType      string
	SchedulerType    string
	CreateTime       string
	SecurityGroupId  string
	VSwitchId        string
	VolumeType       string
	VolumeId         string
	VolumeProtocol   string
	VolumeMountpoint string
	RemoteDirectory  string
	HaEnable         bool
	EcsChargeType    string
	KeyPairName      string
	SccClusterId     string
	ClientVersion    string
	ImageOwnerAlias  string
	ImageId          string
	ApplicationInfo  DescribeClusterApplicationInfoItemList
	EcsInfo          DescribeClusterEcsInfo
}

type DescribeClusterApplicationInfoItem struct {
	Tag     string
	Name    string
	Version string
}

type DescribeClusterEcsInfo struct {
	Manager DescribeClusterManager
	Compute DescribeClusterCompute
	Login   DescribeClusterLogin
}

type DescribeClusterManager struct {
	Count        int
	InstanceType string
}

type DescribeClusterCompute struct {
	Count        int
	InstanceType string
}

type DescribeClusterLogin struct {
	Count        int
	InstanceType string
}

type DescribeClusterApplicationInfoItemList []DescribeClusterApplicationInfoItem

func (list *DescribeClusterApplicationInfoItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterApplicationInfoItem)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
