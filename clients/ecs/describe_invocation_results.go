package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	Invocation DescribeInvocationResultsInvocation
}

type DescribeInvocationResultsInvocation struct {
	PageSize          common.Long
	PageNumber        common.Long
	TotalCount        common.Long
	InvocationResults DescribeInvocationResultsInvocationResultList
}

type DescribeInvocationResultsInvocationResult struct {
	CommandId          common.String
	InvokeId           common.String
	InstanceId         common.String
	FinishedTime       common.String
	Output             common.String
	InvokeRecordStatus common.String
	ExitCode           common.Long
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
