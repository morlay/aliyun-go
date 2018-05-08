package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateKeyRequest struct {
	requests.RpcRequest
	KeyUsage    string `position:"Query" name:"KeyUsage"`
	Origin      string `position:"Query" name:"Origin"`
	Description string `position:"Query" name:"Description"`
	STSToken    string `position:"Query" name:"STSToken"`
}

func (req *CreateKeyRequest) Invoke(client *sdk.Client) (resp *CreateKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "CreateKey", "kms", "")
	resp = &CreateKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateKeyResponse struct {
	responses.BaseResponse
	RequestId   common.String
	KeyMetadata CreateKeyKeyMetadata
}

type CreateKeyKeyMetadata struct {
	CreationDate       common.String
	Description        common.String
	KeyId              common.String
	KeyState           common.String
	KeyUsage           common.String
	DeleteDate         common.String
	Creator            common.String
	Arn                common.String
	Origin             common.String
	MaterialExpireTime common.String
}
