package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachClassicLinkVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AttachClassicLinkVpcRequest) Invoke(client *sdk.Client) (response *AttachClassicLinkVpcResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachClassicLinkVpcRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachClassicLinkVpc", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AttachClassicLinkVpcResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachClassicLinkVpcResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachClassicLinkVpcResponse struct {
	RequestId string
}
