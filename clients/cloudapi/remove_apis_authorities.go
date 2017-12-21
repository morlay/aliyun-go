package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveApisAuthoritiesRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	AppId       int64  `position:"Query" name:"AppId"`
	StageName   string `position:"Query" name:"StageName"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	Description string `position:"Query" name:"Description"`
}

func (req *RemoveApisAuthoritiesRequest) Invoke(client *sdk.Client) (resp *RemoveApisAuthoritiesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveApisAuthorities", "apigateway", "")
	resp = &RemoveApisAuthoritiesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveApisAuthoritiesResponse struct {
	responses.BaseResponse
	RequestId string
}
