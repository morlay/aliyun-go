package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsPublishDeleteRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
}

func (req *OnsPublishDeleteRequest) Invoke(client *sdk.Client) (resp *OnsPublishDeleteResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishDelete", "", "")
	resp = &OnsPublishDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsPublishDeleteResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}
