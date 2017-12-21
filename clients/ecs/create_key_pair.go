package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateKeyPairRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateKeyPairRequest) Invoke(client *sdk.Client) (response *CreateKeyPairResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateKeyPairRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateKeyPair", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateKeyPairResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateKeyPairResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateKeyPairResponse struct {
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
	PrivateKeyBody     string
}
