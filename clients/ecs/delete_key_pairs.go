package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteKeyPairsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairNames         string `position:"Query" name:"KeyPairNames"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteKeyPairsRequest) Invoke(client *sdk.Client) (response *DeleteKeyPairsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteKeyPairsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteKeyPairs", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteKeyPairsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteKeyPairsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteKeyPairsResponse struct {
	RequestId string
}
