package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInvocationsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InvokeStatus         string `position:"Query" name:"InvokeStatus"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	Timed                string `position:"Query" name:"Timed"`
	CommandName          string `position:"Query" name:"CommandName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CommandType          string `position:"Query" name:"CommandType"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

func (req *DescribeInvocationsRequest) Invoke(client *sdk.Client) (resp *DescribeInvocationsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocations", "ecs", "")
	resp = &DescribeInvocationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInvocationsResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Long
	PageNumber  common.Long
	PageSize    common.Long
	Invocations DescribeInvocationsInvocationList
}

type DescribeInvocationsInvocation struct {
	InvokeId        common.String
	CommandId       common.String
	CommandType     common.String
	CommandName     common.String
	Frequency       common.String
	Timed           bool
	InvokeStatus    common.String
	InvokeInstances DescribeInvocationsInvokeInstanceList
}

type DescribeInvocationsInvokeInstance struct {
	InstanceId           common.String
	InstanceInvokeStatus common.String
}

type DescribeInvocationsInvocationList []DescribeInvocationsInvocation

func (list *DescribeInvocationsInvocationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvocation)
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

type DescribeInvocationsInvokeInstanceList []DescribeInvocationsInvokeInstance

func (list *DescribeInvocationsInvokeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvokeInstance)
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
