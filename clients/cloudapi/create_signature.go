package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSignatureRequest struct {
	SignatureName   string `position:"Query" name:"SignatureName"`
	SignatureKey    string `position:"Query" name:"SignatureKey"`
	SignatureSecret string `position:"Query" name:"SignatureSecret"`
}

func (r CreateSignatureRequest) Invoke(client *sdk.Client) (response *CreateSignatureResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateSignatureRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateSignature", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateSignatureResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateSignatureResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateSignatureResponse struct {
	RequestId     string
	SignatureId   string
	SignatureName string
}
