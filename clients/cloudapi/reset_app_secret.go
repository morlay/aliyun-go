package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAppSecretRequest struct {
	requests.RpcRequest
	AppKey string `position:"Query" name:"AppKey"`
}

func (req *ResetAppSecretRequest) Invoke(client *sdk.Client) (resp *ResetAppSecretResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ResetAppSecret", "apigateway", "")
	resp = &ResetAppSecretResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAppSecretResponse struct {
	responses.BaseResponse
	RequestId string
}
