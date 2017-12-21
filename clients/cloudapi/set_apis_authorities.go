package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetApisAuthoritiesRequest struct {
	GroupId     string `position:"Query" name:"GroupId"`
	AppId       int64  `position:"Query" name:"AppId"`
	StageName   string `position:"Query" name:"StageName"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	Description string `position:"Query" name:"Description"`
}

func (r SetApisAuthoritiesRequest) Invoke(client *sdk.Client) (response *SetApisAuthoritiesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetApisAuthoritiesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetApisAuthorities", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetApisAuthoritiesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetApisAuthoritiesResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetApisAuthoritiesResponse struct {
	RequestId string
}
