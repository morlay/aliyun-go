package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAppKeySecretRequest struct {
	AppKey string `position:"Query" name:"AppKey"`
}

func (r ResetAppKeySecretRequest) Invoke(client *sdk.Client) (response *ResetAppKeySecretResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetAppKeySecretRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ResetAppKeySecret", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ResetAppKeySecretResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetAppKeySecretResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetAppKeySecretResponse struct {
	RequestId string
}
