package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	Duration     int    `position:"Query" name:"Duration"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "RenewInstance", "", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	OrderId   common.String
	RequestId common.String
}
