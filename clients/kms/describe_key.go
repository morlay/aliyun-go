package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeKeyRequest struct {
	requests.RpcRequest
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (req *DescribeKeyRequest) Invoke(client *sdk.Client) (resp *DescribeKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "DescribeKey", "kms", "")
	resp = &DescribeKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeKeyResponse struct {
	responses.BaseResponse
	RequestId   common.String
	KeyMetadata DescribeKeyKeyMetadata
}

type DescribeKeyKeyMetadata struct {
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
