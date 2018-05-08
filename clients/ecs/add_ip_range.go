package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddIpRangeRequest struct {
	requests.RpcRequest
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AddIpRangeRequest) Invoke(client *sdk.Client) (resp *AddIpRangeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AddIpRange", "ecs", "")
	resp = &AddIpRangeResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddIpRangeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
