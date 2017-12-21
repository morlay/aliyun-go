package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateHaVipRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateHaVipRequest) Invoke(client *sdk.Client) (response *CreateHaVipResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateHaVipRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateHaVip", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateHaVipResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateHaVipResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateHaVipResponse struct {
	RequestId string
	HaVipId   string
}
