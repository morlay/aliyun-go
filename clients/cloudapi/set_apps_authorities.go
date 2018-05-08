package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetAppsAuthoritiesRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	ApiId       string `position:"Query" name:"ApiId"`
	StageName   string `position:"Query" name:"StageName"`
	AppIds      string `position:"Query" name:"AppIds"`
	Description string `position:"Query" name:"Description"`
}

func (req *SetAppsAuthoritiesRequest) Invoke(client *sdk.Client) (resp *SetAppsAuthoritiesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetAppsAuthorities", "apigateway", "")
	resp = &SetAppsAuthoritiesResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAppsAuthoritiesResponse struct {
	responses.BaseResponse
	RequestId common.String
}
