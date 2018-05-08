package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	KeyId           common.String
	RequestId       common.String
	ImportToken     common.String
	PublicKey       common.String
	TokenExpireTime common.String
}
