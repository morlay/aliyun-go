package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInvocationResultsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InvokeRecordStatus   string `position:"Query" name:"InvokeRecordStatus"`
}

func (req *DescribeInvocationResultsRequest) Invoke(client *sdk.Client) (resp *DescribeInvocationResultsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocationResults", "ecs", "")
	resp = &DescribeInvocationResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInvocationResultsResponse struct {
	responses.BaseResponse
	RequestId  string
	Invocation DescribeInvocationResultsInvocation
}

type DescribeInvocationResultsInvocation struct {
	PageSize          int64
	PageNumber        int64
	TotalCount        int64
	InvocationResults DescribeInvocationResultsInvocationResultList
}

type DescribeInvocationResultsInvocationResult struct {
	CommandId          string
	InvokeId           string
	InstanceId         string
	FinishedTime       string
	Output             string
	InvokeRecordStatus string
	ExitCode           int64
}

type DescribeInvocationResultsInvocationResultList []DescribeInvocationResultsInvocationResult

func (list *DescribeInvocationResultsInvocationResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationResultsInvocationResult)
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
