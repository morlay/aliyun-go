package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClassicLinkInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             string `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           string `position:"Query" name:"PageNumber"`
}

func (req *DescribeClassicLinkInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeClassicLinkInstancesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeClassicLinkInstances", "ecs", "")
	resp = &DescribeClassicLinkInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClassicLinkInstancesResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Links      DescribeClassicLinkInstancesLinkList
}

type DescribeClassicLinkInstancesLink struct {
	InstanceId string
	VpcId      string
}

type DescribeClassicLinkInstancesLinkList []DescribeClassicLinkInstancesLink

func (list *DescribeClassicLinkInstancesLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClassicLinkInstancesLink)
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
