package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
	Id               common.String
	RegionId         common.String
	Name             common.String
	Description      common.String
	Status           common.String
	OsTag            common.String
	AccountType      common.String
	SchedulerType    common.String
	CreateTime       common.String
	SecurityGroupId  common.String
	VSwitchId        common.String
	VolumeType       common.String
	VolumeId         common.String
	VolumeProtocol   common.String
	VolumeMountpoint common.String
	RemoteDirectory  common.String
	HaEnable         bool
	EcsChargeType    common.String
	KeyPairName      common.String
	SccClusterId     common.String
	ClientVersion    common.String
	ImageOwnerAlias  common.String
	ImageId          common.String
	ApplicationInfo  DescribeClusterApplicationInfoItemList
	EcsInfo          DescribeClusterEcsInfo
}

type DescribeClusterApplicationInfoItem struct {
	Tag     common.String
	Name    common.String
	Version common.String
}

type DescribeClusterEcsInfo struct {
	Manager DescribeClusterManager
	Compute DescribeClusterCompute
	Login   DescribeClusterLogin
}

type DescribeClusterManager struct {
	Count        common.Integer
	InstanceType common.String
}

type DescribeClusterCompute struct {
	Count        common.Integer
	InstanceType common.String
}

type DescribeClusterLogin struct {
	Count        common.Integer
	InstanceType common.String
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
