package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	KeyMetadata DescribeKeyKeyMetadata
}

type DescribeKeyKeyMetadata struct {
	CreationDate       string
	Description        string
	KeyId              string
	KeyState           string
	KeyUsage           string
	DeleteDate         string
	Creator            string
	Arn                string
	Origin             string
	MaterialExpireTime string
}
