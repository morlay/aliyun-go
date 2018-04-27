package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteKeyMaterialRequest struct {
	requests.RpcRequest
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (req *DeleteKeyMaterialRequest) Invoke(client *sdk.Client) (resp *DeleteKeyMaterialResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "DeleteKeyMaterial", "kms", "")
	resp = &DeleteKeyMaterialResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteKeyMaterialResponse struct {
	responses.BaseResponse
	RequestId string
}
