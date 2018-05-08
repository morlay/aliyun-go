package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveAppsAuthoritiesRequest struct {
	requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
	AppIds    string `position:"Query" name:"AppIds"`
}

func (req *RemoveAppsAuthoritiesRequest) Invoke(client *sdk.Client) (resp *RemoveAppsAuthoritiesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveAppsAuthorities", "apigateway", "")
	resp = &RemoveAppsAuthoritiesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveAppsAuthoritiesResponse struct {
	responses.BaseResponse
	RequestId common.String
}
