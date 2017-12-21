package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportKeyPairRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PublicKeyBody        string `position:"Query" name:"PublicKeyBody"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ImportKeyPairRequest) Invoke(client *sdk.Client) (response *ImportKeyPairResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ImportKeyPairRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ImportKeyPair", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ImportKeyPairResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImportKeyPairResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImportKeyPairResponse struct {
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
}
