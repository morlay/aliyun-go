package hsm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId    string
	RegionId      string
	ZoneId        string
	HsmStatus     int
	HsmOem        string
	HsmDeviceType string
	VpcId         string
	VswitchId     string
	Ip            string
	Remark        string
	ExpiredTime   int64
	CreateTime    int64
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

type DescribeInstancesWhiteListList []string

func (list *DescribeInstancesWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
