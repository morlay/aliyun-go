package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsSubscriptionUpdateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ReadEnable   string `position:"Query" name:"ReadEnable"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
}

func (req *OnsSubscriptionUpdateRequest) Invoke(client *sdk.Client) (resp *OnsSubscriptionUpdateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionUpdate", "", "")
	resp = &OnsSubscriptionUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsSubscriptionUpdateResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}
