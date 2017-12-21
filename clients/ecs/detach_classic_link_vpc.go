package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachClassicLinkVpcRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DetachClassicLinkVpcRequest) Invoke(client *sdk.Client) (resp *DetachClassicLinkVpcResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachClassicLinkVpc", "ecs", "")
	resp = &DetachClassicLinkVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachClassicLinkVpcResponse struct {
	responses.BaseResponse
	RequestId string
}
