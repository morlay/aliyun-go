package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyApiRequest struct {
	requests.RpcRequest
	GroupId              string `position:"Query" name:"GroupId"`
	ApiId                string `position:"Query" name:"ApiId"`
	ApiName              string `position:"Query" name:"ApiName"`
	Description          string `position:"Query" name:"Description"`
	Visibility           string `position:"Query" name:"Visibility"`
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
	AllowSignatureMethod string `position:"Query" name:"AllowSignatureMethod"`
}

func (req *ModifyApiRequest) Invoke(client *sdk.Client) (resp *ModifyApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApi", "apigateway", "")
	resp = &ModifyApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyApiResponse struct {
	responses.BaseResponse
	RequestId common.String
}
