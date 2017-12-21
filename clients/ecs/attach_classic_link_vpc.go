package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachClassicLinkVpcRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AttachClassicLinkVpcRequest) Invoke(client *sdk.Client) (resp *AttachClassicLinkVpcResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachClassicLinkVpc", "ecs", "")
	resp = &AttachClassicLinkVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachClassicLinkVpcResponse struct {
	responses.BaseResponse
	RequestId string
}
