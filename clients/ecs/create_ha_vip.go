package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateHaVipRequest struct {
	requests.RpcRequest
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateHaVipRequest) Invoke(client *sdk.Client) (resp *CreateHaVipResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateHaVip", "ecs", "")
	resp = &CreateHaVipResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateHaVipResponse struct {
	responses.BaseResponse
	RequestId common.String
	HaVipId   common.String
}
