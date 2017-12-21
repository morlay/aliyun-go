package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySignatureRequest struct {
	SignatureId     string `position:"Query" name:"SignatureId"`
	SignatureName   string `position:"Query" name:"SignatureName"`
	SignatureKey    string `position:"Query" name:"SignatureKey"`
	SignatureSecret string `position:"Query" name:"SignatureSecret"`
}

func (r ModifySignatureRequest) Invoke(client *sdk.Client) (response *ModifySignatureResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySignatureRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifySignature", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ModifySignatureResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySignatureResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySignatureResponse struct {
	RequestId     string
	SignatureId   string
	SignatureName string
}
