package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchApiRequest struct {
	requests.RpcRequest
	GroupId        string `position:"Query" name:"GroupId"`
	ApiId          string `position:"Query" name:"ApiId"`
	StageName      string `position:"Query" name:"StageName"`
	Description    string `position:"Query" name:"Description"`
	HistoryVersion string `position:"Query" name:"HistoryVersion"`
}

func (req *SwitchApiRequest) Invoke(client *sdk.Client) (resp *SwitchApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SwitchApi", "apigateway", "")
	resp = &SwitchApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type SwitchApiResponse struct {
	responses.BaseResponse
	RequestId string
}
