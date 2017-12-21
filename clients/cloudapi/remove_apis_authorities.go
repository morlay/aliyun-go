package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveApisAuthoritiesRequest struct {
	GroupId     string `position:"Query" name:"GroupId"`
	AppId       int64  `position:"Query" name:"AppId"`
	StageName   string `position:"Query" name:"StageName"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	Description string `position:"Query" name:"Description"`
}

func (r RemoveApisAuthoritiesRequest) Invoke(client *sdk.Client) (response *RemoveApisAuthoritiesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveApisAuthoritiesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveApisAuthorities", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveApisAuthoritiesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveApisAuthoritiesResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveApisAuthoritiesResponse struct {
	RequestId string
}
