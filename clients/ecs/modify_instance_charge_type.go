package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceChargeTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	DryRun               string `position:"Query" name:"DryRun"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	IncludeDataDisks     string `position:"Query" name:"IncludeDataDisks"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
}

func (req *ModifyInstanceChargeTypeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceChargeTypeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceChargeType", "ecs", "")
	resp = &ModifyInstanceChargeTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceChargeTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
}
