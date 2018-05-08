package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceStatus       string `position:"Query" name:"InstanceStatus"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ChargeType           string `position:"Query" name:"ChargeType"`
}

func (req *DescribeInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstances", "redisa", "")
	resp = &DescribeInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	Instances  DescribeInstancesKVStoreInstanceList
}

type DescribeInstancesKVStoreInstance struct {
	InstanceId          common.String
	InstanceName        common.String
	ConnectionDomain    common.String
	Port                common.Long
	UserName            common.String
	InstanceStatus      common.String
	RegionId            common.String
	Capacity            common.Long
	InstanceClass       common.String
	QPS                 common.Long
	Bandwidth           common.Long
	Connections         common.Long
	ZoneId              common.String
	Config              common.String
	ChargeType          common.String
	NetworkType         common.String
	VpcId               common.String
	VSwitchId           common.String
	PrivateIp           common.String
	CreateTime          common.String
	EndTime             common.String
	HasRenewChangeOrder common.String
	IsRds               bool
	InstanceType        common.String
	ArchitectureType    common.String
	NodeType            common.String
	PackageType         common.String
}

type DescribeInstancesKVStoreInstanceList []DescribeInstancesKVStoreInstance

func (list *DescribeInstancesKVStoreInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesKVStoreInstance)
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
