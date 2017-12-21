package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteKeyPairsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairNames         string `position:"Query" name:"KeyPairNames"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteKeyPairsRequest) Invoke(client *sdk.Client) (resp *DeleteKeyPairsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteKeyPairs", "ecs", "")
	resp = &DeleteKeyPairsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteKeyPairsResponse struct {
	responses.BaseResponse
	RequestId string
}
