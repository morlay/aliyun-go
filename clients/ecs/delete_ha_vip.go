package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteHaVipRequest) Invoke(client *sdk.Client) (response *DeleteHaVipResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteHaVipRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteHaVip", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteHaVipResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteHaVipResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteHaVipResponse struct {
	RequestId string
}
