package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveTrafficControlApisRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	GroupId          string `position:"Query" name:"GroupId"`
	ApiIds           string `position:"Query" name:"ApiIds"`
	StageName        string `position:"Query" name:"StageName"`
}

func (req *RemoveTrafficControlApisRequest) Invoke(client *sdk.Client) (resp *RemoveTrafficControlApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveTrafficControlApis", "apigateway", "")
	resp = &RemoveTrafficControlApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveTrafficControlApisResponse struct {
	responses.BaseResponse
	RequestId string
}
