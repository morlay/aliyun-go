package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceNetworkSpecRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay                 string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken             string `position:"Query" name:"ClientToken"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	InternetMaxBandwidthOut int    `position:"Query" name:"InternetMaxBandwidthOut"`
	EndTime                 string `position:"Query" name:"EndTime"`
	StartTime               string `position:"Query" name:"StartTime"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	InstanceId              string `position:"Query" name:"InstanceId"`
	NetworkChargeType       string `position:"Query" name:"NetworkChargeType"`
	InternetMaxBandwidthIn  int    `position:"Query" name:"InternetMaxBandwidthIn"`
	AllocatePublicIp        string `position:"Query" name:"AllocatePublicIp"`
}

func (r ModifyInstanceNetworkSpecRequest) Invoke(client *sdk.Client) (response *ModifyInstanceNetworkSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceNetworkSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceNetworkSpec", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceNetworkSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceNetworkSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceNetworkSpecResponse struct {
	RequestId string
	OrderId   string
}
