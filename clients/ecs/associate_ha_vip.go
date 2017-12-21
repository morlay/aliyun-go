package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociateHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AssociateHaVipRequest) Invoke(client *sdk.Client) (response *AssociateHaVipResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AssociateHaVipRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AssociateHaVip", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AssociateHaVipResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AssociateHaVipResponse

	err = client.DoAction(&req, &resp)
	return
}

type AssociateHaVipResponse struct {
	RequestId string
}
