package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetParametersForImportRequest struct {
	requests.RpcRequest
	KeyId             string `position:"Query" name:"KeyId"`
	STSToken          string `position:"Query" name:"STSToken"`
	WrappingAlgorithm string `position:"Query" name:"WrappingAlgorithm"`
	WrappingKeySpec   string `position:"Query" name:"WrappingKeySpec"`
}

func (req *GetParametersForImportRequest) Invoke(client *sdk.Client) (resp *GetParametersForImportResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "GetParametersForImport", "kms", "")
	resp = &GetParametersForImportResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetParametersForImportResponse struct {
	responses.BaseResponse
	KeyId           string
	RequestId       string
	ImportToken     string
	PublicKey       string
	TokenExpireTime string
}
