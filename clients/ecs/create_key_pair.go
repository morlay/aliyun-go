package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateKeyPairRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateKeyPairRequest) Invoke(client *sdk.Client) (resp *CreateKeyPairResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateKeyPair", "ecs", "")
	resp = &CreateKeyPairResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateKeyPairResponse struct {
	responses.BaseResponse
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
	PrivateKeyBody     string
}
