package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetApisAuthoritiesRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	AppId       int64  `position:"Query" name:"AppId"`
	StageName   string `position:"Query" name:"StageName"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	Description string `position:"Query" name:"Description"`
}

func (req *SetApisAuthoritiesRequest) Invoke(client *sdk.Client) (resp *SetApisAuthoritiesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetApisAuthorities", "apigateway", "")
	resp = &SetApisAuthoritiesResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetApisAuthoritiesResponse struct {
	responses.BaseResponse
	RequestId common.String
}
