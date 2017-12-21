package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindIpRangeRequest struct {
	requests.RpcRequest
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UnbindIpRangeRequest) Invoke(client *sdk.Client) (resp *UnbindIpRangeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "UnbindIpRange", "ecs", "")
	resp = &UnbindIpRangeResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnbindIpRangeResponse struct {
	responses.BaseResponse
	RequestId string
}
