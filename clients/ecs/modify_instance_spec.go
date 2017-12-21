package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceSpecRequest struct {
	ResourceOwnerId                  int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount             string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                      string `position:"Query" name:"ClientToken"`
	AllowMigrateAcrossZone           string `position:"Query" name:"AllowMigrateAcrossZone"`
	OwnerAccount                     string `position:"Query" name:"OwnerAccount"`
	InternetMaxBandwidthOut          int    `position:"Query" name:"InternetMaxBandwidthOut"`
	OwnerId                          int64  `position:"Query" name:"OwnerId"`
	TemporaryInternetMaxBandwidthOut int    `position:"Query" name:"TemporaryInternetMaxBandwidthOut"`
	SystemDiskCategory               string `position:"Query" name:"SystemDiskCategory"`
	TemporaryStartTime               string `position:"Query" name:"TemporaryStartTime"`
	Async                            string `position:"Query" name:"Async"`
	InstanceId                       string `position:"Query" name:"InstanceId"`
	InstanceType                     string `position:"Query" name:"InstanceType"`
	TemporaryEndTime                 string `position:"Query" name:"TemporaryEndTime"`
	InternetMaxBandwidthIn           int    `position:"Query" name:"InternetMaxBandwidthIn"`
}

func (r ModifyInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceSpec", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceSpecResponse struct {
	RequestId string
}
