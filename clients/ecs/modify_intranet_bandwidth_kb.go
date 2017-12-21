package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyIntranetBandwidthKbRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	IntranetMaxBandwidthOut int    `position:"Query" name:"IntranetMaxBandwidthOut"`
	InstanceId              string `position:"Query" name:"InstanceId"`
	IntranetMaxBandwidthIn  int    `position:"Query" name:"IntranetMaxBandwidthIn"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyIntranetBandwidthKbRequest) Invoke(client *sdk.Client) (response *ModifyIntranetBandwidthKbResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyIntranetBandwidthKbRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyIntranetBandwidthKb", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyIntranetBandwidthKbResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyIntranetBandwidthKbResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyIntranetBandwidthKbResponse struct {
	RequestId string
}
