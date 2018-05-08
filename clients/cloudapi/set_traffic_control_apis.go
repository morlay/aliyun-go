package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetTrafficControlApisRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	GroupId          string `position:"Query" name:"GroupId"`
	ApiIds           string `position:"Query" name:"ApiIds"`
	StageName        string `position:"Query" name:"StageName"`
}

func (req *SetTrafficControlApisRequest) Invoke(client *sdk.Client) (resp *SetTrafficControlApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetTrafficControlApis", "apigateway", "")
	resp = &SetTrafficControlApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetTrafficControlApisResponse struct {
	responses.BaseResponse
	RequestId common.String
}
