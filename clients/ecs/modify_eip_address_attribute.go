package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyEipAddressAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyEipAddressAttributeRequest) Invoke(client *sdk.Client) (response *ModifyEipAddressAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyEipAddressAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyEipAddressAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyEipAddressAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyEipAddressAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyEipAddressAttributeResponse struct {
	RequestId string
}
