package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsTopicUpdateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Perm         int    `position:"Query" name:"Perm"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsTopicUpdateRequest) Invoke(client *sdk.Client) (resp *OnsTopicUpdateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicUpdate", "", "")
	resp = &OnsTopicUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicUpdateResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}
