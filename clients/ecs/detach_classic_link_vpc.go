package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachClassicLinkVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DetachClassicLinkVpcRequest) Invoke(client *sdk.Client) (response *DetachClassicLinkVpcResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachClassicLinkVpcRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachClassicLinkVpc", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DetachClassicLinkVpcResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachClassicLinkVpcResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachClassicLinkVpcResponse struct {
	RequestId string
}
