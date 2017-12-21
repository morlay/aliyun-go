package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveTrafficControlApisRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	GroupId          string `position:"Query" name:"GroupId"`
	ApiIds           string `position:"Query" name:"ApiIds"`
	StageName        string `position:"Query" name:"StageName"`
}

func (r RemoveTrafficControlApisRequest) Invoke(client *sdk.Client) (response *RemoveTrafficControlApisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveTrafficControlApisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveTrafficControlApis", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveTrafficControlApisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveTrafficControlApisResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveTrafficControlApisResponse struct {
	RequestId string
}
