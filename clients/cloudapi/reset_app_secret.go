package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAppSecretRequest struct {
	AppKey string `position:"Query" name:"AppKey"`
}

func (r ResetAppSecretRequest) Invoke(client *sdk.Client) (response *ResetAppSecretResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetAppSecretRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ResetAppSecret", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ResetAppSecretResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetAppSecretResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetAppSecretResponse struct {
	RequestId string
}
