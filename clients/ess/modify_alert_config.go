package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAlertConfigRequest struct {
	SuccessConfig        int    `position:"Query" name:"SuccessConfig"`
	RejectConfig         int    `position:"Query" name:"RejectConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FailConfig           int    `position:"Query" name:"FailConfig"`
}

func (r ModifyAlertConfigRequest) Invoke(client *sdk.Client) (response *ModifyAlertConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyAlertConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyAlertConfig", "ess", "")

	resp := struct {
		*responses.BaseResponse
		ModifyAlertConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyAlertConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyAlertConfigResponse struct {
	RequestId string
}
