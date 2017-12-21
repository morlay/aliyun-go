package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancePhysicalAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeInstancePhysicalAttributeRequest) Invoke(client *sdk.Client) (response *DescribeInstancePhysicalAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstancePhysicalAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstancePhysicalAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstancePhysicalAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstancePhysicalAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstancePhysicalAttributeResponse struct {
	RequestId        string
	InstanceId       string
	VlanId           string
	NodeControllerId string
	RackId           string
}
