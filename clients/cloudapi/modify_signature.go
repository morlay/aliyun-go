package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySignatureRequest struct {
	requests.RpcRequest
	SignatureId     string `position:"Query" name:"SignatureId"`
	SignatureName   string `position:"Query" name:"SignatureName"`
	SignatureKey    string `position:"Query" name:"SignatureKey"`
	SignatureSecret string `position:"Query" name:"SignatureSecret"`
}

func (req *ModifySignatureRequest) Invoke(client *sdk.Client) (resp *ModifySignatureResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifySignature", "apigateway", "")
	resp = &ModifySignatureResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySignatureResponse struct {
	responses.BaseResponse
	RequestId     common.String
	SignatureId   common.String
	SignatureName common.String
}
