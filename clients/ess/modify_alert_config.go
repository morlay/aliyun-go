package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyAlertConfigRequest struct {
	requests.RpcRequest
	SuccessConfig        int    `position:"Query" name:"SuccessConfig"`
	RejectConfig         int    `position:"Query" name:"RejectConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FailConfig           int    `position:"Query" name:"FailConfig"`
}

func (req *ModifyAlertConfigRequest) Invoke(client *sdk.Client) (resp *ModifyAlertConfigResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyAlertConfig", "ess", "")
	resp = &ModifyAlertConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAlertConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
