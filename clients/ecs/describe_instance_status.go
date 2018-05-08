package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	ClusterId            string `position:"Query" name:"ClusterId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeInstanceStatusRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceStatusResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceStatus", "ecs", "")
	resp = &DescribeInstanceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceStatusResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalCount       common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	InstanceStatuses DescribeInstanceStatusInstanceStatusList
}

type DescribeInstanceStatusInstanceStatus struct {
	InstanceId common.String
	Status     common.String
}

type DescribeInstanceStatusInstanceStatusList []DescribeInstanceStatusInstanceStatus

func (list *DescribeInstanceStatusInstanceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceStatusInstanceStatus)
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
