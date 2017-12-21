package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportKeyPairRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PublicKeyBody        string `position:"Query" name:"PublicKeyBody"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ImportKeyPairRequest) Invoke(client *sdk.Client) (resp *ImportKeyPairResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ImportKeyPair", "ecs", "")
	resp = &ImportKeyPairResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportKeyPairResponse struct {
	responses.BaseResponse
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
}
