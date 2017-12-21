package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceVncPasswdRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceVncPasswdRequest) Invoke(client *sdk.Client) (response *DescribeInstanceVncPasswdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceVncPasswdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceVncPasswd", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceVncPasswdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceVncPasswdResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceVncPasswdResponse struct {
	RequestId string
	VncPasswd string
}
