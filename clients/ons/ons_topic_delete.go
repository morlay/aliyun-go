package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsTopicDeleteRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsTopicDeleteRequest) Invoke(client *sdk.Client) (resp *OnsTopicDeleteResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicDelete", "", "")
	resp = &OnsTopicDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicDeleteResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}
