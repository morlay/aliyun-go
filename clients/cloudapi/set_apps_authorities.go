package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetAppsAuthoritiesRequest struct {
	GroupId     string `position:"Query" name:"GroupId"`
	ApiId       string `position:"Query" name:"ApiId"`
	StageName   string `position:"Query" name:"StageName"`
	AppIds      string `position:"Query" name:"AppIds"`
	Description string `position:"Query" name:"Description"`
}

func (r SetAppsAuthoritiesRequest) Invoke(client *sdk.Client) (response *SetAppsAuthoritiesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetAppsAuthoritiesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetAppsAuthorities", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetAppsAuthoritiesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetAppsAuthoritiesResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetAppsAuthoritiesResponse struct {
	RequestId string
}
