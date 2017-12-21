package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyHaVipAttributeRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyHaVipAttributeRequest) Invoke(client *sdk.Client) (response *ModifyHaVipAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyHaVipAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyHaVipAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyHaVipAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyHaVipAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyHaVipAttributeResponse struct {
	RequestId string
}
