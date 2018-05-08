package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateSignatureRequest struct {
	requests.RpcRequest
	SignatureName   string `position:"Query" name:"SignatureName"`
	SignatureKey    string `position:"Query" name:"SignatureKey"`
	SignatureSecret string `position:"Query" name:"SignatureSecret"`
}

func (req *CreateSignatureRequest) Invoke(client *sdk.Client) (resp *CreateSignatureResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateSignature", "apigateway", "")
	resp = &CreateSignatureResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSignatureResponse struct {
	responses.BaseResponse
	RequestId     common.String
	SignatureId   common.String
	SignatureName common.String
}
