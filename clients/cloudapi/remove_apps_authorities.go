package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveAppsAuthoritiesRequest struct {
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
	AppIds    string `position:"Query" name:"AppIds"`
}

func (r RemoveAppsAuthoritiesRequest) Invoke(client *sdk.Client) (response *RemoveAppsAuthoritiesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveAppsAuthoritiesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveAppsAuthorities", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveAppsAuthoritiesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveAppsAuthoritiesResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveAppsAuthoritiesResponse struct {
	RequestId string
}
