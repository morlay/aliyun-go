package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceAutoRenewAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceAutoRenewAttributeRequest) Invoke(client *sdk.Client) (response *DescribeInstanceAutoRenewAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceAutoRenewAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceAutoRenewAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceAutoRenewAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceAutoRenewAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId               string
	InstanceRenewAttributes DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttribute struct {
	InstanceId       string
	AutoRenewEnabled bool
	Duration         int
	PeriodUnit       string
	RenewalStatus    string
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList []DescribeInstanceAutoRenewAttributeInstanceRenewAttribute

func (list *DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeInstanceRenewAttribute)
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
