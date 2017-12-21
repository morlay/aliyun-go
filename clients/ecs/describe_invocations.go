package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInvocationsRequest struct {
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

func (r DescribeInvocationsRequest) Invoke(client *sdk.Client) (response *DescribeInvocationsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInvocationsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocations", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInvocationsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInvocationsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInvocationsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Invocation DescribeInvocationsInvocationItemList
}

type DescribeInvocationsInvocationItem struct {
	InvokeId     string
	CommandId    string
	CommandType  string
	CommandName  string
	Timed        bool
	InvokeStatus string
	InvokeItem   DescribeInvocationsInvokeItemItemList
}

type DescribeInvocationsInvokeItemItem struct {
	InstanceId string
	Status     string
}

type DescribeInvocationsInvocationItemList []DescribeInvocationsInvocationItem

func (list *DescribeInvocationsInvocationItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvocationItem)
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

type DescribeInvocationsInvokeItemItemList []DescribeInvocationsInvokeItemItem

func (list *DescribeInvocationsInvokeItemItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvokeItemItem)
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
