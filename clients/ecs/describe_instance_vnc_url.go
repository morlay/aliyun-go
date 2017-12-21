package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceVncUrlRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceVncUrlRequest) Invoke(client *sdk.Client) (response *DescribeInstanceVncUrlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceVncUrlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceVncUrl", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceVncUrlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceVncUrlResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceVncUrlResponse struct {
	RequestId string
	VncUrl    string
}
