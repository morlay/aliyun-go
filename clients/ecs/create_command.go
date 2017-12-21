package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCommandRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string `position:"Query" name:"WorkingDir"`
	Description          string `position:"Query" name:"Description"`
	Type                 string `position:"Query" name:"Type"`
	CommandContent       string `position:"Query" name:"CommandContent"`
	Timeout              int64  `position:"Query" name:"Timeout"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
}

func (r CreateCommandRequest) Invoke(client *sdk.Client) (response *CreateCommandResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateCommandRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateCommand", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateCommandResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateCommandResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateCommandResponse struct {
	RequestId string
	CommandId string
}
