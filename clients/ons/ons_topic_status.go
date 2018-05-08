package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsTopicStatusRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsTopicStatusRequest) Invoke(client *sdk.Client) (resp *OnsTopicStatusResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicStatus", "", "")
	resp = &OnsTopicStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsTopicStatusData
}

type OnsTopicStatusData struct {
	TotalCount    common.Long
	LastTimeStamp common.Long
}
