package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSignatureRequest struct {
	SignatureId string `position:"Query" name:"SignatureId"`
}

func (r DeleteSignatureRequest) Invoke(client *sdk.Client) (response *DeleteSignatureResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSignatureRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteSignature", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSignatureResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSignatureResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSignatureResponse struct {
	RequestId string
}
