package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceStatusRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	ClusterId            string `position:"Query" name:"ClusterId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeInstanceStatusRequest) Invoke(client *sdk.Client) (response *DescribeInstanceStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceStatus", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceStatusResponse struct {
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	InstanceStatuses DescribeInstanceStatusInstanceStatusList
}

type DescribeInstanceStatusInstanceStatus struct {
	InstanceId string
	Status     string
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
