package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	Instances  DescribeInstancesKVStoreInstanceList
}

type DescribeInstancesKVStoreInstance struct {
	InstanceId          string
	InstanceName        string
	ConnectionDomain    string
	Port                int64
	UserName            string
	InstanceStatus      string
	RegionId            string
	Capacity            int64
	InstanceClass       string
	QPS                 int64
	Bandwidth           int64
	Connections         int64
	ZoneId              string
	Config              string
	ChargeType          string
	NetworkType         string
	VpcId               string
	VSwitchId           string
	PrivateIp           string
	CreateTime          string
	EndTime             string
	HasRenewChangeOrder bool
	IsRds               bool
	InstanceType        string
	ArchitectureType    string
	NodeType            string
	PackageType         string
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
