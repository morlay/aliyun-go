package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetTrafficControlApisRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	GroupId          string `position:"Query" name:"GroupId"`
	ApiIds           string `position:"Query" name:"ApiIds"`
	StageName        string `position:"Query" name:"StageName"`
}

func (r SetTrafficControlApisRequest) Invoke(client *sdk.Client) (response *SetTrafficControlApisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetTrafficControlApisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetTrafficControlApis", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetTrafficControlApisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetTrafficControlApisResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetTrafficControlApisResponse struct {
	RequestId string
}
