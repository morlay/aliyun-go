package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateApiRequest struct {
	GroupId              string `position:"Query" name:"GroupId"`
	ApiName              string `position:"Query" name:"ApiName"`
	Visibility           string `position:"Query" name:"Visibility"`
	Description          string `position:"Query" name:"Description"`
	AuthType             string `position:"Query" name:"AuthType"`
	RequestConfig        string `position:"Query" name:"RequestConfig"`
	ServiceConfig        string `position:"Query" name:"ServiceConfig"`
	RequestParameters    string `position:"Query" name:"RequestParameters"`
	SystemParameters     string `position:"Query" name:"SystemParameters"`
	ConstantParameters   string `position:"Query" name:"ConstantParameters"`
	ServiceParameters    string `position:"Query" name:"ServiceParameters"`
	ServiceParametersMap string `position:"Query" name:"ServiceParametersMap"`
	ResultType           string `position:"Query" name:"ResultType"`
	ResultSample         string `position:"Query" name:"ResultSample"`
	FailResultSample     string `position:"Query" name:"FailResultSample"`
	ErrorCodeSamples     string `position:"Query" name:"ErrorCodeSamples"`
	OpenIdConnectConfig  string `position:"Query" name:"OpenIdConnectConfig"`
}

func (r CreateApiRequest) Invoke(client *sdk.Client) (response *CreateApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateApiResponse struct {
	RequestId string
	ApiId     string
}
