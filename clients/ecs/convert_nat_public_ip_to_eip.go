package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConvertNatPublicIpToEipRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

func (req *ConvertNatPublicIpToEipRequest) Invoke(client *sdk.Client) (resp *ConvertNatPublicIpToEipResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ConvertNatPublicIpToEip", "ecs", "")
	resp = &ConvertNatPublicIpToEipResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConvertNatPublicIpToEipResponse struct {
	responses.BaseResponse
	RequestId string
}
