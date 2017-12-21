package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceAutoReleaseTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      string `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyInstanceAutoReleaseTimeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceAutoReleaseTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceAutoReleaseTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAutoReleaseTime", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceAutoReleaseTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceAutoReleaseTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceAutoReleaseTimeResponse struct {
	RequestId string
}
