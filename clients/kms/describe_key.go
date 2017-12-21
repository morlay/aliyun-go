package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (r DescribeKeyRequest) Invoke(client *sdk.Client) (response *DescribeKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "DescribeKey", "kms", "")

	resp := struct {
		*responses.BaseResponse
		DescribeKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeKeyResponse struct {
	RequestId   string
	KeyMetadata DescribeKeyKeyMetadata
}

type DescribeKeyKeyMetadata struct {
	CreationDate string
	Description  string
	KeyId        string
	KeyState     string
	KeyUsage     string
	DeleteDate   string
	Creator      string
	Arn          string
}
