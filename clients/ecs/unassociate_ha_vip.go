package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociateHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnassociateHaVipRequest) Invoke(client *sdk.Client) (response *UnassociateHaVipResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnassociateHaVipRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "UnassociateHaVip", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		UnassociateHaVipResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnassociateHaVipResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnassociateHaVipResponse struct {
	RequestId string
}
