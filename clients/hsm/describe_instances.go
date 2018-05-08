package hsm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	PageSize        int    `position:"Query" name:"PageSize"`
	CurrentPage     int    `position:"Query" name:"CurrentPage"`
	HsmStatus       int    `position:"Query" name:"HsmStatus"`
}

func (req *DescribeInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "DescribeInstances", "hsm", "")
	resp = &DescribeInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId    common.String
	RegionId      common.String
	ZoneId        common.String
	HsmStatus     common.Integer
	HsmOem        common.String
	HsmDeviceType common.String
	VpcId         common.String
	VswitchId     common.String
	Ip            common.String
	Remark        common.String
	ExpiredTime   common.Long
	CreateTime    common.Long
	WhiteList     DescribeInstancesWhiteListList
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

type DescribeInstancesWhiteListList []common.String

func (list *DescribeInstancesWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
