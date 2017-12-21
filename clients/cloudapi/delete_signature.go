package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSignatureRequest struct {
	requests.RpcRequest
	SignatureId string `position:"Query" name:"SignatureId"`
}

func (req *DeleteSignatureRequest) Invoke(client *sdk.Client) (resp *DeleteSignatureResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteSignature", "apigateway", "")
	resp = &DeleteSignatureResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSignatureResponse struct {
	responses.BaseResponse
	RequestId string
}
