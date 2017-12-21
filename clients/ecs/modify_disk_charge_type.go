package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDiskChargeTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDiskChargeTypeRequest) Invoke(client *sdk.Client) (response *ModifyDiskChargeTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDiskChargeTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDiskChargeType", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDiskChargeTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDiskChargeTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDiskChargeTypeResponse struct {
	RequestId string
	OrderId   string
}
